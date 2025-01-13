package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Romma711/ozora_web_ecommerse/server/pkg/auth"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/types"
	"github.com/Romma711/ozora_web_ecommerse/server/pkg/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) GetUsersRoutes(r *mux.Router) {
	r.HandleFunc("/register", h.HandleRegister).Methods(http.MethodPost)
	r.HandleFunc("/login", h.HandleLogin).Methods(http.MethodPost)

	r.HandleFunc("/admin/users/change/{id}", h.HandleUpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/admin/users/drop/{id}", h.HandleDeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/admin/users", h.HandleGetUsers).Methods(http.MethodGet)
	r.HandleFunc("/admin/users/{id}", h.HandleGetUser).Methods(http.MethodGet)
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.UserPayLoad //creates a local variable with the data from the request

	err := json.NewDecoder(r.Body).Decode(&user) //decodes the json object into the local variable
	if err != nil {
		log.Println(err)
	}

	hashedPassword, err := auth.HashPassword(user.Password) //hashes the password
	if err != nil {
		log.Println(err)
		return
	}
	user.Password = hashedPassword //asing the hashed password to the user object

	err = h.store.CreateUser(&user) //creates the user in the DB
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"User created successfully\"}")) //if everything is ok, sends a message to the client
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var login types.Login //creates a local variable

	err := json.NewDecoder(r.Body).Decode(&login) //decodes the json object into the local variable
	if err != nil {
		log.Println(err)
	}

	user, err := h.store.GetUserByEmail(login.Email) //serch in DB for the user
	if err != nil {
		log.Println(err)
		return
	}
	
	err = auth.ComparePassword(login.Password,user.Password) //compares the password
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\":\"Username or password is incorrect\"}"))
	}
	
	token := auth.GenerateToken(*user) //generates the token
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token}) //sends the token to the client
}


///ADMIN FUNCTIONS
func (h *Handler) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if role := auth.RoleUser(token); role != "admin" {
		utils.UnauthorizedUser(w)
		return
	}
	var user types.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {	
		log.Println(err)
	}
	err = h.store.UpdateUser(&user)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"User updated successfully\"}"))
}

func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	token := vars["token"]
	if role := auth.RoleUser(token); role != "admin" {
		utils.UnauthorizedUser(w)
		return
	}
	
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\":\"Error: Invalid ID\"}"))
		return
	}
	err = h.store.DeleteUser(id)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\":\"User deleted successfully\"}"))
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	token := mux.Vars(r)["token"]
	if role := auth.RoleUser(token); role != "admin" {
		utils.UnauthorizedUser(w)
		return
	}

	users, err := h.store.GetUsers()
	if err != nil {
		log.Println(err)
		return
	}
	
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	token := vars["token"]
	if role := auth.RoleUser(token); role != "admin" {
		utils.UnauthorizedUser(w)
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Error: ", err)	
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"message\":\"Error: Invalid ID\"}"))
		return
	}
	user, err := h.store.GetUserByID(id)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(&user)
	if err != nil {
		log.Println(err)
		return
	}
}