package controllers

import (
	"net/http"
	"api/utils"
	"api/models"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	wallets, err := models.GetWallets()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusBadRequest})
		utils.CheckErr(err)
		return
	}
	err = utils.ToJson(w, utils.DefaultResponse{wallets, http.StatusOK})
	utils.CheckErr(err)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	publicKey := params["public_key"]
	wallet, err := models.GetWalletByPublicKey(publicKey)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusBadRequest})
		utils.CheckErr(err)
		return
	}
	err = utils.ToJson(w, utils.DefaultResponse{wallet, http.StatusOK})
	utils.CheckErr(err)
}

func PutWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	publicKey := params["public_key"]
	body, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return
	}
	wallet.PublicKey = publicKey
	rows, err := models.UpdateWallet(wallet)
	err = utils.ToJson(w, rows)
	utils.CheckErr(err)
}

func PutAddCashWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	params := mux.Vars(r)
	publicKey := params["public_key"]
	body, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return
	}
	wallet.PublicKey = publicKey
	rows, err := models.AddCashWallet(wallet)
	err = utils.ToJson(w, rows)
	utils.CheckErr(err)
}