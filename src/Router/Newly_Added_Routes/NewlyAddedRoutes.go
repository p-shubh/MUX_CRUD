package Newly_Added_Routes

import "CEO_ASSIST/Router"

var cfg = Router.Cfg

func NewlyAdded() {
	cfg.Router.HandleFunc("/masterdataBulkUpload", cfg.MasterdataBulkUpload)
}

//var dataConfig = cfg
