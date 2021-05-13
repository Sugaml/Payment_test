package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/pkg/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (server *Server) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	firstday := r.FormValue("datestart")
	lastday := r.FormValue("dateend")
	t := time.Now()
	var datestart time.Time
	var dateend time.Time
	if firstday == "" {
		datestart = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)

	} else {
		datestart, err = time.Parse("2006-01-02", firstday)
	}
	if lastday == "" {
		dateend = time.Now()
	} else {
		dateend, err = time.Parse("2006-01-02", lastday)
	}
	project := models.Project{}
	projectReceived, err := project.FindAllByUser(server.DB, uint(pid), datestart, dateend)
	var total int
	invoiceItems := []models.InvoiceItems{}
	for _, project := range projectReceived {
		createdyear, createdmonth, _ := project.CreatedAt.Date()
		yearnow, monthnow, _ := dateend.Date()
		subscription := models.Subscription{}
		subscriptionReceived, err := subscription.Find(server.DB, project.SubscriptionID)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		projectPrice := int(subscriptionReceived.Price)
		projectStartDate := datestart.Day()
		projectEndDate := dateend.Day()
		if projectStartDate < project.CreatedAt.Day() && createdyear == yearnow && createdmonth == monthnow {
			if project.DeletedAt != nil && projectEndDate > project.DeletedAt.Day() {
				projectStartDate = project.CreatedAt.Day()
				projectEndDate = project.DeletedAt.Day()
			} else {
				projectStartDate = project.CreatedAt.Day()
			}
		}
		effectiveTime := projectEndDate - projectStartDate + 1
		cost := projectPrice / 30 * effectiveTime
		invoiceItem := models.InvoiceItems{}
		invoiceItem.UserID = uint(pid)
		invoiceItem.Particular = project.Name
		invoiceItem.Days = uint(effectiveTime)
		invoiceItem.Rate = uint(projectPrice)
		invoiceItem.Total = uint(cost)
		invoiceItems = append(invoiceItems, invoiceItem)
		total = total + cost

	}
	invoice := models.Invoice{}
	invoice.UserID = uint(pid)
	invoice.TotalCost = uint(total)
	invoice.StartDate = datestart
	invoice.EndDate = dateend
	responses.JSON(w, http.StatusOK, invoice)
	isMonth := true
	if isMonth == true {
		invoiceSaved, err := invoice.Save(server.DB)
		if err != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
		for _, invoiceItem := range invoiceItems {
			invoiceItem.InvoiceID = invoiceSaved.ID
			_, error := invoiceItem.Save(server.DB)
			if error != nil {
				responses.ERROR(w, http.StatusNotFound, err)
				return
			}
		}
		paymentHistory := models.PaymentHistory{}
		paymentHistory.Credit = invoiceSaved.TotalCost
		paymentHistory.InvoiceID = invoiceSaved.ID
		paymentHistory.Balance = invoiceSaved.TotalCost
		paymentHistory.UserID = invoiceSaved.UserID
		_, error := paymentHistory.Save(server.DB)
		if error != nil {
			responses.ERROR(w, http.StatusNotFound, err)
			return
		}
	}

}
func (server *Server) GetInvoice(w http.ResponseWriter, r *http.Request) {
	data := models.Invoice{}
	datas, err := data.FindAll(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, datas)
}

func (server *Server) GetInvoiceById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Invoice{}
	dataReceived, err := data.Find(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}
	responses.JSON(w, http.StatusOK, dataReceived)
}
