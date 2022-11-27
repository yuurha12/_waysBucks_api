package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysbucks_BE/dto/result"
	toppingdto "waysbucks_BE/dto/topping"
	"waysbucks_BE/models"
	"waysbucks_BE/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerTopping struct {
	ToppingRepository repositories.ToppingRepository
}

// Create `path_file` Global variable here ...
var topping_file = os.Getenv("PATH_FILE")

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *handlerTopping {
	return &handlerTopping{ToppingRepository}
}

func (h *handlerTopping) FindToppings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	toppings, err := h.ToppingRepository.FindToppings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	for i, p := range toppings {
		toppings[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: toppings}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) GetTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var topping models.Topping
	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	topping.Image = os.Getenv("PATH_FILE") + topping.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertResponseTopping(topping)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//admin token condition
	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not ADMIN"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	// qty, _ := strconv.Atoi(r.FormValue("qty"))
	request := toppingdto.ToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
		// Qty:   qty,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping := models.Topping{
		Title: request.Title,
		Price: request.Price,
		Image: filename,
		// Qty:    request.Qty,
	}

	topping, err = h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ = h.ToppingRepository.GetTopping(topping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) DeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not ADMIN"}
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	_, err = h.ToppingRepository.DeleteTopping(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data := user.ID

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerTopping) UpdateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userRole := userInfo["role"]
	userId := int(userInfo["id"].(float64))

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	request := toppingdto.UpdateTopping{
		Title: r.FormValue("title"),
		Price: price,
		Qty:   qty,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	//admin token condition
	if userId != id && userRole != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "not ADMIN"}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, err := h.ToppingRepository.GetTopping(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		topping.Title = request.Title
	}

	if request.Price != 0 {
		topping.Price = request.Price
	}
	if filename != "false" {
		topping.Image = filename
	}

	if qty != 0 {
		topping.Qty = request.Qty
	}

	data, err := h.ToppingRepository.UpdateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: convertResponseTopping(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTopping(u models.Topping) toppingdto.ToppingResponse {
	return toppingdto.ToppingResponse{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
		// Qty:   u.Qty,
		// User:     u.User,
	}
}
