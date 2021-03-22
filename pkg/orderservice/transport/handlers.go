package transport

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

type Order struct {
	ID        uuid.UUID  `json:"id"`
	MenuItems []MenuItem `json:"menu_items"`
}

type OrderFullInfo struct {
	ID                 uuid.UUID  `json:"id"`
	MenuItems          []MenuItem `json:"menu_items"`
	OrderedAtTimestamp time.Time  `json:"ordered_at_timestamp"`
	Cost               int        `json:"cost"`
}

type MenuItem struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}

type CreateOrderRequest struct {
	MenuItems []MenuItem `json:"menu_items"`
}

type CreateOrderResponse struct {
	ID string `json:"id"`
}

func Router() http.Handler {
	router := mux.NewRouter()

	subRouter := router.PathPrefix("/api/v1").Subrouter()
	subRouter.Handle("/orders", ordersHandler).Methods(http.MethodGet)
	subRouter.Handle("/order/{id}", orderHandler).Methods(http.MethodGet)
	subRouter.Handle("/order", createOrderHandler).Methods(http.MethodPost)

	return logMiddleware(setContentTypeMiddleware("application/json", router))
}

var ordersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	rawResp, _ := json.Marshal([]Order{
		{
			ID: uuid.New(),
			MenuItems: []MenuItem{
				{ID: "0ff-23-21", Quantity: 3},
				{ID: "0ff-23-22", Quantity: 4},
				{ID: "0ff-23-23", Quantity: 10},
			},
		},
		{
			ID: uuid.New(),
			MenuItems: []MenuItem{
				{ID: "0ff-23-21", Quantity: 3},
				{ID: "0ff-23-22", Quantity: 4},
				{ID: "0ff-23-23", Quantity: 10},
			},
		},
		{
			ID: uuid.New(),
			MenuItems: []MenuItem{
				{ID: "0ff-23-21", Quantity: 3},
				{ID: "0ff-23-22", Quantity: 4},
				{ID: "0ff-23-23", Quantity: 10},
			},
		},
	})

	if _, err := w.Write(rawResp); err != nil {
		log.WithFields(log.Fields{"query": r.URL.Query()}).Error(err)
	}
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
		ID: id,
		MenuItems: []MenuItem{
			{
				ID:       "0ff-23-23",
				Quantity: 3,
			},
		},
		OrderedAtTimestamp: time.Now(),
		Cost:               999,
	}

	rawRespData, _ := json.Marshal(respData)
	if _, err := w.Write(rawRespData); err != nil {
		log.WithFields(log.Fields{"query": r.URL.Query(), "response": rawRespData}).Error(err)
	}
})

var createOrderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func() {
		if err := r.Body.Close(); err != nil {
			log.WithFields(log.Fields{"query": r.URL.Query()}).Error(err)
		}
	}()

	var msg CreateOrderRequest
	if err := json.Unmarshal(b, &msg); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	uid := uuid.New()
	response, err := json.Marshal(CreateOrderResponse{ID: uid.String()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
})
