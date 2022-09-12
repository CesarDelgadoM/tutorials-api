package handlers

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CesarDelgadoM/tutorials-api/api/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TutorialHandlers struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTutorialHandler(db *mongo.Database) *TutorialHandlers {

	collection := db.Collection(os.Getenv("COLLECTION_TUTORIAL"))

	return &TutorialHandlers{
		db:         db,
		collection: collection,
	}
}

//var collection *mongo.Collection = database.ConnectMongoDB().Collection(os.Getenv("COLLECTION_TUTORIAL"))

func (h *TutorialHandlers) GetAllTutorials(c *fiber.Ctx) error {

	log.Println("Create context...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tutorials []models.Tutorial

	log.Println("Find data in the database...")
	result, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	defer result.Close(ctx)

	log.Println("Decode all data...")
	for result.Next(ctx) {

		var tutorial models.Tutorial
		if err = result.Decode(&tutorial); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Decode Model Error")
		}

		tutorials = append(tutorials, tutorial)
	}

	log.Println("Generating reponse...")
	return c.Status(fiber.StatusOK).JSON(tutorials)
}

func (h *TutorialHandlers) GetTutorialById(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tutorial models.Tutorial

	id := c.Params("id")
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Parse ObjectIDFromHex Error")
	}

	err = h.collection.FindOne(ctx, bson.M{"_id": idhex}).Decode(&tutorial)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tutorial Not Found")
	}

	return c.Status(fiber.StatusOK).JSON(tutorial)
}

func (h *TutorialHandlers) CreateTutorial(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tutorial models.Tutorial

	err := c.BodyParser(&tutorial)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Bad Request")
	}

	newTutorial := models.NewTutorial(primitive.NewObjectID(), tutorial.Title,
		tutorial.Description, tutorial.Published)
	newTutorial.CreateAt = primitive.NewDateTimeFromTime(time.Now())
	newTutorial.UpdateAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := h.collection.InsertOne(ctx, newTutorial)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *TutorialHandlers) UpdateTutorial(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tutorial models.Tutorial

	id := c.Params("id")
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	err = c.BodyParser(&tutorial)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "BodyParser error")
	}

	tutorialUpdate := models.NewTutorial(idhex, tutorial.Title, tutorial.Description, tutorial.Published)
	tutorialUpdate.UpdateAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := h.collection.UpdateOne(ctx, bson.M{"_id": idhex}, bson.M{"$set": tutorialUpdate})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *TutorialHandlers) DeleteTutorialById(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id := c.Params("id")
	idhex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Parse ObjectIDFromHex Error")
	}

	result, err := h.collection.DeleteOne(ctx, bson.M{"_id": idhex})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Delete Tutorial Error")
	}

	if result.DeletedCount < 1 {
		return fiber.NewError(fiber.StatusNotFound, "Not Found Tutorial to Delete")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *TutorialHandlers) DeleteAllTutorials(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := h.collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Delete Many Error")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}

func (h *TutorialHandlers) GetTutorialByTitle(c *fiber.Ctx) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var tutorials []models.Tutorial
	title := c.Params("title")

	cursor, err := h.collection.Find(ctx, bson.D{
		{
			Key: "title", Value: primitive.Regex{
				Pattern: title,
				Options: "",
			},
		},
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Query Find Title Error")
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {

		var tutorial models.Tutorial
		if err := cursor.Decode(&tutorial); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Decode Model Error")
		}

		tutorials = append(tutorials, tutorial)
	}

	return c.Status(fiber.StatusOK).JSON(tutorials)

}
