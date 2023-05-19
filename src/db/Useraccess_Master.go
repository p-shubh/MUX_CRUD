package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type UserMaster struct {
	UserID               int
	FirstName            string
	LastName             string
	LoginName            string
	Mobile               string
	EmailID              string
	Password             string
	Createdby            int
	Createdon            string
	Updatedby            int
	Updatedon            string
	MC_userrolesInfo     []MC_userroles
	MC_userdivisionsInfo []MC_userdivisions
}

type MC_userroles struct {
	MC_userrolesID        int
	UserID                int
	RoleID                int
	MC_userrolesRow_Order int
}

type MC_userdivisions struct {
	MC_userdivisionsID        int
	UserID                    int
	DepartmentID              int
	MC_userdivisionsRow_Order int
}

type User_Access_Master struct {
	Userid                int
	UserAccessMasterItems []UserAccessMasterInfo
}

type UserAccessMasterInfo struct {
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

func (db *DB_manager) InsertUserMasterRecords(rec UserMaster, AddRowCountUserRoles int, AddRowCountUserDivision int) (int, error) {
	var UserID int

	query := `INSERT INTO 
    m_users(firstname,lastname,loginname,mobile,emailid,password,createdby,createdon,updatedby,updatedon) 
    VALUES($1, $2, $3,$4,$5,$6,$7,$8,$9,$10)RETURNING UserID`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return UserID, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(
		rec.FirstName,
		rec.LastName,
		rec.LoginName,
		rec.Mobile,
		rec.EmailID,
		rec.Password,
		rec.Createdby,
		rec.Createdon,
		rec.Updatedby,
		rec.Updatedon,
	).Scan(&UserID)

	if err != nil {
		log.Fatal(err)
		return UserID, err
	}

	var i int
	for i = 0; i < AddRowCountUserRoles; i++ {

		query = `INSERT INTO 
    mc_userroles(userid,roleid,row_order) 
    VALUES($1, $2,$3)`

		stmt, err = db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return UserID, err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			UserID,
			rec.MC_userrolesInfo[i].RoleID,
			rec.MC_userrolesInfo[i].MC_userrolesRow_Order,
		)
		if err != nil {
			log.Fatal(err)
			return UserID, err
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(affect, "record added")
	}

	var j int
	for j = 0; j < AddRowCountUserDivision; j++ {

		query = `INSERT INTO
		    mc_userdivisions(userid,departmentid,row_order)
		    VALUES($1, $2, $3)`

		stmt, err = db.Prepare(query)
		if err != nil {
			log.Fatal(err)
			return UserID, err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			UserID,
			rec.MC_userdivisionsInfo[j].DepartmentID,
			rec.MC_userdivisionsInfo[j].MC_userdivisionsRow_Order,
		)
		if err != nil {
			log.Fatal(err)
			return UserID, err
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(affect, "record added")
	}
	return UserID, err
}

func (db *DB_manager) UpdateUserMasterRecords(rec UserMaster, AddRowCountUserRoles int, AddRowCountUserDivision int) error {

	// update query for selected user id ( Section 1)
	query := fmt.Sprintf(`UPDATE m_users SET firstname = '%s',lastname = '%s',loginname = '%s'	,
mobile= '%s',emailid = '%s',password = '%s',updatedby = %d,updatedon ='%s' Where userid = %d`, rec.FirstName, rec.LastName, rec.LoginName, rec.Mobile, rec.EmailID, rec.Password, rec.Updatedby, rec.Updatedon, rec.UserID)

	stmt, err := db.Query(query)

	if err != nil {
		fmt.Println("failed to update m_users record. Continue with the next operation ", rec)
		fmt.Println("error ", err)

	}
	defer stmt.Close()

	if err != nil {
		log.Println(err)

	}
	fmt.Println("1 record Updated")

	var i int
	for i = 0; i < AddRowCountUserRoles; i++ {

		if rec.MC_userrolesInfo[i].MC_userrolesID > 0 {
			query = fmt.Sprintf(`UPDATE mc_userroles SET roleid = %d,row_order = %d Where id = %d`, rec.MC_userrolesInfo[i].RoleID, rec.MC_userrolesInfo[i].MC_userrolesRow_Order, rec.MC_userrolesInfo[i].MC_userrolesID)

			stmt, err = db.Query(query)

			if err != nil {
				fmt.Println("failed to update mc_userroles status record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Println(err)

			}
			fmt.Println("1 record Updated")
		} else {

			query = `INSERT INTO 
    mc_userroles(userid,roleid,row_order) 
    VALUES($1, $2,$3)`

			stmt, err := db.Prepare(query)
			if err != nil {
				log.Fatal(err)
				return err
			}

			defer stmt.Close()

			res, err := stmt.Exec(
				rec.UserID,
				rec.MC_userrolesInfo[i].RoleID,
				rec.MC_userrolesInfo[i].MC_userrolesRow_Order,
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
	for j = 0; j < AddRowCountUserDivision; j++ {

		if rec.MC_userdivisionsInfo[j].MC_userdivisionsID > 0 {
			query = fmt.Sprintf(`UPDATE mc_userdivisions SET departmentid = %d,row_order = %d Where id = %d`, rec.MC_userdivisionsInfo[j].DepartmentID, rec.MC_userdivisionsInfo[j].MC_userdivisionsRow_Order, rec.MC_userdivisionsInfo[j].MC_userdivisionsID)

			stmt, err = db.Query(query)

			if err != nil {
				fmt.Println("failed to update mc_userdivisions status record. Continue with the next operation ", rec)
				fmt.Println("error ", err)

			}
			defer stmt.Close()

			if err != nil {
				log.Println(err)

			}
			fmt.Println("1 record Updated")
		} else {

			query = `INSERT INTO
		    mc_userdivisions(userid,departmentid,row_order)
		    VALUES($1, $2, $3)`

			stmt, err := db.Prepare(query)
			if err != nil {
				log.Fatal(err)
				return err
			}

			defer stmt.Close()

			res, err := stmt.Exec(
				rec.UserID,
				rec.MC_userdivisionsInfo[j].DepartmentID,
				rec.MC_userdivisionsInfo[j].MC_userdivisionsRow_Order,
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

func (db *DB_manager) InsertUserAccessMasterRecords(rec User_Access_Master, Userid int) error {

	var i int
	for i = 0; i < len(rec.UserAccessMasterItems); i++ {

		query := `INSERT INTO
     m_useraccess(userid,companyid,divisionid,moduleid,submoduleid,pageid,isshow,isshowaddpage,isadd,iseditself,isedit,isview,isdeleteself,isdelete,isprint)
    VALUES($1, $2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)`
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			Userid,
			1,
			0,
			rec.UserAccessMasterItems[i].Moduleid,
			rec.UserAccessMasterItems[i].SubModuleid,
			rec.UserAccessMasterItems[i].Pageid,
			rec.UserAccessMasterItems[i].Isshow,
			rec.UserAccessMasterItems[i].Isshowaddpage,
			rec.UserAccessMasterItems[i].Isadd,
			rec.UserAccessMasterItems[i].Iseditself,
			rec.UserAccessMasterItems[i].Isedit,
			rec.UserAccessMasterItems[i].Isview,
			rec.UserAccessMasterItems[i].Isdeleteself,
			rec.UserAccessMasterItems[i].Isdelete,
			rec.UserAccessMasterItems[i].Isprint,
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

func (db *DB_manager) DeleteUserAccessMaster(UserID int) ([]User_Access_Master, error) {
	query := fmt.Sprintf(`DELETE FROM M_useraccess Where UserID = %d`, UserID)
	stmt, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to delete User_Access_Master record. Continue with the next operation ", err)
		fmt.Println("error ", err)
	}
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("records deleted successfully")
	return nil, err
}
