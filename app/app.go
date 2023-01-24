package app

import (
	"bank/model"
	"encoding/json"
	"net/http"
)

type Repository interface {
	GetUser(name string) (*model.User, error)
	InsertUser(user *model.User) error
}
type Application struct {
	repository Repository
}

func (a Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data model.User
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.repository.InsertUser(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Success"))
}

type GetUser struct {
	Name string `json:"name"`
}

func (a Application) GetUserByName(w http.ResponseWriter, r *http.Request) {
	var request GetUser
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := a.repository.GetUser(request.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
func NewApplication(repository Repository) *Application {
	return &Application{
		repository: repository,
	}
}
