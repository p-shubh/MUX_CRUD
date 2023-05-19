package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

type ItemType string

const (
	ItemType_NORM       ItemType = "NORM"
	ItemType_AMR        ItemType = "AMR"
	ItemType_Stock      ItemType = "Stock"
	ItemType_Norm_CD    ItemType = "Norm_CD"
	ItemType_Import_AMR ItemType = "Import_AMR"
	ItetmType_Invalid   ItemType = "Invalid"
)

func StringToItemType(itemType string) ItemType {
	switch itemType {
	case "NORM":
		return ItemType_NORM
	case "AMR":
		return ItemType_AMR
	case "Stock":
		return ItemType_Stock
	case "Norm_CD":
		return ItemType_Norm_CD
	case "Import_AMR":
		return ItemType_Import_AMR
	}
	return ItetmType_Invalid
}

type Norm struct {
	Date             string `json:"date"`
	Department       string
	Sub_category     string
	Part_number      string
	Part_description string
	Item_class       string
	AMR              float64
	Norm             float64
	Created_By       int
	Created_At       string
	Updated_By       int
	Updated_At       string
}

type Norm_CD struct {
	Date          string
	Department    string
	Sub_category  string
	Part_number   string
	NOD           float64
	Lead_time     float64
	Safety_factor float64
	Item_class    string
	Buyer         string
	Supplier      string
	Rate          float64
	Created_By    int
	Created_At    string
	Updated_By    int
	Updated_At    string
}

type BPR struct {
	Date                  string
	Department            string
	FG_Assembly_Code      string
	Part_number           string
	Description           *string `json:"Description,omitempty"`
	Category              *string
	Sub_category          string
	Item_class            *string `json:"Item_class,omitempty"`
	AMR                   float64
	Norm                  float64
	Stock                 float64
	Stock_2               float64
	Gap                   float64
	Penetration           float64
	Colour                string
	Colour_2              string
	Total_Coverage        float64
	Desired_inventory     float64
	Actual_inventory      float64
	White_inventory       float64
	Buyer                 *string
	Supplier              *string
	Alternate_Part_number *string
	Created_By            int
	Created_At            string
	Updated_By            int
	Updated_At            string
}

type ImportBPR struct {
	Date                     string
	Department               string
	Part_number              string
	Description              string
	Sub_category             string
	Item_class               string
	Supplier                 string
	Plant                    string
	AMR_1                    float64
	AMR_2                    float64
	AMR_3                    float64
	Todate_aval              float64
	Total_Stock_1            float64
	Colour                   string
	Total_coverage_in_months float64
	Store_Stock              float64
	Under_inspection         float64
	Total_Stock_2            float64
	Colour_2                 string
	Desired_inventory        float64
	Actual_inventory         float64
	White_inventory          float64
}

type Part struct {
	Part_number string `json:"part_number"`
	Description string `json:"description"`
}

type Part_details struct {
	Part_number  string `json:"part_number"`
	Department   string
	Sub_category string
	Item_class   string
	Rate         float64
}

type Part_supplier_output struct {
	Department   string
	Buyer        string
	Supplier     string
	CountofRed   int
	CountofBlack int
	Created_By   int
	Created_At   string
	Updated_By   int
	Updated_At   string
}

type Part_supplier_details struct {
	Part_number string
	Department  string
	Buyer       string
	Supplier    string
}

type Subassembly struct {
	SANumber    string         `json:"sa_number"`
	Assembly    string         `json: "assembly"` // this is also equal to part number in few data models
	Description sql.NullString `json:"description"`
}

type AMR struct {
	Date               string `json:"date"`
	Department         string
	Sub_category       string
	File_Type          string
	Part_number        string
	Part_Description   string
	To_Date_Aval       float64
	Net_Month_Shortage float64
	AMR                float64
	Created_By         int
	Created_At         string
	Updated_By         int
	Updated_At         string
}

type Import_AMR struct {
	Date             string `json:"date"`
	Department       string
	Part_number      string
	Part_Description string
	AMR_1            float64
	AMR_2            float64
	AMR_3            float64
	Item_class       string
	Created_By       int
	Created_At       string
	Updated_By       int
	Updated_At       string
}

type Import_Norm struct {
	Date             string `json:"date"`
	Department       string
	Sub_category     string
	Part_number      string
	Part_description string
	Item_class       string
	AMR_1            float64
	AMR_2            float64
	AMR_3            float64
	Norm_1           float64
	Norm_2           float64
	Norm_3           float64
	Created_By       int
	Created_At       string
	Updated_By       int
	Updated_At       string
}

type Stock struct {
	Date                 string `json:"date"`
	Department           string `json:"department"`
	Part_number          string
	Storage_location     string
	ToDate_Aval          float64
	Under_inspection     float64
	Store_stock          float64
	VMI_stock            float64
	Sub_Contractor_stock float64
	Total_stock          float64
	Total_stock_2        float64
	Created_By           int
	Created_At           string
	Updated_By           int
	Updated_At           string
}

type Stock_file_log struct {
	Date       string `json:"date"`
	Department string `json:"department"`
	File_type  string
	File_name  string
	Created_By int
	Created_At string
	Updated_By int
	Updated_At string
}

type Stock_file_map struct {
	Department      string `json:"department"`
	Stock_file_type string
}

type Stock_file_Status struct {
	Stock_file_type string
	Status          string
}

type Sub_category struct {
	Sub_categoryID int
	DepartmentID   int
	Department     string `json:"department"`
	Sub_category   string
	Display_Order  int
	Created_By     *int
	Created_At     *string
	Updated_By     *int
	Updated_At     *string
}
type Department struct {
	DepartmentID  int
	Department    string `json:"department"`
	Display_Order int
	Created_By    *int
	Created_At    *string
	Updated_By    *int
	Updated_At    *string
}

type Item_class struct {
	Item_classID   int
	Sub_categoryID int
	DepartmentID   int
	Department     string `json:"department"`
	Sub_category   string
	Item_class     string
	Created_By     *int
	Created_At     *string
	Updated_By     *int
	Updated_At     *string
}

type Main struct {
	Status                    string
	Status_code               string
	Count                     int
	Department                []DepartmentwiseInfo
	BuyerwiseTrendInfo        []BuyerwiseTrend
	BPRSummaryInfo            []BPRSummary
	Supplier                  []SupplierInfo
	DisplayBPRInfo            []BPR
	DisplayLineWiseBPR        []BPR
	DisplaySubcategorywiseBPR []BPR
	DisplayImportBPR          []ImportBPR
	ColourTrend               []ColourTrendInfo
	FullKitPartWise           []FullKitPartWiseInfo
	SupplierSummaryInfo       []SupplierSummary
}

type ColourTrendInfo struct {
	Department     string
	Part_number    string
	Sub_category   string
	Description    string
	BPRColourTrend []BPRColourTrendInfo
}

type BPRColourTrendInfo struct {
	Date        string
	Penetration float64
	Colour      string
}

type DepartmentwiseInfo struct {
	DepartmentID   int
	DepartmentName string
	Sub_category   string
}

type BPRSummary struct {
	Date                   string
	DepartmentName         string
	Sub_category           string
	Item_class             string
	Display_order          int
	Line                   string
	Total                  float64
	B                      int
	R                      int
	Y                      int
	G                      int
	W                      int
	Percent_of_Red_Items   float64
	Percnet_of_White_Items float64
	NoCriticalPartListFlag int
}

type Full_kit_master struct {
	Date                    string `json:"date"`
	Department              string `json:"department"`
	FG_assembly_code        string
	FG_assembly_description string
	Item_class              string
	Remark                  int
}

type Full_kit struct {
	FG_assembly_code        string
	FG_assembly_description string
	Sub_category            string
	Total                   int
	G                       int
	R                       int
	W                       int
	Y                       int
}

type Full_kit_output struct {
	Date             string `json:"date"`
	Department       string `json:"department"`
	FG_assembly_code string
	Remark           int
	Item_class       string
	Created_By       int
	Created_At       string
	Updated_By       int
	Updated_At       string
}

type SupplierInfo struct {
	Date         string
	Sub_category string
	Part_number  string
	Description  string
	Lead_time    float64
	Norm         float64
	Stock        float64
	Gap          float64
	Penetration  float64
	Colour       string
	Buyer        string
	Supplier     string
}

type BuyerwiseTrend struct {
	Date                   string
	Department             string
	Buyer                  string
	Total                  int
	B                      int
	R                      int
	Y                      int
	G                      int
	W                      int
	Percent_of_Red_Items   float64
	Percnet_of_White_Items float64
}

type SupplierName struct {
	Part_number  string
	Department   string
	Sub_category string
	Buyer        string
	Supplier     string
}

type BPRLinewise struct {
	Department              string
	FG_assembly_code        string
	FG_assembly_description string
	Child_part_code         string
	Line                    string
}

type FullKitPartWiseInfo struct {
	FG_assembly_code string
	Part_number      string
	Description      string
	Item_class       string
	Norm             float64
	Stock            float64
	Gap              float64
	Penetration      float64
	Colour           string
}

type SupplierSummary struct {
	Date                   string
	DepartmentName         string
	VendorName             string
	Total                  int
	B                      int
	R                      int
	Y                      int
	G                      int
	W                      int
	Percent_of_Red_Items   float64
	Percnet_of_White_Items float64
}

