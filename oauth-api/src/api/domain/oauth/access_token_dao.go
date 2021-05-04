package oauth

import (
	"fmt"
	"github.com/voicurobert/golang-microservices/src/api/utils/errors"
)

var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserID)
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {
	at := tokens[accessToken]
	fmt.Println(at)
	if at == nil || at.IsExpired() {
		return nil, errors.NewNotFoundApiError("no access token found")
	}
	return at, nil
}
