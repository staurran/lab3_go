package app

import (
	"github.com/gin-gonic/gin"
	"lab3/internal/app/ds"
	"log"
	"net/http"
	"strconv"
)

func (a *Application) StartServer() {
	log.Println("Server start up")
	log.Println("Server start up")
	r := gin.Default()

	r.GET("/goods", a.GetAll)
	r.GET("/goods/:id", a.GetProduct)
	r.POST("/goods", a.PostProduct)
	r.PUT("/goods/:id", a.ChangePrice)
	r.DELETE("goods/:id", a.DeleteProduct)
	r.GET("/ping/:name", a.Ping)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println("Run failed")
	}
	log.Println("Server down")
}

type AnswerJSON struct {
	Status      string `json:"status"`
	Description string `json:"description"`
}

type pingResp struct {
	Status string `json:"status"`
}

// Ping godoc
// @Summary      Show hello text
// @Description  very very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  string
// @Router       /ping/{name} [get]
func (a *Application) Ping(gCtx *gin.Context) {
	name := gCtx.Param("name")
	gCtx.String(http.StatusOK, "Hello %s", name)
}

//type GoodsListResponse = []ds.Goods

// GetAll godoc
// @Summary      Show all rows in db
// @Description  Return all product and info about rows
// @Tags         Tests
// @Produce      json
// @Success      200  {object} []ds.Goods
// @Router       /goods [get]
func (a *Application) GetAll(gCtx *gin.Context) {
	all_rows, err := a.repo.GetAllProducts()
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusInternalServerError, all_rows)

}

// GetProduct godoc
// @Summary      Show product info by id
// @Description  Return all info of one product by id
//@Parameters	id
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  ds.Goods
// @Router       /goods/{id} [get]
func (a *Application) GetProduct(gCtx *gin.Context) {
	id_product := gCtx.Param("id")
	id_product_int, err := strconv.Atoi(id_product)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	product, err := a.repo.GetProductByID(uint(id_product_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, &product)
}

// ChangePrice godoc
// @Summary      Change price of product by id
// @Description  Change price of product by id. Price can't be 0
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  ds.Goods
// @Router       /goods/{id} [put]
func (a *Application) ChangePrice(gCtx *gin.Context) {
	var params ds.Goods
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	if params.Price == 0 {
		answer := AnswerJSON{Status: "error", Description: "product cant cost 0"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	id_product := gCtx.Param("id")
	id_product_int, err := strconv.Atoi(id_product)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}

	err = a.repo.ChangeProduct(uint(id_product_int), params.Price)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant change price"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	product, err := a.repo.GetProductByID(uint(id_product_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, &product)
}

// PostProduct godoc
// @Summary      Add new row
// @Description  add new row with parameters in json
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  ds.Goods
// @Router       /goods [post]
func (a *Application) PostProduct(gCtx *gin.Context) {
	var params ds.Goods
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	err = a.repo.CreateProduct(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, params)
}

// DeleteProduct godoc
// @Summary      Delete row by id
// @Description  Delete row by id. If there is not this id return error
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  AnswerJSON
// @Router       /goods/{id} [delete]
func (a *Application) DeleteProduct(gCtx *gin.Context) {
	id_product := gCtx.Param("id")
	id_product_int, err := strconv.Atoi(id_product)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	err = a.repo.DeleteProduct(uint(id_product_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant delete row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	answer := AnswerJSON{Status: "successful", Description: "row was deleted"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}
