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
// WITHOUT WARRANTSIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	models "qcode/models"
	"strconv"

	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
)

// excelCmd represents the excel command
var excelCmd = &cobra.Command{
	Use:   "excel",
	Short: "excel converts a simple msexcel file to an object then stores it to a database",
	Long: `this is how 

		 qcode msexcwl.xlsx 0 4
				here msexcwl.xlsx represents the filename
				here 0 represents the sheet number
				here 4 represents the number of rows to scan
				enjoy!!!
				`,
	Run: func(cmd *cobra.Command, args []string) {
		users := models.Users{
			Store: map[string]*models.User{},
		}
		if len(args) == 0 {
			checkError(fmt.Errorf("please specify the nessery command "))
		}
		excelFileName := args[0]
		sheets, _ := strconv.Atoi(args[1])
		rows, _ := strconv.Atoi(args[2])
		fmt.Println("excel called")
		var user = make([]models.User, rows)
		bee, err := xlsx.FileToSlice(excelFileName)
		if err != nil {
			log.Println(err)
		}

		for i := 1; i < rows; i++ {
			user[i].Name = bee[sheets][i][0]
			user[i].Email = bee[sheets][i][1]
			user[i].PhoneNo = bee[sheets][i][2]
			user[i].Role = bee[sheets][i][3]
			bol, _ := strconv.ParseBool(bee[sheets][i][4])
			user[i].Check = bol
		}
		fmt.Println(user)
		orm, err := xorm.NewEngine("sqlite3", f)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = orm.CreateTables(&models.User{})
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, use := range user {
			err = users.Adduser(use, orm)
		}
		if err != nil {
			log.Println(err)
		}
	},
}

// wd, _ := os.Getwd()
// path := filepath.Join(wd, "assets", "imgs", fName)
// dst, _ := os.Create(path)

func init() {
	rootCmd.AddCommand(excelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// excelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// excelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
