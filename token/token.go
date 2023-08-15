package token

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/dgrijalva/jwt-go"
)

var JWT_PUBLIC_KEY *rsa.PublicKey

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	pubKeyPEM, _ := base64.StdEncoding.DecodeString(os.Getenv("JWT_PUBLIC_KEY"))
	block, _ := pem.Decode(pubKeyPEM)
	if block == nil {
		log.Fatal("Failed to parse PEM block containing the public key")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal("Failed to parse public key: " + err.Error())
	}

	var ok bool
	JWT_PUBLIC_KEY, ok = pub.(*rsa.PublicKey)
	if !ok {
		log.Fatal("Public key of unsupported type")
	}
}

func DecodeToken(tokenString string, target interface{}) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_PUBLIC_KEY, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		recordBytes := claims["record"].(map[string]interface{}) // Assuming "record" is a JSON object
		recordJSON, err := json.Marshal(recordBytes)
		if err != nil {
			return err
		}
		err = json.Unmarshal(recordJSON, &target)
		if err != nil {
			return err
		}
		
		return nil
	} else {
		return err
	}

}
