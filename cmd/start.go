/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"drexel.edu/voter-api/pkg/http/rest"
	"drexel.edu/voter-api/pkg/process"
	"drexel.edu/voter-api/pkg/retrieve"
	"drexel.edu/voter-api/pkg/storage/json"
	"github.com/spf13/cobra"
)

const (
	defaultFilePath = "./Data"
)

var port int
var jsonFilePath string

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the server",
	Long:  `Allows the user to specify the port, otherwise uses 3000 by default`,
	Run: func(cmd *cobra.Command, args []string) {

		repository, err := json.NewJsonDB(jsonFilePath)
		if err != nil {
			panic(err)
		}

		processService := process.NewService(repository)
		retrievalService := retrieve.NewService(repository)

		router := rest.Handler(port, processService, retrievalService)

		fmt.Printf("The Server is started: http://localhost:%d", port)

		log.Fatal(router.Listen(string(fmt.Sprintf(":%d", port))))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().IntVarP(&port, "port", "p", 3000, "The port on which to start the server")
	startCmd.Flags().StringVarP(&jsonFilePath, "filePath", "f", defaultFilePath, "The file path to the Json DB")
}
