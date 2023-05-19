package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
)

type root struct {
	Status      string
	Status_code string
	Count       int
	Department  []DepartmentInfo
	Inventory   []Inventoryinfo
	Location    []LocationwiseBPRinfo
}

type Part_supplier_record struct {
	Department   string
	Buyer        string
	Supplier     string
	CountofRed   int
	CountofBlack int
}

type Supplier struct {
	Buyer        string
	Supplier     string
	Display_name string
	Bar_chart    []string
	R            int
	B            int
	Total        int
}

type DepartmentInfo struct {
	DepartmentID   int
	DepartmentName string
	Sub_category   string
	Type           int
	Order1         int
	Order2         int
	SupplierInfo   []Supplier
	BPRTrendInfo   []BPRTrend
}

type BPRTrend struct {
	Date      string
	Bar_Trend []string
	B         int
	R         int
	W         int
}

type LocationwiseBPRinfo struct {
	Date           string
	DepartmentName string
	Sub_category   string
	Type           int
	Order1         int
	Order2         int
	B              int
	R              int
	Y              int
	G              int
	W              int
	Total          int
}

type Inventoryinfo struct {
	Date              string
	DepartmentName    string
	Sub_category      string
	Order1            int
	Order2            int
	Type              int
	Desired_inventory string
	Actual_inventory  string
	White_inventory   string
	Colour            string
}

func (db *DB_manager) ReadPartSupplierRecords1(DepartmentName string, Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT  part_supplier_details.department,buyer,part_supplier_details.supplier,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B FROM part_supplier_details JOIN bpr ON part_supplier_details.department = bpr.department AND part_supplier_details.part_number = bpr.part_number WHERE part_supplier_details.department = '%s' AND bpr.date = '%s' AND bpr.colour IN ('R','B') GROUP BY part_supplier_details.department,buyer,part_supplier_details.supplier", DepartmentName, Date))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_supplier_details ", err)
		return nil, err
	}

	var PSDArray []Part_supplier_record
	var PSDRec Part_supplier_record
	for rows.Next() {
		err := rows.Scan(&PSDRec.Department, &PSDRec.Buyer, &PSDRec.Supplier, &PSDRec.CountofRed, &PSDRec.CountofBlack)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		PSDArray = append(PSDArray, PSDRec)
	}
	return PSDArray, nil
}

func (db *DB_manager) ReadPartSupplierRecords(Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT department,sub_category FROM (SELECT distinct 2 A,department.department,'' Sub_Category  FROM department JOIN Part_details ON Part_details.department=department.department  WHERE department.department = 'Manesar Child Part'  UNION SELECT distinct 1 A,department.department,Part_details.Sub_Category  FROM department JOIN Part_details ON Part_details.department=department.department WHERE department.department = 'SCM Child Part' AND Part_details.Sub_Category in ('Disc Brake','Drum Brake')) Main ORDER BY A"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from part_supplier_details ", err)
		return nil, err
	}

	var PSDArray []DepartmentInfo
	var departmentinfo DepartmentInfo
	for rows.Next() {
		err := rows.Scan(&departmentinfo.DepartmentName, &departmentinfo.Sub_category)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		/////////////////////////////////////////////////////////////////////////////////////////////
		var query string
		if departmentinfo.DepartmentName != "Manesar Child Part" {
			query = fmt.Sprintf("Select Buyer, Supplier,coalesce(display_name,'') AS display_name, R, B, Total from (SELECT Buyer,supplier,coalesce(display_name,'') AS display_name,SUM(R) R,SUM(B)B, SUM(R) + SUM(B) Total FROM (SELECT SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B ,colour,bpr.part_number,bpr.department,buyer,supplier,display_name,part_supplier_details.sub_category FROM bpr LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department JOIN part_supplier_details ON part_supplier_details.part_number=bpr.part_number AND part_supplier_details.department=bpr.department WHERE date='%s' AND bpr.department='%s' AND COLOUR IN ('B','R') AND part_details.sub_category='%s' GROUP BY Colour,bpr.part_number,buyer,supplier ,display_name,bpr.department,part_supplier_details.sub_category) ABC GROUP BY Buyer,supplier,display_name ORDER BY Total DESC) FinalResult ORDER BY TOTAL DESC,display_name LIMIT 15", Date, departmentinfo.DepartmentName, departmentinfo.Sub_category)
		} else {
			query = fmt.Sprintf("Select Buyer, Supplier,coalesce(display_name,'') AS display_name, R, B, Total from (SELECT Buyer,supplier,coalesce(display_name,'') AS display_name,SUM(R) R,SUM(B)B, SUM(R) + SUM(B) Total FROM (SELECT SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B ,colour,bpr.part_number,bpr.department,buyer,supplier,display_name,part_supplier_details.sub_category FROM bpr JOIN part_supplier_details ON part_supplier_details.part_number=bpr.part_number AND part_supplier_details.department=bpr.department WHERE date='%s' AND bpr.department='%s' AND COLOUR IN ('B','R') GROUP BY Colour,bpr.part_number,buyer,supplier ,display_name,bpr.department,part_supplier_details.sub_category) ABC GROUP BY Buyer,supplier,display_name ORDER BY Total DESC) FinalResult ORDER BY TOTAL DESC,display_name LIMIT 15", Date, departmentinfo.DepartmentName)
		}
		rows1, err := db.Query(query)
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from part_supplier_details ", err)
			return nil, err
		}
		var PSDArray1 []Supplier
		var supplierinfo Supplier
		for rows1.Next() {
			err = rows1.Scan(&supplierinfo.Buyer, &supplierinfo.Supplier, &supplierinfo.Display_name, &supplierinfo.R, &supplierinfo.B, &supplierinfo.Total)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			//supplierinfo.Bar_chart=append(PSDArray1,supplierinfo)
			PSDArray1 = append(PSDArray1, supplierinfo)
		}
		////////////////////////////////////////////////////////////////////////////////////////////
		departmentinfo.SupplierInfo = PSDArray1
		PSDArray = append(PSDArray, departmentinfo)
	}

	var PSDrootinfo root
	PSDrootinfo.Status = "true"
	PSDrootinfo.Status_code = "200"
	PSDrootinfo.Count = len(PSDArray)
	PSDrootinfo.Department = PSDArray
	return PSDrootinfo, nil
}

