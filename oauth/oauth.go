package oauth

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthConf *oauth2.Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func init() {
	OauthConf = &oauth2.Config{
		ClientID:     "ClientID",
		ClientSecret: "ClientSecret",
		RedirectURL:  "http://localhost:9300/login/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func GetToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawStdEncoding.EncodeToString(b)
}

func GetLoginURL(state string) string {
	return OauthConf.AuthCodeURL(state)
}
