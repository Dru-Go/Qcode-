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

	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed adds some random users to the database",
	Long: ` qcode seed will add the sample datas for experemntation
				qcode seed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("seed called")
		users := models.Users{
			Store: map[string]*models.User{},
		}
		orm, err := SessionSqlite(f)
		if err != nil {
			log.Printf("Error Creating Database Session: %s", err)
		}
		// Create a database tables
		err = Create(orm)
		if err != nil {
			log.Panicf("Error Creating Table: %s", err)
		}
		// Add some seed tasks
		user1 := models.User{
			Name:    "Kigo",
			Email:   "Kigo@mail.com",
			PhoneNo: "+2313432533",
			Role:    "Guest",
		}
		user2 := models.User{
			Name:    "Martin Garices",
			Email:   "Martin@mail.com",
			PhoneNo: "+23994328003",
			Role:    "Guest",
		}
		user3 := models.User{
			Name:    "Alen Walker",
			Email:   "Walker@mail.com",
			PhoneNo: "+25233223223",
			Role:    "Speaker",
		}
		user4 := models.User{
			Name:    "Avici",
			Email:   "Avici@mail.com",
			PhoneNo: "+257656767676",
			Role:    "Speaker",
		}
		// Add and generate QRcodes and seed to the database
		users.Lock()
		orm.ShowSQL(true)
		_, err = orm.Insert(&user1, &user2, &user3, &user4)
		if err != nil {
			log.Fatal(err)
		}
		users.Unlock()
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
