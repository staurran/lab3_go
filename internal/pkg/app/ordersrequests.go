package app

import (
	"github.com/gin-gonic/gin"
	"lab3/internal/app/ds"
	"lab3/internal/app/utils/token"
	"net/http"
	"time"
)

type ReqStruct struct {
	Baskets []uint `json:"baskets"`
}

func (a *Application) AddOrder(gCtx *gin.Context) {
	var params ReqStruct

	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	date := time.Now().Format("31-02-2006")
	user_id, err := token.ExtractTokenID(gCtx)

	order := ds.Orders{Status: 1, Date: date, Id_user: user_id}
	err = a.repo.CreateOrder(&order)

	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create order"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}

	for _, id_basket := range params.Baskets {
		var orderGood ds.GoodOrder
		basket, err := a.repo.GetBasketById(id_basket)
		if err != nil {
			answer := AnswerJSON{Status: "error", Description: "cant add good in order"}
			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
			return
		}
		orderGood.Id_good = basket.Id_good
		orderGood.Id_order = order.Id_order
		orderGood.Quantity = basket.Quantity
		err = a.repo.CreateGoodOrder(&orderGood)
		if err != nil {

			answer := AnswerJSON{Status: "error", Description: "cant add good in order"}
			gCtx.IndentedJSON(http.StatusInternalServerError, answer)
			return
		}
	}

	answer := AnswerJSON{Status: "successful", Description: "good was added to basket"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}

type GoodsOrder struct {
	good     ds.Goods
	quantity int
}

type OrderRes struct {
	Id_order    uint
	Date        string
	Status      string
	Description string
	goods       []GoodsOrder
}

func (a *Application) GetAllOrders(gCtx *gin.Context) {
	id_user, err := token.ExtractTokenID(gCtx)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant extract user_id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	order, err := a.repo.GetOrder(id_user)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get rows in basket"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	var results []OrderRes
	for ind, ord := range order {
		results[ind].Date = ord.Date
		results[ind].Id_order = ord.Id_order
		results[ind].Status = ord.Status
		results[ind].Description = ord.Description
		results[ind].goods, err = a.repo.GetGoodOrder(id_order)
	}
	gCtx.IndentedJSON(http.StatusOK, order)
}
