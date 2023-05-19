package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type H_Modules struct {
	ModuleID         int
	ModuleName       string
	DisplayIndex     int
	DefaultPageIndex int
	ModuleIcon       string
	CompanyID        int
	CreatedBy        int
	CreatedAt        time.Time
	UpdatedBy        string
	UpdatedAt        time.Time
	IsDeleted        int
}

type H_Submodules struct {
	H_SubModuleID int
	SubModuleName string
	Moduleid      int
	Displayindex  int
	Defaultpageid int
}

// Insert H_Modules data
func (db *DB_manager) InsertH_ModulesRecords(rec H_Modules) error {

	query := `INSERT INTO
     h_modules(name,displayindex,defaultpageindex,moduleicon)
    VALUES($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.ModuleName,
		rec.DisplayIndex,
		rec.DefaultPageIndex,
		rec.ModuleIcon,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(affect, "record added")
	return err
}

// Update H_Modules data
func (db *DB_manager) UpdateH_ModulesRecords(rec H_Modules) error {

	query := fmt.Sprintf(`UPDATE H_Modules SET name = '%s',displayindex = %d,defaultpageindex = %d,moduleicon = '%s' Where id = %d
`, rec.ModuleName, rec.DisplayIndex, rec.DefaultPageIndex, rec.ModuleIcon, rec.ModuleID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update h_modules record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")
	return err
}

// Insert H_SubModules data
func (db *DB_manager) InsertH_SubModulesRecords(rec H_Submodules) error {

	query := `INSERT INTO
     h_submodules(name,moduleid,displayindex,defaultpageid)
    VALUES($1, $2, $3, $4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.SubModuleName,
		rec.Moduleid,
		rec.Displayindex,
		rec.Defaultpageid,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(affect, "record added")
	return err
}

// Update H_SubModules data
func (db *DB_manager) UpdateH_SubModulesRecords(rec H_Submodules) error {

	query := fmt.Sprintf(`UPDATE h_submodules SET name = '%s',moduleid = %d,displayindex = %d,defaultpageid = %d Where id = %d
`, rec.SubModuleName, rec.Moduleid, rec.Displayindex, rec.Defaultpageid, rec.H_SubModuleID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update h_submodules record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")
	return err
}