func (db *DB_manager) ReadBPRTrendRecords(Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT * from (SELECT distinct department.department,Part_details.Sub_Category,department.display_order AS ORDER1,sub_category.display_order AS ORDER2, 2 AS Type FROM department JOIN Part_details ON Part_details.department=department.department LEFT JOIN sub_category ON part_details.Sub_category = sub_category.sub_category AND sub_category.department=department.department WHERE department.department IN ('SCM Child Part','Manesar Child Part') UNION SELECT distinct 'Finished Goods' AS department, department.department AS Sub_Category,3 AS ORDER1,department.display_order AS ORDER2, 3 AS Type FROM department JOIN Part_details ON Part_details.department=department.department  WHERE department.department in  ('Drum Brake Line','Disc Brake Line','Foundry') UNION SELECT distinct department.department,'' Sub_Category ,department.display_order AS ORDER1,1 AS ORDER2, 1 AS Type FROM department JOIN Part_details ON Part_details.department=department.department WHERE department.department in  ('Import')) MAIN ORDER BY 3,4"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from bpr_trend_records", err)
		return nil, err
	}

	var BPRTArray []DepartmentInfo
	var departmentinfo DepartmentInfo
	for rows.Next() {
		err := rows.Scan(&departmentinfo.DepartmentName, &departmentinfo.Sub_category, &departmentinfo.Order1, &departmentinfo.Order2, &departmentinfo.Type)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		/////////////////////////////////////////////////////////////////////////////////////////////

		var query string
		if departmentinfo.Type == 2 {
			query = fmt.Sprintf("SELECT to_char((to_date(bpr.date, 'yyyy-mm-dd')),'DD-MM'),SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department =bpr.department AND COLOUR IN ('B','R','W') WHERE to_date(bpr.date, 'yyyy-mm-dd') > (to_date('%s','yyyy-mm-dd')  - 7) AND to_date(bpr.date, 'yyyy-mm-dd') <= (to_date('%s','yyyy-mm-dd') - 0) AND bpr.department = '%s' AND part_details.sub_category = '%s' GROUP BY bpr.date ORDER BY bpr.date ", Date, Date, departmentinfo.DepartmentName, departmentinfo.Sub_category)
		} else if departmentinfo.Type == 1 {
			query = fmt.Sprintf("SELECT to_char((to_date(bpr.date, 'yyyy-mm-dd')),'DD-MM'),SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department =bpr.department AND COLOUR IN ('B','R','W') WHERE to_date(bpr.date, 'yyyy-mm-dd') > (to_date('%s','yyyy-mm-dd')  - 7) AND to_date(bpr.date, 'yyyy-mm-dd') <= (to_date('%s','yyyy-mm-dd') - 0) AND bpr.department = '%s' GROUP BY bpr.date ORDER BY bpr.date ", Date, Date, departmentinfo.DepartmentName)
		} else if departmentinfo.Type == 3 {
			query = fmt.Sprintf("SELECT to_char((to_date(bpr.date, 'yyyy-mm-dd')),'DD-MM'),SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department =bpr.department AND COLOUR IN ('B','R','W') WHERE to_date(bpr.date, 'yyyy-mm-dd') > (to_date('%s','yyyy-mm-dd')  - 7) AND to_date(bpr.date, 'yyyy-mm-dd') <= (to_date('%s','yyyy-mm-dd') - 0) AND bpr.department = '%s' GROUP BY bpr.date ORDER BY bpr.date ", Date, Date, departmentinfo.Sub_category)
		}
		rows1, err := db.Query(query)
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from bpr_trend_records  ", err)
			return nil, err
		}
		var BPRTArray1 []BPRTrend
		var BPRTRec BPRTrend
		for rows1.Next() {
			err := rows1.Scan(&BPRTRec.Date, &BPRTRec.B, &BPRTRec.R, &BPRTRec.W)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			BPRTArray1 = append(BPRTArray1, BPRTRec)
		}
		departmentinfo.BPRTrendInfo = BPRTArray1
		BPRTArray = append(BPRTArray, departmentinfo)
	}
	var pqrrootinfo root
	pqrrootinfo.Status = "true"
	pqrrootinfo.Status_code = "200"
	pqrrootinfo.Count = len(BPRTArray)
	pqrrootinfo.Department = BPRTArray
	return pqrrootinfo, nil
}

