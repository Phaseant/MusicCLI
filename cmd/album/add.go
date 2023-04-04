/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package album

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Phaseant/MusicCLI/cmd/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	filepath string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add albums to Music Server",
	Long:  `Long description`,
	Run: func(cmd *cobra.Command, args []string) {
		addAlbum()
	},
}

func init() {
	AlbumCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&filepath, "file", "f", "", "path to file with json albums")

	if err := addCmd.MarkFlagRequired("file"); err != nil {
		fmt.Println(err)
	}
}

func addAlbum() error {
	dataToSend, err := openFile()
	if err != nil {
		log.Errorf("unable to open file: %v", err)
		return err
	}

	// Create a new request using http
	req, _ := http.NewRequest("POST", constants.Url+"/api/album/", bytes.NewBuffer(dataToSend))

	// add authorization header to the req
	req.Header.Add("Authorization", constants.BearerToken)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error on response adding admin")
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error while reading the response bytes adding admin")
		return err
	}

	jsonResp := []map[string]string{}

	if err := json.Unmarshal(body, &jsonResp); err != nil {
		log.Errorf("unable to unmarshal json: %v", err)
		return err
	}
	for _, album := range jsonResp {

		if album["Error"] != "" {
			log.Errorf("unable to add album: %v", album["error"])
			return errors.New(album["error"])
		}

		fmt.Printf("Added album with id: %s", album["id"])
	}
	return nil
}

func openFile() ([]byte, error) {
	//find file in folder
	data, err := os.ReadFile(filepath)

	if err != nil {
		return []byte{}, err
	}

	return data, nil
}
