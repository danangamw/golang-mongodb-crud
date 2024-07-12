package handlers

import (
	"context"
	"encoding/json"
	"os"

	"github.com/danangamw/go_mongo_crud/config"
	"github.com/danangamw/go_mongo_crud/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create a new user
func CreateUser(c *fiber.Ctx) error {
	db_name := os.Getenv("DATABASE_NAME")
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer client.Disconnect(context.Background())

	var user models.User
	if err := json.Unmarshal(c.Body(), &user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	user.ID = primitive.NewObjectID()

	collection := client.Database(db_name).Collection("users")
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}

// Get all users
func GetAllUsers(c *fiber.Ctx) error {
	db_name := os.Getenv("DATABASE_NAME")
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer client.Disconnect(context.Background())

	collection := client.Database(db_name).Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer cursor.Close(context.Background())

	var users []models.User
	for cursor.Next(context.Background()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		users = append(users, user)
	}

	return c.JSON(users)
}

// Get an user by ID
func GetUserByID(c *fiber.Ctx) error {
	db_name := os.Getenv("DATABASE_NAME")
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := client.Database(db_name).Collection("users")
	var user models.User
	if err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(user)
}

// Update a user
func UpdateUser(c *fiber.Ctx) error {
	db_name := os.Getenv("DATABASE_NAME")
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	var updatedUser models.User
	if err := json.Unmarshal(c.Body(), &updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := client.Database(db_name).Collection("users")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updatedUser}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}

func DeleteUser(c *fiber.Ctx) error {
	db_name := os.Getenv("DATABASE_NAME")
	client, err := config.ConnectToMongoDB()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	defer client.Disconnect(context.Background())

	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := client.Database(db_name).Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(result)
}
