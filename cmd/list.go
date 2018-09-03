// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	models "qcode/models"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shows all the user contents in the database",
	Long: `Here is how u use it
				qcode list`,
	Run: func(cmd *cobra.Command, args []string) {
		users := models.Users{
			Store: map[string]*models.User{},
		}
		fmt.Println("list called")
		orm, err := SessionSqlite(f)
		if err != nil {
			log.Printf("Error Creating Database Session: %s", err)
		}

		orm.ShowSQL(true)
		user, err := users.Getuser(orm)
		if err != nil {
			checkError(fmt.Errorf("Unable to Get users from the database: %s", err))
		}
		fmt.Println(len(user))
		fmt.Println(user)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
