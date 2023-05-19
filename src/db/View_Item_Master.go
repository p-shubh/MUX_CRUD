package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type ItemMasterroot struct {
	Status                        string
	Status_code                   string
	Count                         int
	AllItem_Master_DetailsInfo    []Item_Master_Details
	Item_Master_DetailsInfo       []Item_Master_Details_WithID
	Norm_Calculation_details_Info []Norm_Calculation_details
}

type Item_Master_Details struct {
	PartID                int
	Part_DetailsID        int
	Part_number           *string
	Description           *string
	Department            *string
	Sub_category          *string
	Item_class            *string
	Rate                  *float64
	Alternate_part_number *string
	Plant                 *string
	Buyer                 *string
	Supplier              *string
	Lead_time             *float64
	Safety_factor         *float64
	Row_Order             *int
}

type Item_Master_Details_WithID struct {
	PartID                int
	Part_DetailsID        int
	Part_number           *string
	Description           *string
	Department            *string
	Sub_category          *string
	Item_class            *string
	Rate                  *float64
	Alternate_part_number *string
	PlantID               *string
	Plant                 *string
	BuyerID               *string
	Buyer                 *string
	SupplierID            *string
	Supplier              *string
	Lead_time             *float64
	Safety_factor         *float64
	Row_Order             *int
}

type Norm_Calculation_details struct {
	Date          string
	Department    string
	Part_number   string
	NOD           float64
	Lead_time     float64
	Safety_factor float64
	Buyer         string
	Supplier      string
	Sub_category  string
	Rate          float64
}

func (db *DB_manager) GETAllItemMasterRecords() (ItemMasterroot, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT part.part_id,part.part_number,part.description,part_details.department,part_details.sub_category,coalesce(part_details.item_class,'') AS item_class,coalesce(part_details.rate,0.0) AS rate,coalesce(part_details.alternate_part_number,'') AS alternate_part_number,coalesce(part_details.plant,'') AS plant,coalesce(part_supplier_details.buyer,'') AS buyer,coalesce(part_supplier_details.supplier,'') AS supplier,coalesce(norm_cd.lead_time,0.0) AS lead_time,coalesce(norm_cd.safety_factor,0.0) AS safety_factor FROM Part LEFT JOIN part_details ON part.part_number = part_details.part_number LEFT JOIN part_supplier_details ON part.part_number = part_supplier_details.part_number AND part_details.department = part_supplier_details.department LEFT JOIN norm_cd ON part.part_number = norm_cd.part_number AND part_details.department = norm_cd.department ORDER BY part_details.row_order ASC"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Item_Master_Details", err)
		return ItemMasterroot{}, err
	}
	var ItemMasterArray []Item_Master_Details
	var ItemMasterRec Item_Master_Details
	for rows.Next() {
		err := rows.Scan(&ItemMasterRec.PartID, &ItemMasterRec.Part_number, &ItemMasterRec.Description, &ItemMasterRec.Department, &ItemMasterRec.Sub_category, &ItemMasterRec.Item_class, &ItemMasterRec.Rate, &ItemMasterRec.Alternate_part_number, &ItemMasterRec.Plant, &ItemMasterRec.Buyer, &ItemMasterRec.Supplier, &ItemMasterRec.Lead_time, &ItemMasterRec.Safety_factor)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		ItemMasterArray = append(ItemMasterArray, ItemMasterRec)
	}

	var item_master_details_info ItemMasterroot
	item_master_details_info.Status = "true"
	item_master_details_info.Status_code = "200"
	item_master_details_info.Count = len(ItemMasterArray)
	item_master_details_info.AllItem_Master_DetailsInfo = ItemMasterArray
	return item_master_details_info, nil
}

