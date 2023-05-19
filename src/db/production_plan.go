package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ProductionReport struct {
	Operator       string
	Date           string
	Shift          string
	TimeSlot       string
	PartNumber     string
	SubPartNumber  string
	Operation      string
	Machine        string
	Planned        int
	Produced       int
	Rejected       int
	Downtime       int
	DowntimeReason string
	Remarks        string
}

func (db *DB_manager) InsertPeriodicProductionRecord(rec ProductionReport) {
	query := `INSERT INTO 
    production_data(operator, date, shift, timeslot, partnumber, subpartnumber, operation, machine, planned, produced, rejected, downtime, downtimereason, remarks) 
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.Operator,
		rec.Date,
		rec.Shift,
		rec.TimeSlot,
		rec.PartNumber,
		rec.SubPartNumber,
		rec.Operation,
		rec.Machine,
		rec.Planned,
		rec.Produced,
		rec.Rejected,
		rec.Downtime,
		rec.DowntimeReason,
		rec.Remarks,
	)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect, "record added")
}
