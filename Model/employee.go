package Model

type Position struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	Id         int64    `json:"id"`
	FirstName  string   `json:"firstName"`
	LastName   string   `json:"lastName"`
	PositionId int64    `json:"positionId"`
	Position   Position `json:"position"`
}