func (db *DB_manager) ReadLocationwiseBPRRecords(Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT * from (SELECT distinct department.department,Part_details.Sub_Category,department.display_order AS ORDER1,sub_category.display_order AS ORDER2, 2 AS Type FROM department JOIN Part_details ON Part_details.department=department.department LEFT JOIN sub_category ON part_details.Sub_category = sub_category.sub_category AND sub_category.department=department.department WHERE department.department IN ('Warehouse','SCM Child Part','Manesar Child Part') UNION SELECT distinct 'Finished Goods' AS department, department.department AS Sub_Category,3 AS ORDER1,department.display_order AS ORDER2, 3 AS Type FROM department JOIN Part_details ON Part_details.department=department.department  WHERE department.department in  ('Drum Brake Line','Disc Brake Line','Foundry') UNION SELECT distinct department.department,'' Sub_Category ,department.display_order AS ORDER1,1 AS ORDER2, 1 AS Type FROM department JOIN Part_details ON Part_details.department=department.department WHERE department.department in  ('Import')) MAIN ORDER BY 3,4"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from location_wise_bpr ", err)
		return nil, err
	}

	var departmentinfo LocationwiseBPRinfo
	var LocationArray1 []LocationwiseBPRinfo
	for rows.Next() {
		err := rows.Scan(&departmentinfo.DepartmentName, &departmentinfo.Sub_category, &departmentinfo.Order1, &departmentinfo.Order2, &departmentinfo.Type)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		/////////////////////////////////////////////////////////////////////////////////////////////
		var query string
		if departmentinfo.Type == 2 {
			query = fmt.Sprintf("SELECT date,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G, SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category = '%s' GROUP BY date,bpr.department,part_details.sub_category ", Date, departmentinfo.DepartmentName, departmentinfo.Sub_category)
		} else if departmentinfo.Type == 1 {
			query = fmt.Sprintf("SELECT date,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G, SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' GROUP BY date,bpr.department", Date, departmentinfo.DepartmentName)
		} else if departmentinfo.Type == 3 {
			query = fmt.Sprintf("SELECT date,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G, SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category != 'Child Parts' GROUP BY date,bpr.department", Date, departmentinfo.Sub_category)
		}
		rows1, err := db.Query(query)
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from location_wise_bpr  ", err)
			return nil, err
		}

		var BPRLRec LocationwiseBPRinfo
		for rows1.Next() { // hopefully we are lucky enough to find only 1 bpr record.
			err := rows1.Scan(&BPRLRec.Date, &BPRLRec.B, &BPRLRec.R, &BPRLRec.Y, &BPRLRec.G, &BPRLRec.W, &BPRLRec.Total)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			BPRLRec.DepartmentName = departmentinfo.DepartmentName
			BPRLRec.Sub_category = departmentinfo.Sub_category
			BPRLRec.Type = departmentinfo.Type
			BPRLRec.Order1 = departmentinfo.Order1
			BPRLRec.Order2 = departmentinfo.Order2
			LocationArray1 = append(LocationArray1, BPRLRec)
		}

	}
	var PSDrootinfo root
	PSDrootinfo.Status = "true"
	PSDrootinfo.Status_code = "200"
	PSDrootinfo.Count = len(LocationArray1)
	PSDrootinfo.Location = LocationArray1
	return PSDrootinfo, nil
}

