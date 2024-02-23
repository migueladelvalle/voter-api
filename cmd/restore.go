/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"drexel.edu/voter-api/pkg/storage/json"
	"github.com/spf13/cobra"
)

const (
	defaultTargetFilePath = "./Data"
	defaultBackupFilePath = "./Data.Bak"
)

var backupRoute string

var targetFilePath string

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restores the database to a backup file",
	Long: `Restores the database to a specified backup file. 
	if no file is provided then a default with sample data is 
	used`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restore called")

		db, err := json.NewJsonDB(jsonFilePath)
		if err != nil {
			panic(err)
		}

		err = db.RestoreDB(backupRoute)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restoreCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	restoreCmd.Flags().StringVarP(&backupRoute, "target", "t", defaultBackupFilePath, "target a specific backup file")
	restoreCmd.Flags().StringVarP(&targetFilePath, "destination", "f", defaultTargetFilePath, "The file path to the Json DB")
}
