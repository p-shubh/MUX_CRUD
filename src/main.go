package main

import (
	"CEO_ASSIST/Router"
	"CEO_ASSIST/db"
	"flag"
)

func init() {

	/*RUN BEFORE THE MAIN FUNCTION RUNS*/

}

func main() {

	flag.Parse()

	/*DATABASE CONNECTION*/
	db.DATABASE_CONNECTION()

	/*ROUTER ENGINE*/
	Router.Router_Engine()

}