func (db *DB_manager) InsertCalculateNorm(rec Norm) (int, int, error) {
	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM norm where department='%s' AND date='%s' AND part_number = '%s'", rec.Department, rec.Date, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from norm ", err)
		fmt.Println("failed_query ", qry)
		return 0, 0, err
	}

	var CountRec int
	InsertedCount := 0
	UpdatedCount := 0
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec == 0 {
			query := `INSERT INTO 
    norm (date, department, part_number,item_class,amr,norm,created_by,created_at,updated_by,updated_at) 
    VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
			stmt, err := db.Prepare(query)
			if err != nil {
				fmt.Println("failed to insert norm record. Continue with the next operation ", rec)
				return 0, 0, err
			}

			defer stmt.Close()

			res, err := stmt.Exec(rec.Date,
				rec.Department,
				rec.Part_number,
				rec.Item_class,
				rec.AMR,
				rec.Norm,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)
			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			InsertedCount++
			fmt.Println(affect, "record added")
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE  norm SET amr = %f,norm=%f,updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' AND date='%s'`, rec.AMR, rec.Norm, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number, rec.Date)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update norm record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			UpdatedCount++
			fmt.Println("1 record Updated")
		}
		return InsertedCount, UpdatedCount, nil
	}
	return 0, 0, errors.New("failed_to_get_norm_record")
}

func (db *DB_manager) InsertAMR(rec AMR) (int, int, error) {
	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM amr where department='%s' AND date='%s' AND part_number = '%s'", rec.Department, rec.Date, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from amr", err)
		fmt.Println("failed_query ", qry)
		return 0, 0, err
	}

	var CountRec int
	InsertedAMRCount := 0
	UpdatedAMRCount := 0
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec == 0 {
			query := `INSERT INTO AMR(date,department,file_type,part_number,part_description,to_date_aval,net_month_shortage,AMR,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`
			stmt, err := db.Prepare(query)
			if err != nil {
				fmt.Println("failed to insert amr record. Continue with the next operation ", rec)
				return 0, 0, err
			}

			defer stmt.Close()

			res, err := stmt.Exec(rec.Date,
				rec.Department,
				rec.File_Type,
				rec.Part_number,
				rec.Part_Description,
				rec.To_Date_Aval,
				rec.Net_Month_Shortage,
				rec.AMR,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)

			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			InsertedAMRCount++
			fmt.Println(affect, "record added")
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE amr SET amr = %f,file_type = '%s',updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' AND date='%s'`, rec.AMR, rec.File_Type, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number, rec.Date)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update amr record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			UpdatedAMRCount++
			fmt.Println("1 record Updated")
		}
		return InsertedAMRCount, UpdatedAMRCount, nil
	}
	return 0, 0, errors.New("failed_to_get_amr_record")
}

func (db *DB_manager) InsertImport_AMR(rec Import_AMR) (int, int, error) {
	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM import_amr where department='%s' AND date='%s' AND part_number = '%s'", rec.Department, rec.Date, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from import_amr ", err)
		fmt.Println("failed_query ", qry)
		return 0, 0, err
	}

	var CountRec int
	InsertedImportAMRCount := 0
	UpdatedImportAMRCount := 0
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec == 0 {
			query := `INSERT INTO Import_AMR(date,department,part_number,part_description,AMR_1,AMR_2,AMR_3,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
			stmt, err := db.Prepare(query)
			if err != nil {
				fmt.Println("failed to insert import_amr record. Continue with the next operation ", rec)
				return 0, 0, err
			}

			defer stmt.Close()

			res, err := stmt.Exec(rec.Date,
				rec.Department,
				rec.Part_number,
				rec.Part_Description,
				rec.AMR_1,
				rec.AMR_2,
				rec.AMR_3,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)

			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			InsertedImportAMRCount++
			fmt.Println(affect, "record added")
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE import_amr SET amr_1 = %f,amr_2 = %f,amr_3 = %f,updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' AND date='%s'`, rec.AMR_1, rec.AMR_2, rec.AMR_3, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number, rec.Date)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update import_amr record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			UpdatedImportAMRCount++
			fmt.Println("1 record Updated")
		}
		return InsertedImportAMRCount, UpdatedImportAMRCount, nil
	}
	return 0, 0, errors.New("failed_to_get_import_amr_record")
}

func (db *DB_manager) InsertImport_Norm(rec Import_Norm) (int, int, error) {
	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM import_norm where department='%s' AND date='%s' AND part_number = '%s'", rec.Department, rec.Date, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from import_norm ", err)
		fmt.Println("failed_query ", qry)
		return 0, 0, err
	}

	var CountRec int
	InsertedImportNormCount := 0
	UpdatedImportNormCount := 0
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec == 0 {
			query := `INSERT INTO Import_Norm(date,department,part_number,item_class,AMR_1,AMR_2,AMR_3,Norm_1,Norm_2,Norm_3,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`
			stmt, err := db.Prepare(query)
			if err != nil {
				fmt.Println("failed to insert import_norm record. Continue with the next operation ", rec)
				return 0, 0, err
			}

			defer stmt.Close()

			res, err := stmt.Exec(rec.Date,
				rec.Department,
				rec.Part_number,
				rec.Item_class,
				rec.AMR_1,
				rec.AMR_2,
				rec.AMR_3,
				rec.Norm_1,
				rec.Norm_2,
				rec.Norm_3,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)

			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			InsertedImportNormCount++
			fmt.Println(affect, "record added")
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE import_norm SET amr_1 = %f,amr_2 = %f,amr_3 = %f,norm_1 = %f,norm_2 = %f,norm_3 = %f,updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' AND date='%s'`, rec.AMR_1, rec.AMR_2, rec.AMR_3, rec.Norm_1, rec.Norm_2, rec.Norm_3, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number, rec.Date)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update import_norm record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			UpdatedImportNormCount++
			fmt.Println("1 record Updated")
		}
		return InsertedImportNormCount, UpdatedImportNormCount, nil
	}
	return 0, 0, errors.New("failed_to_get_import_norm_record")
}

func (db *DB_manager) InsertUpdateNorm_CD(rec Norm_CD) (int, int, error) {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM norm_cd where department='%s' AND part_number = '%s'", rec.Department, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from norm_cd ", err)
		fmt.Println("failed_query ", qry)
		return 0, 0, err
	}

	var CountRec int
	InsertedCount := 0
	UpdatedCount := 0
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE norm_cd SET date= '%s',department= '%s',part_number = '%s',nod = %f ,lead_time = %f,safety_factor = %f,buyer = '%s',supplier = '%s',updated_by = %d,updated_at= '%s' WHERE department='%s' AND part_number='%s'`, rec.Date, rec.Department, rec.Part_number, rec.NOD, rec.Lead_time, rec.Safety_factor, rec.Buyer, rec.Supplier, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update norm_cd record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			UpdatedCount++
			fmt.Println("1 record Updated")
		} else {
			stmt, err := db.Prepare("INSERT INTO Norm_CD(date,department,part_number,NOD,lead_time,safety_factor,buyer,supplier,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(rec.Date,
				rec.Department,
				rec.Part_number,
				rec.NOD,
				rec.Lead_time,
				rec.Safety_factor,
				rec.Buyer,
				rec.Supplier,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)

			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			InsertedCount++
			fmt.Println(affect, "record added")
		} // else end norm_cd
		return InsertedCount, UpdatedCount, nil
	}
	return 0, 0, errors.New("failed_to_get_norm_cd_record")
}

func (db *DB_manager) InsertUpdatePartDetails(rec Norm_CD) error {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM part_details where department='%s' AND part_number = '%s'", rec.Department, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_details ", err)
		fmt.Println("failed_query ", qry)
		return err
	}

	var CountRec int
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE part_details SET part_number = '%s',department = '%s',sub_category = '%s',rate = %f,updated_by = %d,updated_at = '%s' WHERE department='%s' AND part_number='%s'`, rec.Part_number, rec.Department, rec.Sub_category, rec.Rate, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update part_details record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("1 record Updated")
		} else {
			stmt, err := db.Prepare("INSERT INTO part_details(part_number,department,sub_category,rate,created_by,created_at,updated_by,updated_at) VALUES($1, $2, $3,$4,$5, $6, $7, $8)")
			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(rec.Part_number,
				rec.Department,
				rec.Sub_category,
				rec.Rate,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)
			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(affect, "record added")
		} // else end part_details
		return err
	}
	return err
}

