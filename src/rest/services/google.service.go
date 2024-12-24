package services

import (
	"Hackathon-Management-System/src/internal/constants"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type GoogleService struct {
}

func NewGoogleServices() *GoogleService {
	return &GoogleService{}
}

func (s GoogleService) init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func (s GoogleService) ProcessOAuth(c *gin.Context) (map[string]interface{}, error) {
	s.init()
	code := c.Query("code")
	accessToken, err := s.GetAccessToken(code)
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return nil, err
	}

	user, err := s.GetGoogleUser(accessToken)
	if err != nil {
		fmt.Println("Error getting user details:", err)
		return nil, err
	}
	name := user["name"].(string)
	email := user["email"].(string)
	avatar := user["picture"].(string)
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(avatar)
	return user, nil
}

func (s GoogleService) GetAccessToken(code string) (string, error) {
	backgroundUrl := os.Getenv("BACKEND_URL") + "/auth/google/callback"
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	data.Set("client_secret", os.Getenv("GOOGLE_CLIENT_SECRET"))
	data.Set("redirect_uri", backgroundUrl)
	data.Set("grant_type", "authorization_code")

	resp, err := http.PostForm(constants.GoogleAccessToken, data)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", err
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", errors.New("Error getting access token")
	}
	return accessToken, nil
}

func (s GoogleService) GetGoogleUser(accessToken string) (map[string]interface{}, error) {
	req, err := http.NewRequest("GET", constants.GoogleUserInfo, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user map[string]interface{}
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return user, nil
}
