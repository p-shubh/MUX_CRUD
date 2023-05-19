package db

import "fmt"

import (
	_ "github.com/lib/pq"
)

type CriticalSupplierPartListCount struct {
	Status                       string
	Status_code                  string
	Count                        int
	CriticalSupplierPartListInfo []CriticalSupplierPartList
}

type CriticalSupplierPartList struct {
	Date                     string
	Department               string
	Part_number              string
	Description              *string `json:"Description,omitempty"`
	Sub_category             string
	Norm                     float64
	Stock                    float64
	Gap                      float64
	Penetration              float64
	Total_coverage_in_months float64
	Colour                   string
}

func (db *DB_manager) ReadItemClassWiseBPRSummaryRecords(DepartmentName string, Date string) (interface{}, error) {

	var BPR_Total_ZeroCount float64
	rows, err := db.Query(fmt.Sprintf("SELECT department,item_class FROM item_class WHERE department = '%s' AND item_class <> 'FG' ORDER BY item_classid ASC", DepartmentName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from item class wise bpr summary records", err)
		return nil, err
	}

	var itemclassinfo BPRSummary
	var BPRSummaryitemclassArray []BPRSummary
	var j = 0
	for rows.Next() {
		err := rows.Scan(&itemclassinfo.DepartmentName, &itemclassinfo.Item_class)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		/////////////////////////////////////////////////////////////////////////////////////////////
		rows1, err := db.Query(fmt.Sprintf("SELECT date,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND part_details.item_class = '%s' GROUP BY date,bpr.department,part_details.item_class", Date, itemclassinfo.DepartmentName, itemclassinfo.Item_class))
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from item class wise bpr summary", err)
			return nil, err
		}

		var BPRSummICRec BPRSummary

		for rows1.Next() {
			j++
			err := rows1.Scan(&BPRSummICRec.Date, &BPRSummICRec.Total, &BPRSummICRec.B, &BPRSummICRec.R, &BPRSummICRec.Y, &BPRSummICRec.G, &BPRSummICRec.W, &BPRSummICRec.Percent_of_Red_Items, &BPRSummICRec.Percnet_of_White_Items)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}

		}

		BPRSummICRec.Date = Date
		BPRSummICRec.DepartmentName = itemclassinfo.DepartmentName
		BPRSummICRec.Item_class = itemclassinfo.Item_class
		BPRSummaryitemclassArray = append(BPRSummaryitemclassArray, BPRSummICRec)
	}

	rows2, err := db.Query(fmt.Sprintf("SELECT date,SUM(Total) AS Total,SUM(B) AS B,SUM(R) AS R,SUM(Y) AS Y,SUM(G) AS G,SUM(W) AS W,SUM(percent_of_red_items) AS percent_of_red_items,SUM(percent_of_white_items) AS percent_of_white_items FROM(SELECT date,SUM(CASE WHEN colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS Total,SUM(CASE WHEN colour = 'B' THEN 1 ELSE 0 END) AS B,SUM(CASE WHEN colour = 'R' THEN 1 ELSE 0 END) AS R,SUM(CASE WHEN colour = 'Y' THEN 1 ELSE 0 END) AS Y,SUM(CASE WHEN colour = 'G' THEN 1 ELSE 0 END) AS G,SUM(CASE WHEN colour = 'W' THEN 1 ELSE 0 END) AS W,SUM(CASE WHEN bpr.colour IN ('R')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_red_items,SUM(CASE WHEN bpr.colour IN ('W')  THEN 1 ELSE 0 END) * 100 /SUM(CASE WHEN bpr.colour IN ('B','R','Y','G','W')  THEN 1 ELSE 0 END) AS percent_of_white_items FROM bpr JOIN part_details ON part_details.part_number=bpr.part_number AND part_details.department=bpr.department AND COLOUR IN ('B','R','Y','G','W') WHERE bpr.date = '%s' AND bpr.department = '%s' AND item_class <> 'FG' GROUP BY date,bpr.department,part_details.item_class) AS ABC GROUP BY date", Date, DepartmentName))
	defer rows2.Close()
	if err != nil {
		fmt.Println("failed to get data from item class wise bpr summary total", err)
		return nil, err
	}

	var BPRSummTotalRec BPRSummary
	for rows2.Next() {
		j++
		err := rows2.Scan(&BPRSummTotalRec.Date, &BPRSummTotalRec.Total, &BPRSummTotalRec.B, &BPRSummTotalRec.R, &BPRSummTotalRec.Y, &BPRSummTotalRec.G, &BPRSummTotalRec.W, &BPRSummTotalRec.Percent_of_Red_Items, &BPRSummTotalRec.Percnet_of_White_Items)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

	}

	// If j = 0 there are no rows in result
	if j == 0 {
		BPR_Total_ZeroCount = -1
	}
	BPRSummTotalRec.Item_class = "Total"
	BPRSummaryitemclassArray = append(BPRSummaryitemclassArray, BPRSummTotalRec)

	var BPRSICmaininfo Main
	BPRSICmaininfo.Status = "true"
	BPRSICmaininfo.Status_code = "200"
	if BPR_Total_ZeroCount == -1 {
		BPRSICmaininfo.Count = 0
	} else {
		BPRSICmaininfo.Count = len(BPRSummaryitemclassArray)
	}
	BPRSICmaininfo.BPRSummaryInfo = BPRSummaryitemclassArray
	return BPRSICmaininfo, nil
}