func (db *DB_manager) InsertUpdatePartSupplierDetails(rec Norm_CD) error {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM part_supplier_details where department='%s' AND part_number = '%s'", rec.Department, rec.Part_number)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_supplier_details ", err)
		fmt.Println("failed_query ", qry)
		return err
	}

	var CountRec int
	for rows.Next() {
		err := rows.Scan(&CountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		if CountRec > 0 {
			query := fmt.Sprintf(`UPDATE part_supplier_details SET part_number = '%s',department = '%s',sub_category = '%s',buyer = '%s',supplier = '%s',updated_by = %d,updated_at = '%s'
WHERE department='%s' AND part_number='%s'`, rec.Part_number, rec.Department, rec.Sub_category, rec.Buyer, rec.Supplier, rec.Updated_By, rec.Updated_At, rec.Department, rec.Part_number)

			stmt, err := db.Query(query)

			if err != nil {
				fmt.Println("failed to update part_supplier_details record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("1 record Updated")
		} else {
			stmt, err := db.Prepare("INSERT INTO part_supplier_details(part_number,department,sub_category,buyer,supplier,created_by,created_at,updated_by,updated_at) VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9)")

			if err != nil {
				log.Fatal(err)
			}

			res, err := stmt.Exec(rec.Part_number,
				rec.Department,
				rec.Sub_category,
				rec.Buyer,
				rec.Supplier,
				rec.Created_By,
				rec.Created_At,
				rec.Updated_By,
				rec.Updated_At,
			)
			if err != nil {
				log.Fatal(err)
			}

			affect, err := res.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(affect, "record added")
		} // else end part_supplier_details
		return err
	}
	return err
}

func (db *DB_manager) InsertBPR(rec BPR) {
	stmt, err := db.Prepare("INSERT INTO bpr(date,department,part_number,norm,stock,stock_2,gap,penetration,colour,colour_2,total_coverage,desired_inventory,actual_inventory,white_inventory,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4, $5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(rec.Date,
		rec.Department,
		rec.Part_number,
		rec.Norm,
		rec.Stock,
		rec.Stock_2,
		rec.Gap,
		rec.Penetration,
		rec.Colour,
		rec.Colour_2,
		rec.Total_Coverage,
		rec.Desired_inventory,
		rec.Actual_inventory,
		rec.White_inventory,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

}

func (db *DB_manager) InsertStock(rec Stock) {
	stmt, err := db.Prepare("INSERT INTO Stock(date,department,part_number, todate_aval,under_inspection, store_stock, VMI_stock, Sub_Contractor_stock, Total_stock,total_stock_2,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(rec.Date,
		rec.Department,
		rec.Part_number,
		rec.ToDate_Aval,
		rec.Under_inspection,
		rec.Store_stock,
		rec.VMI_stock,
		rec.Sub_Contractor_stock,
		rec.Total_stock,
		rec.Total_stock_2,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB_manager) InsertStock_file_log(rec Stock_file_log) {
	stmt, err := db.Prepare("INSERT INTO Stock_file_log(date,department,file_type,file_name,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(rec.Date,
		rec.Department,
		rec.File_type,
		rec.File_name,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB_manager) InsertFull_kit_output(rec Full_kit_output) {
	stmt, err := db.Prepare("INSERT INTO Full_kit_output(date,department,fg_assembly_code,remark,item_class,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(rec.Date,
		rec.Department,
		rec.FG_assembly_code,
		rec.Remark,
		rec.Item_class,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB_manager) InsertPart_supplier_output(rec Part_supplier_output) {
	stmt, err := db.Prepare("INSERT INTO Part_supplier_output(department,buyer,supplier,count_of_red,count_of_black,created_by,created_at,updated_by,updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(rec.Department,
		rec.Buyer,
		rec.Supplier,
		rec.CountofRed,
		rec.CountofBlack,
		rec.Created_By,
		rec.Created_At,
		rec.Updated_By,
		rec.Updated_At,
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
}

func (db *DB_manager) UpdateNorm(rec1 Norm) {
	query := fmt.Sprintf(`UPDATE  norm SET norm=%f,updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' AND date='%s'`, rec1.Norm, rec1.Updated_By, rec1.Updated_At, rec1.Department, rec1.Part_number, rec1.Date)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update norm record. Continue with the next operation ", rec1)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1 record Updated")
}

func (db *DB_manager) UpdateImportNorm(rec1 Import_Norm) {
	query := fmt.Sprintf(`UPDATE  import_norm SET norm_1=%f,norm_2=%f,norm_3=%f,updated_by = %d ,updated_at = '%s' WHERE department='%s' AND part_number='%s' 
AND date='%s'`, rec1.Norm_1, rec1.Norm_2, rec1.Norm_3, rec1.Updated_By, rec1.Updated_At, rec1.Department, rec1.Part_number, rec1.Date)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update import_norm record. Continue with the next operation ", rec1)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("1 record Updated")
}

func (db *DB_manager) ReadNormRecord(DepartmentName string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM norm WHERE date = (SELECT MAX(date) FROM norm WHERE department = '%s') AND department = '%s'", DepartmentName, DepartmentName))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Norm ", err)

		return nil, err
	}
	var NormArray []Norm
	var NormRec Norm
	for rows.Next() {
		err := rows.Scan(&NormRec.Date, &NormRec.Department, &NormRec.Part_number, &NormRec.Item_class, &NormRec.AMR, &NormRec.Norm, &NormRec.Created_By, &NormRec.Created_At, &NormRec.Updated_By, &NormRec.Updated_At)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		NormArray = append(NormArray, NormRec)
	}
	return NormArray, nil
}

func (db *DB_manager) ReadPartRecords(partNumberRegex string) ([]string, error) {
	partNumbers := []string{}
	rows, err := db.Query("SELECT * FROM part WHERE part_number LIKE '" + partNumberRegex + "%'")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part ", err)
		return partNumbers, err
	}

	var part Part
	for rows.Next() { // hopefully we are lucky enough to find only 1 part record.
		err := rows.Scan(&part.Part_number, &part.Description)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		partNumbers = append(partNumbers, part.Part_number)
	}

	return partNumbers, nil
}

// ReadSubassemblyRecords reads all subassembly associated with an assembly == part number.
// returns list of all subassembly parts in return
func (db *DB_manager) ReadSubassemblyRecords(assembly string) ([]string, error) {
	subAssemblies := []string{}
	rows, err := db.Query("SELECT * FROM subassembly where assembly = '" + assembly + "'")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from subassembly  ", err)
		return subAssemblies, err
	}

	var saRec Subassembly
	for rows.Next() {
		err := rows.Scan(&saRec.SANumber, &saRec.Assembly, &saRec.Description)
		if err != nil {
			fmt.Println("Subassembly: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		subAssemblies = append(subAssemblies, saRec.SANumber)
	}

	return subAssemblies, nil
}

func (db *DB_manager) ReadstockfilemapRecord(DepartmentName string, Date string) ([]Stock_file_Status, error) {

	rows, err := db.Query("SELECT stock_file_map.stock_file_type, CASE when stock_file_log.file_name IS NULL then 'Pending' else 'Uploaded' END AS Status FROM stock_file_map left join stock_file_log ON stock_file_map.department = stock_file_log.department AND stock_file_map.stock_file_type = stock_file_log.file_type AND stock_file_log.date = '" + Date + "'where stock_file_map.department = '" + DepartmentName + "'")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Stock_file_map  ", err)
		return nil, err
	}
	var saRecArray []Stock_file_Status
	var saRec Stock_file_Status
	for rows.Next() {
		err := rows.Scan(&saRec.Stock_file_type, &saRec.Status)
		if err != nil {
			fmt.Println("Stock_file_map: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		saRecArray = append(saRecArray, saRec)
	}

	return saRecArray, nil
}

func (db *DB_manager) ReadItemclassRecord(DepartmentName string, sub_category string) ([]string, error) {
	itemclass := []string{}
	rows, err := db.Query("SELECT * FROM item_class where department = '" + DepartmentName + "' AND sub_category = '" + sub_category + "' ")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from item_class  ", err)
		return itemclass, err
	}

	var icRec Item_class
	for rows.Next() {
		err := rows.Scan(&icRec.Item_classID, &icRec.Sub_categoryID, &icRec.DepartmentID, &icRec.Department, &icRec.Sub_category, &icRec.Item_class, &icRec.Created_By, &icRec.Created_At, &icRec.Updated_By, &icRec.Updated_At)
		if err != nil {
			fmt.Println("Sub_category: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		itemclass = append(itemclass, icRec.Item_class)
	}

	return itemclass, nil
}

func (db *DB_manager) ReadSubcategoryRecord(DepartmentName string) ([]string, error) {
	subcategory := []string{}
	rows, err := db.Query("SELECT * FROM Sub_category where department = '" + DepartmentName + "'")
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Sub_category  ", err)
		return subcategory, err
	}

	var saRec Sub_category
	for rows.Next() {
		err := rows.Scan(&saRec.Sub_categoryID, &saRec.Department, &saRec.Sub_category, &saRec.Display_Order, &saRec.DepartmentID, &saRec.Created_By, &saRec.Created_At, &saRec.Updated_By, &saRec.Updated_At)
		if err != nil {
			fmt.Println("Sub_category: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		subcategory = append(subcategory, saRec.Sub_category)
	}

	return subcategory, nil
}

func (db *DB_manager) ReadDepartmentRecord(UserID int) ([]string, error) {
	department := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM department where departmentid IN (SELECT departmentid FROM mc_userdivisions WHERE userid = %d)", UserID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from department  ", err)
		return department, err
	}

	var daRec Department
	for rows.Next() {
		err := rows.Scan(&daRec.DepartmentID, &daRec.Department, &daRec.Display_Order, &daRec.Created_By, &daRec.Created_At, &daRec.Updated_By, &daRec.Updated_At)
		if err != nil {
			fmt.Println("Department: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		department = append(department, daRec.Department)
	}

	return department, nil
}

func (db *DB_manager) ReadAMRRecord(DepartmentName string, Date string, File_Type string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT amr.date,amr.department,coalesce(part_details.sub_category,'') AS sub_category,amr.file_type,amr.part_number,amr.part_description,amr.to_date_aval,amr.net_month_shortage,amr.amr FROM amr LEFT JOIN part_details ON amr.part_number = part_details .part_number AND amr.department = part_details .department where amr.department='%s' AND amr.date='%s' AND amr.file_type = '%s'", DepartmentName, Date, File_Type))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from AMR ", err)
		return AMR{}, err
	}
	var AmrArray []AMR
	var AMRRec AMR
	for rows.Next() {
		err := rows.Scan(&AMRRec.Date, &AMRRec.Department, &AMRRec.Sub_category, &AMRRec.File_Type, &AMRRec.Part_number, &AMRRec.Part_Description, &AMRRec.To_Date_Aval, &AMRRec.Net_Month_Shortage, &AMRRec.AMR)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		AmrArray = append(AmrArray, AMRRec)
	}

	return AmrArray, nil
}

func (db *DB_manager) GETAMRRecord(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT amr.date,amr.department,part_details.sub_category,amr.file_type,amr.part_number,amr.part_description,amr.to_date_aval,amr.net_month_shortage,amr.amr FROM amr JOIN part_details ON amr.part_number = part_details .part_number AND amr.department = part_details .department where amr.department='%s' AND amr.date='%s'", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from AMR ", err)
		return AMR{}, err
	}
	var AmrArray []AMR
	var AMRRec AMR
	for rows.Next() {
		err := rows.Scan(&AMRRec.Date, &AMRRec.Department, &AMRRec.Sub_category, &AMRRec.File_Type, &AMRRec.Part_number, &AMRRec.Part_Description, &AMRRec.To_Date_Aval, &AMRRec.Net_Month_Shortage, &AMRRec.AMR)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		AmrArray = append(AmrArray, AMRRec)
	}

	return AmrArray, nil
}

func (db *DB_manager) ReadImport_NormRecord(DepartmentName string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT * FROM import_norm WHERE date = (SELECT MAX(date) FROM import_norm WHERE department = '%s') AND department = '%s'", DepartmentName, DepartmentName))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Import_Norm ", err)
		return nil, err
	}
	var Import_NormArray []Import_Norm
	var Import_NormRec Import_Norm
	for rows.Next() {
		err := rows.Scan(&Import_NormRec.Date, &Import_NormRec.Department, &Import_NormRec.Part_number, &Import_NormRec.Item_class, &Import_NormRec.AMR_1, &Import_NormRec.AMR_2, &Import_NormRec.AMR_3, &Import_NormRec.Norm_1, &Import_NormRec.Norm_2, &Import_NormRec.Norm_3, &Import_NormRec.Created_By, &Import_NormRec.Created_At, &Import_NormRec.Updated_By, &Import_NormRec.Updated_At)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		Import_NormArray = append(Import_NormArray, Import_NormRec)
	}
	return Import_NormArray, nil
}

func (db *DB_manager) ReadImport_AMRRecord(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT import_amr.date,import_amr.department,import_amr.part_number,import_amr.amr_1,import_amr.amr_2,import_amr.amr_3,COALESCE(part_details.item_class,'') FROM import_amr LEFT JOIN part_details ON import_amr.part_number = part_details.part_number AND part_details.department = import_amr.department where import_amr.department='%s' AND import_amr.date='%s'", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Import_AMR ", err)
		return AMR{}, err
	}
	var Import_AmrArray []Import_AMR
	var Import_AMRRec Import_AMR
	for rows.Next() {
		err := rows.Scan(&Import_AMRRec.Date, &Import_AMRRec.Department, &Import_AMRRec.Part_number, &Import_AMRRec.AMR_1, &Import_AMRRec.AMR_2, &Import_AMRRec.AMR_3, &Import_AMRRec.Item_class)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		Import_AmrArray = append(Import_AmrArray, Import_AMRRec)
	}

	return Import_AmrArray, nil
}

