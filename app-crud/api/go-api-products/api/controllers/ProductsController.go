package controllers

import (
	"net/http"
	"encoding/json"
	"strconv"
	"go-api-products/api/utils"
	"go-api-products/api/models"
	"github.com/gorilla/mux"
)

func PostProduct(w http.ResponseWriter, r *http.Request) {
	body := utils.BodyParser(r)
	var product models.Product 
	err := json.Unmarshal(body, &product)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	rs, err := models.NewProduct(product)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rs, http.StatusCreated)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := models.GetProducts()
	utils.ToJson(w, products, http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	product := models.GetProductById(id)
	if product.Id == 0 {
		utils.ToJson(w, "Product not found", http.StatusBadRequest)
		return
	}
	utils.ToJson(w, product, http.StatusOK)
}

func PutProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	body := utils.BodyParser(r)
	var product models.Product 
	err := json.Unmarshal(body, &product)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	product.Id = id
	rs, err := models.UpdateProduct(product)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rs, http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseUint(vars["id"], 10, 64)
	rows, err := models.DeleteProduct(id)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusOK)
}