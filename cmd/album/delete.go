/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package album

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Phaseant/MusicCLI/cmd/constants"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	id string
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete album",
	Long:  `delete album`,
	Run: func(cmd *cobra.Command, args []string) {
		deleteAlbum()
	},
}

func init() {
	AlbumCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringVar(&id, "id", "", "admin id to delete")

	if err := deleteCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
	}
}

func deleteAlbum() {
	// Create a new request using http
	req, _ := http.NewRequest("DELETE", constants.Url+"/api/album/"+id, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", constants.BearerToken)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error("error on response adding admin")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("error while reading the response bytes adding admin")
		return
	}

	var jsonResp errDel

	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		log.Error(err)
		return
	}

	if !jsonResp.Deleted {
		log.Errorf("unable to delete album: %s", jsonResp.Error)
		return
	}

	log.Println("Album was deleted")

}

type errDel struct {
	Deleted bool   `json:"deleted"`
	Error   string `json:"error"`
}
