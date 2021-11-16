package services

import (
	"context"
	"e-commerce-transaction/models/database"
	"e-commerce-transaction/models/response"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllProducts() ([]database.Products, error) {
	var data []database.Products

	productsCollection := mgm.Coll(&database.Products{})

	err := productsCollection.SimpleFind(&data, bson.M{})
	if err != nil {
		return data, err
	}

	return data, nil
}

func GetAllProductsByCategory(slugCategory string) (response.ProductsByCategory, error) {
	var data response.ProductsByCategory

	productsCollection := mgm.Coll(&database.Products{})

	err := productsCollection.SimpleFind(&data.Products, bson.M{"product_category_name": slugCategory})
	if err != nil {
		return data, err
	}

	sum := float32(0)
	for i := 0; i < len(data.Products); i++ {
		sum += (data.Products[i].MeanPrice)
	}

	data.Total = len(data.Products)
	data.MeanPriceCategory = float32(sum) / float32(len(data.Products))

	return data, nil

}

func GetMeanProductsPriceAggregateByCategory() ([]primitive.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	group := bson.D{
		{
			"$group", bson.D{
				{"_id", bson.D{
					{"product_category_name", "$product_category_name"},
				}},
				{"mean_price_category", bson.D{{"$avg", "$mean_price"}}},
			},
		},
	}

	sort := bson.D{{
		"$sort", bson.M{"mean_price_category": -1},
	}}

	cursor, err := mgm.Coll(&database.Products{}).Aggregate(ctx, mongo.Pipeline{group, sort})
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// func GetAllMeanPricesCategory(state string) (response.ProductsByCategory, error) {

// 	matchCase := bson.D{{"$match"}}

// }
