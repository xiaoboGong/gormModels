package models

type ChosenSku struct{
    Id int `gorm:"primaryKey; column:id;not null"`
    ProductId int `gorm:" column:product_id;not null"`
    FirstName string `gorm:" column:first_name;not null"`
    SecondName string `gorm:" column:second_name;not null"`
    ThirdName string `gorm:" column:third_name;not null"`
    SkuItemId string `gorm:" column:sku_item_id;not null"`
    SkuSpecId string `gorm:" column:sku_spec_id;not null"`
    SkuTitle string `gorm:" column:sku_title;not null"`
    SkuCover string `gorm:" column:sku_cover;not null"`
    SkuLink string `gorm:" column:sku_link;not null"`
    SkuPrice int `gorm:" column:sku_price;not null"`
    SkuStock int `gorm:" column:sku_stock;not null"`
    SkuLimit string `gorm:" column:sku_limit;not null"`
    CreateTime int `gorm:" column:create_time;not null"`
}
