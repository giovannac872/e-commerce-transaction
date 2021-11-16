package response

import "e-commerce-transaction/models/database"

type ProductsByCategory struct {
	Total             int                 `json:"total"`
	MeanPriceCategory float32             `json:"mean_price_category"`
	Products          []database.Products `json:"products"`
}

type MeanPriceCategory struct {
	Name              string  `json:"category_name"`
	MeanPriceCategory float32 `json:"mean_price_category"`
}

type MeanPriceCategories struct {
	Total      int                   `json:"total"`
	Categories []MeanPriceCategories `json:"categories"`
}

type TotalSells struct {
	SellerState string `json:"seller_state"`
	TotalValue  string `json:"total_value"`
}
