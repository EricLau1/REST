package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"api/utils"
	"api/models"
)

func ConfirmAccount(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	nickname, err := utils.B64Decode(params["nickname"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJson(w, utils.DefaultError{"Url inválida", http.StatusBadRequest})
		return
	}
	user, err := models.GetUserByNickname(nickname)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJson(w, utils.DefaultError{"Nickname não existe", http.StatusBadRequest})
		return
	}
	_, err = models.ConfirmAccount(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.ToJson(w, utils.DefaultError{"Conta não foi confirmada.", http.StatusBadRequest})
		return
	}
	utils.ToJson(w, utils.DefaultResponse{
		user.Nickname + ", seu email foi confirmado com sucesso! :)",
		http.StatusOK})
}