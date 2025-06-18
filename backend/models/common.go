package models

type ResultAPI struct {
	StatusCode int         `json:"nStatusCode"`
	Message    string      `json:"sMessage,omitempty"`
	Result     interface{} `json:"objResult,omitempty"`
}

type Role struct {
	ID       uint
	RoleName string `gorm:"column:rolename"`
}

type Position struct {
	ID           uint
	PositionName string `gorm:"column:positionname"`
}

type Employee struct {
	ID         uint
	FirstName  string `gorm:"column:firstname"`
	LastName   string `gorm:"column:lastname"`
	RoleID     uint
	PositionID uint
	Role       Role     `gorm:"foreignKey:RoleID"`
	Position   Position `gorm:"foreignKey:PositionID"`
}
