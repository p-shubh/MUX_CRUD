package Utils

type request struct {
	ID          string `json:"ID"`
	Domain      string `json:"Domain"`
	File        string `json:"File"`
	Description string `json:"Description"`
}

type response struct {
	Domain      string `json:"Domain"`
	Status      string `json:"Status"`
	File        string `json:"File"`
	Description string `json:"Description"`
}
