package db

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Excel_Upload_Master_Data struct {
	Part_ID             int
	Part_Number         string
	Part_Description    string
	Department          string
	Sub_category        string
	Item_class          string
	Item_rate           float64
	AlternatePartNumber string
	Plant               string
	Number_of_days      float64
	Lead_Time           float64
	Safety_factor       float64
	Buyer               string
	Supplier            string
	Display_Name        string
	Created_By          int
	Created_At          string
	Updated_By          int
	Updated_At          string
	Excel_line_no       int
}

type ItemMasterInsertUpdateCount struct {
	PartNumbers           string
	DepartmentName        string
	Sub_categoryName      string
	TableName             string
	InsertUpdateParameter string
	Message               string
	Excel_row_number      int
}
type ItemOfMasterData struct {
	Part_number           string
	Description           string
	Department            string
	Sub_category          string
	Item_class            string
	Rate                  float64
	Alternate_part_number string
	Plant                 string
	Nod                   float64
	Lead_time             float64
	Safety_factor         float64
	Buyer                 string
	Supplier              string
}
type AllListData struct {
	Details []ItemOfMasterData
}

func (db *DB_manager) InsertItemMasterExcelUpload(rec []Excel_Upload_Master_Data) ([]ItemMasterInsertUpdateCount, error) {

	var ItemMasterInsertUpdateCountArray []ItemMasterInsertUpdateCount

	var i int
	for i = 0; i < len(rec); i++ {

		var Part_ID int
		ItemMasterRec, _ := db.CheckDuplicatesPartsValidationRecords(rec[i].Part_Number)
		Part_ID = ItemMasterRec.Part_ID
		fmt.Println(ItemMasterRec.Part_ID)
		if ItemMasterRec.Count == 0 {

			query := `INSERT INTO 
    part(part_number,description,created_by,created_at,updated_by,updated_at) 
    VALUES($1, $2, $3,$4,$5,$6)RETURNING Part_ID`

			stmt, err := db.Prepare(query)
			if err != nil {
				fmt.Println("failed to insert part records. ", err)
				ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, "", "Part", "insert", "Error in Part Insert", rec[i].Excel_line_no))
			}
			defer stmt.Close()

			err = stmt.QueryRow(
				rec[i].Part_Number,
				rec[i].Part_Description,
				rec[i].Created_By,
				rec[i].Created_At,
				rec[i].Updated_By,
				rec[i].Updated_At,
			).Scan(&Part_ID)

			if err != nil {
				fmt.Println("failed to scan part records. ", err)
				ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, "", "Part", "insert", "Error in Part Insert", rec[i].Excel_line_no))

			}

		} // if ItemMasterRec.Count == 0

		// Check department name that maintain in our back end data
		departmentCount, _ := db.CheckDepartmentName(rec[i].Department) // Validation

		// Check sub category name that maintain in our back end data
		subcategoryCount, _ := db.CheckSubCategoryName(rec[i].Department, rec[i].Sub_category) // Validation

		if departmentCount == 0 {
			ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "", "", "Department "+rec[i].Department+" is not maintain in master data.", rec[i].Excel_line_no))

		} else if subcategoryCount == 0 {
			ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "", "", "Sub_Category "+rec[i].Sub_category+" is not maintain in master data.", rec[i].Excel_line_no))
		} else {
			qry := fmt.Sprintf("SELECT count(*) FROM part_details where department='%s' AND part_number = '%s'", rec[i].Department, rec[i].Part_Number)

			rows, err := db.Query(qry)
			defer rows.Close()
			if err != nil {
				fmt.Println("failed to get data from part_details ", err)
				fmt.Println("failed_query ", qry)
				return ItemMasterInsertUpdateCountArray, nil
			}

			var CountRec int
			for rows.Next() {
				err := rows.Scan(&CountRec)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

				if CountRec > 0 {

					query := fmt.Sprintf(`UPDATE part_details SET part_number = '%s',department = '%s',sub_category = '%s',item_class = '%s',rate = %f,alternate_part_number = '%s',plant = '%s',updated_by = %d,updated_at = '%s' WHERE department='%s' AND part_number='%s'`, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, rec[i].Item_class, rec[i].Item_rate, rec[i].AlternatePartNumber, rec[i].Plant, rec[i].Updated_By, rec[i].Updated_At, rec[i].Department, rec[i].Part_Number)

					stmt, err := db.Exec(query)

					if err != nil {
						fmt.Println("failed to update part_details record. Continue with the next operation ", rec)
						fmt.Println("error ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Details", "update", "Error in part_details update", rec[i].Excel_line_no))
					}
					defer stmt.RowsAffected()

					fmt.Println("1 record Updated")
				} else {

					query := `INSERT INTO
				    part_details(part_id,part_number,department,sub_category,item_class,rate,alternate_part_number,plant,created_by,created_at,updated_by,updated_at)
				    VALUES($1, $2, $3,$4,$5, $6, $7, $8, $9, $10, $11,$12)`

					stmt, err := db.Prepare(query)

					if err != nil {
						fmt.Println("failed to insert part_details record ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Details", "insert", "Error in part_details insert", rec[i].Excel_line_no))
					}

					defer stmt.Close()

					res, err := stmt.Exec(
						Part_ID,
						rec[i].Part_Number,
						rec[i].Department,
						rec[i].Sub_category,
						rec[i].Item_class,
						rec[i].Item_rate,
						rec[i].AlternatePartNumber,
						rec[i].Plant,
						rec[i].Created_By,
						rec[i].Created_At,
						rec[i].Updated_By,
						rec[i].Updated_At,
					)
					if err != nil {
						fmt.Println("failed to scan part_details record", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Details", "insert", "Error in part_details insert", rec[i].Excel_line_no))
					}

					affect, err := res.RowsAffected()

					fmt.Println(affect, "record added")

				} // else end part_details

			} // for CountRec

			qry = fmt.Sprintf("SELECT count(*) FROM norm_cd where department='%s' AND part_number = '%s'", rec[i].Department, rec[i].Part_Number)

			rows, err = db.Query(qry)
			defer rows.Close()
			if err != nil {
				fmt.Println("failed to get data from norm_cd ", err)
				fmt.Println("failed_query ", qry)
				return ItemMasterInsertUpdateCountArray, nil
			}

			for rows.Next() {
				err := rows.Scan(&CountRec)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

				if CountRec > 0 {

					query := fmt.Sprintf(`UPDATE norm_cd SET department = '%s',part_number = '%s',nod = %f,lead_time = %f,safety_factor = %f,buyer = '%s',supplier ='%s',updated_by = %d,updated_at = '%s' Where department='%s' AND part_number='%s'`, rec[i].Department, rec[i].Part_Number, rec[i].Number_of_days, rec[i].Lead_Time, rec[i].Safety_factor, rec[i].Buyer, rec[i].Supplier, rec[i].Updated_By, rec[i].Updated_At, rec[i].Department, rec[i].Part_Number)

					stmt, err := db.Exec(query)

					if err != nil {
						fmt.Println("failed to update norm calculation details record. Continue with the next operation ", rec)
						fmt.Println("error ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Norm_Calculation_Details", "update", "Error in norm_calculation_details update", rec[i].Excel_line_no))
					}
					defer stmt.RowsAffected()

					fmt.Println("1 record Updated")
				} else {
					query := `INSERT INTO Norm_CD(part_id,department,part_number,nod,lead_time,safety_factor,buyer,supplier,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

					stmt, err := db.Prepare(query)
					if err != nil {
						fmt.Println("failed to insert norm_calculation_details record ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Norm_Calculation_Details", "insert", "Error in norm_calculation_details insert", rec[i].Excel_line_no))
					}
					defer stmt.Close()

					res, err := stmt.Exec(
						Part_ID,
						rec[i].Department,
						rec[i].Part_Number,
						rec[i].Number_of_days,
						rec[i].Lead_Time,
						rec[i].Safety_factor,
						rec[i].Buyer,
						rec[i].Supplier,
						rec[i].Created_By,
						rec[i].Created_At,
						rec[i].Updated_By,
						rec[i].Updated_At,
					)

					if err != nil {
						fmt.Println("failed to scan norm_calculation_details record ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Norm_Calculation_Details", "insert", "Error in norm_calculation_details insert", rec[i].Excel_line_no))
					}

					affect, err := res.RowsAffected()

					fmt.Println(affect, "record added")

				} // else end norm_cd

			} // for CountRec

			qry = fmt.Sprintf("SELECT count(*) FROM part_supplier_details where department='%s' AND part_number = '%s'", rec[i].Department, rec[i].Part_Number)

			rows, err = db.Query(qry)
			defer rows.Close()
			if err != nil {
				fmt.Println("failed to get data from part_supplier_details ", err)
				fmt.Println("failed_query ", qry)
				return ItemMasterInsertUpdateCountArray, nil
			}

			for rows.Next() {
				err := rows.Scan(&CountRec)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

				rec[i].Display_Name = rec[i].Supplier
				if CountRec > 0 {
					query := fmt.Sprintf(`UPDATE part_supplier_details SET part_number = '%s',department = '%s',sub_category = '%s',buyer = '%s',supplier = '%s',display_name = '%s',updated_by = %d,updated_at = '%s' Where department='%s' AND part_number='%s'`, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, rec[i].Buyer, rec[i].Supplier, rec[i].Display_Name, rec[i].Updated_By, rec[i].Updated_At, rec[i].Department, rec[i].Part_Number)

					stmt, err := db.Exec(query)

					if err != nil {
						fmt.Println("failed to update part supplier details record. Continue with the next operation ", rec)
						fmt.Println("error ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Supplier_details", "update", "Error in part supplier details update", rec[i].Excel_line_no))
					}
					defer stmt.RowsAffected()

					fmt.Println("1 record Updated")
				} else {
					query := `INSERT INTO
				    part_supplier_details(part_id,part_number,department,sub_category,buyer,supplier,display_name,created_by,created_at,updated_by,updated_at)
				    VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11)`

					stmt, err := db.Prepare(query)

					if err != nil {
						fmt.Println("failed to insert part_supplier_details record ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Supplier_details", "insert", "Error in part supplier details insert", rec[i].Excel_line_no))
					}
					defer stmt.Close()

					res, err := stmt.Exec(
						Part_ID,
						rec[i].Part_Number,
						rec[i].Department,
						rec[i].Sub_category,
						rec[i].Buyer,
						rec[i].Supplier,
						rec[i].Display_Name,
						rec[i].Created_By,
						rec[i].Created_At,
						rec[i].Updated_By,
						rec[i].Updated_At,
					)
					if err != nil {
						fmt.Println("failed to scan part_supplier_details record ", err)
						ItemMasterInsertUpdateCountArray = append(ItemMasterInsertUpdateCountArray, AdditemmasterinsertErrorDetails(ItemMasterInsertUpdateCountArray, rec[i].Part_Number, rec[i].Department, rec[i].Sub_category, "Part_Supplier_details", "insert", "Error in part supplier details insert", rec[i].Excel_line_no))
					}

					affect, err := res.RowsAffected()

					fmt.Println(affect, "record added")
				} // else end part_supplier_details

			} // for CountRec

		}

	} // for i = 0; i < len(rec); i++

	return ItemMasterInsertUpdateCountArray, nil
}

func AdditemmasterinsertErrorDetails(array []ItemMasterInsertUpdateCount, PartNumbers string, DepartmentName string, Sub_Category string, TableName string, InsertUpdateParameter string, Message string, Excel_row_number int) ItemMasterInsertUpdateCount {
	var ItemMasterInsertUpdateCountRec ItemMasterInsertUpdateCount
	ItemMasterInsertUpdateCountRec.PartNumbers = PartNumbers
	ItemMasterInsertUpdateCountRec.DepartmentName = DepartmentName
	ItemMasterInsertUpdateCountRec.Sub_categoryName = Sub_Category
	ItemMasterInsertUpdateCountRec.TableName = TableName
	ItemMasterInsertUpdateCountRec.InsertUpdateParameter = InsertUpdateParameter
	ItemMasterInsertUpdateCountRec.Message = Message
	ItemMasterInsertUpdateCountRec.Excel_row_number = Excel_row_number
	return ItemMasterInsertUpdateCountRec
	//array = append(array, ItemMasterInsertUpdateCountRec)
}

func (db *DB_manager) CheckDepartmentName(DepartmentName string) (int, error) {

	qry := fmt.Sprintf("SELECT count(*) FROM department where department='%s'", DepartmentName)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from department ", err)
		fmt.Println("failed_query ", qry)
		return -1, err
	}

	var CountRec int
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		return CountRec, nil
	}
	return -1, errors.New("failed_to_get_department_record")
}

func (db *DB_manager) CheckSubCategoryName(DepartmentName string, Sub_category string) (int, error) {

	qry := fmt.Sprintf("SELECT count(*) FROM sub_category where department='%s' AND sub_category  = '%s'", DepartmentName, Sub_category)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from sub_category ", err)
		fmt.Println("failed_query ", qry)
		return -1, err
	}

	var CountRec int
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		return CountRec, nil
	}
	return -1, errors.New("failed_to_get_subcategory_record")
}

func (db *DB_manager) GetAllDataListBaseOnDepartmentNameAndAll(isALL bool, departmentName string) (AllListData, error) {
	var ItemOfMasterDataList ItemOfMasterData
	var AllListDataDetails AllListData

	query_main := "SELECT part.part_number,part.description,coalesce(part_details.department,'') AS department, coalesce(part_details.sub_category,'') AS sub_category, coalesce(part_details.item_class,'') AS item_class, coalesce(part_details.rate,0.0) AS rate,coalesce(part_details.alternate_part_number,'') AS alternate_part_number, coalesce(part_details.plant,'') AS plant,coalesce(norm_cd.nod,0.0) AS nod, coalesce(norm_cd.lead_time,0.0) AS lead_time, coalesce(norm_cd.safety_factor,0.0) AS safety_factor , coalesce(part_supplier_details.buyer,'') AS buyer, coalesce(part_supplier_details.supplier,'') AS supplier FROM Part LEFT JOIN part_details ON part.part_number = part_details.part_number LEFT JOIN part_supplier_details ON part.part_number = part_supplier_details.part_number AND part_details.department = part_supplier_details.department LEFT JOIN norm_cd ON part.part_number = norm_cd.part_number AND part_details.department = norm_cd.department"
	if isALL != true {
		query_main = query_main + " WHERE part_details.department in (" + departmentName + ")"
	}

	query := fmt.Sprintf(query_main)
	rows, err := db.Query(query)

	for rows.Next() {
		err = rows.Scan(&ItemOfMasterDataList.Part_number, &ItemOfMasterDataList.Description, &ItemOfMasterDataList.Department, &ItemOfMasterDataList.Sub_category, &ItemOfMasterDataList.Item_class, &ItemOfMasterDataList.Rate, &ItemOfMasterDataList.Alternate_part_number, &ItemOfMasterDataList.Plant, &ItemOfMasterDataList.Nod, &ItemOfMasterDataList.Lead_time, &ItemOfMasterDataList.Safety_factor, &ItemOfMasterDataList.Buyer, &ItemOfMasterDataList.Supplier)
		AllListDataDetails.Details = append(AllListDataDetails.Details, ItemOfMasterDataList)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	if err != nil {
		return AllListDataDetails, err
	}
	return AllListDataDetails, err
}
