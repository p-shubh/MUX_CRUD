package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type UserAccessMaster struct {
	Status         string
	Status_code    string
	Count          int
	UserMasterInfo []Usermaster
}

type Usermaster struct {
	Userid    int
	Firstname string
	Lastname  string
	Loginname string
	Mobile_No *string
	Emailid   *string
	Password  string
}

type Useraccessmaster struct {
	Userid               int
	Firstname            string
	Lastname             string
	Loginname            string
	Mobile_No            string
	Emailid              string
	Password             string
	MC_UserRolesInfo     []MC_UserRoles
	MC_UserDivisionsInfo []MC_UserDivisions
}

type MC_UserRoles struct {
	MC_UserRoleID int
	UserID        int
	RoleID        int
	RoleName      string
	Row_Order     int
}

type MC_UserDivisions struct {
	MC_UserDivisionID int
	UserID            int
	DepartmentID      int
	DepartmentName    string
	Row_Order         int
}

type User_Access_Master_Info struct {
	UserAccessMasterID int
	UserID             int
	UserName           string
	CompanyID          int
	DivisionID         int
	ModuleID           int
	ModuleName         string
	SubModuleID        int
	SubModuleName      string
	PageID             int
	PageName           string
	Isshow             int
	Isshowaddpage      int
	Isadd              int
	Iseditself         int
	Isedit             int
	Isview             int
	Isdeleteself       int
	Isdelete           int
	Isprint            int
}

type CheckDuplicateRecValidation struct {
	Count  int
	Status string
}

// GET User Access Master List API
func (db *DB_manager) GETUserMasterListRecords() (UserAccessMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT userid,firstname,lastname,loginname,mobile,emailid,password FROM m_users"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from user access master", err)
		return UserAccessMaster{}, err
	}
	var UserMasterArray []Usermaster
	var UserMasterRec Usermaster
	for rows.Next() {
		err := rows.Scan(&UserMasterRec.Userid, &UserMasterRec.Firstname, &UserMasterRec.Lastname, &UserMasterRec.Loginname, &UserMasterRec.Mobile_No, &UserMasterRec.Emailid, &UserMasterRec.Password)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		UserMasterArray = append(UserMasterArray, UserMasterRec)
	}

	var user_access_info UserAccessMaster
	user_access_info.Status = "true"
	user_access_info.Status_code = "200"
	user_access_info.Count = len(UserMasterArray)
	user_access_info.UserMasterInfo = UserMasterArray
	return user_access_info, nil
}

// GET User Access Master Display/Edit API
func (db *DB_manager) GETUserMasterAccessRecords(UserID int) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT m_users.userid,m_users.firstname,m_users.lastname,m_users.loginname,coalesce(m_users.mobile,'') AS mobile,m_users.emailid,m_users.password FROM m_users Where m_users.userid = %d", UserID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from production form records", err)
		return nil, err
	}
	var UserAccessMasterRec Useraccessmaster
	for rows.Next() {
		err := rows.Scan(&UserAccessMasterRec.Userid, &UserAccessMasterRec.Firstname, &UserAccessMasterRec.Lastname, &UserAccessMasterRec.Loginname, &UserAccessMasterRec.Mobile_No, &UserAccessMasterRec.Emailid, &UserAccessMasterRec.Password)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		////////////////////
		var query string

		query = fmt.Sprintf("SELECT mc_userroles.id,mc_userroles.userid,mc_userroles.roleid,m_roles.name,coalesce(mc_userroles.row_order,0) AS row_order FROM mc_userroles JOIN m_roles ON mc_userroles.roleid = m_roles.id Where mc_userroles .userid = %d ORDER BY mc_userroles.row_order ASC", UserID)

		rows1, err := db.Query(query)
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from mc_userroles", err)
			return nil, err
		}
		var MC_UserRolesArray []MC_UserRoles
		var MC_UserRolesRec MC_UserRoles
		for rows1.Next() {
			err = rows1.Scan(&MC_UserRolesRec.MC_UserRoleID, &MC_UserRolesRec.UserID, &MC_UserRolesRec.RoleID, &MC_UserRolesRec.RoleName, &MC_UserRolesRec.Row_Order)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

			MC_UserRolesArray = append(MC_UserRolesArray, MC_UserRolesRec)
		}
		var query1 string

		query1 = fmt.Sprintf("SELECT mc_userdivisions.id,mc_userdivisions.userid,mc_userdivisions.departmentid,department.department,coalesce(mc_userdivisions.row_order,0) AS row_order FROM mc_userdivisions JOIN department ON mc_userdivisions.departmentid = department.departmentid Where mc_userdivisions.userid = %d ORDER BY mc_userdivisions.row_order ASC", UserID)

		rows2, err := db.Query(query1)
		defer rows2.Close()
		if err != nil {
			fmt.Println("failed to get data from mc_userdivisions", err)
			return nil, err
		}
		var MC_UserDivisionsArray []MC_UserDivisions
		var MC_UserDivisionsRec MC_UserDivisions
		for rows2.Next() {
			err = rows2.Scan(&MC_UserDivisionsRec.MC_UserDivisionID, &MC_UserDivisionsRec.UserID, &MC_UserDivisionsRec.DepartmentID, &MC_UserDivisionsRec.DepartmentName, &MC_UserDivisionsRec.Row_Order)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			MC_UserDivisionsArray = append(MC_UserDivisionsArray, MC_UserDivisionsRec)
		}
		////////////////////////////////////////////////////////////////////////////////////////////
		UserAccessMasterRec.MC_UserRolesInfo = MC_UserRolesArray
		UserAccessMasterRec.MC_UserDivisionsInfo = MC_UserDivisionsArray

	}

	return UserAccessMasterRec, nil
}

