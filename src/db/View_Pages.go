package db

import (
	"fmt"
	_ "github.com/lib/pq"
)

type H_PagesMaster struct {
	H_PagesID           int
	PageName            string
	NameOnMenu          string
	H_PagesDescription  string
	H_PagesDisplayIndex int
	DefaultModuleID     int
	H_ModuleName        string
	DefaultSubModuleID  int
	H_SubModuleName     string
	ActualPath          string
	PageType            string
	H_PagesPageID       int
}

type H_SubPagesMaster struct {
	H_SubPagesID   int
	Pageid         int
	PageName       string
	H_SubPagesName string
}

// GET H_Pages List API
func (db *DB_manager) GETH_PagesMasterListRecords() ([]H_PagesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT H_Pages.id,H_Pages.name,H_Pages.nameonmenu,H_Pages.description,H_Pages.displayindex,H_Pages.defaultmoduleid,H_modules.name AS H_modulesname,H_Pages.defaultsubmoduleid,coalesce(H_submodules.name,'') AS H_submodulesname,H_Pages.actualpagepath,H_Pages.pagetype,H_Pages.pageid FROM H_Pages JOIN H_modules ON H_Pages.defaultmoduleid = H_modules.id LEFT JOIN H_submodules ON H_Pages.defaultsubmoduleid = H_submodules.id ORDER BY H_Pages.id"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_PagesMaster ", err)
		return nil, err
	}
	var H_PageslistArray []H_PagesMaster
	var H_PageslistRec H_PagesMaster
	for rows.Next() {
		err := rows.Scan(&H_PageslistRec.H_PagesID, &H_PageslistRec.PageName, &H_PageslistRec.NameOnMenu, &H_PageslistRec.H_PagesDescription, &H_PageslistRec.H_PagesDisplayIndex, &H_PageslistRec.DefaultModuleID, &H_PageslistRec.H_ModuleName, &H_PageslistRec.DefaultSubModuleID, &H_PageslistRec.H_SubModuleName, &H_PageslistRec.ActualPath, &H_PageslistRec.PageType, &H_PageslistRec.H_PagesPageID)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_PageslistArray = append(H_PageslistArray, H_PageslistRec)
	}
	return H_PageslistArray, nil
}

// GET H_Pages Display/Edit API
func (db *DB_manager) GETH_PagesMasterRecords(H_pagesID int) ([]H_PagesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT H_Pages.id,H_Pages.name,H_Pages.nameonmenu,H_Pages.description,H_Pages.displayindex,H_Pages.defaultmoduleid,H_modules.name AS H_modulesname,H_Pages.defaultsubmoduleid,coalesce(H_submodules.name,'') AS H_submodulesname,H_Pages.actualpagepath,H_Pages.pagetype,H_Pages.pageid FROM H_Pages JOIN H_modules ON H_Pages.defaultmoduleid = H_modules.id LEFT JOIN H_submodules ON H_Pages.defaultsubmoduleid = H_submodules.id Where H_Pages.id = %d", H_pagesID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_PagesMaster ", err)
		return nil, err
	}
	var H_pageseditArray []H_PagesMaster
	var H_pageseditRec H_PagesMaster
	for rows.Next() {
		err := rows.Scan(&H_pageseditRec.H_PagesID, &H_pageseditRec.PageName, &H_pageseditRec.NameOnMenu, &H_pageseditRec.H_PagesDescription, &H_pageseditRec.H_PagesDisplayIndex, &H_pageseditRec.DefaultModuleID, &H_pageseditRec.H_ModuleName, &H_pageseditRec.DefaultSubModuleID, &H_pageseditRec.H_SubModuleName, &H_pageseditRec.ActualPath, &H_pageseditRec.PageType, &H_pageseditRec.H_PagesPageID)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_pageseditArray = append(H_pageseditArray, H_pageseditRec)
	}

	return H_pageseditArray, nil
}

// GET H_SubPages List API
func (db *DB_manager) GETH_SubPagesMasterListRecords() ([]H_SubPagesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT h_subpages.id,h_subpages.pageid,coalesce(h_pages.name,'') AS pagename,h_subpages.name FROM h_subpages LEFT JOIN h_pages ON h_subpages.pageid = h_pages.id ORDER BY h_subpages.id"))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_SubPagesMaster ", err)
		return nil, err
	}
	var H_SubPageslistArray []H_SubPagesMaster
	var H_SubPageslistRec H_SubPagesMaster
	for rows.Next() {
		err := rows.Scan(&H_SubPageslistRec.H_SubPagesID, &H_SubPageslistRec.Pageid, &H_SubPageslistRec.PageName, &H_SubPageslistRec.H_SubPagesName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_SubPageslistArray = append(H_SubPageslistArray, H_SubPageslistRec)
	}
	return H_SubPageslistArray, nil
}

// GET H_SubPages Display/Edit API
func (db *DB_manager) GETH_SubPagesMasterRecords(H_subpagesID int) ([]H_SubPagesMaster, error) {

	rows, err := db.Query(fmt.Sprintf("SELECT h_subpages.id,h_subpages.pageid,coalesce(h_pages.name,'') AS pagename,h_subpages.name FROM h_subpages LEFT JOIN h_pages ON h_subpages.pageid = h_pages.id Where h_subpages.id = %d", H_subpagesID))
	defer rows.Close()
	if err != nil {
		fmt.Println("failed to get data from H_SubPagesMaster ", err)
		return nil, err
	}
	var H_subpageseditArray []H_SubPagesMaster
	var H_subpageseditRec H_SubPagesMaster
	for rows.Next() {
		err := rows.Scan(&H_subpageseditRec.H_SubPagesID, &H_subpageseditRec.Pageid, &H_subpageseditRec.PageName, &H_subpageseditRec.H_SubPagesName)
		if err != nil {
			fmt.Println("failed to scan the record.. continue with the next.. ", err)
			continue
		}
		H_subpageseditArray = append(H_subpageseditArray, H_subpageseditRec)
	}

	return H_subpageseditArray, nil
}
