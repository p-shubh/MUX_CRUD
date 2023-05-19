package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type M_rolesMaster struct {
	M_RoleID    int
	RoleName    string
	Description string
	IsActive    string
	PagePath    string
}

type Role_Access_Master_Info struct {
	RoleAccessMasterID int
	RoleID             int
	RoleName           string
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

// GET M_Roles List API
func (db *DB_manager) GETM_RoleMasterListRecords() ([]M_rolesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id,name,coalesce(description,'') AS description,coalesce(isactive,'') AS isactive,coalesce(page_path,'') AS page_path FROM m_roles ORDER BY id"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from M_RoleMaster ", err)
		return nil, err
	}
	var M_RoleslistArray []M_rolesMaster
	var M_RoleslistRec M_rolesMaster
	for rows.Next() {
		err := rows.Scan(&M_RoleslistRec.M_RoleID, &M_RoleslistRec.RoleName, &M_RoleslistRec.Description, &M_RoleslistRec.IsActive, &M_RoleslistRec.PagePath)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		M_RoleslistArray = append(M_RoleslistArray, M_RoleslistRec)
	}

	return M_RoleslistArray, nil
}

// GET M_Roles Display/Edit API
func (db *DB_manager) GETM_RoleMasterRecords(M_roleID int) ([]M_rolesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id,name,coalesce(description,'') AS description,coalesce(isactive,'') AS isactive,coalesce(page_path,'') AS page_path FROM m_roles Where id = %d", M_roleID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from M_RoleMaster ", err)
		return nil, err
	}
	var M_RoleseditArray []M_rolesMaster
	var M_RoleseditRec M_rolesMaster
	for rows.Next() {
		err := rows.Scan(&M_RoleseditRec.M_RoleID, &M_RoleseditRec.RoleName, &M_RoleseditRec.Description, &M_RoleseditRec.IsActive, &M_RoleseditRec.PagePath)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		M_RoleseditArray = append(M_RoleseditArray, M_RoleseditRec)
	}

	return M_RoleseditArray, nil
}

// GET M_RoleAccessMaster Display/Edit API
func (db *DB_manager) GETM_RoleAccessMasterRecords(RoleID int) ([]Role_Access_Master_Info, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT m_roleaccess.id,m_roleaccess.roleid,m_roles.name,m_roleaccess.companyid,m_roleaccess.divisionid,m_roleaccess.moduleid,h_modules.name,m_roleaccess.submoduleid,CASE WHEN m_roleaccess.submoduleid = 0 then 'No Sub Module' ELSE h_submodules.name END,m_roleaccess.pageid,h_pages.name,m_roleaccess.isshow,m_roleaccess.isshowaddpage,m_roleaccess.isadd,m_roleaccess.iseditself,m_roleaccess.isedit,m_roleaccess.isview,m_roleaccess.isdeleteself,m_roleaccess.isdelete,m_roleaccess.isprint FROM m_roleaccess LEFT JOIN m_roles ON m_roleaccess.roleid = m_roles.id LEFT JOIN h_modules ON m_roleaccess.moduleid = h_modules.id LEFT JOIN h_submodules ON m_roleaccess.submoduleid = h_submodules.id LEFT JOIN h_pages ON m_roleaccess.pageid = h_pages.id Where m_roleaccess.roleid = %d ORDER BY m_roleaccess.moduleid,m_roleaccess.submoduleid,m_roleaccess.pageid", RoleID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from RoleAccessMaster ", err)
		return nil, err
	}
	var roleaccessMastereditArray []Role_Access_Master_Info
	var roleaccessMastereditRec Role_Access_Master_Info
	for rows.Next() {
		err := rows.Scan(&roleaccessMastereditRec.RoleAccessMasterID, &roleaccessMastereditRec.RoleID, &roleaccessMastereditRec.RoleName, &roleaccessMastereditRec.CompanyID, &roleaccessMastereditRec.DivisionID, &roleaccessMastereditRec.ModuleID, &roleaccessMastereditRec.ModuleName, &roleaccessMastereditRec.SubModuleID, &roleaccessMastereditRec.SubModuleName, &roleaccessMastereditRec.PageID, &roleaccessMastereditRec.PageName, &roleaccessMastereditRec.Isshow, &roleaccessMastereditRec.Isshowaddpage, &roleaccessMastereditRec.Isadd, &roleaccessMastereditRec.Iseditself, &roleaccessMastereditRec.Isedit, &roleaccessMastereditRec.Isview, &roleaccessMastereditRec.Isdeleteself, &roleaccessMastereditRec.Isdelete, &roleaccessMastereditRec.Isprint)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		roleaccessMastereditArray = append(roleaccessMastereditArray, roleaccessMastereditRec)
	}

	return roleaccessMastereditArray, nil
}