func (db *DB_manager) ReadInventoryRecords(Date string) (interface{}, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT * from (SELECT distinct department.department,Part_details.Sub_Category,department.display_order AS ORDER1,sub_category.display_order AS ORDER2, 2 AS Type FROM department JOIN Part_details ON Part_details.department=department.department LEFT JOIN sub_category ON part_details.Sub_category = sub_category.sub_category AND sub_category.department=department.department WHERE department.department IN ('Warehouse') UNION SELECT distinct 'Finished Goods' AS department, department.department AS Sub_Category,3 AS ORDER1,department.display_order AS ORDER2, 3 AS Type FROM department JOIN Part_details ON Part_details.department=department.department  WHERE department.department in  ('Drum Brake Line','Disc Brake Line','Foundry') UNION SELECT distinct department.department,'' Sub_Category ,department.display_order AS ORDER1,1 AS ORDER2, 1 AS Type FROM department JOIN Part_details ON Part_details.department=department.department WHERE department.department in  ('SCM Child Part','Manesar Child Part','Import')) MAIN ORDER BY 3,4"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from inventory", err)
		return nil, err
	}

	var departmentinfo Inventoryinfo
	var InventoryArray1 []Inventoryinfo
	for rows.Next() {
		err := rows.Scan(&departmentinfo.DepartmentName, &departmentinfo.Sub_category, &departmentinfo.Order1, &departmentinfo.Order2, &departmentinfo.Type)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		/////////////////////////////////////////////////////////////////////////////////////////////
		var query1 string
		if departmentinfo.Type == 2 {
			query1 = fmt.Sprintf("SELECT bpr.date,round((sum(bpr.desired_inventory)/1000000 ) :: numeric ,2) AS ToG,round((sum(bpr.actual_inventory) / 1000000) :: numeric,2) AS Actual,round((sum(bpr.white_inventory) / 1000000) :: numeric,2) AS White from bpr JOIN part_details ON part_details.department = bpr.department AND part_details.part_number = bpr.part_number WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category = '%s' GROUP BY bpr.date", Date, departmentinfo.DepartmentName, departmentinfo.Sub_category)
		} else if departmentinfo.Type == 1 {
			query1 = fmt.Sprintf("SELECT bpr.date,round((sum(bpr.desired_inventory)/1000000 ) :: numeric ,2) AS ToG,round((sum(bpr.actual_inventory) / 1000000) :: numeric,2) AS Actual,round((sum(bpr.white_inventory) / 1000000) :: numeric,2) AS White  from bpr JOIN part_details ON part_details.department = bpr.department AND part_details.part_number = bpr.part_number WHERE bpr.date = '%s' AND bpr.department = '%s' GROUP BY bpr.date", Date, departmentinfo.DepartmentName)
		} else if departmentinfo.Type == 3 {
			query1 = fmt.Sprintf("SELECT bpr.date,round((sum(bpr.desired_inventory)/1000000 ) :: numeric ,2) AS ToG,round((sum(bpr.actual_inventory) / 1000000) :: numeric,2) AS Actual,round((sum(bpr.white_inventory) / 1000000) :: numeric,2) AS White from bpr JOIN part_details ON part_details.department = bpr.department AND part_details.part_number = bpr.part_number WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.sub_category != 'Child Parts' GROUP BY bpr.date", Date, departmentinfo.Sub_category)
		}
		rows1, err := db.Query(query1)
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from part_supplier_details ", err)
			return nil, err
		}

		var InventoryRec Inventoryinfo
		var desired_inventory float64
		var actual_inventory float64
		for rows1.Next() {
			err := rows1.Scan(&InventoryRec.Date, &InventoryRec.Desired_inventory, &InventoryRec.Actual_inventory, &InventoryRec.White_inventory)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			InventoryRec.DepartmentName = departmentinfo.DepartmentName
			InventoryRec.Sub_category = departmentinfo.Sub_category
			InventoryRec.Type = departmentinfo.Type
			InventoryRec.Order1 = departmentinfo.Order1
			InventoryRec.Order2 = departmentinfo.Order2

			desired_inventory, _ = strconv.ParseFloat(InventoryRec.Desired_inventory, 64)
			actual_inventory, _ = strconv.ParseFloat(InventoryRec.Actual_inventory, 64)
			// Conditions add
			if desired_inventory >= actual_inventory {
				InventoryRec.Colour = "Green"
			} else if desired_inventory < actual_inventory {
				InventoryRec.Colour = "Red"
			}

			//	InventoryRec.Desired_inventory =  InventoryRec.Desired_inventory + "M"
			//	InventoryRec.Actual_inventory = InventoryRec.Actual_inventory + "M"
			//	InventoryRec.White_inventory =  InventoryRec.White_inventory + "M"

			InventoryArray1 = append(InventoryArray1, InventoryRec)
		}
	}
	var abcrootinfo root
	abcrootinfo.Status = "true"
	abcrootinfo.Status_code = "200"
	abcrootinfo.Count = len(InventoryArray1)
	abcrootinfo.Inventory = InventoryArray1
	return abcrootinfo, nil
}
