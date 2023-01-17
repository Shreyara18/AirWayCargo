package handlers

import (
	"Air_Way_Cargo/internal/daos"
	"Air_Way_Cargo/internal/daos/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wiz-freight-org/adapters/utils"

	"github.com/wiz-freight-org/adapters/utils/db"
)

func SetApiRoutes(v1 *gin.RouterGroup) {
	v1.POST("/v1/create-user", CreateUser)
	v1.POST("/v1/login", Login)
	v1.GET("/v1/account/:aid/get-quote", GetQuote)
	v1.POST("/v1/account/:aid/create-quote/:qid", CreateQuote)
	v1.GET("/v1/account/:aid/get-quote/:qid", ListQuote)
	v1.GET("/v1/account/:aid/get-booking ", GetBooking)
	v1.PUT("/v1/account/:aid/create-booking/:qid", CreateBooking)
	v1.GET("/v1/account/:aid/get-booking/:qid", ListBooking)
}

func CreateUser(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Config{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	err = n.NewCreateUser(dto_reqs)
	c.JSON(http.StatusCreated, "")

}
func Login(c *gin.Context) {

	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Config{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	Login, err := n.NewLogin(dto_reqs.Email, dto_reqs.Password)
	c.JSON(http.StatusCreated, Login)

}

func GetQuote(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Quotes{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	getquote, err := n.NewGetQuote(dto_reqs.ConfigId)
	c.JSON(http.StatusCreated, getquote)

}
func CreateQuote(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Quotes{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	err = n.NewCreateQuote(dto_reqs)
	c.JSON(http.StatusCreated, "")

}
func ListQuote(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Quotes{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	listquote, err := n.NewListQuote(dto_reqs.ID)
	c.JSON(http.StatusCreated, listquote)
}
func GetBooking(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Booking{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	getbooking, err := n.NewGetBooking(dto_reqs.ConfigId)
	c.JSON(http.StatusCreated, getbooking)
}
func CreateBooking(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Booking{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	err = n.NewCreateBooking(dto_reqs)
	c.JSON(http.StatusCreated, "")

}
func ListBooking(c *gin.Context) {
	dbConn := db.New(&utils.Logger{})
	decoder := json.NewDecoder(c.Request.Body)
	dto_reqs := &models.Task{}
	err := decoder.Decode(dto_reqs)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"err": "Inavlid Decode request ",
		})

	}
	n := daos.NewDB(dbConn)
	listbooking, err := n.NewListBooking(dto_reqs.ID)
	c.JSON(http.StatusCreated, listbooking)
}
