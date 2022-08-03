package controllers

import (
	"net/http"
	"time"

	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/models"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/tokenization"
	"github.com/jawohlCodeTeam/jwt-rest-aldagel/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var NewUser models.User

func Login(w http.ResponseWriter, r *http.Request) {
	creds := &models.UserCredentials{}
	if err := utils.ParseBody(r, creds); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	claims, err := models.Login(creds)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	tokenString, err := tokenization.GenerateToken(claims.StandardClaims)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	cookie := tokenization.GenerateCookie(tokenString, expiresAt)
	http.SetCookie(w, cookie)
	utils.RespondWithJSON(w, http.StatusOK, utils.GenerateResponseMap("Success: Login successful"))
}

func Register(w http.ResponseWriter, r *http.Request) {
	NewUser := &models.User{}
	if err := utils.ParseBody(r, NewUser); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(NewUser.Password), 14)
	NewUser.Password = string(encryptedPassword)
	user, err := models.Register(NewUser)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, user)
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	_, err := tokenization.ParseToken(r, "auth_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Error: Invalid token")
		return
	}
	user := &models.User{}
	if err := utils.ParseBody(r, user); err != nil {
		errorMessage := "Error: " + err.Error()
		response := map[string]string{"message": errorMessage}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	friends, err := models.GetFriends(user)
	if err != nil {
		errorMessage := "Error: " + err.Error()
		response := map[string]string{"message": errorMessage}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, friends)
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	_, err := tokenization.ParseToken(r, "auth_token")
	if err != nil {
		utils.RespondWithError(w, http.StatusForbidden, "Error: Invalid token")
		return
	}
	userSlice := &[]models.User{}
	if err := utils.ParseBody(r, userSlice); err != nil {
		errorMessage := "Error: " + err.Error()
		response := map[string]string{"message": errorMessage}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}
	if userAndFriend := len(*userSlice); userAndFriend == 2 {
		err := models.AddFriend(&(*userSlice)[0], &(*userSlice)[1])
		if err != nil {
			utils.RespondWithJSON(w, http.StatusBadRequest, utils.GenerateResponseMap("Error: "+err.Error()))
			return
		}
		userName := (*userSlice)[1].FullName
		response2 := utils.GenerateResponseMap("Success: Friend added: " + userName)
		utils.RespondWithJSON(w, http.StatusOK, response2)
		return
	} else {
		response := utils.GenerateResponseMap("Error: Not valid data for adding friend")
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}
