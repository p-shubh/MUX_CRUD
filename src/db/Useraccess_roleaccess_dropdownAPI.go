package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type M_Roles struct {
	ID   int
	Name string
}

type DepartmentWiseAccess struct {
	DepartmentID   int
	DepartmentName string
}

type H_module struct {
	ID   int
	Name string
}

type H_pages struct {
	ID   int
	Name string
}

type M_Users struct {
	UserID   int
	UserName string
}

type H_submodule struct {
	ID   int
	Name string
}

func (db *DB_manager) ReadM_RolesRecords() ([]M_Roles, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT ID,Name From m_roles"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from M_Roles", err)
		return nil, err
	}
	var M_RolesArray []M_Roles
	var M_RolesRec M_Roles
	for rows.Next() {
		err := rows.Scan(&M_RolesRec.ID, &M_RolesRec.Name)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		M_RolesArray = append(M_RolesArray, M_RolesRec)
	}

	return M_RolesArray, nil
}

func (db *DB_manager) ReadDepartmentWiseAccessRecords() ([]DepartmentWiseAccess, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT departmentid,department From department ORDER BY departmentid"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from department", err)
		return nil, err
	}
	var departmentArray []DepartmentWiseAccess
	var deparmtmentRec DepartmentWiseAccess
	for rows.Next() {
		err := rows.Scan(&deparmtmentRec.DepartmentID, &deparmtmentRec.DepartmentName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		departmentArray = append(departmentArray, deparmtmentRec)
	}

	return departmentArray, nil
}

func (db *DB_manager) ReadH_ModulesRecords() ([]H_module, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT ID,Name From H_modules"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_Modules ", err)
		return nil, err
	}
	var H_moduleArray []H_module
	var H_moduleRec H_module
	for rows.Next() {
		err := rows.Scan(&H_moduleRec.ID, &H_moduleRec.Name)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_moduleArray = append(H_moduleArray, H_moduleRec)
	}

	return H_moduleArray, nil
}

func (db *DB_manager) ReadH_PagesRecords() ([]H_pages, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT ID,Name From H_pages"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_Pages", err)
		return nil, err
	}
	var H_pagesArray []H_pages
	var H_pagesRec H_pages
	for rows.Next() {
		err := rows.Scan(&H_pagesRec.ID, &H_pagesRec.Name)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_pagesArray = append(H_pagesArray, H_pagesRec)
	}

	return H_pagesArray, nil
}

// To Populate dropdown API form h_pages by using ModuleID & SubModuleID
func (db *DB_manager) GETH_PagesRecordsByModuleIDandSubModuleID(ModuleID int, SubModuleID int) ([]H_pages, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT ID,Name From H_pages Where defaultmoduleid= %d AND defaultsubmoduleid = %d", ModuleID, SubModuleID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_Pages", err)
		return nil, err
	}
	var H_pagesArr []H_pages
	var H_pagesinfo H_pages
	for rows.Next() {
		err := rows.Scan(&H_pagesinfo.ID, &H_pagesinfo.Name)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_pagesArr = append(H_pagesArr, H_pagesinfo)
	}

	return H_pagesArr, nil
}

// Dropdown to populate API from user access master page
func (db *DB_manager) ReadM_UsersRecords() ([]M_Users, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT userid,loginname as username FROM m_users ORDER BY userid"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from M_users", err)
		return nil, err
	}
	var M_usersArray []M_Users
	var M_usersRec M_Users
	for rows.Next() {
		err := rows.Scan(&M_usersRec.UserID, &M_usersRec.UserName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		M_usersArray = append(M_usersArray, M_usersRec)
	}

	return M_usersArray, nil
}

func (db *DB_manager) ReadH_SubModulesRecords(ModuleID int) ([]H_submodule, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT ID,Name From H_submodules Where ModuleID = %d UNION SELECT 0, 'No Sub Module'", ModuleID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_SubModules ", err)
		return nil, err
	}
	var H_submoduleArray []H_submodule
	var H_submoduleRec H_submodule
	for rows.Next() {
		err := rows.Scan(&H_submoduleRec.ID, &H_submoduleRec.Name)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_submoduleArray = append(H_submoduleArray, H_submoduleRec)
	}

	return H_submoduleArray, nil
}
