package services

import (
	"context"
	"e-commerce-transaction/models/database"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllSellers() ([]database.Sellers, error) {
	var data []database.Sellers

	sellersCollection := mgm.Coll(&database.Sellers{})

	err := sellersCollection.SimpleFind(&data, bson.M{})
	if err != nil {
		return data, err
	}

	return data, nil
}

func GetTotalSellsAggregateByState() ([]primitive.M, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	group := bson.D{
		{
			"$group", bson.D{
				{"_id", bson.D{
					{"seller_state", "$seller_state"},
				}},
				{"total_value", bson.D{{"$sum", "$total_value"}}},
			},
		},
	}

	sort := bson.D{{
		"$sort", bson.M{"total_value": -1},
	}}

	cursor, err := mgm.Coll(&database.Sellers{}).Aggregate(ctx, mongo.Pipeline{group, sort})
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetSellersByState(state string) ([]database.Sellers, error) {
	var data []database.Sellers

	sellersCollection := mgm.Coll(&database.Sellers{})

	err := sellersCollection.SimpleFind(&data, bson.M{"seller_state": state})
	if err != nil {
		return data, err
	}

	return data, nil
}

// pipeline := []bson.M{
// 	{
// 		"$group": bson.M{
// 			"_id":         "$seller_state",
// 			"total_value": bson.M{"$sum": "$total_value"},
// 			"maxRange":    bson.M{"$max": "$max_shipping_date"},
// 			"minRange":    bson.M{"$min": "$min_shipping_date"},
// 			"avgPrice":    bson.M{"$avg": "$total_value"},
// 		},
// 	},
// 	{
// 		"$sort": bson.M{"sumQuantity": -1},
// 	},
// }

// results := []bson.M{}
// err := mongodb.GetMongoDB().C("sellers").Pipe(pipeline).All(&results)
