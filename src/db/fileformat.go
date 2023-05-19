package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type FileFormatInfo struct {
	Department        string
	FunctionName      string
	FileName          string
	ExcelFileName     string
	ColumnDetailsInfo []ColumnDetails
}

type ColumnDetails struct {
	ColumnName   string
	ColumnOrder  int
	ColumnNumber int
}

func (db *DB_manager) InsertFileFormatRecords(rec FileFormatInfo, Department string, FunctionName string, FileName string, ExcelFileName string) error {
	var i int
	for i = 0; i < len(rec.ColumnDetailsInfo); i++ {

		query := `INSERT INTO 
   file_format(department,functionname,filename,excel_file_name,columnname,column_order,column_number) 
    VALUES($1, $2, $3,$4,$5,$6,$7)`
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Println(err)
			return err
		}

		defer stmt.Close()

		res, err := stmt.Exec(
			Department,
			FunctionName,
			FileName,
			ExcelFileName,
			rec.ColumnDetailsInfo[i].ColumnName,
			rec.ColumnDetailsInfo[i].ColumnOrder,
			rec.ColumnDetailsInfo[i].ColumnNumber,
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

func (db *DB_manager) GETFunctionNameBySelectedDepartment(Department string) ([]string, error) {
	functionname := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct functionName FROM file_format Where department = '%s'", Department))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from file format", err)
		return nil, err
	}

	var FuncRec FileFormatInfo
	for rows.Next() {
		err := rows.Scan(&FuncRec.FunctionName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		functionname = append(functionname, FuncRec.FunctionName)
	}

	return functionname, nil
}

func (db *DB_manager) GETFileTypeBySelectedDepartmentANDFunctionName(Department string, FunctionName string) ([]string, error) {
	filetype := []string{}
	rows, err := db.Query(fmt.Sprintf("SELECT distinct filename FROM file_format Where department = '%s' AND functionname = '%s'", Department, FunctionName))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from file format", err)
		return nil, err
	}

	var FileTypeRec FileFormatInfo
	for rows.Next() {
		err := rows.Scan(&FileTypeRec.FileName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		filetype = append(filetype, FileTypeRec.FileName)
	}

	return filetype, nil
}

func (db *DB_manager) DeleteFileFormatRecords(Department string, FunctionName string, FileName string) ([]FileFormatInfo, error) {
	query := fmt.Sprintf(`DELETE FROM file_format Where department = '%s' AND functionname = '%s'
AND filename = '%s'`, Department, FunctionName, FileName)
	stmt, err := db.Query(query)
	if err != nil {
		fmt.Println("failed to delete file format record. Continue with the next operation ", err)
		fmt.Println("error ", err)
	}
	defer stmt.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("records deleted successfully")
	return nil, err
}
