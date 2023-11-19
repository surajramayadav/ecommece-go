package services

import (
	"context"
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func AddProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	validatorErr := Validate.Struct(product)
	if validatorErr != nil {
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}
	count, err := productCollection.CountDocuments(context.TODO(), bson.M{"name": product.Name})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count > 0 {
		response.SendErrorResponse(c, 400, "Product already exists")
		return
	}

	result, err := productCollection.InsertOne(context.TODO(), product)
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	if result.InsertedID == "" {
		response.SendErrorResponse(c, 500, "something went wrong!!!")
		return
	}

	var ProductAsArray []models.Product
	if err := productCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&product); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	ProductAsArray = append(ProductAsArray, *&product)

	response.SendSuccessResponse(c, 201, ProductAsArray)

}

func ViewProduct(c *gin.Context) {
	var product []models.Product

	cursor, err := productCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var singleProduct models.Product
		if err = cursor.Decode(&singleProduct); err != nil {
			response.SendErrorResponse(c, 400, err.Error())
			return
		}
		product = append(product, singleProduct)
	}

	response.SendSuccessResponse(c, 200, product)
}

func ViewProductById(c *gin.Context) {

	var product []models.Product
	var singleProduct models.Product

	id, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		response.SendErrorResponse(c, 400, e)
		return
	}

	if err := productCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&singleProduct); err != nil {
		utils.LogData(err.Error())
		response.SendErrorResponse(c, 400, "product is not found")
		return
	}
	product = append(product, singleProduct)

	response.SendSuccessResponse(c, 200, product)

}

func UpdateProduct(c *gin.Context) {
	var product []models.Product
	var singleProduct models.Product

	if err := c.BindJSON(&singleProduct); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
	}

	id, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		response.SendErrorResponse(c, 400, e)
		return
	}

	count, err := productCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "product is not found")
		return
	}

	updateCount, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": singleProduct})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		response.SendErrorResponse(c, 400, "product is not updated")
		return
	}

	response.SendSuccessResponse(c, 200, product, "message", "product updated successfully")

}

func DeleteProduct(c *gin.Context) {
	var product []models.Product

	id, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		response.SendErrorResponse(c, 400, e)
		return
	}
	count, err := productCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "product is not found")
		return
	}

	deleteCount, err := productCollection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err != nil {
		utils.LogData("18" + err.Error())

		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if deleteCount.DeletedCount == 0 {
		response.SendErrorResponse(c, 400, "product is not deleted")
		return
	}

	response.SendSuccessResponse(c, 200, product, "message", "product deleted successfully")

}
