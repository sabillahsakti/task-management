package taskcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabillahsakti/task-management/config"
	"github.com/sabillahsakti/task-management/helper"
	"github.com/sabillahsakti/task-management/models"
	"gorm.io/gorm"
)

func GetByID(w http.ResponseWriter, r *http.Request) {
	//Mengambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.ResponseError(w, http.StatusNotFound, "Task not found")
			return
		}
		helper.ResponseError(w, http.StatusInternalServerError, "Error finding task")
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Success", task)

}

func GetByUser(w http.ResponseWriter, r *http.Request) {
	var task []models.Task

	// Ambil user_id dari context
	userID := r.Context().Value("user_id").(int)

	// Ambil query parameter untuk sorting
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	// Ambil parameter query untuk filtering
	status := r.URL.Query().Get("status")
	priority := r.URL.Query().Get("priority")

	// Untuk sorting
	if sortBy == "" {
		sortBy = "due_date"
	}
	if order == "" {
		order = "desc"
	}

	// Untuk filter
	query := config.DB.Where("user_id = ?", userID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if priority != "" {
		query = query.Where("priority = ?", priority)
	}

	if order != "asc" && order != "desc" {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid order parameter, must be 'asc' or 'desc'")
		return
	}

	if err := query.Order(sortBy + " " + order).Find(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.ResponseError(w, http.StatusNotFound, "Task not found")
			return
		}
		helper.ResponseError(w, http.StatusInternalServerError, "Error finding tasks")
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Success", task)

}

func Create(w http.ResponseWriter, r *http.Request) {
	//Mengambil inputan json dari body

	var task models.Task
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	// Ambil user_id dari context
	userID := r.Context().Value("user_id").(int)
	task.UserID = int64(userID)

	// Insert ke database
	if err := config.DB.Create(&task).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Input Success", task)
}

func Update(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	//Amil dari body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	//Mengambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	// Cek apakah task dengan ID tersebut ada
	var existingTask models.Task
	if err := config.DB.First(&existingTask, id).Error; err != nil {
		helper.ResponseError(w, http.StatusNotFound, "Task not found")
		return
	}

	// Memperbarui task
	if err := config.DB.Model(&existingTask).Updates(task).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, "Error updating task")
		return
	}

	helper.ResponseJson(w, http.StatusOK, "Task updated successfully", existingTask)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	//Mengambil ID dari URL
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var task models.Task
	if err := config.DB.First(&task, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.ResponseError(w, http.StatusNotFound, "Task not found")
			return
		}
		helper.ResponseError(w, http.StatusInternalServerError, "Error finding task")
		return
	}

	// Hapus dari database
	if err := config.DB.Delete(&task).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, "Error deleting task")
		return
	}

	// Kirim respons berhasil
	helper.ResponseJson(w, http.StatusOK, "Task deleted successfully", nil)

}
