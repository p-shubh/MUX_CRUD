package db

import (
	"database/sql"
	"fmt"
	"log"
)

func (db *DB_manager) PostModulePageToDatabase(rec H_Modules) (bool, error) {

	count := 0

	//Count the module name
	row := db.QueryRow(`SELECT COUNT(id) From h_modules WHERE module_name = '` + rec.ModuleName + `' ;`)
	errr := row.Scan(&count)
	if errr != nil {
		log.Println("Failed to Count module_name in PostModulePage ")
		return false, fmt.Errorf("query error in Count the module name %q", errr.Error())
	} else if count >= 1 {
		log.Println("Duplicate module name :", count)
		return false, fmt.Errorf("Sorry please try another module name its already added")
	}

	// Prepare the insert statement
	stmt, err := db.Prepare("INSERT INTO h_modules (module_name, display_index, default_page_index, module_icon, company_id, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		log.Println("Failed to prepare the query in PostModulePage", err.Error())
		return false, fmt.Errorf("Failed to prepare the Query")
	}

	defer stmt.Close()

	// Execute the insert statement
	_, err = stmt.Exec(rec.ModuleName, rec.DisplayIndex, rec.DefaultPageIndex, rec.ModuleIcon, rec.CompanyID, rec.CreatedBy, rec.CreatedAt)
	if err != nil {
		log.Println("Failed to execute the query in PostModulePage", err.Error())
		return false, err
	}
	return true, nil
}

func (db *DB_manager) UpdateH_ModulesRecordsToDatabase(rec H_Modules) (bool, error) {

	count := 0

	//Count the module name
	row := db.QueryRow(`SELECT COUNT(id) From h_modules WHERE module_name = '` + rec.ModuleName + `' ;`)
	errr := row.Scan(&count)
	if errr != nil {
		log.Println("Failed to Count module_name in PostModulePage ")
		return false, fmt.Errorf("query error in Count the module name %q", errr.Error())
	} else if count >= 1 {
		log.Println("Duplicate module name :", count)
		return false, fmt.Errorf("Sorry please try another module name its already added")
	}

	query := fmt.Sprintf(`UPDATE H_Modules SET module_name = '%s',display_index = %d,default_page_index = %d,module_icon = '%s',updated_by = %d, updated_at = NOW() Where id = %d`, rec.ModuleName, rec.DisplayIndex, rec.DefaultPageIndex, rec.ModuleIcon, rec.UpdatedBy, rec.ModuleID)

	stmt, err := db.Query(query)

	if err != nil {
		log.Println("failed to update h_modules record. Continue with the next operation ", rec, "query : ", query)
		log.Println("error ", err)
		return false, fmt.Errorf("query error in Count the module name %q", errr.Error())
	}
	defer func(stmt *sql.Rows) {
		err := stmt.Close()
		if err != nil {
			log.Println("error ", err)
		}
	}(stmt)

	log.Println("1 record Updated")
	return true, err
}
