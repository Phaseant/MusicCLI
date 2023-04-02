/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Phaseant/MusicCLI/cmd/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	username string
	password string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new admin to Music Server",
	Long:  `Add new admin to your server`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := addAdmin()
		if err == nil {
			fmt.Printf("Token for new admin:\n%s", token)
		}
	},
}

func init() {
	AdminCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&username, "username", "u", "", "username for new admin")
	addCmd.Flags().StringVarP(&password, "password", "p", "", "password for new admin")

	if err := addCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}
	if err := addCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println(err)
	}

}

func addAdmin() (string, error) {

	id, err := addUser()

	if err != nil {
		log.Error(err)
		return "", err
	}

	// Create a new request using http
	req, _ := http.NewRequest("POST", constants.Url+"/api/admin", bytes.NewBuffer(id))

	// add authorization header to the req
	req.Header.Add("Authorization", constants.BearerToken)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error on response adding admin")
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error while reading the response bytes adding admin")
		return "", err
	}

	answ := make(map[string]bool, 1)

	json.Unmarshal(body, &answ)

	if !answ["Added"] {
		errorAdd := make(map[string]string, 1)
		json.Unmarshal(body, &errorAdd)
		return "", fmt.Errorf("error adding admin: %s", errorAdd["Error"])
	}

	token, err := login()
	if err != nil {
		log.Error(err)
		return "", err
	}

	return token, nil
}

func addUser() ([]byte, error) {
	if username == "" || password == "" {
		return []byte{}, fmt.Errorf("username or password is not valid")
	}

	user := user{Username: username, Password: password}

	json_data, _ := json.Marshal(user)

	// Create a new request using http
	req, _ := http.NewRequest("POST", constants.Url+"/auth/register", bytes.NewBuffer(json_data))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, fmt.Errorf("unable to read response adding new user")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error while reading the response bytes adding new user")
	}

	regErr := make(map[string]string)

	json.Unmarshal(body, &regErr)

	if regErr["Error"] != "" {
		return []byte{}, fmt.Errorf("error creating new user: %s", regErr["Error"])
	}

	return body, nil
}

func login() (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf("username of password is not valid")
	}

	user := user{Username: username, Password: password}

	json_data, _ := json.Marshal(user)

	// Create a new request using http
	req, _ := http.NewRequest("POST", constants.Url+"/auth/login", bytes.NewBuffer(json_data))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("unable to read response logging in")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading the response bytes logging in")
	}

	tokenRaw := make(map[string]string, 1)
	json.Unmarshal(body, &tokenRaw)

	if tokenRaw["token"] == "" {
		return "", fmt.Errorf("unable to log in")
	}

	return tokenRaw["token"], nil
}
