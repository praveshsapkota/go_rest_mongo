package routes

import (
	"example/hello/pkg/models"
	"example/hello/pkg/utils"
	"log"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ADDItemGroup(app *fiber.App) {
	bookGroup := app.Group("/items")

	bookGroup.Get("/", getItems)
	bookGroup.Get("/:id", getItem)
	bookGroup.Get("/scan/:barcode", searchBarcode)
	bookGroup.Post("/", createItem)
	// bookGroup.Put("/:id", updateBook)
	// bookGroup.Delete("/:id", deleteBook)
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
	message     string
}

var Validator = validator.New()

// type Handler struct {
// 	validator *validator.Validate
// }

// // NewHandler creates a new handler with a validator instance
// func NewHandler() *Handler {
// 	h := &Handler{
// 		validator: validator.New(),
// 	}
// 	// register English translations for validator
// 	en := en.New()
// 	_ = en_translations.RegisterDefaultTranslations(h.validator, en)
// 	return h
// }

func getItems(c *fiber.Ctx) error {
	print("cvsdv--------------------")
	coll := utils.GetDBCollection(os.Getenv("collectionName"))

	// find all books
	items := make([]models.Item, 0)
	cursor, err := coll.Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// iterate over the cursor
	for cursor.Next(c.Context()) {
		book := models.Item{}
		err := cursor.Decode(&book)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		items = append(items, book)
	}

	return c.Status(200).JSON(fiber.Map{"data": items})
}

func getItem(c *fiber.Ctx) error {
	coll := utils.GetDBCollection(os.Getenv("collectionName"))

	// find the book
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	item := models.Item{}

	err = coll.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&item)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{"data": item})
}

func searchBarcode(c *fiber.Ctx) error {
	barcode := c.Params("barcode")
	print("-----------------------------------------")
	i, err := strconv.ParseInt(barcode, 10, 64)
	coll := utils.GetDBCollection(os.Getenv("collectionName"))
	print(i,
		"+++++++++++++++++++", barcode)

	filter := bson.M{"sku.barcode": i}
	item := models.Search{}
	projection := bson.M{"name": 1, "thumbnail": 1, "image": 1, "sku": bson.M{"$elemMatch": bson.M{"barcode": i}}}

	err = coll.FindOne(c.Context(), filter, options.FindOne().SetProjection(projection)).Decode(&item)
	// print(item)
	if err != nil {
		log.Fatal(err)
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(200).JSON(fiber.Map{"data": item})
}

// type createDTO struct {
// 	Title  string `json:"title" bson:"title"`
// 	Author string `json:"author" bson:"author"`
// 	Year   string `json:"year" bson:"year"`
// }

func createItem(c *fiber.Ctx) error {
	// validate the body
	var errors []ErrorResponse
	print("----------------------------------")
	b := new(models.Create_item)
	err := c.BodyParser(b)
	print(b.Name, b.Thumbnail)
	// print(b.Skus[0])
	if err != nil {
		print(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err = Validator.Struct(b)
	if err != nil {
		// get all the validation errors as a slice of ErrorResponse structs
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, ErrorResponse{
				Error:       true,
				FailedField: e.Field(),
				Tag:         e.Tag(),
				Value:       e.Value(),
				message:     "Invalid Input",
			})
			print(errors, "-----------")
		}
		print(errors)
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// create the book
	coll := utils.GetDBCollection(os.Getenv("collectionName"))
	result, err := coll.InsertOne(c.Context(), b)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create book",
			"message": err.Error(),
		})
	}
	print(result)
	// return the book
	return c.Status(201).JSON(fiber.Map{
		"result": result,
	})
}

// type updateDTO struct {
// 	Title  string `json:"title,omitempty" bson:"title,omitempty"`
// 	Author string `json:"author,omitempty" bson:"author,omitempty"`
// 	Year   string `json:"year,omitempty" bson:"year,omitempty"`
// }

// func updateBook(c *fiber.Ctx) error {
// 	// validate the body
// 	b := new(updateDTO)
// 	if err := c.BodyParser(b); err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid body",
// 		})
// 	}

// 	// get the id
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid id",
// 		})
// 	}

// 	// update the book
// 	coll := common.GetDBCollection("books")
// 	result, err := coll.UpdateOne(c.Context(), bson.M{"_id": objectId}, bson.M{"$set": b})
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Failed to update book",
// 			"message": err.Error(),
// 		})
// 	}

// 	// return the book
// 	return c.Status(200).JSON(fiber.Map{
// 		"result": result,
// 	})
// }

// func deleteBook(c *fiber.Ctx) error {
// 	// get the id
// 	id := c.Params("id")
// 	if id == "" {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "id is required",
// 		})
// 	}
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "invalid id",
// 		})
// 	}

// 	// delete the book
// 	coll := common.GetDBCollection("books")
// 	result, err := coll.DeleteOne(c.Context(), bson.M{"_id": objectId})
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{
// 			"error":   "Failed to delete book",
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"result": result,
// 	})
// }
