package photo

type Photo struct {
	ID   string `json:"id"`
	URLs struct {
		Regular string `json:"regular"`
	} `json:"urls"`
}
