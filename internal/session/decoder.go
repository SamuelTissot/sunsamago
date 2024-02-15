package session

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// claims structure
// 	UserId       string `json:"userId"`
// 	SessionId    string `json:"sessionId"`
// 	SessionNonce string `json:"sessionNonce"`
// 	GroupId      string `json:"groupId"`
// 	GroupUrlPath string `json:"groupUrlPath"`
// 	Iat          int    `json:"iat"`
// 	Exp          int    `json:"exp"`

type Decoder struct {
	token  string
	claims jwt.Claims
}

func NewDecoder(token string) Decoder {
	return Decoder{token: token}
}

func UserID(sessionID string) (string, error) {
	claims, err := parse(sessionID)
	if err != nil {
		return "", err
	}

	return parseString(claims, "userId")
}

func GroupID(sessionID string) (string, error) {
	claims, err := parse(sessionID)
	if err != nil {
		return "", err
	}

	return parseString(claims, "groupId")
}

func parse(sessionID string) (map[string]any, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(sessionID, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func parseString(m map[string]any, key string) (string, error) {
	var (
		ok  bool
		raw interface{}
		iss string
	)
	raw, ok = m[key]
	if !ok {
		return "", nil
	}

	iss, ok = raw.(string)
	if !ok {
		return "", fmt.Errorf("%s is invalid", key)
	}

	return iss, nil
}
