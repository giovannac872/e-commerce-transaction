package main

import (
	"e-commerce-transaction/controllers"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv() {
	godotenv.Load()

}

func GetMongoUri() string {
	HOST := os.Getenv("DB_MONGO_HOST")
	PORT := os.Getenv("DB_MONGO_PORT")
	USER := os.Getenv("DB_MONGO_USER")
	PASS := url.QueryEscape(os.Getenv("DB_MONGO_PASS"))

	return fmt.Sprintf("mongodb://%v:%v@%v:%v", USER, PASS, HOST, PORT)
}

func ConfigMongoDB() {
	uri := GetMongoUri()

	err := mgm.SetDefaultConfig(nil, os.Getenv("DB_MONGO_DATABASE"), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error setting mgm default config")
		os.Exit(1)
	}
}
func ConfigRouterAndStartHTTPServer(port string) {
	router := fiber.New()

	// Add middlewares
	router.Use(cors.New())

	productsController := controllers.ProductsController{}
	sellersController := controllers.SellersController{}

	router.Get("/products", productsController.GetAllProducts)
	router.Get("/products/categories/all", productsController.GetMeanProductsPriceAggregateByCategory)
	router.Get("/products/:category/category", productsController.GetProductsByCategory)

	router.Get("/sellers", sellersController.GetAllSellers)
	router.Get("/sellers/sells/all", sellersController.GetTotalSellsAggregateByState)
	router.Get("/sellers/:state", sellersController.GetSellersByState)

	log.Println("Listening on port " + port)
	router.Listen(":" + port)
}

func main() {

	LoadEnv()

	ConfigMongoDB()

	port := os.Getenv("HTTP_PORT")

	ConfigRouterAndStartHTTPServer(port)
}
