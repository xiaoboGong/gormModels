package models

type ChosenProductLinks struct{
    Id int `gorm:"primaryKey; column:id;not null"`
    Link string `gorm:" column:link;not null"`
    ProductId int `gorm:" column:product_id;not null"`
    CreateTime int `gorm:" column:create_time;not null"`
}
