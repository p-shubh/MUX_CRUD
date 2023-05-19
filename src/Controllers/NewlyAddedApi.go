package Controllers

import (
	"CEO_ASSIST/db"
	"net/http"
)

func (cfg *Config) MasterdataBulkUpload(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title                            string
		Message                          string
		MessageColor                     string
		ItemMasterInsertUpdateCountArray []db.ItemMasterInsertUpdateCount
	}{
		Title:                            "Intellixente",
		Message:                          "",
		MessageColor:                     "",
		ItemMasterInsertUpdateCountArray: []db.ItemMasterInsertUpdateCount{},
	}

	render(w, "masterdataBulkUpload", data)
}
