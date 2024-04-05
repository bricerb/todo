package entities

type ToDo struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Complete bool   `json:"complete,omitempty"`
}