func (db *DB_manager) ReadStockRecord(PartNumber string, entryTime string, department string) (Stock, error) {

	qry := fmt.Sprintf("SELECT * FROM stock where part_number='%s' AND date='%s' AND department='%s'", PartNumber, entryTime, department)
	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from stock ", err)
		fmt.Println("failed_query ", qry)
		return Stock{}, err
	}

	var StockRec Stock
	for rows.Next() {
		err := rows.Scan(&StockRec.Date, &StockRec.Department, &StockRec.Part_number, &StockRec.ToDate_Aval, &StockRec.Under_inspection, &StockRec.Store_stock, &StockRec.VMI_stock, &StockRec.Sub_Contractor_stock, &StockRec.Total_stock, &StockRec.Total_stock_2, &StockRec.Created_By, &StockRec.Created_At, &StockRec.Updated_By, &StockRec.Updated_At)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		return StockRec, nil
	}

	return Stock{}, errors.New("failed_to_get_stock_record")
}

func (db *DB_manager) ReadPart_detailsRecord(PartNumber string, department string) (Part_details, error) {
	qry := fmt.Sprintf("SELECT part_number,department,sub_category,item_class,rate FROM part_details where part_number ='%s' AND department='%s'", PartNumber, department)
	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_details ", err)
		fmt.Println("failed_query ", qry)
		return Part_details{}, err
	}

	var PartRec Part_details
	for rows.Next() {
		err := rows.Scan(&PartRec.Part_number, &PartRec.Department, &PartRec.Sub_category, &PartRec.Item_class, &PartRec.Rate)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		return PartRec, nil
	}

	return Part_details{}, errors.New("failed_to_get_part_details_record")
}

func (db *DB_manager) ReadNorm_CDRecord(PartNumber string, department string, Calculation_Parameter int, Sub_category string) (Norm_CD, error) {

	var query_where string

	var query string

	query_main := "SELECT coalesce(norm_cd.date,'') AS date,norm_cd.department,coalesce(part_details.sub_category,'') AS sub_category,norm_cd.part_number,norm_cd.NOD,norm_cd.lead_time,norm_cd.safety_factor,part_details.item_class FROM norm_cd JOIN part_details ON norm_cd.part_number = part_details.part_number AND part_details.department = norm_cd.department Where" + query_where

	if Calculation_Parameter == 1 {
		query_main = query_main + " norm_cd.part_number = '%s' AND norm_cd.department = '%s'"
		query = fmt.Sprintf(query_main, PartNumber, department)
	} else {
		query_main = query_main + " norm_cd.part_number = '%s' AND norm_cd.department = '%s' AND part_details.sub_category = '%s'"
		query = fmt.Sprintf(query_main, PartNumber, department, Sub_category)
	}

	rows, err := db.Query(query)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from norm_cd ", err)
		fmt.Println("failed_query ", query)
		return Norm_CD{}, err
	}

	var Norm_CDRec Norm_CD
	for rows.Next() {
		err := rows.Scan(&Norm_CDRec.Date, &Norm_CDRec.Department, &Norm_CDRec.Sub_category, &Norm_CDRec.Part_number, &Norm_CDRec.NOD, &Norm_CDRec.Lead_time, &Norm_CDRec.Safety_factor, &Norm_CDRec.Item_class)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		return Norm_CDRec, nil
	}

	return Norm_CD{}, errors.New("failed_to_get_norm_record")
}

func (db *DB_manager) ReadBPRRecords(DepartmentName string, Date string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,bpr.part_number,part.description ,part_details.sub_category,part_details.Item_class,(select coalesce(amr,0.0) AS amr from norm where department=bpr.department and part_number=bpr.part_number order by date desc limit 1),bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour,bpr.desired_inventory,bpr.actual_inventory,bpr.white_inventory,part_supplier_details.buyer,part_supplier_details.supplier, part_details.alternate_part_number FROM bpr LEFT JOIN part ON part.part_number = bpr.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department LEFT JOIN part_supplier_details ON bpr.part_number = part_supplier_details.part_number AND  bpr.department = part_supplier_details.department Where bpr.department = '%s' AND bpr.date = '%s' AND part_details.sub_category <> 'Child Parts' ORDER BY bpr.penetration DESC", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr ", err)
		return Main{}, err
	}

	var BPRArray []BPR
	var BPRRec BPR

	for rows.Next() {
		err := rows.Scan(&BPRRec.Date, &BPRRec.Department, &BPRRec.Part_number, &BPRRec.Description, &BPRRec.Sub_category, &BPRRec.Item_class, &BPRRec.AMR, &BPRRec.Norm, &BPRRec.Stock, &BPRRec.Gap, &BPRRec.Penetration, &BPRRec.Colour, &BPRRec.Desired_inventory, &BPRRec.Actual_inventory, &BPRRec.White_inventory, &BPRRec.Buyer, &BPRRec.Supplier, &BPRRec.Alternate_Part_number)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		BPRArray = append(BPRArray, BPRRec)

	}

	var BPRmaininfo Main
	BPRmaininfo.Status = "true"
	BPRmaininfo.Status_code = "200"
	BPRmaininfo.Count = len(BPRArray)
	BPRmaininfo.DisplayBPRInfo = BPRArray
	return BPRmaininfo, nil
}

