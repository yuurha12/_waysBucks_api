package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	orderdto "waysbucks_BE/dto/order"
	dto "waysbucks_BE/dto/result"
	"waysbucks_BE/models"
	"waysbucks_BE/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerOrder struct {
	OrderRepository repositories.OrderRepository
}

func HandlerOrder(OrderRepository repositories.OrderRepository) *handlerOrder {
	return &handlerOrder{OrderRepository}
}

//show all order
func (h *handlerOrder) FindOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	orders, err := h.OrderRepository.FindOrders()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

//get order by id
func (h *handlerOrder) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: order}
	json.NewEncoder(w).Encode(response)
}

//createorder
func (h *handlerOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(orderdto.CreateOrder)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	buyerID := int(userInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	product, err := h.OrderRepository.GetProductOrder(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Product Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	topings, err := h.OrderRepository.GetToppingOrder(request.ToppingID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "Toping Not Found!"}
		json.NewEncoder(w).Encode(response)
		return
	}

	var priceTopings = 0
	for _, i := range topings {
		priceTopings += i.Price
	}
	var subTotal = request.Qty * (product.Price + priceTopings)

	order := models.Order{
		Qty:       request.Qty,
		BuyerID:   buyerID,
		ProductID: product.ID,
		Topping:   topings,
		Price:     subTotal,
	}

	data, err := h.OrderRepository.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	orders, _ := h.OrderRepository.GetOrder(data.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: orders}
	json.NewEncoder(w).Encode(response)
}

//update
// func (h *handlerOrder) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	request := new(orderdto.UpdateOrder)
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	order, err := h.OrderRepository.GetOrder(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	if request.Fullname != "" {
// 		order.Fullname = request.Fullname
// 	}

// 	if request.Email != "" {
// 		order.Email = request.Email
// 	}

// 	if request.Password != "" {
// 		order.Password = request.Password
// 	}

// 	if request.Image != "" {
// 		order.Image = request.Image
// 	}

// 	data, err := h.OrderRepository.UpdateOrder(order)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Status: "success", Data: data}
// 	json.NewEncoder(w).Encode(response)
// }

//delete order
func (h *handlerOrder) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	order, err := h.OrderRepository.GetOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.OrderRepository.DeleteOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertResponseOrder(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseOrder(u models.Order) orderdto.OrderResponse {
	return orderdto.OrderResponse{
		ID: u.ID,
	}
}
