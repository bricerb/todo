package entities

type ToDo struct {
	ID       string `json:"id,omitempty" db:"id"`
	Name     string `json:"name,omitempty" form:"name" db:"name"`
	Complete bool   `json:"complete" form:"complete" db:"complete"`
}
