package helper

import (
	"github/dev-hack95/Textflow/utilities/logs"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var sampleSecretKey = []byte("Hello World")

func CreateToken(Firstname, Lastname, email string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"first_name": Firstname,
			"last_name":  Lastname,
			"email":      email,
			"exp":        time.Now().Add(time.Hour * 24 * 45).Unix(),
			"admin":      true,
		})

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "Error Occured while Hashing passsword", err
	}
	return string(hashPassword), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check := true
	if err != nil {
		check = false
	}
	return check
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return sampleSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		logs.Error("Error: ", "Token is invalid!")
	}

	return nil
}

func StringGenerator(data string) (string, error) {
	rand.Seed(time.Now().UnixNano())

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 7)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return "video/" + string(b) + "_" + data, nil
}
