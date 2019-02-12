package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"api/utils"
	"api/models"
	"api/validation"
	"api/services/email"
	"strconv"
	"github.com/gorilla/mux"
)

func registerMessage(nickname string) string {
	nickname = utils.B64Encode(nickname)
	return `
		<p> Clique no link abaixo para ativar sua conta! </p>
		<a href="http://localhost:3000/api/account/` + nickname +`"> Clique Aqui</a>
		<p><strong> Atenção! </strong> Não envie esse link pra ninguém. </p>
	`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	users, err := models.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusBadRequest})
		utils.CheckErr(err)
		return
	}
	err = utils.ToJson(w, utils.DefaultResponse{users, http.StatusOK})
	utils.CheckErr(err)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return	
	}
	user, err = validation.VerifyPostUser(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return	
	}
	_, err = models.NewUser(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return		
	}
	email.Send(email.Email{ "ericlau.oliveira@gmail.com", "Confirmação de Cadastro", registerMessage(user.Nickname)})
	err = utils.ToJson(w, utils.DefaultResponse{"Ative sua conta no email: " + user.Email, http.StatusCreated})	
	utils.CheckErr(err)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	user, err := models.GetUserById(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusBadRequest})
		utils.CheckErr(err)
		return		
	}
	if user.UUID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{"Usuário não existe", http.StatusBadRequest})
		utils.CheckErr(err)
		return		
	}
	err = utils.ToJson(w, utils.DefaultResponse{user, http.StatusOK})
	utils.CheckErr(err)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	body, _ := ioutil.ReadAll(r.Body)
	user := models.User{UUID: uint32(id)}
	err := json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		return
	}
	rows, err := models.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		return
	}
	err = utils.ToJson(w, rows)
	utils.CheckErr(err)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	rows, err := models.DeleteUser(uint32(id))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		return
	}
	err = utils.ToJson(w, rows)
}