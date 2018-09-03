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
	"os"

	"github.com/spf13/cobra"
)

// trancateCmd represents the trancate command
var trancateCmd = &cobra.Command{
	Use:   "trancate",
	Short: "Clear the database and images",
	Long: ` To clear all the images 
				qcode trancate img 

			To clear the database 
				qcode trancate db

			To clear all 
				qcode trancate`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			checkError(fmt.Errorf("please specify the nessery command "))
		}
		fmt.Println("trancate called")
		switch args[0] {
		case "img":
			TrancateImg("/images")
		case "db":
			TrancateDB("/db")
		default:
			TrancateAll("/images", "/db")
		}

	},
}

func TrancateImg(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func TrancateDB(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func TrancateAll(imgpath, dbpath string) error {
	err := TrancateImg(imgpath)
	if err != nil {
		return err
	}
	err = TrancateDB(dbpath)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(trancateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trancateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trancateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
