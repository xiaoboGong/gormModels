package models

type ChosenProduct struct{
    Id int `gorm:"primaryKey; column:id;not null"`
    SiteId int8 `gorm:" column:site_id;not null"`
    SourceId string `gorm:" column:source_id;not null"`
    Title string `gorm:" column:title;not null"`
    Images string `gorm:" column:images;not null"`
    HandlingFeeDiscount string `gorm:" column:handling_fee_discount;not null"`
    Tags string `gorm:" column:tags;not null"`
    Type int8 `gorm:" column:type;not null"`
    CustomsType int8 `gorm:" column:customs_type;not null"`
    Price int `gorm:" column:price;not null"`
    ShipFeeInJapan string `gorm:" column:ship_fee_in_Japan;not null"`
    ShipDaysInJapan string `gorm:" column:ship_days_in_Japan;not null"`
    DeliveryAddress string `gorm:" column:delivery_address;not null"`
    Description string `gorm:" column:description;not null"`
    UseStartAt int `gorm:" column:use_start_at;not null"`
    UseEndAt int `gorm:" column:use_end_at;not null"`
    Status int8 `gorm:" column:status;not null"`
    Flag int `gorm:" column:flag;not null"`
    Operator string `gorm:" column:operator;not null"`
    SkuInfo string `gorm:" column:sku_info;not null"`
    CreateTime int `gorm:" column:create_time;not null"`
    UpdateTime int `gorm:" column:update_time;not null"`
}