func (db *DB_manager) ReadBPRLinewiseRecords(Date string, DepartmentName string, LineName string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,full_kit_master.fg_assembly_code,bpr.part_number,part.description,part_details.sub_category,full_kit_master.line AS Category,part_details.Item_class,(select coalesce(amr,0.0) AS amr from norm where department=bpr.department and part_number=bpr.part_number order by date desc limit 1),bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour,bpr.desired_inventory,bpr.actual_inventory,bpr.white_inventory FROM bpr LEFT JOIN part ON bpr.part_number = part.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department = part_details.department LEFT JOIN full_kit_master ON bpr.part_number = full_kit_master.child_part_code AND  full_kit_master.department = bpr.department Where bpr.date = '%s' AND bpr.department = '%s' AND part_details.Sub_category <> 'Finished Goods' AND full_kit_master.line = '%s' ", Date, DepartmentName, LineName))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from linewise bpr ", err)
		return Main{}, err
	}

	var BPRArray1 []BPR
	var BPRRec1 BPR

	for rows.Next() {
		err := rows.Scan(&BPRRec1.Date, &BPRRec1.Department, &BPRRec1.FG_Assembly_Code, &BPRRec1.Part_number, &BPRRec1.Description, &BPRRec1.Sub_category, &BPRRec1.Category, &BPRRec1.Item_class, &BPRRec1.AMR, &BPRRec1.Norm, &BPRRec1.Stock, &BPRRec1.Gap, &BPRRec1.Penetration, &BPRRec1.Colour, &BPRRec1.Desired_inventory, &BPRRec1.Actual_inventory, &BPRRec1.White_inventory)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		BPRArray1 = append(BPRArray1, BPRRec1)

	}
	var LinewiseBPRmaininfo Main
	LinewiseBPRmaininfo.Status = "true"
	LinewiseBPRmaininfo.Status_code = "200"
	LinewiseBPRmaininfo.Count = len(BPRArray1)
	LinewiseBPRmaininfo.DisplayLineWiseBPR = BPRArray1
	return LinewiseBPRmaininfo, nil
}

func (db *DB_manager) ReadBPRSubcategorywiseRecords(Date string, DepartmentName string, SubCategory string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,bpr.part_number,part.description ,part_details.sub_category,part_details.Item_class,(select coalesce(amr,0.0) AS amr from norm where department=bpr.department and part_number=bpr.part_number order by date desc limit 1),bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour,bpr.desired_inventory,bpr.actual_inventory,bpr.white_inventory,coalesce(part_supplier_details.buyer,'') AS buyer,coalesce(part_supplier_details.supplier,'') AS supplier FROM bpr LEFT JOIN part ON part.part_number = bpr.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department LEFT JOIN part_supplier_details ON bpr.part_number = part_supplier_details.part_number AND  bpr.department = part_supplier_details.department Where bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category = '%s' ORDER BY bpr.penetration DESC", Date, DepartmentName, SubCategory))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from sub category wise ", err)
		return Main{}, err
	}

	var BPRArray2 []BPR
	var BPRRec2 BPR

	for rows.Next() {
		err := rows.Scan(&BPRRec2.Date, &BPRRec2.Department, &BPRRec2.Part_number, &BPRRec2.Description, &BPRRec2.Sub_category, &BPRRec2.Item_class, &BPRRec2.AMR, &BPRRec2.Norm, &BPRRec2.Stock, &BPRRec2.Gap, &BPRRec2.Penetration, &BPRRec2.Colour, &BPRRec2.Desired_inventory, &BPRRec2.Actual_inventory, &BPRRec2.White_inventory, &BPRRec2.Buyer, &BPRRec2.Supplier)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		BPRArray2 = append(BPRArray2, BPRRec2)

	}
	var SubcategorywiseBPRmaininfo Main
	SubcategorywiseBPRmaininfo.Status = "true"
	SubcategorywiseBPRmaininfo.Status_code = "200"
	SubcategorywiseBPRmaininfo.Count = len(BPRArray2)
	SubcategorywiseBPRmaininfo.DisplaySubcategorywiseBPR = BPRArray2
	return SubcategorywiseBPRmaininfo, nil
}

func (db *DB_manager) ReadImportBPRRecords(DepartmentName string, Date string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,bpr.part_number,coalesce(part.description,'') AS description,part_details.sub_category,coalesce(part_details.Item_class,'') AS Item_class,coalesce(part_supplier_details.supplier,'') AS supplier,coalesce(part_details.plant,'') AS plant,import_norm.amr_1,import_norm.amr_2,import_norm.amr_3,stock.todate_aval,stock.under_inspection,stock.total_stock,bpr.colour,round(bpr.total_coverage :: numeric,2) AS Total_Coverage_In_Months,stock.store_stock,stock.under_inspection,stock.total_stock_2,bpr.colour_2,bpr.desired_inventory,bpr.actual_inventory,bpr.white_inventory FROM bpr LEFT JOIN part ON part.part_number = bpr.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department LEFT JOIN part_supplier_details ON bpr.part_number = part_supplier_details.part_number AND  bpr.department = part_supplier_details.department LEFT JOIN stock ON bpr.part_number = stock.part_number AND bpr.department = stock.department AND bpr.date = stock.date LEFT JOIN import_norm ON bpr.part_number = import_norm .part_number AND bpr.department = import_norm.department AND import_norm .date = (SELECT MAX(date) FROM import_norm) Where bpr.department = '%s' AND bpr.date = '%s' ORDER BY bpr.penetration DESC", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from import_bpr ", err)
		return Main{}, err
	}

	var IBPRArray []ImportBPR
	var IBPRRec ImportBPR

	for rows.Next() {
		err := rows.Scan(&IBPRRec.Date, &IBPRRec.Department, &IBPRRec.Part_number, &IBPRRec.Description, &IBPRRec.Sub_category, &IBPRRec.Item_class, &IBPRRec.Supplier, &IBPRRec.Plant, &IBPRRec.AMR_1, &IBPRRec.AMR_2, &IBPRRec.AMR_3, &IBPRRec.Todate_aval, &IBPRRec.Under_inspection, &IBPRRec.Total_Stock_1, &IBPRRec.Colour, &IBPRRec.Total_coverage_in_months, &IBPRRec.Store_Stock, &IBPRRec.Under_inspection, &IBPRRec.Total_Stock_2, &IBPRRec.Colour_2, &IBPRRec.Desired_inventory, &IBPRRec.Actual_inventory, &IBPRRec.White_inventory)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		IBPRArray = append(IBPRArray, IBPRRec)

	}

	var IBPRmaininfo Main
	IBPRmaininfo.Status = "true"
	IBPRmaininfo.Status_code = "200"
	IBPRmaininfo.Count = len(IBPRArray)
	IBPRmaininfo.DisplayImportBPR = IBPRArray
	return IBPRmaininfo, nil
}

func (db *DB_manager) ReadStockFileMaplogRecord(DepartmentName string, Date string) (int, int, error) {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM stock_file_map where department='%s'", DepartmentName)
	qry1 := fmt.Sprintf("SELECT DISTINCT count(*) FROM stock_file_log where department='%s' AND date='%s'", DepartmentName, Date)

	rows, err := db.Query(qry)
	rows1, err := db.Query(qry1)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from stockfilemaplog ", err)
		fmt.Println("failed_query ", qry)
		fmt.Println("failed_query ", qry1)
		return 0, 0, err
	}

	var ReqCountRec int

	var CurrentCountRec int
	for rows.Next() {
		err := rows.Scan(&ReqCountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
	}
	for rows1.Next() {
		err := rows1.Scan(&CurrentCountRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		return ReqCountRec, CurrentCountRec, nil
	}
	return 0, 0, errors.New("failed_to_get_stockfilemaplog_record")
}

func (db *DB_manager) ReadNormRecords(DepartmentName string, Date string, Sub_category string) ([]Norm, error) {

	if DepartmentName == "SCM Child Part" {
		rows, err := db.Query(fmt.Sprintf("SELECT norm.date,norm.department,part_details.sub_category,norm.part_number,part.description,norm.item_class,norm.amr,norm.norm,norm.created_at,norm.updated_at FROM norm JOIN part ON norm.part_number = part .part_number JOIN part_details ON norm.part_number = part_details .part_number AND norm.department = part_details .department WHERE norm.date = (SELECT MAX(date) FROM norm WHERE norm.department = '%s' ) AND norm.department = '%s' AND part_details.sub_category = '%s'", DepartmentName, DepartmentName, Sub_category))

		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from Norm ", err)
			return nil, err
		}
		var NormArray []Norm
		var NormRec Norm
		for rows.Next() {
			err := rows.Scan(&NormRec.Date, &NormRec.Department, &NormRec.Sub_category, &NormRec.Part_number, &NormRec.Part_description, &NormRec.Item_class, &NormRec.AMR, &NormRec.Norm, &NormRec.Created_At, &NormRec.Updated_At)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			NormArray = append(NormArray, NormRec)
		}
		return NormArray, nil
	} else {
		rows1, err := db.Query(fmt.Sprintf("SELECT norm.date,norm.department,part_details.sub_category,norm.part_number,part.description,norm.item_class,norm.amr,norm.norm,norm.created_at,norm.updated_at FROM norm JOIN part ON norm.part_number = part .part_number JOIN part_details ON norm.part_number = part_details .part_number AND norm.department = part_details .department WHERE norm.date = (SELECT MAX(date) FROM norm WHERE norm.department = '%s' ) AND norm.department = '%s'", DepartmentName, DepartmentName))

		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from Norm ", err)
			return nil, err
		}
		var NormArr []Norm
		var normRec Norm
		for rows1.Next() {
			err := rows1.Scan(&normRec.Date, &normRec.Department, &normRec.Sub_category, &normRec.Part_number, &normRec.Part_description, &normRec.Item_class, &normRec.AMR, &normRec.Norm, &normRec.Created_At, &normRec.Updated_At)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			NormArr = append(NormArr, normRec)
		}

		return NormArr, nil
	}
}

