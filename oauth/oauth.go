package oauth

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthConf *oauth2.Config

const OauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthInit() {
	OauthConf = &oauth2.Config{
		ClientID:     "237352887062-e47geq5e4haje4q7a0d6ra9e9ul13lle.apps.googleusercontent.com",
		ClientSecret: "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",
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