func (db *DB_manager) GETCriticalPartListBPRRecords(DepartmentName string, Date string) (interface{}, error) {

	if DepartmentName == "Import" {
		rows, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,bpr.part_number,part.description,part_details.sub_category,bpr.norm,bpr.stock,bpr.gap,bpr.penetration,ROUND((bpr.stock/import_norm.norm_1) :: numeric,2) AS Total_Coverage_In_Months,bpr.colour FROM bpr LEFT JOIN part ON part.part_number = bpr.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department LEFT JOIN import_norm ON bpr.part_number = import_norm .part_number AND bpr.department = import_norm.department AND import_norm .date = (SELECT MAX(date) FROM import_norm) Where bpr.department = '%s' AND bpr.date = '%s' AND (bpr.colour = 'B' OR bpr.colour = 'R') AND import_norm.norm_1 > 0 ORDER BY bpr.penetration DESC", DepartmentName, Date))
		defer rows.Close()
		if err != nil {
			fmt.Println("failed to get data from critical part list ", err)
			return nil, err
		}

		var CritialPartlistArray []CriticalSupplierPartList
		var CriticalPartlistRec CriticalSupplierPartList

		for rows.Next() {
			err := rows.Scan(&CriticalPartlistRec.Date, &CriticalPartlistRec.Department, &CriticalPartlistRec.Part_number, &CriticalPartlistRec.Description, &CriticalPartlistRec.Sub_category, &CriticalPartlistRec.Norm, &CriticalPartlistRec.Stock, &CriticalPartlistRec.Gap, &CriticalPartlistRec.Penetration, &CriticalPartlistRec.Total_coverage_in_months, &CriticalPartlistRec.Colour)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			CritialPartlistArray = append(CritialPartlistArray, CriticalPartlistRec)

		}
		var CritialPartlistinfo CriticalSupplierPartListCount
		CritialPartlistinfo.Status = "true"
		CritialPartlistinfo.Status_code = "200"
		CritialPartlistinfo.Count = len(CritialPartlistArray)
		CritialPartlistinfo.CriticalSupplierPartListInfo = CritialPartlistArray
		return CritialPartlistinfo, nil
	} else {

		rows1, err := db.Query(fmt.Sprintf("SELECT bpr.date,bpr.department,bpr.part_number,part.description ,part_details.sub_category,bpr.norm,bpr.stock,bpr.gap,bpr.penetration,bpr.colour FROM bpr LEFT JOIN part ON part.part_number = bpr.part_number LEFT JOIN part_details ON bpr.part_number = part_details.part_number AND bpr.department=part_details.department Where bpr.department = '%s' AND bpr.date = '%s' AND (bpr.colour = 'B' OR bpr.colour = 'R') ORDER BY bpr.penetration DESC", DepartmentName, Date))
		defer rows1.Close()
		if err != nil {
			fmt.Println("failed to get data from critical part list ", err)
			return nil, err
		}

		var CritialPartlistArray1 []CriticalSupplierPartList
		var CriticalPartlistRec1 CriticalSupplierPartList

		for rows1.Next() {
			err := rows1.Scan(&CriticalPartlistRec1.Date, &CriticalPartlistRec1.Department, &CriticalPartlistRec1.Part_number, &CriticalPartlistRec1.Description, &CriticalPartlistRec1.Sub_category, &CriticalPartlistRec1.Norm, &CriticalPartlistRec1.Stock, &CriticalPartlistRec1.Gap, &CriticalPartlistRec1.Penetration, &CriticalPartlistRec1.Colour)
			if err != nil {
				fmt.Println("failed to scan the record.. continue with the next.. ", err)
				continue
			}
			CritialPartlistArray1 = append(CritialPartlistArray1, CriticalPartlistRec1)

		}

		var CritialPartlistinfo1 CriticalSupplierPartListCount
		CritialPartlistinfo1.Status = "true"
		CritialPartlistinfo1.Status_code = "200"
		CritialPartlistinfo1.Count = len(CritialPartlistArray1)
		CritialPartlistinfo1.CriticalSupplierPartListInfo = CritialPartlistArray1
		return CritialPartlistinfo1, nil
	}
}
