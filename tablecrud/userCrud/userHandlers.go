package userComp

import (
	db "AdvDashBoard/goAdvDashBoard/datastore/dataBase"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	//"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	hostName string
)

// SetHost ...
func SetHost(host string) { hostName = host }

var (
	muxRouterPort string
)

// SetMuxPortAddr ...
func SetMuxPortAddr(addr string) {
	muxRouterPort = addr
}

var StatusHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is up and running, ListenAndServe on " + muxRouterPort))
	})

/* Set up a global string for our secret */
var mySigningKey = []byte("secret")

// AuthUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func AuthUser(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>userHandlers.AuthUser()")

	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["admin"] = true
	claims["name"] = "Ado Kukic"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	log.Println("   Encoding tokenString=%v", tokenString)

	json.NewEncoder(w).Encode(tokenString)
}

// GetAuthUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetAuthUser(w http.ResponseWriter, r *http.Request) {
	log.Println("===>userHandlers.GetAuthUser()")
	w.Header().Add("Content-Type", "application/json")

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var aUser User
	json.Unmarshal([]byte(har), &aUser)
	fmt.Printf("unmarshalled aUser=%v", aUser)

	myUser := &User{UserName: aUser.UserName}
	myDB := db.Connect()
	user, sqlErr := myUser.GetByName(myDB)

	if sqlErr == nil {
		log.Printf("Successful GetByName() call in service.GetAuthUsers()\n")
		log.Printf("user=%v\n", user)
		if myUser.Password == aUser.Password {
			log.Printf("Password Matches")

			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)

			claims["role"] = myUser.RoleName
			claims["name"] = myUser.UserName
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			tokenString, _ := token.SignedString(mySigningKey)
			//log.Println("   Encoding tokenString=%v", tokenString)
			user.JwtToken = tokenString
			log.Printf("Full User with Token=%v\n", user)

			json.NewEncoder(w).Encode(user)
			return
		}
		log.Printf("Password MIS match")
		json.NewEncoder(w).Encode("Your password failed")
		return
	}
	log.Printf("Error with database select in service.GetAuthUser()\n")
	log.Printf("Reason:  %v\n", sqlErr)
}

// GetUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("===>userHandlers.GetUser()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myUser := &User{}
	paramInt, paramErr := strconv.Atoi(params["id"])
	if paramErr == nil {
		myUser.ID = paramInt
		user, sqlErr := myUser.GetByID(myDB)
		if sqlErr == nil {
			log.Printf("Successful select in service.GetUsers() user=%v\n", user)
			json.NewEncoder(w).Encode(user)
			return
		}
		log.Printf("Error with database select in service.GetUser()\n")
		log.Printf("Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in userHandlers.GetUser()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

// RegisterUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>userHandlers.RegisterUser()")
	w.Header().Add("Content-Type", "application/json")

	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var user User
	json.Unmarshal([]byte(har), &user)
	// user.RoleName = "Viewer"
	fmt.Printf("%v", user)

	myDB := db.Connect()

	log.Printf("   userHandlers.Create():=%v\n", user)

	sqlErr := user.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful registration of user userHandlers.RegisterUser()\n")
		log.Printf("user=%v\n", user)
		json.NewEncoder(w).Encode(user)
		return
	}
	log.Printf("Error registering user in userHandlers.User()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

// CreateUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>userHandlers.CreateUser()")
	w.Header().Add("Content-Type", "application/json")

	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)

	var user User
	json.Unmarshal([]byte(har), &user)
	fmt.Printf("%v", user)

	myDB := db.Connect()

	log.Printf("   userHandlers.Create():=%v\n", user)

	sqlErr := user.Create(myDB)
	if sqlErr == nil {
		log.Printf("Successful created user userHandlers.CreateUser()\n")
		log.Printf("user=%v\n", user)
		json.NewEncoder(w).Encode(user)
		return
	}
	log.Printf("Error creating user in userHandlers.CreateUser()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

// UpdateUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("===>userHandlers.UpdateUser()")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	myDB := db.Connect()
	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)

	pUser := params["User"]
	log.Printf("   User=%v\n", pUser)

	har, err := ioutil.ReadAll(r.Body)
	log.Printf("   har=%s\n", har)
	log.Printf("   err=%v\n", err)
	var user User
	json.Unmarshal([]byte(har), &user)
	fmt.Printf("%v", user)

	sqlErr := user.Update(myDB)
	if sqlErr == nil {
		log.Printf("Successful update in userHandlers.UpdateUser() user=%v\n", user)
		json.NewEncoder(w).Encode(user)
		return
	}
	log.Printf("Error updateing user in userHandlers.UpdateUser()\n")
	log.Printf("Reason:%v\n", sqlErr)
}

//DeleteUser ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("\n\n===>userHandlers.DeleteUser()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	params := mux.Vars(r)
	log.Printf("   params=%v\n", params)
	myDB := db.Connect()
	myUser := &User{}
	paramInt, paramErr := strconv.Atoi(params["id"])

	if paramErr == nil {
		myUser.ID = paramInt
		sqlErr := myUser.Delete(myDB)
		if sqlErr == nil {
			log.Printf("Successful delete in service.DeleteUsers() user=%v\n", myUser)
			json.NewEncoder(w).Encode(myUser)
			return
		}
		log.Printf("Error with database delete in service.DeleteUsers()\n")
		log.Printf("Reason:  %v\n", sqlErr)
	}
	log.Printf("Error with Paramerer in userHandlers.DeleteUsers()\n")
	log.Printf("Reason:  %v\n", paramErr)
}

/*=====================================================================*\
	Handlers Functions completed the END 2 END naming conversion

\*=====================================================================*/
// GetUsersFullList ...
/*---------------------------------------------------------------------*\
\*---------------------------------------------------------------------*/
func GetUsersFullList(w http.ResponseWriter, r *http.Request) {
	log.Println("===>userHandlers.GetUsersFullList()")
	w.Header().Add("Content-Type", "application/json")
	log.Printf("   w.Header=%v\n", w.Header())
	log.Printf("   hostName=%v\n", hostName)

	myDB := db.Connect()
	myUser := &User{}
	users, sqlErr := myUser.GetUsersFullList(myDB)
	log.Printf("%d users found inside service.GetUsers()\n", len(users))
	// for i := 0; i < len(users); i++ {
	// 	log.Printf("   %v\n", users[i])
	// }

	if sqlErr == nil {
		log.Printf("Successful select of all users in service.GetUsers()")
		json.NewEncoder(w).Encode(users)
		return
	}
	log.Printf("Error with database select in userHandlers.GetUser()\n")
	log.Printf("Reason:%v\n", sqlErr)
}
