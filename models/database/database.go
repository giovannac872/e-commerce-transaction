package database

import "github.com/kamva/mgm/v3"

type Products struct {
	mgm.DefaultModel    `bson:",inline"`
	ProductId           string   `json:"product_id" bson:"product_id"`
	ProductCategoryName string   `json:"product_category_name" bson:"product_category_name"`
	SellersId           []string `json:"seler_id" bson:"seller_id"`
	MeanPrice           float32  `json:"mean_price" bson:"mean_price,truncate"`
}

type Sellers struct {
	mgm.DefaultModel         `bson:",inline"`
	SellerId                 string  `json:"seller_id" bson:"seller_id"`
	SellerState              string  `json:"seller_state" bson:"seller_state"`
	SellerCity               string  `json:"seller_city" bson:"seller_city"`
	SellerZipCodePrefix      float32 `json:"seller_zip_code_prefix" bson:"seller_zip_code_prefix"`
	GeolocationZipCodePrefix string  `json:"geolocation_zip_code_prefix"`
	GeolocationLat           float32 `json:"geolocation_lat" bson:"geolocation_lat,truncate"`
	TotalValue               float32 `json:"total_value" bson:"total_value,truncate"`
	MinShippingDate          string  `json:"min_shipping_date" bson:"min_shipping_date"`
	MaxShippingDate          string  `json:"max_shipping_date" bson:"max_shipping_date"`
}
