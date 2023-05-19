package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

type Shift struct {
	start_time string
	end_time   string
	index      int
}

func (db *DB_manager) ReadAllShifts(partNumberRegex string) (map[int]string, error) {
	shifts := make(map[int]string)
	rows, err := db.Query("SELECT * FROM shift ORDER BY index ASC")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part ", err)
		return shifts, err
	}
	index := 0
	var shift Shift
	for rows.Next() { // hopefully we are lucky enough to find only 1 norm record.
		err := rows.Scan(&shift.start_time, &shift.end_time, &shift.index)
		if err != nil {
			fmt.Println("shift: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		timeslot := shift.start_time + " to " + shift.end_time
		shifts[index] = timeslot
		index++
	}

	return shifts, nil
}
