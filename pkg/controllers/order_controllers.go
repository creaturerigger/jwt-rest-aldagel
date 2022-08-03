package controllers

import (
	"net/http"

	"strconv"

	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/models"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/tokenization"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/utils"
)

var Order models.Order

func AddOrder(w http.ResponseWriter, r *http.Request) {
	_, err := tokenization.ParseToken(r, "auth_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Error: Invalid token")
		return
	}
	order := &models.Order{}
	if err := utils.ParseBody(r, order); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	user, err := models.GetUser(strconv.FormatUint(uint64(order.UserID), 10))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	models.AddOrder(user, order)
	utils.RespondWithJSON(w, http.StatusOK, order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	_, err := tokenization.ParseToken(r, "auth_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Error: Invalid token")
		return
	}
	user, err := models.GetUser(r.Header.Get("UserID"))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	orders, err := models.GetOrders(user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, orders)
}
