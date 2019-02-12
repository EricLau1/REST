package controllers

import (
	"net/http"
	"api/utils"
	"api/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrInsufficientCash = errors.New("Dinheiro insuficiente")
	ErrInvalidWallet = errors.New("Carteira inválida")
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	transactions, err := models.GetTransactions()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusBadRequest})
		utils.CheckErr(err)
		return
	}
	err = utils.ToJson(w, utils.DefaultResponse{transactions, http.StatusOK})
	utils.CheckErr(err)
}

func PostTransaction(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	transaction, err := verifyTransaction(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return
	}
	message := fmt.Sprintf("%s transferiu: $ %.2f, para %s", transaction.Origin.User.Nickname,
	transaction.Cash, transaction.Target.User.Nickname)
	transaction.Message = message
	_, err = models.NewTransaction(transaction)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err = utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnprocessableEntity})
		utils.CheckErr(err)
		return
	}
	err = utils.ToJson(w, utils.DefaultResponse{"Transação concluída com sucesso!", http.StatusCreated })
	utils.CheckErr(err)
}

func verifyTransaction(r *http.Request) (models.Transaction, error) {
	params := mux.Vars(r)
	publicKeyTarget := params["public_key"]
	targetWallet, err := models.GetWalletByPublicKey(publicKeyTarget)
	if err != nil {
		return models.Transaction{}, err
	}
	body, _ := ioutil.ReadAll(r.Body)
	var originWallet models.Wallet
	err = json.Unmarshal(body, &originWallet)
	if err != nil {
		return models.Transaction{}, err
	}
	verifyOrigin, err := models.GetWalletByPublicKey(originWallet.PublicKey)
	if err != nil {
		return models.Transaction{}, err
	}
	if originWallet.Balance > verifyOrigin.Balance {
		return models.Transaction{}, ErrInsufficientCash
	}
	if targetWallet.PublicKey == "" || verifyOrigin.PublicKey == "" {
		return models.Transaction{}, ErrInvalidWallet	
	}
	return models.Transaction{Origin: verifyOrigin, Target: targetWallet, 
		Cash: originWallet.Balance}, nil
}