func (db *DB_manager) ReadImportNormRecords(DepartmentName string, Date string) ([]Import_Norm, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT import_norm.date,import_norm.department,part_details.sub_category,import_norm.part_number,part.description,import_norm.item_class,import_norm.amr_1,import_norm.amr_2,import_norm.amr_3,import_norm.norm_1,import_norm.norm_2,import_norm.norm_3,import_norm.created_at,import_norm.updated_at FROM import_norm JOIN part ON import_norm .part_number = part.part_number JOIN part_details ON import_norm.part_number = part_details .part_number AND import_norm.department = part_details .department WHERE import_norm .date = (SELECT MAX(date) FROM import_norm  WHERE import_norm .department = '%s' ) AND import_norm .department = '%s'", DepartmentName, DepartmentName))

	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Import_Norm ", err)
		return nil, err
	}
	var INormArray []Import_Norm
	var INormRec Import_Norm
	for rows.Next() {
		err := rows.Scan(&INormRec.Date, &INormRec.Department, &INormRec.Sub_category, &INormRec.Part_number, &INormRec.Part_description, &INormRec.Item_class, &INormRec.AMR_1, &INormRec.AMR_2, &INormRec.AMR_3, &INormRec.Norm_1, &INormRec.Norm_2, &INormRec.Norm_3, &INormRec.Created_At, &INormRec.Updated_At)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		INormArray = append(INormArray, INormRec)
	}
	return INormArray, nil
}

func (db *DB_manager) ReadStockPartRecords(Date string, DepartmentName string) ([]string, error) {
	stock := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT date,department,part_number,todate_aval,under_inspection,store_stock,vmi_stock,sub_contractor_stock,total_stock FROM stock where date = '" + Date + "' AND department = '" + DepartmentName + "'"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from stockpart ", err)
		return nil, err
	}

	var StockRec Stock
	for rows.Next() {
		err := rows.Scan(&StockRec.Date, &StockRec.Department, &StockRec.Part_number, &StockRec.ToDate_Aval, &StockRec.Under_inspection, &StockRec.Store_stock, &StockRec.VMI_stock, &StockRec.Sub_Contractor_stock, &StockRec.Total_stock)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		stock = append(stock, StockRec.Part_number)
	}
	return stock, nil
}

func (db *DB_manager) ReadStockRecords(Date string, DepartmentName string, PartNumber string) ([]Stock, error) {

	partNumberArray := strings.Split(PartNumber, ",")
	fmt.Println(partNumberArray)
	var i int
	var StockArray []Stock
	for i = 0; i < len(partNumberArray); i++ {
		rows, err := db.Query(fmt.Sprintf("SELECT date,department,part_number,todate_aval,under_inspection,store_stock,vmi_stock,sub_contractor_stock,total_stock FROM stock where date='%s' AND department='%s' AND part_number = '%s'", Date, DepartmentName, partNumberArray[i]))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from Stock ", err)
			return nil, err
		}

		var StockRec Stock
		for rows.Next() {
			err := rows.Scan(&StockRec.Date, &StockRec.Department, &StockRec.Part_number, &StockRec.ToDate_Aval, &StockRec.Under_inspection, &StockRec.Store_stock, &StockRec.VMI_stock, &StockRec.Sub_Contractor_stock, &StockRec.Total_stock)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			StockArray = append(StockArray, StockRec)
		}
	}

	return StockArray, nil
}

func (db *DB_manager) ReadBPRRecord(DepartmentName string, Date string) (int, error) {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM bpr where department='%s' AND date='%s'", DepartmentName, Date)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr ", err)
		fmt.Println("failed_query ", qry)
		return 0, err
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
	return 0, errors.New("failed_to_get_bpr_record")
}

func (db *DB_manager) ReadFull_kit_masterRecords(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,full_kit_master.department,fg_assembly_code,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,item_class FROM full_kit_master JOIN bpr ON full_kit_master.department = bpr.department AND full_kit_master.child_part_code = bpr.part_number JOIN part_details ON part_details.part_number  = full_kit_master.child_part_code AND part_details.department  = full_kit_master.department WHERE  full_kit_master.department = '%s' AND bpr.date = '%s' GROUP BY  bpr.date,full_kit_master.department,fg_assembly_code,item_class", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from FKM ", err)
		return nil, err
	}

	var FKMArray []Full_kit_master
	var FKMRec Full_kit_master
	for rows.Next() {
		err := rows.Scan(&FKMRec.Date, &FKMRec.Department, &FKMRec.FG_assembly_code, &FKMRec.Remark, &FKMRec.Item_class)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		FKMArray = append(FKMArray, FKMRec)
	}
	return FKMArray, nil
}

func (db *DB_manager) ReadSupplierInfo(Date string, DepartmetName string, SupplierName string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,part_supplier_details.sub_category,bpr.part_number,part.Description,lead_time,bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour,part_supplier_details.buyer,part_supplier_details.supplier FROM bpr JOIN part_supplier_details ON bpr.part_number = part_supplier_details.part_number AND bpr.department = part_supplier_details.department JOIN part ON bpr.part_number =  part.part_number LEFT JOIN norm_cd ON bpr.part_number = norm_cd.part_number AND bpr.department = norm_cd.department where bpr.date = '%s' AND bpr.department = '%s' AND part_supplier_details.supplier = '%s' ORDER BY bpr.penetration DESC", Date, DepartmetName, SupplierName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from supplierinfo ", err)
		return Main{}, err
	}

	var SupplierArray []SupplierInfo
	var SupplierRec SupplierInfo

	for rows.Next() {
		err := rows.Scan(&SupplierRec.Date, &SupplierRec.Sub_category, &SupplierRec.Part_number, &SupplierRec.Description, &SupplierRec.Lead_time, &SupplierRec.Norm, &SupplierRec.Stock, &SupplierRec.Gap, &SupplierRec.Penetration, &SupplierRec.Colour, &SupplierRec.Buyer, &SupplierRec.Supplier)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		SupplierArray = append(SupplierArray, SupplierRec)

	}
	var supplierinfo Main
	supplierinfo.Status = "true"
	supplierinfo.Status_code = "200"
	supplierinfo.Count = len(SupplierArray)
	supplierinfo.Supplier = SupplierArray
	return supplierinfo, nil
}

func (db *DB_manager) ReadBuyerwiseTrendInfo(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,part_supplier_details.buyer,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G, SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_supplier_details ON part_supplier_details.part_number=bpr.part_number AND part_supplier_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' GROUP BY bpr.date,bpr.department,part_supplier_details.buyer", Date, DepartmentName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from BuyerwiseTrendInfo ", err)
		return nil, err
	}

	var BuyerArray []BuyerwiseTrend
	var BuyerRec BuyerwiseTrend

	for rows.Next() {
		err := rows.Scan(&BuyerRec.Date, &BuyerRec.Department, &BuyerRec.Buyer, &BuyerRec.Total, &BuyerRec.B, &BuyerRec.R, &BuyerRec.Y, &BuyerRec.G, &BuyerRec.W, &BuyerRec.Percent_of_Red_Items, &BuyerRec.Percnet_of_White_Items)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		BuyerArray = append(BuyerArray, BuyerRec)

	}

	var buyermaininfo Main
	buyermaininfo.Status = "true"
	buyermaininfo.Status_code = "200"
	buyermaininfo.Count = len(BuyerArray)
	buyermaininfo.BuyerwiseTrendInfo = BuyerArray
	return buyermaininfo, nil
}

