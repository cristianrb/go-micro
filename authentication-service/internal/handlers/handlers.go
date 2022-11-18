package handlers

import (
	"authentication/internal/services"
	"authentication/pkg/utils"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type handlerConfig struct {
	repo   services.UsersService
	client *http.Client
}

func NewHandlerConfig(repo services.UsersService, client *http.Client) *handlerConfig {
	return &handlerConfig{
		repo:   repo,
		client: client,
	}
}

func (u *handlerConfig) authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	// validate the user against the database
	user, err := u.repo.GetByEmail(requestPayload.Email)
	if err != nil {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := u.repo.PasswordMatches(requestPayload.Password, *user)
	if err != nil || !valid {
		utils.ErrorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	// log authentication
	err = u.logRequest("authentication", fmt.Sprintf("%s logged in", user.Email))
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	payload := utils.JsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in user %s", user.Email),
		Data:    user,
	}

	utils.WriteJSON(w, http.StatusAccepted, payload)
}

func (u *handlerConfig) logRequest(name, data string) error {
	var entry struct {
		Name string `json:"name"`
		Data string `json:"data"`
	}

	entry.Name = name
	entry.Data = data

	jsonData, _ := json.MarshalIndent(entry, "", "\t")
	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	_, err = u.client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
