package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Trigger struct {
	Date                string
	Location            string
	Organization        string
	Item_code           string
	Description         string
	Item_location       string
	Drg_rev             string
	Priority            string
	Bp                  int
	PoNumber            string
	ReleaseNumber       string
	ReleaseMethod       string
	ReleaseDate         string
	OpenPendingQuantity int
	ASNQty              string
}

func (db *DB_manager) InsertTrigger(rec Trigger) {

	stmt, err := db.Prepare("INSERT INTO fleetguard_trigger(date, location,organization, item_code,description,item_location,drg_rev ,priority ,bp ,po_no ,release_no,release_method ,release_date,open_pending_quantity,asn_qty) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12, $13, $14, $15)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.Date,
		rec.Location,
		rec.Organization,
		rec.Item_code,
		rec.Description,
		rec.Item_location,
		rec.Drg_rev,
		rec.Priority,
		rec.Bp,
		rec.PoNumber,
		rec.ReleaseNumber,
		rec.ReleaseMethod,
		rec.ReleaseDate,
		rec.OpenPendingQuantity,
		rec.ASNQty)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect, "record added")

}
