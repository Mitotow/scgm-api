package services

import (
	"encoding/json"
	"fmt"
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/models"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	Scope       = "identify"
	GrantType   = "authorization_code"
	ContentType = "application/x-www-form-urlencoded"
)

type DiscordTokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type AuthService interface {
	Login() string
	Callback(code string) (*DiscordTokenResponse, *models.ErrorResponse)
}

type AuthServiceImpl struct {
	env      *config.EnvironmentVariables
	messages *config.Messages
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{env: config.GetEnv(), messages: config.GetMessages()}
}

func (s AuthServiceImpl) Login() string {
	authUrl := fmt.Sprintf(
		"%s/oauth2/authorize?client_id=%s&redirect_uri=%s&response_type=code&scope=%s",
		s.env.DiscordApiEndpoint,
		s.env.DiscordClientId,
		url.QueryEscape(s.env.RedirectUri),
		url.QueryEscape(Scope),
	)

	return authUrl
}

func (s AuthServiceImpl) Callback(code string) (*DiscordTokenResponse, *models.ErrorResponse) {
	if code == "" {
		return &DiscordTokenResponse{}, &models.ErrorResponse{
			Status: http.StatusBadRequest,
			Error:  s.messages.InvalidDiscordCode,
		}
	}

	data := url.Values{
		"client_id":     {s.env.DiscordClientId},
		"client_secret": {s.env.DiscordClientSecret},
		"grant_type":    {GrantType},
		"code":          {code},
		"redirect_uri":  {s.env.RedirectUri},
	}

	res, err := http.Post(
		fmt.Sprintf("%s/oauth2/token", s.env.DiscordApiEndpoint),
		ContentType,
		strings.NewReader(data.Encode()))

	if err != nil {
		return &DiscordTokenResponse{}, &models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Error:  s.messages.InternalServerError,
		}
	}

	defer func() {
		if closeErr := res.Body.Close(); closeErr != nil {
			log.Println("ERROR: Failed to close response body :", closeErr.Error())
		}
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR: io.ReadAll caused an error in callback auth function :", err.Error())
		return &DiscordTokenResponse{}, models.CreateInternalServerError()
	}

	var response DiscordTokenResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("ERROR: Unmarshal body inside response caused an error in callback auth function :", err.Error())
		return &DiscordTokenResponse{}, models.CreateInternalServerError()
	}

	return &response, nil
}
