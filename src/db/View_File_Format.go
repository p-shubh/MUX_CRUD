package db

import (
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type FileFormat struct {
	ColumnName   string
	ColumnNumber int
}

type FileFormatMain struct {
	Status         string
	Status_code    string
	Count          int
	FileFormatInfo []FileFormatRec
}

type FileFormatRec struct {
	Department      string
	Functionname    string
	Filename        string
	Excel_File_Name string
	Columnname      string
	Column_order    int
	Column_number   int
}

func (db *DB_manager) GETFileFormatRecord(Department string, FunctionName string, FileName string) ([]FileFormat, error) {

	qry := fmt.Sprintf("SELECT columnname,column_number FROM file_format WHERE department = '%s' AND functionname = '%s' AND filename = '%s'", Department, FunctionName, FileName)
	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from file format", err)
		fmt.Println("failed_query ", qry)
		return nil, err
	}
	var FileFormatArray []FileFormat
	var FileFormatRec FileFormat
	for rows.Next() {
		err := rows.Scan(&FileFormatRec.ColumnName, &FileFormatRec.ColumnNumber)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		//var remove_space = strings.Join(strings.Fields(FileFormatRec.ColumnName), " ")
		//var s_lowercase = strings.ToLower(remove_space)
		FileFormatArray = append(FileFormatArray, FileFormatRec)
	}

	return FileFormatArray, nil
}

func (db *DB_manager) ReadFileFormatInfoRecords(Department string, FunctionName string, FileName string) (interface{}, error) {

	qry := fmt.Sprintf("SELECT department,functionname,filename,coalesce(excel_file_name,'') AS excel_file_name,columnname,column_order,column_number FROM file_format Where department = '%s' AND functionname = '%s' AND filename = '%s' ORDER BY column_number", Department, FunctionName, FileName)
	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from file format rec", err)
		fmt.Println("failed_query ", qry)
		return nil, err
	}
	var FileFormatArray []FileFormatRec
	var fileformat_recs FileFormatRec
	for rows.Next() {
		err := rows.Scan(&fileformat_recs.Department, &fileformat_recs.Functionname, &fileformat_recs.Filename, &fileformat_recs.Excel_File_Name, &fileformat_recs.Columnname, &fileformat_recs.Column_order, &fileformat_recs.Column_number)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		FileFormatArray = append(FileFormatArray, fileformat_recs)
	}

	var FileFormatinfo FileFormatMain
	FileFormatinfo.Status = "true"
	FileFormatinfo.Status_code = "200"
	FileFormatinfo.Count = len(FileFormatArray)
	FileFormatinfo.FileFormatInfo = FileFormatArray
	return FileFormatinfo, nil

}

func (db *DB_manager) GetFileNameFromExcelName(Excel_file_name string, Department string, FunctionName string) (string, error) {

	qry := fmt.Sprintf("SELECT distinct filename FROM file_format where position(excel_file_name in '%s') > 0 and department = '%s' and functionname = '%s'", Excel_file_name, Department, FunctionName)
	rows, err := db.Query(qry)
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from file format", err)
		fmt.Println("failed_query ", qry)
		return "", err
	}

	var filetypeRec string
	for rows.Next() {
		err := rows.Scan(&filetypeRec)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		return filetypeRec, nil
	}

	return "", errors.New("failed_to_get_file_format_record")
}
