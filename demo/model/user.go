package model

import "github.com/jinzhu/gorm"

//	type User struct {
//		gorm.Model
//		Name      string `gorm:"varchar(20);not null" binding:"required"`
//		Telephone string `gorm:"varchar(20);not null;unique" binding:"required,len=11"`
//		Password  string `gorm:"size:255;not null" binding:"required,min=6"`
//	}
type User struct {
	gorm.Model
	Name      string `gorm:"varchar(20);not null" form:"name" json:"name" binding:"required"`
	Telephone string `gorm:"varchar(20);not null;unique" form:"telephone" json:"telephone" binding:"required"`
	Password  string `gorm:"size:255;not null" form:"password" json:"password" binding:"required"`
}
