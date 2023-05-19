package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

type H_ModulesMaster struct {
	ModuleID         int
	ModuleName       string
	DisplayIndex     int
	DefaultPageIndex int
	ModuleIcon       string
	CompanyId        int
	Comapany_Name    string
	Created_by       int
	Created_at       time.Time
	Updated_by       int
	Updated_at       time.Time
	Is_deleted       int
}

type H_SubmodulesMaster struct {
	SubModuleNameID int
	SubModuleName   string
	Moduleid        int
	ModuleName      string
	Displayindex    int
	Defaultpageid   int
}

// GET H_Modules List API
func (db *DB_manager) GETH_ModulesMasterListRecords() ([]H_ModulesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT h_modules.id,h_modules.module_name,h_modules.display_index,h_modules.default_page_index,h_modules.module_icon, COALESCE(h_modules.created_by, 0), COALESCE(h_modules.created_at, '0001-01-01 00:00:00'), COALESCE(h_modules.updated_by, 0), COALESCE(h_modules.updated_at, '0001-01-01 00:00:00'),  COALESCE(h_modules.is_deleted, 0),c_companies.name,COALESCE(h_modules.company_id, 0) AS company_id FROM h_modules JOIN c_companies on h_modules.company_id = c_companies.id ORDER BY id ASC"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_ModulesMaster ", err)
		return nil, err
	}
	var H_modulelistArray []H_ModulesMaster
	var H_modulelistRec H_ModulesMaster
	for rows.Next() {
		err := rows.Scan(&H_modulelistRec.ModuleID, &H_modulelistRec.ModuleName, &H_modulelistRec.DisplayIndex, &H_modulelistRec.DefaultPageIndex, &H_modulelistRec.ModuleIcon, &H_modulelistRec.Created_by, &H_modulelistRec.Created_at, &H_modulelistRec.Updated_by, &H_modulelistRec.Updated_at, &H_modulelistRec.Is_deleted, &H_modulelistRec.Comapany_Name, &H_modulelistRec.CompanyId)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_modulelistArray = append(H_modulelistArray, H_modulelistRec)
	}

	return H_modulelistArray, nil
}

// GET H_Modules Display/Edit API
func (db *DB_manager) GETH_ModulesMasterRecords(h_modulesID int) ([]H_ModulesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT id,name,displayindex,defaultpageindex,moduleicon FROM h_modules Where id = %d", h_modulesID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_ModulesMaster ", err)
		return nil, err
	}
	var H_moduleeditArray []H_ModulesMaster
	var H_moduleeditRec H_ModulesMaster
	for rows.Next() {
		err := rows.Scan(&H_moduleeditRec.ModuleID, &H_moduleeditRec.ModuleName, &H_moduleeditRec.DisplayIndex, &H_moduleeditRec.DefaultPageIndex, &H_moduleeditRec.ModuleIcon)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_moduleeditArray = append(H_moduleeditArray, H_moduleeditRec)
	}

	return H_moduleeditArray, nil
}

// GET H_SubModules List API
func (db *DB_manager) GETH_SubModulesMasterListRecords() ([]H_SubmodulesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT h_submodules.id,h_submodules.name,h_submodules.moduleid,h_modules.name AS modulename,h_submodules.displayindex,h_submodules.defaultpageid FROM h_submodules JOIN h_modules ON h_submodules.moduleid = h_modules.id ORDER BY h_submodules.id ASC"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_SubModulesMaster ", err)
		return nil, err
	}
	var H_SubmodulelistArray []H_SubmodulesMaster
	var H_SubmodulelistRec H_SubmodulesMaster
	for rows.Next() {
		err := rows.Scan(&H_SubmodulelistRec.SubModuleNameID, &H_SubmodulelistRec.SubModuleName, &H_SubmodulelistRec.Moduleid, &H_SubmodulelistRec.ModuleName, &H_SubmodulelistRec.Displayindex, &H_SubmodulelistRec.Defaultpageid)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_SubmodulelistArray = append(H_SubmodulelistArray, H_SubmodulelistRec)
	}
	return H_SubmodulelistArray, nil
}

// GET H_SubModules Display/Edit API
func (db *DB_manager) GETH_SubModulesMasterRecords(h_submodulesID int) ([]H_SubmodulesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT h_submodules.id,h_submodules.name,h_submodules.moduleid,h_modules.name AS modulename,h_submodules.displayindex,h_submodules.defaultpageid FROM h_submodules JOIN h_modules ON h_submodules.moduleid = h_modules.id Where h_submodules.id = %d", h_submodulesID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_SubModulesMaster ", err)
		return nil, err
	}
	var H_submoduleeditArray []H_SubmodulesMaster
	var H_submoduleeditRec H_SubmodulesMaster
	for rows.Next() {
		err := rows.Scan(&H_submoduleeditRec.SubModuleNameID, &H_submoduleeditRec.SubModuleName, &H_submoduleeditRec.Moduleid, &H_submoduleeditRec.ModuleName, &H_submoduleeditRec.Displayindex, &H_submoduleeditRec.Defaultpageid)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_submoduleeditArray = append(H_submoduleeditArray, H_submoduleeditRec)
	}

	return H_submoduleeditArray, nil
}
