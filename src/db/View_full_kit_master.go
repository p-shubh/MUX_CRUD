package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type Full_kit_root struct {
	Status                   string
	Status_code              string
	Count                    int
	Fullkitmasterdisplayinfo []Full_Kit_Master_View
	ProductionLineInfo       []ProductionLineDetails
}

type Full_Kit_Master_View struct {
	Fg_assembly_code        string
	Fg_assembly_description string
	Child_part_code         string
	Child_Part_Description  string
}

type ProductionLineDetails struct {
	LineName   string
	Department string
}

func (db *DB_manager) GETFullKitMasterRecordsByLine(Line string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT fg_assembly_code,fg_assembly_description,child_part_code,part.description FROM full_kit_master LEFT JOIN part ON full_kit_master.child_part_code = part.part_number Where full_kit_master.line = '%s'", Line))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Full Kit Master", err)
		return nil, err
	}
	var Full_kitArray []Full_Kit_Master_View
	var Full_kitRec Full_Kit_Master_View
	for rows.Next() {
		err := rows.Scan(&Full_kitRec.Fg_assembly_code, &Full_kitRec.Fg_assembly_description, &Full_kitRec.Child_part_code, &Full_kitRec.Child_Part_Description)
		if err != nil {
			fmt.Println(err)
		}
		Full_kitArray = append(Full_kitArray, Full_kitRec)
	}
	var fullkitdisplayinfo Full_kit_root
	fullkitdisplayinfo.Status = "true"
	fullkitdisplayinfo.Status_code = "200"
	fullkitdisplayinfo.Count = len(Full_kitArray)
	fullkitdisplayinfo.Fullkitmasterdisplayinfo = Full_kitArray
	return fullkitdisplayinfo, nil
}

func (db *DB_manager) GETProductionLineRecords(DepartmentName string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT linename,department FROM Line Where department = '%s' ORDER BY linename", DepartmentName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Line ", err)
		return nil, err
	}

	var LineInfo []ProductionLineDetails
	var LineRec ProductionLineDetails

	for rows.Next() {
		err := rows.Scan(&LineRec.LineName, &LineRec.Department)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		LineInfo = append(LineInfo, LineRec)

	}

	var lineinfo Full_kit_root
	lineinfo.Status = "true"
	lineinfo.Status_code = "200"
	lineinfo.Count = len(LineInfo)
	lineinfo.ProductionLineInfo = LineInfo
	return lineinfo, nil
}
