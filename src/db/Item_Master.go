package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Item_Master struct {
	Part_ID           int
	Part_Number       string
	Part_Description  string
	Created_By        int
	Created_At        string
	Updated_By        int
	Updated_At        string
	Count             int
	Status            string
	Part_DetailsItems []Part_Details
}

type Part_Details struct {
	PartDetails_ID      int
	Part_ID             int
	Part_Number         string
	Department          string
	Sub_category        string
	Item_class          string
	Item_rate           float64
	AlternatePartNumber string
	Plant               string
	Buyer               string
	Supplier            string
	Display_Name        string
	Date                string
	Lead_Time           float64
	Safety_factor       float64
	Row_Order           int
	Created_By          int
	Created_At          string
	Updated_By          int
	Updated_At          string
}

type DropdownBuyer struct {
	BuyerID   string
	BuyerName string
}

type DropdownSupplier struct {
	SupplierID   string
	SupplierName string
}

type DropdownPlant struct {
	PlantID   string
	PlantName string
}

func (db *DB_manager) InsertItemMasterRecords(rec Item_Master) (int, error) {
	var Part_ID int

	query := `INSERT INTO 
    part(part_number,description,created_by,created_at,updated_by,updated_at) 
    VALUES($1, $2, $3,$4,$5,$6)RETURNING Part_ID`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return Part_ID, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		rec.Part_Number,
		rec.Part_Description,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	).Scan(&Part_ID)

	if err != nil {
		log.Fatal(err)
		return Part_ID, err
	}

	var i int
	for i = 0; i < len(rec.Part_DetailsItems); i++ {

		query = `INSERT INTO 
    part_details(part_id,part_number,department,sub_category,item_class,rate,alternate_part_number,plant,created_by,created_at,updated_by,updated_at,row_order) 
    VALUES($1, $2, $3,$4,$5, $6, $7, $8, $9, $10, $11,$12,$13)`

		stmt, err = db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			Part_ID,
			rec.Part_DetailsItems[i].Part_Number,
			rec.Part_DetailsItems[i].Department,
			rec.Part_DetailsItems[i].Sub_category,
			rec.Part_DetailsItems[i].Item_class,
			rec.Part_DetailsItems[i].Item_rate,
			rec.Part_DetailsItems[i].AlternatePartNumber,
			rec.Part_DetailsItems[i].Plant,
			rec.Part_DetailsItems[i].Created_By,
			rec.Part_DetailsItems[i].Created_At,
			rec.Part_DetailsItems[i].Updated_By,
			rec.Part_DetailsItems[i].Updated_At,
			rec.Part_DetailsItems[i].Row_Order,
		)
		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(affect, "record added")
	}

	var j int
	for j = 0; j < len(rec.Part_DetailsItems); j++ {
		rec.Part_DetailsItems[j].Display_Name = rec.Part_DetailsItems[j].Supplier
		query = `INSERT INTO 
    part_supplier_details(part_id,part_number,department,sub_category,buyer,supplier,display_name,created_by,created_at,updated_by,updated_at,row_order) 
    VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			Part_ID,
			rec.Part_DetailsItems[j].Part_Number,
			rec.Part_DetailsItems[j].Department,
			rec.Part_DetailsItems[j].Sub_category,
			rec.Part_DetailsItems[j].Buyer,
			rec.Part_DetailsItems[j].Supplier,
			rec.Part_DetailsItems[j].Display_Name,
			rec.Part_DetailsItems[j].Created_By,
			rec.Part_DetailsItems[j].Created_At,
			rec.Part_DetailsItems[j].Updated_By,
			rec.Part_DetailsItems[j].Updated_At,
			rec.Part_DetailsItems[j].Row_Order,
		)
		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(affect, "record added")
	}

	var k int
	for k = 0; k < len(rec.Part_DetailsItems); k++ {

		stmt, err = db.Prepare("INSERT INTO Norm_CD(part_id,date,department,part_number,lead_time,safety_factor,buyer,supplier,created_by,created_at,updated_by,updated_at,row_order) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)")
		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		res, err := stmt.Exec(
			Part_ID,
			rec.Part_DetailsItems[k].Date,
			rec.Part_DetailsItems[k].Department,
			rec.Part_DetailsItems[k].Part_Number,
			rec.Part_DetailsItems[k].Lead_Time,
			rec.Part_DetailsItems[k].Safety_factor,
			rec.Part_DetailsItems[k].Buyer,
			rec.Part_DetailsItems[k].Supplier,
			rec.Part_DetailsItems[k].Created_By,
			rec.Part_DetailsItems[k].Created_At,
			rec.Part_DetailsItems[k].Updated_By,
			rec.Part_DetailsItems[k].Updated_At,
			rec.Part_DetailsItems[k].Row_Order,
		)

		if err != nil {
			log.Fatal(err)
			return Part_ID, err
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(affect, "record added")
	}
	return Part_ID, err
}

func (db *DB_manager) UpdateItemMasterRecords(rec Item_Master) error {

	query := fmt.Sprintf(`UPDATE Part SET part_number = '%s',description = '%s',updated_by = %d,	
updated_at = '%s' Where Part_ID = %d`, rec.Part_Number, rec.Part_Description, rec.Updated_By, rec.Updated_At, rec.Part_ID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update part record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")

	var i int
	for i = 0; i < len(rec.Part_DetailsItems); i++ {

		if rec.Part_DetailsItems[i].PartDetails_ID > 0 {

			query = fmt.Sprintf(`UPDATE part_details SET part_number = '%s',department = '%s',sub_category = '%s',item_class = '%s',rate = %f,alternate_part_number = '%s',plant = '%s',updated_by = %d,updated_at = '%s',row_order = %d Where part_details_id = %d`, rec.Part_DetailsItems[i].Part_Number, rec.Part_DetailsItems[i].Department, rec.Part_DetailsItems[i].Sub_category, rec.Part_DetailsItems[i].Item_class, rec.Part_DetailsItems[i].Item_rate, rec.Part_DetailsItems[i].AlternatePartNumber, rec.Part_DetailsItems[i].Plant, rec.Part_DetailsItems[i].Updated_By, rec.Part_DetailsItems[i].Updated_At, rec.Part_DetailsItems[i].Row_Order, rec.Part_DetailsItems[i].PartDetails_ID)

			stmt, err = db.Query(query)

			if err != nil {
				fmt.Println("failed to update part details record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {

				log.Println(err)

			}
			fmt.Println("1 record Updated")
		} else {
			query = `INSERT INTO 
    part_details(part_id,part_number,department,sub_category,item_class,rate,alternate_part_number,plant,created_by,created_at,updated_by,updated_at,row_order) 
    VALUES($1, $2, $3,$4,$5, $6, $7, $8, $9, $10, $11,$12,$13)`

			stmt, err := db.Prepare(query)
			if err != nil {
				log.Fatal(err)
				return err
			}

			defer stmt.Close()

			res, err := stmt.Exec(
				rec.Part_ID,
				rec.Part_DetailsItems[i].Part_Number,
				rec.Part_DetailsItems[i].Department,
				rec.Part_DetailsItems[i].Sub_category,
				rec.Part_DetailsItems[i].Item_class,
				rec.Part_DetailsItems[i].Item_rate,
				rec.Part_DetailsItems[i].AlternatePartNumber,
				rec.Part_DetailsItems[i].Plant,
				rec.Part_DetailsItems[i].Created_By,
				rec.Part_DetailsItems[i].Created_At,
				rec.Part_DetailsItems[i].Updated_By,
				rec.Part_DetailsItems[i].Updated_At,
				rec.Part_DetailsItems[i].Row_Order,
			)
			if err != nil {
				log.Fatal(err)
				return err
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(affect, "record added")
		}
	}

	var j int
	for j = 0; j < len(rec.Part_DetailsItems); j++ {
		rec.Part_DetailsItems[j].Display_Name = rec.Part_DetailsItems[j].Supplier
		if rec.Part_DetailsItems[j].PartDetails_ID > 0 {

			query = fmt.Sprintf(`UPDATE part_supplier_details SET part_number = '%s',department = '%s',sub_category = '%s',buyer = '%s',supplier = '%s',display_name = '%s',updated_by = %d,updated_at = '%s',row_order = %d Where 
part_number = (Select part_number from part_details where part_details_id  = %d) AND department = (Select department from part_details where part_details_id  = %d)`, rec.Part_DetailsItems[j].Part_Number, rec.Part_DetailsItems[j].Department, rec.Part_DetailsItems[j].Sub_category, rec.Part_DetailsItems[j].Buyer, rec.Part_DetailsItems[j].Supplier, rec.Part_DetailsItems[j].Display_Name, rec.Part_DetailsItems[j].Updated_By, rec.Part_DetailsItems[j].Updated_At, rec.Part_DetailsItems[j].Row_Order, rec.Part_DetailsItems[j].PartDetails_ID, rec.Part_DetailsItems[j].PartDetails_ID)

			stmt, err = db.Query(query)

			if err != nil {
				fmt.Println("failed to update part supplier details record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Println(err)

			}
			fmt.Println("1 record Updated")
		} else {
			query = `INSERT INTO 
    part_supplier_details(part_id,part_number,department,sub_category,buyer,supplier,display_name,created_by,created_at,updated_by,updated_at,row_order) 
    VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`

			stmt, err := db.Prepare(query)
			if err != nil {
				log.Fatal(err)
				return err
			}

			defer stmt.Close()

			res, err := stmt.Exec(
				rec.Part_ID,
				rec.Part_DetailsItems[j].Part_Number,
				rec.Part_DetailsItems[j].Department,
				rec.Part_DetailsItems[j].Sub_category,
				rec.Part_DetailsItems[j].Buyer,
				rec.Part_DetailsItems[j].Supplier,
				rec.Part_DetailsItems[j].Display_Name,
				rec.Part_DetailsItems[j].Created_By,
				rec.Part_DetailsItems[j].Created_At,
				rec.Part_DetailsItems[j].Updated_By,
				rec.Part_DetailsItems[j].Updated_At,
				rec.Part_DetailsItems[j].Row_Order,
			)
			if err != nil {
				log.Fatal(err)
				return err
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(affect, "record added")
		}
	}

	var k int
	for k = 0; k < len(rec.Part_DetailsItems); k++ {

		if rec.Part_DetailsItems[k].PartDetails_ID > 0 {

			query = fmt.Sprintf(`UPDATE norm_cd SET department = '%s',part_number = '%s',lead_time = %f,safety_factor = %f,buyer = '%s',supplier ='%s',updated_by = %d,updated_at = '%s',row_order = %d Where part_number = (Select part_number from part_details where part_details_id  = %d) AND department = (Select department from part_details where part_details_id  = %d)`, rec.Part_DetailsItems[k].Department, rec.Part_DetailsItems[k].Part_Number, rec.Part_DetailsItems[k].Lead_Time, rec.Part_DetailsItems[k].Safety_factor, rec.Part_DetailsItems[k].Buyer, rec.Part_DetailsItems[k].Supplier, rec.Part_DetailsItems[k].Updated_By, rec.Part_DetailsItems[k].Updated_At, rec.Part_DetailsItems[k].Row_Order, rec.Part_DetailsItems[k].PartDetails_ID, rec.Part_DetailsItems[k].PartDetails_ID)

			stmt, err = db.Query(query)

			if err != nil {
				fmt.Println("failed to update norm calculation details record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Println(err)

			}
			fmt.Println("1 record Updated")
		} else {
			stmt, err := db.Prepare("INSERT INTO Norm_CD(part_id,date,department,part_number,lead_time,safety_factor,buyer,supplier,created_by,created_at,updated_by,updated_at,row_order) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)")
			if err != nil {
				log.Fatal(err)
				return err
			}

			res, err := stmt.Exec(
				rec.Part_ID,
				rec.Part_DetailsItems[k].Date,
				rec.Part_DetailsItems[k].Department,
				rec.Part_DetailsItems[k].Part_Number,
				rec.Part_DetailsItems[k].Lead_Time,
				rec.Part_DetailsItems[k].Safety_factor,
				rec.Part_DetailsItems[k].Buyer,
				rec.Part_DetailsItems[k].Supplier,
				rec.Part_DetailsItems[k].Created_By,
				rec.Part_DetailsItems[k].Created_At,
				rec.Part_DetailsItems[k].Updated_By,
				rec.Part_DetailsItems[k].Updated_At,
				rec.Part_DetailsItems[k].Row_Order,
			)

			if err != nil {
				log.Fatal(err)
				return err
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(affect, "record added")
		}
	}

	return err
}

func (db *DB_manager) CheckDuplicatesPartsValidationRecords(PartNo string) (Item_Master, error) {
	rows, err := db.Query(fmt.Sprintf("Select DISTINCT count(*),part_id from Part Where part_number= '%s' GROUP BY part_id", PartNo))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from check page access ", err)
		return Item_Master{}, err
	}

	var ItemMasterRec Item_Master
	for rows.Next() {
		err := rows.Scan(&ItemMasterRec.Count, &ItemMasterRec.Part_ID)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		if ItemMasterRec.Count > 0 {
			ItemMasterRec.Status = "true"
		} else {
			ItemMasterRec.Status = "False"
		}
	}
	return ItemMasterRec, nil
}

//Dropdown to populate buyer names
func (db *DB_manager) GETBuyerNameRecords() ([]DropdownBuyer, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT BuyerName AS ID,BuyerName AS Name FROM Buyer"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Program Number", err)
		return nil, err
	}
	var BuyerArray []DropdownBuyer
	var BuyerRec DropdownBuyer
	for rows.Next() {
		err := rows.Scan(&BuyerRec.BuyerID, &BuyerRec.BuyerName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		BuyerArray = append(BuyerArray, BuyerRec)
	}

	return BuyerArray, nil
}

//Dropdown to populate supplier names
func (db *DB_manager) GETSupplierNameRecords() ([]DropdownSupplier, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT Suppliername AS ID,Suppliername AS Name FROM Supplier"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Program Number", err)
		return nil, err
	}
	var SupplierArray []DropdownSupplier
	var SupplierRec DropdownSupplier
	for rows.Next() {
		err := rows.Scan(&SupplierRec.SupplierID, &SupplierRec.SupplierName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		SupplierArray = append(SupplierArray, SupplierRec)
	}

	return SupplierArray, nil
}

//Dropdown to populate plant names
func (db *DB_manager) GETPlantNameRecords() ([]DropdownPlant, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT DISTINCT Plantname AS ID,Plantname AS Name FROM Plant"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Program Number", err)
		return nil, err
	}
	var PlantArray []DropdownPlant
	var PlantRec DropdownPlant
	for rows.Next() {
		err := rows.Scan(&PlantRec.PlantID, &PlantRec.PlantName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		PlantArray = append(PlantArray, PlantRec)
	}

	return PlantArray, nil
}
