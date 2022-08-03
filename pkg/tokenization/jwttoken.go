package tokenization

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

func GenerateClaims(issuer string, expirationTime *jwt.Time) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: expirationTime, //jwt.At(time.Now().Add(time.Hour * 24 * 7))
	}
}

func GenerateToken(claims *jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func GenerateCookie(token string, expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Expires:  expiresAt,
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
}

func ParseToken(r *http.Request, cookieName string) (*jwt.Token, error) {
	token, err := r.Cookie(cookieName)
	if err != nil {
		return nil, err
	}
	return jwt.ParseWithClaims(token.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
}
