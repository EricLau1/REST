package controllers

import (
	"net/http"
	"api/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	err := utils.ToJson(w, utils.DefaultResponse{
		[]string{"Api OK :)"},
		http.StatusOK,
	})
	utils.CheckErr(err)
}

  