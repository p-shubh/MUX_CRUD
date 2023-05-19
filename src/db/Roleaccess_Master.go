package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type M_roles struct {
	M_rolesID   int
	Name        string
	Description string
	IsActive    string
	PagePath    string
}

type RoleAccessMaster struct {
	Roleid                int
	RoleAccessMasterItems []RoleAccessMasterInfo
}

type RoleAccessMasterInfo struct {
	Moduleid      int
	SubModuleid   int
	Pageid        int
	Isshow        int
	Isshowaddpage int
	Isadd         int
	Iseditself    int
	Isedit        int
	Isview        int
	Isdeleteself  int
	Isdelete      int
	Isprint       int
}

// Insert M_roles data
func (db *DB_manager) InsertM_rolesRecords(rec M_roles) error {

	query := `INSERT INTO
     m_roles(name,description,isactive,page_path)
    VALUES($1,$2,$3,$4)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(
		rec.Name,
		rec.Description,
		rec.IsActive,
		rec.PagePath,
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

	fmt.Println(affect, "record added")
	return err
}

// Update M_roles data
func (db *DB_manager) UpdateM_rolesRecords(rec M_roles) error {

	query := fmt.Sprintf(`UPDATE m_roles SET name = '%s',description = '%s',isactive = '%s',page_path = '%s' Where id = %d
`, rec.Name, rec.Description, rec.IsActive, rec.PagePath, rec.M_rolesID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update m_roles record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")
	return err
}

func (db *DB_manager) InsertRoleAccessMasterRecords(rec RoleAccessMaster, Roleid int) error {

	var i int
	for i = 0; i < len(rec.RoleAccessMasterItems); i++ {

		query := `INSERT INTO
     m_roleaccess(roleid,companyid,divisionid,moduleid,submoduleid,pageid,isshow,isshowaddpage,isadd,iseditself,isedit,isview,isdeleteself,isdelete,isprint)
    VALUES($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			Roleid,
			1,
			0,
			rec.RoleAccessMasterItems[i].Moduleid,
			rec.RoleAccessMasterItems[i].SubModuleid,
			rec.RoleAccessMasterItems[i].Pageid,
			rec.RoleAccessMasterItems[i].Isshow,
			rec.RoleAccessMasterItems[i].Isshowaddpage,
			rec.RoleAccessMasterItems[i].Isadd,
			rec.RoleAccessMasterItems[i].Iseditself,
			rec.RoleAccessMasterItems[i].Isedit,
			rec.RoleAccessMasterItems[i].Isview,
			rec.RoleAccessMasterItems[i].Isdeleteself,
			rec.RoleAccessMasterItems[i].Isdelete,
			rec.RoleAccessMasterItems[i].Isprint,
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

func (db *DB_manager) DeleteRoleAccessMaster(RoleID int) ([]RoleAccessMaster, error) {
	query := fmt.Sprintf(`DELETE FROM M_roleaccess Where RoleID = %d`, RoleID)
	stmt, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to delete Role_Access_Master record. Continue with the next operation ", err)
		fmt.Println("error ", err)
	}
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("records deleted successfully")
	return nil, err
}
