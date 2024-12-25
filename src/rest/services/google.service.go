package services

import (
	"Hackathon-Management-System/src/auth/utils"
	"Hackathon-Management-System/src/graph/model"
	"Hackathon-Management-System/src/internal/constants"
	service "Hackathon-Management-System/src/internal/services"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	appConfig "Hackathon-Management-System/src/internal/config"
)

type GoogleService struct {
	Appconfig   *appConfig.AppConfig
	UserService service.UserService
}

func NewGoogleServices(appConfig *appConfig.AppConfig) *GoogleService {
	return &GoogleService{
		Appconfig:   appConfig,
		UserService: *service.NewUserService(appConfig),
	}
}

func (s GoogleService) ProcessOAuth(c *gin.Context) (*model.User, bool, error) {
	code := c.Query("code")
	newSignUp := false
	accessToken, err := s.GetAccessToken(code)
	if err != nil {
		fmt.Println("Error getting access token:", err)
		return nil, newSignUp, err
	}

	user, err := s.GetGoogleUser(accessToken)
	if err != nil {
		fmt.Println("Error getting user details:", err)
		return nil, newSignUp, err
	}
	userByEmailID, _ := s.UserService.GetUserByEmail(c, user.Email)
	if userByEmailID != nil {
		user = userByEmailID
	} else {
		newSignUp = true
		input := model.CreateUserInput{
			Name:  user.Name,
			Email: user.Email,
		}

		user, err = s.UserService.CreateUser(c, input)
		if err != nil {
			return nil, newSignUp, err
		}
	}
	jwtToken, err := utils.GenerateJWTToken(user)
	if err != nil {
		return nil, newSignUp, err
	}

	fmt.Println("JWT Token :: ", jwtToken)
	return user, newSignUp, nil
}

func (s GoogleService) GetAccessToken(code string) (string, error) {
	backgroundUrl := appConfig.NewConfig().Server.BACKEND_URL + "/auth/google/callback"
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", appConfig.NewConfig().GoogleConfig.GOOGLE_CLIENT_ID)
	data.Set("client_secret", appConfig.NewConfig().GoogleConfig.GOOGLE_SECRET_KEY)
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

func (s GoogleService) GetGoogleUser(accessToken string) (*model.User, error) {
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

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	user := &model.User{
		Name:  response["name"].(string),
		Email: response["email"].(string),
	}

	return user, nil
}
