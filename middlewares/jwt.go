package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sabillahsakti/task-management/config"
	"github.com/sabillahsakti/task-management/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			helper.ResponseError(w, http.StatusUnauthorized, "Unauthorize")
			return
		}

		// Menghapus prefix "Bearer " jika ada
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		} else {
			helper.ResponseError(w, http.StatusUnauthorized, "Unauthorize")
			return
		}

		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			log.Println("Error parsing token:", err)
			helper.ResponseError(w, http.StatusUnauthorized, "Unauthorize")
			return
		}

		if !token.Valid {
			helper.ResponseError(w, http.StatusUnauthorized, "Unauthorize")
			return
		}

		// Asumsikan claims.UserID adalah string, maka kita konversi ke int
		userID, err := strconv.Atoi(claims.ID)
		if err != nil {
			helper.ResponseError(w, http.StatusBadRequest, "Invalid user id")
			return
		}
		ctx := context.WithValue(r.Context(), "user_id", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
