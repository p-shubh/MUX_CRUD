package config

import (
	db "CEO_ASSIST/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
const (
	Db_host = "localhost" //"host.docker.internal"
	Db_name = "postgres"
	Db_user = "postgres"
	Db_pw   = "postgres"
	Db_port = 5432
) */

// config file

type DataCollector struct {
	Db_host        string `json:"db_host"`
	Db_name        string `json:"db_name"`
	Db_user        string `json:"db_user"`
	Db_pw          string `json:"db_pw"`
	Db_port        int    `json:"db_port"`
	DataDirectory  string `json:"data_directory"`
	*db.DB_manager `json:_`
}

func NewConfig(filePath string) *DataCollector {
	fmt.Println("parse config file at ", filePath)
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("failed_to_parse_config_file", err)
		return nil
	}

	data := &DataCollector{}
	err = json.Unmarshal([]byte(file), data)
	if err != nil {
		fmt.Println("failed_to_parse_json_config_file ", err)
		return nil
	}
	fmt.Println("parsed_config_file ", data)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		data.Db_host, data.Db_port, data.Db_user, data.Db_pw, data.Db_name)
	fmt.Println("connecting to db ", psqlInfo)
	ck, err := db.NewDB(psqlInfo)
	if err != nil {
		fmt.Println("failed_to_connect_to_database", err)
		return nil
	}

	data.DB_manager = ck
	return data
}