// GET M_UserAccessMaster Display/Edit API
func (db *DB_manager) GETM_UserAccessMasterRecords(UserID int) ([]User_Access_Master_Info, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT m_useraccess.id,m_useraccess.userid,m_users.loginname,m_useraccess.companyid,m_useraccess.divisionid,m_useraccess.moduleid,h_modules.name,m_useraccess.submoduleid,CASE WHEN m_useraccess.submoduleid = 0 then 'No Sub Module' ELSE h_submodules.name END,m_useraccess.pageid,h_pages.name,m_useraccess.isshow,m_useraccess.isshowaddpage,m_useraccess.isadd,m_useraccess.iseditself,m_useraccess.isedit,m_useraccess.isview,m_useraccess.isdeleteself,m_useraccess.isdelete,m_useraccess.isprint FROM m_useraccess LEFT JOIN m_users ON m_useraccess.userid = m_users.userid LEFT JOIN h_modules ON m_useraccess.moduleid = h_modules.id LEFT JOIN h_submodules ON m_useraccess.submoduleid = h_submodules.id LEFT JOIN h_pages ON m_useraccess.pageid = h_pages.id Where m_useraccess.userid = %d ORDER BY m_useraccess.moduleid,m_useraccess.submoduleid,m_useraccess.pageid", UserID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from UserAccessMaster ", err)
		return nil, err
	}
	var useraccessMastereditArray []User_Access_Master_Info
	var useraccessMastereditRec User_Access_Master_Info
	for rows.Next() {
		err := rows.Scan(&useraccessMastereditRec.UserAccessMasterID, &useraccessMastereditRec.UserID, &useraccessMastereditRec.UserName, &useraccessMastereditRec.CompanyID, &useraccessMastereditRec.DivisionID, &useraccessMastereditRec.ModuleID, &useraccessMastereditRec.ModuleName, &useraccessMastereditRec.SubModuleID, &useraccessMastereditRec.SubModuleName, &useraccessMastereditRec.PageID, &useraccessMastereditRec.PageName, &useraccessMastereditRec.Isshow, &useraccessMastereditRec.Isshowaddpage, &useraccessMastereditRec.Isadd, &useraccessMastereditRec.Iseditself, &useraccessMastereditRec.Isedit, &useraccessMastereditRec.Isview, &useraccessMastereditRec.Isdeleteself, &useraccessMastereditRec.Isdelete, &useraccessMastereditRec.Isprint)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		useraccessMastereditArray = append(useraccessMastereditArray, useraccessMastereditRec)
	}
	return useraccessMastereditArray, nil
}

func (db *DB_manager) CheckDuplicatesUserAccessRoleAccessValidation(ModuleID int, SubModuleID int, PageID int, ValidateParameterID string, ID int) (CheckDuplicateRecValidation, error) {
	if ValidateParameterID == "1" {
		rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT count(*) FROM m_useraccess Where moduleid = %d AND submoduleid = %d AND pageid = %d AND userid = %d", ModuleID, SubModuleID, PageID, ID))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data check duplicates records from useraccess table", err)
			return CheckDuplicateRecValidation{}, err
		}

		var UserAccessStatusRec CheckDuplicateRecValidation
		for rows.Next() {
			err := rows.Scan(&UserAccessStatusRec.Count)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

			if UserAccessStatusRec.Count > 0 {
				UserAccessStatusRec.Status = "true"
			} else {
				UserAccessStatusRec.Status = "False"
			}
		}
		return UserAccessStatusRec, nil
	}
	if ValidateParameterID == "2" {
		rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT count(*) FROM m_roleaccess Where moduleid = %d AND submoduleid = %d AND pageid = %d AND roleid = %d", ModuleID, SubModuleID, PageID, ID))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data check duplicates records from roleaccess table", err)
			return CheckDuplicateRecValidation{}, err
		}

		var RoleAccessStatusRec CheckDuplicateRecValidation
		for rows.Next() {
			err := rows.Scan(&RoleAccessStatusRec.Count)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

			if RoleAccessStatusRec.Count > 0 {
				RoleAccessStatusRec.Status = "true"
			} else {
				RoleAccessStatusRec.Status = "False"
			}
		}
		return RoleAccessStatusRec, nil
	}
	return CheckDuplicateRecValidation{}, nil
}
