package services

import (
	"context"
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/response"
	"ecommerce/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type OrderResponseStruct struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User           models.User        `bson:"user" json:"user"`
	Product        models.Product     `bson:"product" json:"product"`
	Status         bool               `bson:"status" json:"status"`
	Payment_status bool               `bson:"payment_status" json:"payment_status"`
}

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "orders")

func AddOrder(c *gin.Context) {

	var order models.Order
	var singleProduct models.Product

	if err := c.ShouldBindJSON(&order); err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	validatorErr := Validate.Struct(order)
	if validatorErr != nil {
		response.SendErrorResponse(c, 400, validatorErr.Error())
		return
	}

	// Update Product Quantity
	id, err := primitive.ObjectIDFromHex(order.Product)
	if err != nil {
		response.SendErrorResponse(c, 400, "product id is invalid")
		return
	}
	if id == primitive.NilObjectID {
		response.SendErrorResponse(c, 400, "product id is empty")
		return
	}

	// Chceking Product is there or not

	count, err := productCollection.CountDocuments(context.TODO(), bson.M{"_id": id})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "product not found")
		return
	}

	// get product data To Update The Quantity

	if err := productCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&singleProduct); err != nil {
		response.SendErrorResponse(c, 400, "product not found")
		return
	}

	quantity := singleProduct.Quantity - 1

	updateQuantity := bson.M{"quantity": quantity}

	// Update Product
	updateCount, err := productCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": updateQuantity})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		response.SendErrorResponse(c, 400, "product is not updated")
		return
	}

	// Add Data in orderCollection

	result, err := orderCollection.InsertOne(context.TODO(), order)
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	if result.InsertedID == "" {
		response.SendErrorResponse(c, 500, "something went wrong!!!")
		return
	}

	var OrderAsArray []models.Order
	if err := orderCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&order); err != nil {
		response.SendErrorResponse(c, 500, err.Error())
		return
	}

	OrderAsArray = append(OrderAsArray, *&order)

	response.SendSuccessResponse(c, 201, OrderAsArray)
}

func ViewOrder(c *gin.Context) {

	var orderArray []models.Order
	var responseOrder []OrderResponseStruct

	var product models.Product
	var user models.User
	var singleOrder models.Order

	cursor, err := orderCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		if err = cursor.Decode(&singleOrder); err != nil {
			response.SendErrorResponse(c, 400, err.Error())
			return
		}
		orderArray = append(orderArray, singleOrder)
	}

	for _, value := range orderArray {

		// convert id to objectId
		userId, err := utils.ConverIntoObject(value.User)
		if err != "" {
			response.SendErrorResponse(c, 400, err)
			return
		}
		// get User data
		if err := userCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user); err != nil {
			response.SendErrorResponse(c, 400, "user not found")
			return
		}

		// convert normal id to object id
		productId, err := utils.ConverIntoObject(value.Product)
		if err != "" {
			response.SendErrorResponse(c, 400, err)
			return
		}
		// get product data
		if err := productCollection.FindOne(context.TODO(), bson.M{"_id": productId}).Decode(&product); err != nil {
			response.SendErrorResponse(c, 400, "product not found")
			return
		}
		resp := OrderResponseStruct{
			Id:             value.Id,
			User:           user,
			Product:        product,
			Status:         value.Status,
			Payment_status: value.Payment_status,
		}
		responseOrder = append(responseOrder, resp)

	}

	response.SendSuccessResponse(c, 200, responseOrder)

}

func ViewOrderById(c *gin.Context) {

	var order models.Order
	var responseOrder []OrderResponseStruct

	var product models.Product
	var user models.User

	id, err := utils.ConverIntoObject(c.Param("id"))
	if err != "" {
		response.SendErrorResponse(c, 400, err)
		return
	}

	if err := orderCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&order); err != nil {
		utils.LogData(err.Error())
		response.SendErrorResponse(c, 400, "order not found")
		return
	}

	// convert id to objectId
	productId, err := utils.ConverIntoObject(order.Product)
	if err != "" {
		response.SendErrorResponse(c, 400, err)
		return
	}
	// get product data
	if err := productCollection.FindOne(context.TODO(), bson.M{"_id": productId}).Decode(&product); err != nil {
		response.SendErrorResponse(c, 400, "product not found")
		return
	}

	// convert id to objectId
	userId, err := utils.ConverIntoObject(order.User)
	if err != "" {
		response.SendErrorResponse(c, 400, err)
		return
	}
	// get User data
	if err := userCollection.FindOne(context.TODO(), bson.M{"_id": userId}).Decode(&user); err != nil {
		response.SendErrorResponse(c, 400, "user not found")
		return
	}
	res := OrderResponseStruct{
		Id:             order.Id,
		User:           user,
		Product:        product,
		Status:         order.Status,
		Payment_status: order.Payment_status,
	}

	responseOrder = append(responseOrder, res)

	response.SendSuccessResponse(c, 200, responseOrder)

}

func UpdateOrder(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}
	orderId, e := utils.ConverIntoObject(c.Param("id"))
	if e != "" {
		response.SendErrorResponse(c, 400, e)
		return
	}
	count, err := orderCollection.CountDocuments(context.TODO(), bson.M{"_id": orderId})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "order is not found")
		return
	}
	updateOrder := bson.M{"status": order.Status, "payment_status": order.Payment_status}
	updateCount, err := orderCollection.UpdateOne(context.TODO(), bson.M{"_id": orderId}, bson.M{"$set": updateOrder})

	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if updateCount.ModifiedCount == 0 {
		response.SendErrorResponse(c, 400, "order is not updated")
		return
	}

	response.SendSuccessResponse(c, 200, order, "message", "order updated successfully")

}

func DeleteOrder(c *gin.Context) {
	var order models.Order

	orderId, e := utils.ConverIntoObject(c.Param("id"))

	if e != "" {
		response.SendErrorResponse(c, 400, e)
		return
	}

	count, err := orderCollection.CountDocuments(context.TODO(), bson.M{"_id": orderId})
	if err != nil {
		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if count == 0 {
		response.SendErrorResponse(c, 400, "order is not found")
		return
	}

	deleteCount, err := orderCollection.DeleteOne(context.TODO(), bson.M{"_id": orderId})

	if err != nil {

		response.SendErrorResponse(c, 400, err.Error())
		return
	}

	if deleteCount.DeletedCount == 0 {
		response.SendErrorResponse(c, 400, "order is not deleted")
		return
	}

	response.SendSuccessResponse(c, 200, order, "message", "order deleted successfully")

}
