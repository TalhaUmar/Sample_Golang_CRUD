package controllers

import (
	"SchmidtRouter/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserController : composing an empty user controller
type UserController struct {
	session *mgo.Session
}

// NewUserController return pointer to the value to user controller struct
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Index : index function
func (uc UserController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Julien Schmidt HTTPRouter")
}

// GetUser : get an existing user
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("go-serve").C("developers").FindId(oid).One(&u); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser : create a new User
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	u.ID = bson.NewObjectId()

	uc.session.DB("go-serve").C("developers").Insert(u)

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

// DeleteUser : delete an existing user
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)

	custom := make(chan int)

	go func() {
		if err := uc.session.DB("serve-go").C("developers").RemoveId(oid); err != nil {
			custom <- 1
			w.WriteHeader(http.StatusNotFound)
		}
	}()

	select {
	case <-custom:
		break
	case <-custom:
		break
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User", oid, "\n")
}
