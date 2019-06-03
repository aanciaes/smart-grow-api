package handlers

import (
	"crypto/rsa"
	"fmt"
	"github.com/aanciaes/smart-grow-api/model"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func readPublicKey() (*rsa.PublicKey, error) {
	publicKeyBytes, err := ioutil.ReadFile("keys/app.rsa.pub")
	if err != nil {
		return nil, fmt.Errorf("error reading public key")
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes);
	if err != nil {
		return nil, fmt.Errorf("error reading private key")
	}

	return publicKey, nil
}

func readPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBytes, err := ioutil.ReadFile("keys/app.rsa")
	if err != nil {
		return nil, fmt.Errorf("error reading private key")
	}

	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes);
	if err != nil {
		return nil, fmt.Errorf("error reading private key")
	}

	return privKey, nil
}

func validateJwt(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		publicKey, err := readPublicKey()
		if err != nil {
			return nil, fmt.Errorf("error reading public key")
		}

		return publicKey, nil
	})

	if err == nil && token.Valid {
		return true
	} else {
		log.Println(err)
		return false
	}
}

func generateJwt(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"userId": user.Id,
		"username": user.Name,
		"isAdmin": user.IsAdmin,
		"exp": time.Now().Add(time.Minute * 1).Unix(),
	})

	privKey, err := readPrivateKey()
	if err == nil {
		tokenString, err := token.SignedString(privKey)
		if err == nil {
			return tokenString, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}

func SecureEndpoint(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")
		if authToken != "" {
			bearer := strings.Split(authToken, " ")

			if bearer[1] != "" {
				if validateJwt(bearer[1]) == true {
					h.ServeHTTP(w, r)
				} else {
					http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				}
			} else {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			}

		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}

	})
}