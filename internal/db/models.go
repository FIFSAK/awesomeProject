package db

type DB struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Size int64  `json:"size"`
}