func (db *DB_manager) ReadBPRSummaryRecords(DepartmentName string, Date string, DisplayParameterID string) (Main, error) {

	if DisplayParameterID == "1" { // Display bpr summary in view-bpr
		rows, err := db.Query(fmt.Sprintf("SELECT department,sub_category,display_order FROM sub_category WHERE department = '%s' ORDER BY 3 ASC", DepartmentName))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from bpr_summary_records", err)
			return Main{}, err
		}

		var departmentinfo BPRSummary
		var BPRSummaryArray []BPRSummary

		for rows.Next() {
			err := rows.Scan(&departmentinfo.DepartmentName, &departmentinfo.Sub_category, &departmentinfo.Display_order)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

			/////////////////////////////////////////////////////////////////////////////////////////////
			var query string
			var query1 string
			if departmentinfo.Sub_category != "Child Parts" {
				query = fmt.Sprintf("SELECT date,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category = '%s' AND part_details.sub_category <> 'Child Parts' GROUP BY date,bpr.department,part_details.sub_category", Date, departmentinfo.DepartmentName, departmentinfo.Sub_category)
			} else {
				query1 = fmt.Sprintf("SELECT date,full_kit_master.line,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department = part_details.department JOIN full_kit_master ON full_kit_master.child_part_code =  bpr.part_number AND  full_kit_master.department = bpr.department WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category <> 'Finished Goods' AND COLOUR IN ('B','R','Y','G','W') GROUP BY date,full_kit_master.line order by full_kit_master.line ASC", Date, departmentinfo.DepartmentName)
			}
			rows1, err := db.Query(query)
			rows2, err := db.Query(query1)
			defer rows.Close()
			if err != nil {
				fmt.Println("failed to get data from bpr_summary_records ", err)
				return Main{}, err
			}
			var BPRsummaryinfo BPRSummary
			var BPRsummaryinfo1 BPRSummary
			for rows1.Next() {
				err := rows1.Scan(&BPRsummaryinfo.Date, &BPRsummaryinfo.Total, &BPRsummaryinfo.B, &BPRsummaryinfo.R, &BPRsummaryinfo.Y, &BPRsummaryinfo.G, &BPRsummaryinfo.W, &BPRsummaryinfo.Percent_of_Red_Items, &BPRsummaryinfo.Percnet_of_White_Items)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

				BPRsummaryinfo.DepartmentName = departmentinfo.DepartmentName
				BPRsummaryinfo.Sub_category = departmentinfo.Sub_category
				BPRsummaryinfo.Display_order = departmentinfo.Display_order
				BPRSummaryArray = append(BPRSummaryArray, BPRsummaryinfo)
			}

			for rows2.Next() {
				err := rows2.Scan(&BPRsummaryinfo1.Date, &BPRsummaryinfo1.Line, &BPRsummaryinfo1.Total, &BPRsummaryinfo1.B, &BPRsummaryinfo1.R, &BPRsummaryinfo1.Y, &BPRsummaryinfo1.G, &BPRsummaryinfo1.W, &BPRsummaryinfo1.Percent_of_Red_Items, &BPRsummaryinfo1.Percnet_of_White_Items)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}
				BPRsummaryinfo1.DepartmentName = departmentinfo.DepartmentName
				BPRsummaryinfo1.Sub_category = BPRsummaryinfo1.Line
				BPRSummaryArray = append(BPRSummaryArray, BPRsummaryinfo1)
			}
		}

		if (departmentinfo.DepartmentName != "Drum Brake Line") && (departmentinfo.DepartmentName != "Disc Brake Line") {

			rows3, err := db.Query(fmt.Sprintf("SELECT date,SUM(Total) AS Total,SUM(B) AS B,SUM(R) AS R,SUM(Y) AS Y,SUM(G) AS G,SUM(W) AS W,SUM(percent_of_red_items) AS percent_of_red_items,SUM(percent_of_white_items) AS percent_of_white_items FROM(SELECT date,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category <> 'Child Parts' GROUP BY date,bpr.department,part_details.sub_category) AS ABC GROUP BY date", Date, DepartmentName))
			defer rows3.Close()
			if err != nil {
				fmt.Println("failed to get data from bpr summary total", err)
				return Main{}, err
			}

			var BPRSummTotalRec BPRSummary
			for rows3.Next() {
				err := rows3.Scan(&BPRSummTotalRec.Date, &BPRSummTotalRec.Total, &BPRSummTotalRec.B, &BPRSummTotalRec.R, &BPRSummTotalRec.Y, &BPRSummTotalRec.G, &BPRSummTotalRec.W, &BPRSummTotalRec.Percent_of_Red_Items, &BPRSummTotalRec.Percnet_of_White_Items)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

			}
			BPRSummTotalRec.Sub_category = "Total"
			if BPRSummTotalRec.B == 0 && BPRSummTotalRec.R == 0 {
				BPRSummTotalRec.NoCriticalPartListFlag = 1
			}
			BPRSummaryArray = append(BPRSummaryArray, BPRSummTotalRec)
		} else {
			rows4, err := db.Query(fmt.Sprintf("SELECT date,SUM(Total) AS Total,SUM(B) AS B,SUM(R) AS R,SUM(Y) AS Y,SUM(G) AS G,SUM(W) AS W,SUM(percent_of_red_items) AS percent_of_red_items,SUM(percent_of_white_items) AS percent_of_white_items FROM(SELECT date,full_kit_master.line,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department = part_details.department JOIN full_kit_master ON full_kit_master.child_part_code =  bpr.part_number AND  full_kit_master.department = bpr.department WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category <> 'Finished Goods' AND COLOUR IN ('B','R','Y','G','W') GROUP BY date,full_kit_master.line) AS ABC GROUP BY date", Date, DepartmentName))
			defer rows4.Close()
			if err != nil {
				fmt.Println("failed to get data from bpr summary child part total", err)
				return Main{}, err
			}

			var BPRSummTotalRec1 BPRSummary
			for rows4.Next() {
				err := rows4.Scan(&BPRSummTotalRec1.Date, &BPRSummTotalRec1.Total, &BPRSummTotalRec1.B, &BPRSummTotalRec1.R, &BPRSummTotalRec1.Y, &BPRSummTotalRec1.G, &BPRSummTotalRec1.W, &BPRSummTotalRec1.Percent_of_Red_Items, &BPRSummTotalRec1.Percnet_of_White_Items)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}
			}

			BPRSummTotalRec1.Sub_category = "Child Part Total"
			BPRSummaryArray = append(BPRSummaryArray, BPRSummTotalRec1)
		}

		var BPRSmaininfo Main
		BPRSmaininfo.Status = "true"
		BPRSmaininfo.Status_code = "200"
		BPRSmaininfo.Count = len(BPRSummaryArray)
		BPRSmaininfo.BPRSummaryInfo = BPRSummaryArray
		return BPRSmaininfo, nil
	}
	if DisplayParameterID == "2" { // Display bpr summary in FBML dashboard.
		rows, err := db.Query(fmt.Sprintf("SELECT department,sub_category,display_order FROM sub_category WHERE department = '%s' ORDER BY 3 ASC", DepartmentName))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from bpr_summary_records", err)
			return Main{}, err
		}

		var departmentinfo1 BPRSummary
		var BPRSummaryArray1 []BPRSummary

		for rows.Next() {
			err := rows.Scan(&departmentinfo1.DepartmentName, &departmentinfo1.Sub_category, &departmentinfo1.Display_order)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

			/////////////////////////////////////////////////////////////////////////////////////////////
			rows1, err := db.Query(fmt.Sprintf("SELECT date,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category = '%s' AND part_details.sub_category <> 'Child Parts' GROUP BY date,bpr.department,part_details.sub_category", Date, departmentinfo1.DepartmentName, departmentinfo1.Sub_category))
			defer rows1.Close()
			if err != nil {
				fmt.Println("failed to get data from bpr summary", err)
				return Main{}, err
			}

			var BPRsummaryinfo2 BPRSummary
			for rows1.Next() {
				err := rows1.Scan(&BPRsummaryinfo2.Date, &BPRsummaryinfo2.Total, &BPRsummaryinfo2.B, &BPRsummaryinfo2.R, &BPRsummaryinfo2.Y, &BPRsummaryinfo2.G, &BPRsummaryinfo2.W, &BPRsummaryinfo2.Percent_of_Red_Items, &BPRsummaryinfo2.Percnet_of_White_Items)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

				BPRsummaryinfo2.DepartmentName = departmentinfo1.DepartmentName
				BPRsummaryinfo2.Sub_category = departmentinfo1.Sub_category
				BPRsummaryinfo2.Display_order = departmentinfo1.Display_order
				BPRSummaryArray1 = append(BPRSummaryArray1, BPRsummaryinfo2)
			}
		}

		rows2, err := db.Query(fmt.Sprintf("SELECT date,SUM(Total) AS Total,SUM(B) AS B,SUM(R) AS R,SUM(Y) AS Y,SUM(G) AS G,SUM(W) AS W,SUM(percent_of_red_items) AS percent_of_red_items,SUM(percent_of_white_items) AS percent_of_white_items FROM(SELECT date,full_kit_master.line,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department = part_details.department JOIN full_kit_master ON full_kit_master.child_part_code =  bpr.part_number AND  full_kit_master.department = bpr.department WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category <> 'Finished Goods' AND COLOUR IN ('B','R','Y','G','W') GROUP BY date,full_kit_master.line) AS ABC GROUP BY date", Date, DepartmentName))
		defer rows2.Close()
		if err != nil {
			fmt.Println("failed to get data from bpr summary child part total", err)
			return Main{}, err
		}

		var BPRSummTotalRec2 BPRSummary
		for rows2.Next() {
			err := rows2.Scan(&BPRSummTotalRec2.Date, &BPRSummTotalRec2.Total, &BPRSummTotalRec2.B, &BPRSummTotalRec2.R, &BPRSummTotalRec2.Y, &BPRSummTotalRec2.G, &BPRSummTotalRec2.W, &BPRSummTotalRec2.Percent_of_Red_Items, &BPRSummTotalRec2.Percnet_of_White_Items)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
		}
		BPRSummTotalRec2.Sub_category = "Child Part Total"
		BPRSummaryArray1 = append(BPRSummaryArray1, BPRSummTotalRec2)

		var BPRSmaininfo1 Main
		BPRSmaininfo1.Status = "true"
		BPRSmaininfo1.Status_code = "200"
		BPRSmaininfo1.Count = len(BPRSummaryArray1)
		BPRSmaininfo1.BPRSummaryInfo = BPRSummaryArray1
		return BPRSmaininfo1, nil
	}
	return Main{}, nil
}