func (db *DB_manager) GETItemMasterRecords(PartID int) (ItemMasterroot, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT part.part_id,part_details.part_details_id,part.part_number,part.description,part_details.department,part_details.sub_category,coalesce(part_details.item_class,'') AS item_class,coalesce(part_details.rate,0.0) AS rate,coalesce(part_details.alternate_part_number,'') AS alternate_part_number,coalesce(plant.plantname,'') AS ID,coalesce(part_details.plant,'') AS plant,coalesce(buyer.buyername,'') AS ID,coalesce(part_supplier_details.buyer,'') AS buyer,coalesce(supplier.suppliername,'') AS ID,coalesce(part_supplier_details.supplier,'') AS supplier,coalesce(norm_cd.lead_time,0.0) AS lead_time,coalesce(norm_cd.safety_factor,0.0) AS safety_factor,norm_cd.row_order FROM Part LEFT JOIN part_details ON part.part_number = part_details.part_number LEFT JOIN part_supplier_details ON part.part_number = part_supplier_details.part_number AND part_details.department = part_supplier_details.department LEFT JOIN norm_cd ON part.part_number = norm_cd.part_number AND part_details.department = norm_cd.department LEFT JOIN plant ON part_details.plant = plant.plantname LEFT JOIN buyer ON part_supplier_details.buyer = buyer.buyername LEFT JOIN supplier ON part_supplier_details.supplier = supplier.suppliername Where part.part_id = %d ORDER BY part_details.row_order ASC", PartID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from Item_Master_Details", err)
		return ItemMasterroot{}, err
	}
	var ItemMasterArray1 []Item_Master_Details_WithID
	var ItemMasterRec1 Item_Master_Details_WithID
	for rows.Next() {
		err := rows.Scan(&ItemMasterRec1.PartID, &ItemMasterRec1.Part_DetailsID, &ItemMasterRec1.Part_number, &ItemMasterRec1.Description, &ItemMasterRec1.Department, &ItemMasterRec1.Sub_category, &ItemMasterRec1.Item_class, &ItemMasterRec1.Rate, &ItemMasterRec1.Alternate_part_number, &ItemMasterRec1.PlantID, &ItemMasterRec1.Plant, &ItemMasterRec1.BuyerID, &ItemMasterRec1.Buyer, &ItemMasterRec1.SupplierID, &ItemMasterRec1.Supplier, &ItemMasterRec1.Lead_time, &ItemMasterRec1.Safety_factor, &ItemMasterRec1.Row_Order)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

		ItemMasterArray1 = append(ItemMasterArray1, ItemMasterRec1)
	}

	var item_master_details_info ItemMasterroot
	item_master_details_info.Status = "true"
	item_master_details_info.Status_code = "200"
	item_master_details_info.Count = len(ItemMasterArray1)
	item_master_details_info.Item_Master_DetailsInfo = ItemMasterArray1
	return item_master_details_info, nil
}

// Fetch data from Norm Calculation table for download & upload excel file format
func (db *DB_manager) DownloadNormCalculationDetailsMasterData(DepartmentName string) (ItemMasterroot, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT norm_cd.date,norm_cd.department,norm_cd.part_number,norm_cd.NOD,norm_cd.lead_time,norm_cd.safety_factor,coalesce(norm_cd.buyer,'') AS buyer,coalesce(norm_cd.supplier) AS supplier,coalesce(part_details.sub_category,'') AS sub_category,coalesce(part_details.rate,0.0) AS rate FROM norm_cd LEFT JOIN part_details ON norm_cd.part_number = part_details.part_number AND part_details.department = norm_cd.department Where norm_cd.department = '%s'", DepartmentName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from norm calculation details ", err)
		return ItemMasterroot{}, err
	}

	var Norm_CDArray []Norm_Calculation_details
	var Norm_CDRec Norm_Calculation_details
	for rows.Next() {
		err := rows.Scan(&Norm_CDRec.Date, &Norm_CDRec.Department, &Norm_CDRec.Part_number, &Norm_CDRec.NOD, &Norm_CDRec.Lead_time, &Norm_CDRec.Safety_factor, &Norm_CDRec.Buyer, &Norm_CDRec.Supplier, &Norm_CDRec.Sub_category, &Norm_CDRec.Rate)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		Norm_CDArray = append(Norm_CDArray, Norm_CDRec)

	}
	var norm_cd_info ItemMasterroot
	norm_cd_info.Status = "true"
	norm_cd_info.Status_code = "200"
	norm_cd_info.Count = len(Norm_CDArray)
	norm_cd_info.Norm_Calculation_details_Info = Norm_CDArray
	return norm_cd_info, nil
}
