package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recaptcha/models"

	"github.com/gin-gonic/gin"
)

func verifyRecaptchaToken(token string) (*models.RecaptchaResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://www.google.com/recaptcha/api/siteverify", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("secret", "6LfMimspAAAAACW8eXUHGQZMCJV13zxeOdHi9wTP")
	query.Add("response", token)
	req.URL.RawQuery = query.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var recaptchaResponse models.RecaptchaResponse
	err = json.NewDecoder(resp.Body).Decode(&recaptchaResponse)
	if err != nil {
		return nil, err
	}

	return &recaptchaResponse, nil
}
func SubmitHandler(c *gin.Context) {
	token := c.PostForm("token")
	recaptchaResponse, err := verifyRecaptchaToken(token)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error verifying reCAPTCHA token")
		return
	}

	if recaptchaResponse.Success && recaptchaResponse.Score >= 0.5 {
		c.String(http.StatusOK, fmt.Sprintf("reCAPTCHA token successfully verified with score %f!,you are a human", recaptchaResponse.Score))
	} else {
		c.String(http.StatusBadRequest, "reCAPTCHA token verification failed or score < 0.5,you are a bot")
	}
}