func (db *DB_manager) ReadFull_kit_masterRecord(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("Select fg_assembly_code,fg_assembly_description,line AS Sub_category,COALESCE(G,0)+ COALESCE(R,0)+ COALESCE(W,0)+COALESCE(Y,0) Total,COALESCE(G,0)  G, COALESCE(R,0)  R,COALESCE(W,0)  W,COALESCE(Y,0)  Y FROM crosstab('SELECT fg_assembly_code,fg_assembly_description,line,item_class,SUM(CASE WHEN colour in( ''R'',''B'') THEN 1 ELSE 0 END) AS RB FROM full_kit_master JOIN bpr ON full_kit_master.department = bpr.department AND full_kit_master.child_part_code = bpr.part_number JOIN part_details ON part_details.part_number  = full_kit_master.child_part_code AND part_details.department  = full_kit_master.department WHERE  full_kit_master.department = ''%s'' AND bpr.date = ''%s'' GROUP BY fg_assembly_code,fg_assembly_description,line,item_class order by 1,2','Select distinct item_class from part_details Where item_class in (''G'',''R'',''W'',''Y'') ORDER BY 1 ASC') AS FinalResult(fg_assembly_code TEXT,fg_assembly_description TEXT,line TEXT,G NUMERIC, R NUMERIC,W NUMERIC,Y NUMERIC)", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from FKM ", err)
		return nil, err
	}

	var FKMArray1 []Full_kit
	var FKMRec1 Full_kit
	for rows.Next() {
		err := rows.Scan(&FKMRec1.FG_assembly_code, &FKMRec1.FG_assembly_description, &FKMRec1.Sub_category, &FKMRec1.Total, &FKMRec1.G, &FKMRec1.R, &FKMRec1.W, &FKMRec1.Y)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		FKMArray1 = append(FKMArray1, FKMRec1)
	}
	return FKMArray1, nil
}

func (db *DB_manager) ReadSupplierNameRecords(DepartmentName string) ([]string, error) {
	supplier := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct supplier FROM part_supplier_details where department = '" + DepartmentName + "'"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_supplier_details ", err)
		return nil, err
	}

	var SuRec SupplierName
	for rows.Next() {
		err := rows.Scan(&SuRec.Supplier)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		supplier = append(supplier, SuRec.Supplier)
	}
	return supplier, nil
}

func (db *DB_manager) ReadBPRLineNamesRecords(DepartmentName string) ([]string, error) {
	bpr_line := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct line FROM full_kit_master where department = '" + DepartmentName + "' order by line ASC"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from full_kit_master line wise records ", err)
		return nil, err
	}

	var BLRec BPRLinewise
	for rows.Next() {
		err := rows.Scan(&BLRec.Line)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		bpr_line = append(bpr_line, BLRec.Line)
	}
	return bpr_line, nil
}

func (db *DB_manager) ReadBPRColourTrendRecords(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT A.date from (SELECT distinct bpr.date from bpr WHERE bpr.department = '%s' AND bpr.date <= '%s' order by date DESC limit 7) A Order by date ASC", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr  ", err)
		return nil, err
	}
	var dateArray []string
	var dateRec string
	for rows.Next() {
		err := rows.Scan(&dateRec)
		if err != nil {
			fmt.Println("Location: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		dateArray = append(dateArray, dateRec)
	}

	////////////////////////////////////////////////
	var daRecArray []ColourTrendInfo
	var daRec ColourTrendInfo
	rows1, err := db.Query(fmt.Sprintf("SELECT DISTINCT bpr.department,part_details.sub_category,bpr.part_number,part.description from bpr JOIN part_details ON bpr.part_number = part_details.part_number JOIN part ON bpr.part_number = part.part_number AND bpr.department = part_details.department WHERE bpr.department = '%s'", DepartmentName))
	defer rows1.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr  ", err)
		return nil, err
	}

	for rows1.Next() {
		err := rows1.Scan(&daRec.Department, &daRec.Sub_category, &daRec.Part_number, &daRec.Description)
		if err != nil {
			fmt.Println("Location: failed to scan the record.. continue with the next.. ", err)
			continue
		}
		//////////////////////////////////////////////////////////////////
		var BPRCTArray1 []BPRColourTrendInfo

		for i := 0; i < len(dateArray); i++ {
			var BPRCTrendinfo BPRColourTrendInfo

			rows2, err := db.Query(fmt.Sprintf("SELECT  round(bpr.penetration) AS penetration,bpr.colour FROM bpr WHERE bpr.date = '%s' AND bpr.department = '%s' AND bpr.part_number = '%s' GROUP BY bpr.date,bpr.penetration,bpr.colour ORDER BY bpr.date", dateArray[i], daRec.Department, daRec.Part_number))
			//var BPRCTrendinfo BPRColourTrendInfo
			defer rows2.Close()
			BPRCTrendinfo.Date = dateArray[i]
			if err != nil {

				BPRCTrendinfo.Penetration = 0
				BPRCTrendinfo.Colour = ""
				fmt.Println("failed to get data from bpr_colour_trend_records ", err)
				//	return nil, err
			}

			for rows2.Next() {
				err = rows2.Scan(&BPRCTrendinfo.Penetration, &BPRCTrendinfo.Colour)
				if err != nil {
					fmt.Println("failed to scan the record.. continue with the next.. ", err)
					continue
				}

			}
			BPRCTArray1 = append(BPRCTArray1, BPRCTrendinfo)

		}
		daRec.BPRColourTrend = BPRCTArray1
		daRecArray = append(daRecArray, daRec)

	}

	var BPRCTrootinfo Main
	BPRCTrootinfo.Status = "true"
	BPRCTrootinfo.Status_code = "200"
	BPRCTrootinfo.Count = len(daRecArray)
	BPRCTrootinfo.ColourTrend = daRecArray
	return BPRCTrootinfo, nil
}

func (db *DB_manager) ReadFullkitPartwiseRecords(FG_assembly_code string, DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT full_kit_master.fg_assembly_code,bpr.part_number,part.description,part_details.Item_class,bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour FROM bpr JOIN part ON bpr.part_number = part.part_number JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department = part_details.department JOIN full_kit_master ON bpr.department = full_kit_master.department AND bpr.part_number = full_kit_master.child_part_code WHERE fg_assembly_code = '%s' AND bpr.department = '%s' AND bpr.date = '%s' AND colour IN ('B','R') ORDER BY bpr.penetration DESC", FG_assembly_code, DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr ", err)
		return nil, err
	}

	var FKArray []FullKitPartWiseInfo
	var FKRec FullKitPartWiseInfo

	for rows.Next() {
		err := rows.Scan(&FKRec.FG_assembly_code, &FKRec.Part_number, &FKRec.Description, &FKRec.Item_class, &FKRec.Norm, &FKRec.Stock, &FKRec.Gap, &FKRec.Penetration, &FKRec.Colour)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		FKArray = append(FKArray, FKRec)

	}

	var fkmmaininfo Main
	fkmmaininfo.Status = "true"
	fkmmaininfo.Status_code = "200"
	fkmmaininfo.Count = len(FKArray)
	fkmmaininfo.FullKitPartWise = FKArray
	return fkmmaininfo, nil
}

func (db *DB_manager) ReadSupplierName(DepartmentName string) ([]string, error) {
	supplierName := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct supplier FROM part_supplier_details where department = '" + DepartmentName + "'"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_supplier_details ", err)
		return nil, err
	}

	var SuRec1 SupplierName
	for rows.Next() {
		err := rows.Scan(&SuRec1.Supplier)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		supplierName = append(supplierName, SuRec1.Supplier)
	}
	return supplierName, nil
}

func (db *DB_manager) ReadSupplierSummaryRecords(Date string, DepartmentName string, SupplierName string) (Main, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,part_supplier_details.supplier AS Vendor_Name,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G, SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_supplier_details ON bpr.part_number = part_supplier_details.part_number AND bpr.department = part_supplier_details.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_supplier_details.supplier = '%s' GROUP BY bpr.date,bpr.department,part_supplier_details.supplier", Date, DepartmentName, SupplierName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Supplierwise bpr Summary", err)
		return Main{}, err
	}

	var SupplierSummaryArray []SupplierSummary
	var SupplierSummaryRec SupplierSummary

	for rows.Next() {
		err := rows.Scan(&SupplierSummaryRec.Date, &SupplierSummaryRec.DepartmentName, &SupplierSummaryRec.VendorName, &SupplierSummaryRec.Total, &SupplierSummaryRec.B, &SupplierSummaryRec.R, &SupplierSummaryRec.Y, &SupplierSummaryRec.G, &SupplierSummaryRec.W, &SupplierSummaryRec.Percent_of_Red_Items, &SupplierSummaryRec.Percnet_of_White_Items)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		SupplierSummaryArray = append(SupplierSummaryArray, SupplierSummaryRec)

	}

	var summarymaininfo Main
	summarymaininfo.Status = "true"
	summarymaininfo.Status_code = "200"
	summarymaininfo.Count = len(SupplierSummaryArray)
	summarymaininfo.SupplierSummaryInfo = SupplierSummaryArray
	return summarymaininfo, nil
}

func (db *DB_manager) ReadBPRSubCategoryRecords(DepartmentName string) ([]string, error) {
	bpr_subcategory := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct sub_category,display_order FROM sub_category WHERE department = '" + DepartmentName + "' ORDER BY 2 ASC"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from sub category records ", err)
		return nil, err
	}

	var BSRec Sub_category
	for rows.Next() {
		err := rows.Scan(&BSRec.Sub_category, &BSRec.Display_Order)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		bpr_subcategory = append(bpr_subcategory, BSRec.Sub_category)
	}
	return bpr_subcategory, nil
}

func (db *DB_manager) ReadStockFileMapCountRecord(Date string, DepartmentName string, FileName string) (int, error) {

	qry := fmt.Sprintf("SELECT DISTINCT count(*) FROM stock_file_log where date='%s' AND department = '%s' AND file_name = '%s'", Date, DepartmentName, FileName)

	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from stock file map", err)
		fmt.Println("failed_query ", qry)
		return 0, err
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
	return 0, errors.New("failed_to_get_bpr_record")
}
