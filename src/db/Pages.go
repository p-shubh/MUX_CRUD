package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type H_Pages struct {
	H_PagesID          int
	PageName           string
	Nameonmenu         string
	Description        string
	Displayindex       int
	Defaultmoduleid    int
	Defaultsubmoduleid int
	Actualpagepath     string
	Pagetype           string
	Pageid             int
}

type H_SubPages struct {
	H_SubPagesid int
	Pageid       int
	PageName     string
}

// Insert H_Pages data
func (db *DB_manager) InsertH_PagesRecords(rec H_Pages) error {

	query := `INSERT INTO
     h_pages(name,nameonmenu,description,displayindex,defaultmoduleid,defaultsubmoduleid,actualpagepath,pagetype,pageid)
    VALUES($1, $2, $3, $4,$5,$6,$7,$8,$9)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.PageName,
		rec.Nameonmenu,
		rec.Description,
		rec.Displayindex,
		rec.Defaultmoduleid,
		rec.Defaultsubmoduleid,
		rec.Actualpagepath,
		rec.Pagetype,
		rec.Pageid,
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

// Update H_Pages data
func (db *DB_manager) UpdateH_PagesRecords(rec H_Pages) error {

	query := fmt.Sprintf(`UPDATE h_pages SET name = '%s',nameonmenu = '%s',description = '%s',displayindex= %d,defaultmoduleid	= %d,defaultsubmoduleid = %d,actualpagepath	= '%s',pagetype = '%s'	,pageid = %d Where id = %d
`, rec.PageName, rec.Nameonmenu, rec.Description, rec.Displayindex, rec.Defaultmoduleid, rec.Defaultsubmoduleid, rec.Actualpagepath, rec.Pagetype, rec.Pageid, rec.H_PagesID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update h_pages record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")
	return err
}

// Insert H_SubPages data
func (db *DB_manager) InsertH_SubPagesRecords(rec H_SubPages) error {

	query := `INSERT INTO
     h_subpages(pageid,name)
    VALUES($1, $2)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.Pageid,
		rec.PageName,
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

// Update H_Subpages data
func (db *DB_manager) UpdateH_SubPagesRecords(rec H_SubPages) error {

	query := fmt.Sprintf(`UPDATE h_subpages SET pageid = %d,name = '%s' Where id = %d
`, rec.Pageid, rec.PageName, rec.H_SubPagesid)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update h_subpages record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")
	return err
}
