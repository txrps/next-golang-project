// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTmSubdistrict = "TmSubdistrict"

// TmSubdistrict mapped from table <TmSubdistrict>
type TmSubdistrict struct {
	SubdistrictID   int32  `gorm:"column:SubdistrictID;primaryKey;autoIncrement:true" json:"SubdistrictID"`
	SubdistrictName string `gorm:"column:SubdistrictName;not null" json:"SubdistrictName"`
	DistrictID      int32  `gorm:"column:DistrictID;not null" json:"DistrictID"`
}

// TableName TmSubdistrict's table name
func (*TmSubdistrict) TableName() string {
	return TableNameTmSubdistrict
}
