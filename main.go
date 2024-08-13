package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Description string             `json:"description"`
	Value       int                `json:"value"`
	Type        int                `json:"type"`
	// CreateDate  primitive.DateTime `json:"create_date,omitempty"`
	// CreatedBy   string             `json:"create_by,omitempty"`
	// UpdateDate  primitive.DateTime `json:"update_date,omitempty"`
	// UpdatedBy   string             `json:"updated_by,omitempty"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello Universe")
	if os.Getenv("ENV") != "producion" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("CONNECTED TO MONGODB ATLAS: SFPC_DB")
	collection = client.Database("sfpc_db").Collection("transactions")

	app := fiber.New()
	app.Get("/api/transaction", getTransactions)
	app.Post("/api/transaction", newTransaction)
	app.Patch("/api/transaction/:id", updateTransaction)
	app.Delete("/api/transaction/:id", deleteTransaction)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "4000"
	}
	app.Listen(":" + PORT)
}

func getTransactions(c *fiber.Ctx) error {
	var transactions []Transaction

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var transaction Transaction
		if err := cursor.Decode(&transaction); err != nil {
			return err
		}
		transactions = append(transactions, transaction)
	}

	return c.JSON(transactions)
}

func newTransaction(c *fiber.Ctx) error {
	transaction := new(Transaction)

	if err := c.BodyParser(transaction); err != nil {
		return err
	}

	if transaction.Description == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Transaction cannot be empty"})
	}

	if transaction.Value == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Transaction cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), transaction)
	if err != nil {
		return err
	}

	transaction.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(200).JSON(transaction)
}

func updateTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	transaction := new(Transaction)

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Transaction ID"})
	}

	err = c.BodyParser(transaction)
	if err != nil {
		return err
	}

	if transaction.Value == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Value"})
	}

	if transaction.Description == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Description"})
	}

	fmt.Println(transaction.Type)
	if !(transaction.Type == 1 || transaction.Type == 0) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Type"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"value":       transaction.Value,
			"description": transaction.Description,
			"type":        transaction.Type,
		}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

func deleteTransaction(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"success": true})
}
