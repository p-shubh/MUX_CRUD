package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Master struct {
	Status         string
	Status_code    string
	Count          int
	RoleAccessInfo []Module
}

type Module struct {
	Name           string
	ID             int
	Moduleicon     string
	DisplayIndex   int
	PageCount      int
	SubModuleCOunt int
	SubModule      []SubModuleInfo
	Pages          []PagesInfo
}

type SubModuleInfo struct {
	SUbModuleName         string
	SubModuleID           int
	Submoduledisplayindex *int
	PageCount             int
	Pages                 []PagesInfo
}

type PagesInfo struct {
	PageName        string
	PageID          int
	DisplayIndex    string
	ActualPagePath  string
	Relatedpagepath string
	Relatedpageid   int
	Isshow          bool
	Isshowaddpage   bool
	Isadd           bool
	Isedit          bool
	Iseditself      bool
	Isview          bool
	Isdeleteself    bool
	Isdelete        bool
	Isprint         bool
}

type PageAccess struct {
	Count int
}

func (db *DB_manager) ReadRoleaccessRecords(UserID int, CompanyID int) (Master, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT H_Modules.Name,H_Modules.Moduleicon,H_Modules.ID , H_Modules.DisplayIndex FROM M_RoleAccess JOIN H_Modules on M_RoleAccess.ModuleID = H_Modules.ID LEFT JOIN H_SubModules on M_RoleAccess.SubModuleID = H_SubModules.ID JOIN H_Pages on M_RoleAccess.PageID = H_Pages.ID LEFT JOIN (SELECT * FROM H_Pages) A ON A.ID=H_Pages.PageID WHERE M_RoleAccess.RoleID in (SELECT  RoleID FROM MC_UserRoles WHERE UserID = %d) AND H_Pages.ID not in(SELECT PageID FROM M_UserAccess WHERE UserID = %d) AND CompanyID=%d Order By H_Modules.DisplayIndex", UserID, UserID, CompanyID))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from role access records", err)
		return Master{}, err
	}

	var RoleaccessArray []Module
	var RoleaccessRec Module
	for rows.Next() {
		err := rows.Scan(&RoleaccessRec.Name, &RoleaccessRec.Moduleicon, &RoleaccessRec.ID, &RoleaccessRec.DisplayIndex)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		////////////////////////////////////
		var SubModuleArray []SubModuleInfo
		var SubModulerec SubModuleInfo

		rows1, err := db.Query(fmt.Sprintf("SELECT DISTINCT COALESCE(H_SubModules.Name,'') SubModuleName, COALESCE(H_SubModules.ID,0) SubModuleID, H_SubModules.DisplayIndex SubModuleDisplayIndex FROM M_RoleAccess JOIN H_Modules on M_RoleAccess.ModuleID = H_Modules.ID LEFT JOIN H_SubModules on M_RoleAccess.SubModuleID = H_SubModules.ID JOIN H_Pages on M_RoleAccess.PageID = H_Pages.ID LEFT JOIN (SELECT * FROM H_Pages) A ON A.ID=H_Pages.PageID WHERE M_RoleAccess.RoleID in (SELECT  RoleID FROM MC_UserRoles WHERE UserID = %d) AND H_Pages.ID not in(SELECT PageID FROM M_UserAccess WHERE UserID = %d) AND CompanyID=%d and H_Modules.ID=%d Order By H_SubModules.DisplayIndex", UserID, UserID, CompanyID, RoleaccessRec.ID))
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from role access sub module records ", err)
			return Master{}, err
		}

		for rows1.Next() {
			err = rows1.Scan(&SubModulerec.SUbModuleName, &SubModulerec.SubModuleID, &SubModulerec.Submoduledisplayindex)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			if SubModulerec.SubModuleID != 0 {

				/////////////////////////////////////////sub module pages//////////////////////////////////////////
				var SubModulePage []PagesInfo
				var SubModulepagerec PagesInfo

				rows2, err := db.Query(fmt.Sprintf("SELECT DISTINCT H_Pages.Name, H_Pages.ID,H_Pages.DisplayIndex , H_Pages.ActualPagePath,COALESCE(A.ActualPagePath,'') RelatedPagePath,COALESCE(A.ID,0) RelatedPageID, IsShow, IsShowAddPage, IsAdd, IsEdit, IsEditSelf, IsView, IsDeleteSelf, IsDelete, IsPrint FROM M_RoleAccess JOIN H_Modules on M_RoleAccess.ModuleID = H_Modules.ID LEFT JOIN H_SubModules on M_RoleAccess.SubModuleID = H_SubModules.ID JOIN H_Pages on M_RoleAccess.PageID = H_Pages.ID LEFT JOIN (SELECT * FROM H_Pages) A ON A.ID=H_Pages.PageID WHERE M_RoleAccess.RoleID in (SELECT  RoleID FROM MC_UserRoles WHERE UserID = %d) AND H_Pages.ID not in(SELECT PageID FROM M_UserAccess WHERE UserID = %d) AND CompanyID=%d and H_Modules.ID=%d AND H_SubModules.ID=%d Order By H_Pages.DisplayIndex", UserID, UserID, CompanyID, RoleaccessRec.ID, SubModulerec.SubModuleID))
				defer rows2.Close()
				if err != nil {
					fmt.Println("failed to get data from role access sub module records ", err)
					return Master{}, err
				}

				for rows2.Next() {
					err = rows2.Scan(&SubModulepagerec.PageName, &SubModulepagerec.PageID, &SubModulepagerec.DisplayIndex, &SubModulepagerec.ActualPagePath, &SubModulepagerec.Relatedpagepath, &SubModulepagerec.Relatedpageid, &SubModulepagerec.Isshow, &SubModulepagerec.Isshowaddpage, &SubModulepagerec.Isadd, &SubModulepagerec.Isedit, &SubModulepagerec.Iseditself, &SubModulepagerec.Isview, &SubModulepagerec.Isdeleteself, &SubModulepagerec.Isdelete, &SubModulepagerec.Isprint)
					if err != nil {
						fmt.Println("failed to scan the record.. continue with the next.. ", err)
						continue
					}
					if SubModulepagerec.PageID != 0 {
						SubModulePage = append(SubModulePage, SubModulepagerec)
					}
				}

				///////////////////////////////////////////////////sub module pages end ///////////////////

				SubModulerec.PageCount = len(SubModulePage)
				SubModulerec.Pages = SubModulePage

				SubModuleArray = append(SubModuleArray, SubModulerec)
			}
		}

		/////////////////////////////////////////////
		///////////////////////////////////////// module pages//////////////////////////////////////////
		var ModulePage []PagesInfo
		var Modulepagerec PagesInfo

		rows3, err := db.Query(fmt.Sprintf("SELECT DISTINCT H_Pages.Name, H_Pages.ID,H_Pages.DisplayIndex , H_Pages.ActualPagePath,COALESCE(A.ActualPagePath,'') RelatedPagePath,COALESCE(A.ID,0) RelatedPageID, IsShow, IsShowAddPage, IsAdd, IsEdit, IsEditSelf, IsView, IsDeleteSelf, IsDelete, IsPrint FROM M_RoleAccess JOIN H_Modules on M_RoleAccess.ModuleID = H_Modules.ID LEFT JOIN H_SubModules on M_RoleAccess.SubModuleID = H_SubModules.ID JOIN H_Pages on M_RoleAccess.PageID = H_Pages.ID LEFT JOIN (SELECT * FROM H_Pages) A ON A.ID=H_Pages.PageID WHERE M_RoleAccess.RoleID in (SELECT  RoleID FROM MC_UserRoles WHERE UserID = %d) AND H_Pages.ID not in(SELECT PageID FROM M_UserAccess WHERE UserID = %d) AND CompanyID=%d and H_Modules.ID=%d  AND M_RoleAccess.SubModuleID<=0 Order By H_Pages.DisplayIndex", UserID, UserID, CompanyID, RoleaccessRec.ID))
		defer rows3.Close()
		if err != nil {
			fmt.Println("failed to get data from role access sub module records ", err)
			return Master{}, err
		}

		for rows3.Next() {
			err = rows3.Scan(&Modulepagerec.PageName, &Modulepagerec.PageID, &Modulepagerec.DisplayIndex, &Modulepagerec.ActualPagePath, &Modulepagerec.Relatedpagepath, &Modulepagerec.Relatedpageid, &Modulepagerec.Isshow, &Modulepagerec.Isshowaddpage, &Modulepagerec.Isadd, &Modulepagerec.Isedit, &Modulepagerec.Iseditself, &Modulepagerec.Isview, &Modulepagerec.Isdeleteself, &Modulepagerec.Isdelete, &Modulepagerec.Isprint)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			if Modulepagerec.PageID != 0 {
				ModulePage = append(ModulePage, Modulepagerec)
			}
		}

		/////////////////////////////////////////////////// module pages end ///////////////////
		RoleaccessRec.SubModuleCOunt = len(SubModuleArray)
		RoleaccessRec.PageCount = len(ModulePage)
		RoleaccessRec.SubModule = SubModuleArray
		RoleaccessRec.Pages = ModulePage
		RoleaccessArray = append(RoleaccessArray, RoleaccessRec)
	}
	var Masterroot Master
	Masterroot.Status = "true"
	Masterroot.Status_code = "200"
	Masterroot.Count = len(RoleaccessArray)
	Masterroot.RoleAccessInfo = RoleaccessArray
	return Masterroot, nil
}

func (db *DB_manager) ReadCheckPageAccessRecords(UserID int, PageID int) int {
	rows, err := db.Query(fmt.Sprintf("Select COUNT(*) FROM(Select m_users.UserID, pageid from m_users JOIN mc_userroles on m_users.UserID = mc_userroles.userid JOIN m_roleaccess on m_roleaccess.roleid = mc_userroles.roleid where m_users.UserID = %d AND pageid = %d UNION Select m_users.UserID,pageid from m_users JOIN m_useraccess on m_users.UserID = m_useraccess.userid where m_users.UserID = %d AND pageid = %d)A", UserID, PageID, UserID, PageID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from check page access ", err)
		return 0
	}

	var PageaccessRec int
	for rows.Next() {
		err := rows.Scan(&PageaccessRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		if PageaccessRec > 0 {
			return 1
		} else {
			return 0
		}
	}
	return 0
}
