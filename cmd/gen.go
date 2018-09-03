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
	"image/png"
	"log"
	"os"
	"path/filepath"
	models "qcode/models"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generates QRcode for all users in the database",
	Long: `Generates QRcode 
		
				- To Generate QRcodes for simple text info

					qcode gen s helloWorld 

				- To Generate QRcodes for users in a sqlite database

					 qcode gen du
					`,
	Run: func(cmd *cobra.Command, args []string) {
		users := models.Users{
			Store: map[string]*models.User{},
		}
		fmt.Println("Generating QRcode")
		if len(args) == 0 {
			checkError(fmt.Errorf("please specify the nessery command "))
		}
		typ := args[0]
		if typ == "simple" || typ == "s" {
			val := args[1]
			if val == "" {
				checkError(fmt.Errorf("No value Specified, Please specify a value %s", val))
			}
			qrCode, _ := qr.Encode(string(val), qr.H, qr.Auto)

			// Scale the barcode to 200x200 pixels
			qrCode, _ = barcode.Scale(qrCode, 200, 200)
			name := val[:5] + ".png"

			wd, _ := os.Getwd()
			path := filepath.Join(wd, "public/images", name)
			dst, _ := os.Create(path)

			defer dst.Close()

			// encode the barcode as png
			png.Encode(dst, qrCode)
			fmt.Println("public/images/" + name)

		} else if typ == "dbuser" || typ == "du" {
			orm, err := SessionSqlite(f)
			if err != nil {
				log.Printf("Error Creating Database Session: %s", err)
			}

			orm.ShowSQL(true)
			var user []models.User
			user, err = users.Getuser(orm)
			if err != nil {
				log.Printf("Error Getting User: %s", err)
			}

			for _, use := range user {
				loc := Mapper(use)
				fmt.Println(loc)
			}
		} else {

			log.Printf("Please Specify the right command %s", typ)
		}

	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
