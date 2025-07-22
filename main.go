package main

import (
	"encoding/json"
	"fmt"
	authMethods "gtk/auth"
	"gtk/req"
	"gtk/types"
	"log"
)

const API_URL = "http://glpi.com/api.php/"
const USERNAME = "usuario"
const PASSWORD = "senha"
const CLIENT_ID = "ID do Client"
const CLIENT_SECRET = "Secret do Client"

func main() {
	var auth types.AuthParams
	authBytes, err := authMethods.GetAuthenticated(API_URL, USERNAME, PASSWORD, CLIENT_ID, CLIENT_SECRET)
	if err != nil {
		log.Panic(err)
	}
	json.Unmarshal(authBytes, &auth)

	req, err := req.MakeReq("GET", API_URL, "/Assistance/Ticket", nil, auth.Token)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(req))
}
