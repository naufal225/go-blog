package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/naufal225/go-blog/config"
	"github.com/naufal225/go-blog/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := mux.Vars(r)["id"]

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		http.Error(w, "Nama dan Email harus diisi", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Gagal menyimpan data ke database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	id := mux.Vars(r)["id"]

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Format JSON tidak valid", http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Email == "" {
		http.Error(w, "Nama dan Email harus diisi", http.StatusBadRequest)
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(w, "Gagal memperbarui data ke database", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result := config.DB.Delete(&models.User{}, id)

	if err := result.Error; err != nil {
		http.Error(w, "Gagal menghapus data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message":"Berhasil menghapus data"})
}