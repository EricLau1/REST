package auth

import (
	"api/models"
	"api/utils"
	"time"
	"errors"
	"fmt"
	"strings"
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	ErrEmailNotFound = errors.New("Email não encontrado")
	ErrInvalidPassword = errors.New("Senha inválida")
)

var secretKey = []byte("S3CR3TK3Y")

type Auth struct {
	User    models.User `json:"user"`
	Token   string `json:"token"`
	IsValid bool `json:"is_valid"`
}

func SignIn(user models.User) (models.User, string, error) {
	password := user.Password
	user, err := models.GetUserByEmail(user.Email)
	if err != nil || user.UUID == 0 {
		return models.User{},"", ErrEmailNotFound
	}
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return models.User{},"", ErrInvalidPassword
	}
	token, err := GenerateJWT(user)
	return user, token, err
}

func GenerateJWT(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = user.Nickname
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token)(interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Oops! Ocorrou um erro. :(")
				}
				return secretKey, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnauthorized})
				return
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Não Autorizado")	
		}
	})
}

func IsAuthorizedByBearerToken(endpoint func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token)(interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Oops! Ocorrou um erro. :(")
					}
					return secretKey, nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					utils.ToJson(w, utils.DefaultError{err.Error(), http.StatusUnauthorized})
					return
				}
				if token.Valid {
					endpoint(w, r)
					return
				}
			}
		} 
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Não Autorizado")	
	})
}