package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreatePaymentSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	users := models.User{}
	user, err := users.FindUserByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	data := models.PaymentSetting{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	data.UserID = user.ID
	dataCreated, err := data.Save(server.DB)
	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

//Get all Payment setting
func (server *Server) GetPaymentSetting(w http.ResponseWriter, r *http.Request) {
	data := models.PaymentSetting{}
	datas, err := data.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}

//Get payment by id
func (server *Server) GetPaymentSettingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.PaymentSetting{}
	dataReceived, err := data.Find(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}

func (server *Server) UpdatePaymentSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	data := models.PaymentSetting{}
	dataReceived, err := data.Find(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()
	dataCreated, err := dataReceived.Update(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusCreated, dataCreated)
}

func (server *Server) DeletePaymentSetting(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	aid := vars["id"]
	id, _ := strconv.ParseUint(aid, 10, 64)
	data := models.PaymentSetting{}
	dataReceived, err := data.Find(server.DB, id)
	if err != nil {
		return
	}
	if dataReceived == nil {
		return
	}
	_, err = dataReceived.Delete(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
