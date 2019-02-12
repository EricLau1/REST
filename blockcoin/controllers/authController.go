package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"api/auth"
	"api/utils"
	"api/models"
)

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnauthorized})
		utils.CheckErr(err)
		return
	}
	user, token, err := auth.SignIn(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnauthorized})
		utils.CheckErr(err)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{auth.Auth{user,token,true}, http.StatusOK})
}