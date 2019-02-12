package routes

import (
	"github.com/gorilla/mux"
	"api/controllers"
	"api/auth"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/users", auth.IsAuthorizedByBearerToken(controllers.GetUsers)).Methods("GET")
	r.HandleFunc("/api/users", controllers.PostUser).Methods("POST")
	r.HandleFunc("/api/users/{uid}", auth.IsAuthorizedByBearerToken(controllers.GetUser)).Methods("GET")
	r.HandleFunc("/api/users/{uid}", auth.IsAuthorizedByBearerToken(controllers.PutUser)).Methods("PUT")
	r.HandleFunc("/api/users/{uid}", auth.IsAuthorizedByBearerToken(controllers.DeleteUser)).Methods("DELETE")
	r.HandleFunc("/api/account/{nickname}", controllers.ConfirmAccount).Methods("GET")
	r.HandleFunc("/api/login", controllers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/api/wallets", controllers.GetWallets).Methods("GET")
	r.HandleFunc("/api/wallets/{public_key}", controllers.GetWallet).Methods("GET")
	r.HandleFunc("/api/wallets/{public_key}", auth.IsAuthorizedByBearerToken(controllers.PutWallet)).Methods("PUT")
	r.HandleFunc("/api/wallets/add/{public_key}", auth.IsAuthorizedByBearerToken(controllers.PutAddCashWallet)).Methods("PUT")
	r.HandleFunc("/api/transactions/{public_key}", auth.IsAuthorizedByBearerToken(controllers.PostTransaction)).Methods("POST")
	r.HandleFunc("/api/transactions", controllers.GetTransactions).Methods("GET")
	return r
}