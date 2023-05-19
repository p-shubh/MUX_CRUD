package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Line struct {
	LineName string
}

type Full_kit_Part struct {
	PartID     string
	PartNumber string
}

type Full_Kit_Master_Info struct {
	DeparmtentName      string
	Line                string
	AssemblyDetailsInfo []AssemblyDetails
}

type AssemblyDetails struct {
	FG_Assembly_Code        string
	FG_Assembly_Description string
	Child_Part_Code         string
	Created_By              int
	Created_At              string
	Updated_By              int
	Updated_At              string
}

type ProdLine_Info struct {
	Department       string
	NewProdLine_Info []NewProdLine
}

type NewProdLine struct {
	Line string
}

func (db *DB_manager) GETFullKitMasterLineDropdownRecords(Department string) ([]string, error) {
	line := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT linename FROM line where department = '" + Department + "'"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from full_kit_master line wise records ", err)
		return nil, err
	}

	var LineRec Line
	for rows.Next() {
		err := rows.Scan(&LineRec.LineName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		line = append(line, LineRec.LineName)
	}
	return line, nil
}

func (db *DB_manager) GETFullKitAssembly_CodeBySelectedDepartmentANDSubcategory(Department string, Sub_category string) ([]Full_kit_Part, error) {
	if Sub_category == "Finished Goods" {

		rows, err := db.Query(fmt.Sprintf("SELECT distinct part_number AS ID,part_number AS Name FROM part_details Where department = '%s' AND sub_category  = '%s'", Department, Sub_category))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from part_details", err)
			return nil, err
		}
		var FG_AssemblyArr []Full_kit_Part
		var FG_AssemblyRec Full_kit_Part
		for rows.Next() {
			err := rows.Scan(&FG_AssemblyRec.PartID, &FG_AssemblyRec.PartNumber)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			FG_AssemblyArr = append(FG_AssemblyArr, FG_AssemblyRec)
		}

		return FG_AssemblyArr, nil
	}
	if Sub_category == "Child Parts" {

		rows, err := db.Query(fmt.Sprintf("SELECT distinct part_number AS ID,part_number AS Name FROM part_details Where department = '%s' AND sub_category  = '%s'", Department, Sub_category))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from part_details", err)
			return nil, err
		}
		var ChildPartArr []Full_kit_Part
		var ChildPartRec Full_kit_Part
		for rows.Next() {
			err := rows.Scan(&ChildPartRec.PartID, &ChildPartRec.PartNumber)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			ChildPartArr = append(ChildPartArr, ChildPartRec)
		}

		return ChildPartArr, nil
	}
	return nil, nil
}

func (db *DB_manager) InsertFull_Kit_MasterRecords(rec Full_Kit_Master_Info, DeparmtentName string, Line string) error {
	var i int
	for i = 0; i < len(rec.AssemblyDetailsInfo); i++ {

		query := `INSERT INTO 
   full_kit_master(department,fg_assembly_code,fg_assembly_description,child_part_code,line,created_by,created_at,updated_by,updated_at) 
    VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9)`
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			DeparmtentName,
			rec.AssemblyDetailsInfo[i].FG_Assembly_Code,
			rec.AssemblyDetailsInfo[i].FG_Assembly_Description,
			rec.AssemblyDetailsInfo[i].Child_Part_Code,
			Line,
			rec.AssemblyDetailsInfo[i].Created_By,
			rec.AssemblyDetailsInfo[i].Created_At,
			rec.AssemblyDetailsInfo[i].Updated_By,
			rec.AssemblyDetailsInfo[i].Updated_At,
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

		fmt.Println(affect, "records added")
	}
	return nil
}

func (db *DB_manager) DeleteFull_Kit_MasterRecords(DeparmtentName string, Line string) ([]Full_Kit_Master_Info, error) {
	query := fmt.Sprintf(`DELETE FROM full_kit_master Where department = '%s' 
AND line = '%s'`, DeparmtentName, Line)
	stmt, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to delete full kit master record. Continue with the next operation ", err)
		fmt.Println("error ", err)
	}
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("records deleted successfully")
	return nil, err
}

func (db *DB_manager) InsertNewProductionLine(rec ProdLine_Info, Department string) {

	var i int
	for i = 0; i < len(rec.NewProdLine_Info); i++ {
		stmt, err := db.Prepare("INSERT INTO line(linename,department) VALUES($1,$2)")
		if err != nil {
			log.Fatal(err)
		}

		res, err := stmt.Exec(rec.NewProdLine_Info[i].Line,
			Department,
		)

		if err != nil {
			log.Fatal(err)
		}

		_, err = res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (db *DB_manager) DeleteProductionLineRecords(DeparmtentName string) ([]ProdLine_Info, error) {
	query := fmt.Sprintf(`DELETE FROM Line Where department = '%s'`, DeparmtentName)
	stmt, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to delete line record. Continue with the next operation ", err)
		fmt.Println("error ", err)
	}
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("records deleted successfully")
	return nil, err
}
