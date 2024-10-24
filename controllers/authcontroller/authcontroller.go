package authcontroller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sabillahsakti/task-management/config"
	"github.com/sabillahsakti/task-management/helper"
	"github.com/sabillahsakti/task-management/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json dari front end
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	// hash pass menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := config.DB.Create(&userInput).Error; err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helper.ResponseJson(w, http.StatusOK, "User registered successfully", userInput)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json dari front end
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()

	//ambil data user berdasarkan uname di database
	var user models.User
	if err := config.DB.Where("username =?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusBadRequest, "Username or password incorrect")
			return
		default:
			helper.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	//Pengecekan password valid atau tidak
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		helper.ResponseError(w, http.StatusBadRequest, "Username or password incorrect")
		return
	}

	// Proses pembuatan token JWT
	expTime := time.Now().Add(time.Minute * 15)
	claims := &config.JWTClaim{
		Username: user.Username,
		ID:       strconv.Itoa(int(user.ID)),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "evoting",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	//Mendeklarasikan algoritma yang akan digunakan untuk signin
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		helper.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	response := map[string]string{
		"token": token, // tambahkan token di response JSON
	}
	helper.ResponseJson(w, http.StatusOK, "Login successfully", response)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Hapus token yg ada di cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	helper.ResponseJson(w, http.StatusOK, "User logout successfully", nil)
	return
}
