package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Mainroot struct {
	Status        string
	Status_code   string
	Count         int
	UserloginInfo Userlogin
}

type Userlogin struct {
	ID         int
	CustomerID int
	FirstName  string
	LastName   string
	RoleID     int
}

type ResetPassword struct {
	UserID   int
	Password string
	Status   string
}

func (db *DB_manager) ReadUserloginRecords(LoginName string, Password string) (Mainroot, error) {
	rows, err := db.Query(fmt.Sprintf("SELECT M_Users.UserID,coalesce(M_Users.CustomerID,0) AS CustomerID,M_Users.FirstName,M_Users.LastName UserName,MC_UserRoles.RoleID  FROM M_Users JOIN MC_UserRoles ON MC_UserRoles.UserID = M_Users.UserID WHERE   (LoginName = '"+LoginName+"' OR EmailID = '"+LoginName+"' OR Mobile = '"+LoginName+"') AND Password= '%s'", Password))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from user login records ", err)
		return Mainroot{}, err
	}

	var UloginRec Userlogin
	for rows.Next() {
		err := rows.Scan(&UloginRec.ID, &UloginRec.CustomerID, &UloginRec.FirstName, &UloginRec.LastName, &UloginRec.RoleID)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

	}
	var loginMainroot Mainroot
	if UloginRec.ID > 0 {
		loginMainroot.Count = 1
	} else {
		loginMainroot.Count = 0
	}

	loginMainroot.Status = "true"
	loginMainroot.Status_code = "200"
	loginMainroot.UserloginInfo = UloginRec
	return loginMainroot, nil
}

func (db *DB_manager) ResetPasswordByUserID(UserID int, OldPassword string, NewPassword string) (ResetPassword, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT UserID,password from m_users Where UserID = %d", UserID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from user login records", err)
		return ResetPassword{}, err
	}

	var ResetPasswordRec ResetPassword
	for rows.Next() {
		err := rows.Scan(&ResetPasswordRec.UserID, &ResetPasswordRec.Password)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}

	}
	if ResetPasswordRec.Password == OldPassword {
		query := fmt.Sprintf(`UPDATE m_users SET password = '%s' Where UserID = %d`, NewPassword, ResetPasswordRec.UserID)

		stmt, err := db.Query(query)

		if err != nil {
			fmt.Println("failed to update user login records. Continue with the next operation", ResetPasswordRec)
			fmt.Println("error ", err)

		}
		defer stmt.Close()

		if err != nil {
			log.Println(err)

		}
		fmt.Println("1 record Updated")
		ResetPasswordRec.Status = "true"
	} else {
		ResetPasswordRec.Status = "False"
	}

	return ResetPasswordRec, nil
}
