package transport

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

type Order struct {
	Id        uuid.UUID  `json:"id"`
	MenuItems []MenuItem `json:"menu_items"`
}

type OrderFullInfo struct {
	Id                 uuid.UUID  `json:"id"`
	MenuItems          []MenuItem `json:"menu_items"`
	OrderedAtTimestamp time.Time  `json:"ordered_at_timestamp"`
	Cost               int        `json:"cost"`
}

type MenuItem struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type CreateOrderRequest struct {
	MenuItems []MenuItem `json:"menu_items"`
}

type CreateOrderResponse struct {
	Id string `json:"id"`
}

func Router() http.Handler {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.Handle("/orders", ordersHandler).Methods(http.MethodGet)
	subRouter.Handle("/order/{id}", orderHandler).Methods(http.MethodGet)
	subRouter.Handle("/order", createOrderHandler).Methods(http.MethodPost)

	return logMiddleware(setContentTypeMiddleware("application/json", router))
}

var ordersHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	rawResp, _ := json.Marshal([]Order{
		{
			Id: uuid.New(),
			MenuItems: []MenuItem{
				{Id: "0ff-23-21", Quantity: 3},
				{Id: "0ff-23-22", Quantity: 4},
				{Id: "0ff-23-23", Quantity: 10},
			},
		},
		{
			Id: uuid.New(),
			MenuItems: []MenuItem{
				{Id: "0ff-23-21", Quantity: 3},
				{Id: "0ff-23-22", Quantity: 4},
				{Id: "0ff-23-23", Quantity: 10},
			},
		},
		{
			Id: uuid.New(),
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

	unparsedId, ok := pathParams["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, err := uuid.Parse(unparsedId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
		OrderedAtTimestamp: time.Now(),
		Cost:               999,
	}

	rawRespData, _ := json.Marshal(respData)
	w.Write(rawRespData)
})

var createOrderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var msg CreateOrderRequest
	if err := json.Unmarshal(b, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	uid := uuid.New()
	response, err := json.Marshal(CreateOrderResponse{Id: uid.String()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})
