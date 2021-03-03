package transport

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Order struct {
	Id        string     `json:"id"`
	MenuItems []MenuItem `json:"menu_items"`
}

type OrderFullInfo struct {
	Id                 string     `json:"id"`
	MenuItems          []MenuItem `json:"menu_items"`
	OrderedAtTimestamp int        `json:"ordered_at_timestamp"`
	Cost               int        `json:"cost"`
}

type MenuItem struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

func Router() http.Handler {
	router := mux.NewRouter()

	router.Handle("/api/v1/orders", setContentTypeMiddleware("application/json", ordersHandler))
	router.Handle("/api/v1/order/{id}", setContentTypeMiddleware("application/json", orderHandler))

	return logMiddleware(router)
}

var ordersHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	rawResp, _ := json.Marshal([]Order{
		{
			Id: "0ff-123-3f1",
			MenuItems: []MenuItem{
				{Id: "0ff-23-21", Quantity: 3},
				{Id: "0ff-23-22", Quantity: 4},
				{Id: "0ff-23-23", Quantity: 10},
			},
		},
		{
			Id: "0ff-123-3f2",
			MenuItems: []MenuItem{
				{Id: "0ff-23-21", Quantity: 3},
				{Id: "0ff-23-22", Quantity: 4},
				{Id: "0ff-23-23", Quantity: 10},
			},
		},
		{
			Id: "0ff-123-3f3",
			MenuItems: []MenuItem{
				{Id: "0ff-23-21", Quantity: 3},
				{Id: "0ff-23-22", Quantity: 4},
				{Id: "0ff-23-23", Quantity: 10},
			},
		},
	})
	w.Write(rawResp)
})

var orderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	id, ok := pathParams["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	respData := OrderFullInfo{
		Id: id,
		MenuItems: []MenuItem{
			{
				Id:       "0ff-23-23",
				Quantity: 3,
			},
		},
		OrderedAtTimestamp: 12345234,
		Cost:               999,
	}

	rawRespData, _ := json.Marshal(respData)
	w.Write(rawRespData)
})
