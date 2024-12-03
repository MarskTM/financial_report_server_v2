package controller

import (
	"context"
	"net/http"
	"net/url"
	"phenikaa/infrastructure"
	"phenikaa/model"

	"time"
)

func SaveHttpCookie(fullDomain string, tokenDetail *model.TokenDetail, w http.ResponseWriter) error {
	domain, err := url.Parse(fullDomain)
	if err != nil {
		return err
	}

	cookie_access := http.Cookie{
		Name:     "AccessToken",
		Domain:   domain.Hostname(),
		Path:     "/",
		Value:    tokenDetail.AccessToken,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * time.Duration(model.AccessTokenTime)),
	}

	cookie_refresh := http.Cookie{
		Name:     "RefreshToken",
		Domain:   domain.Hostname(),
		Path:     "/",
		Value:    tokenDetail.RefreshToken,
		HttpOnly: false,
		Secure:   false,
		Expires:  time.Now().Add(time.Hour * time.Duration(model.RefreshTokenTime)),
	}

	http.SetCookie(w, &cookie_access)
	http.SetCookie(w, &cookie_refresh)
	return nil
}

func GetAndDecodeToken(token string) (map[string]interface{}, error) {
	if token == "" {
		return nil, nil
	}
	decodedToken, err := infrastructure.GetDecodeAuth().Decode(token)
	if err != nil {
		return nil, err
	}
	claims, err := decodedToken.AsMap(context.Background())
	if err != nil {
		return nil, err
	}
	return claims, nil
}
