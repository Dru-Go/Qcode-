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

	"github.com/spf13/cobra"
	"gopkg.in/mail.v2"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send extraxts the value from the specified database then sends main to the given email",
	Long: `qcode send sends email to the collection of users specified in the database
			qcode send path(of the database)
			qcode send ./mydb.db 
			currently only supports sqlite database only
			`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("send called")
		m := mail.NewMessage()
		m.SetHeader("From", "alex@example.com")
		m.SetHeader("To", "bob@example.com", "cora@example.com")
		m.SetAddressHeader("Cc", "dan@example.com", "Dan")
		m.SetHeader("Subject", "Hello!")
		m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
		m.Attach("/home/Alex/lolcat.jpg")

		d := mail.NewDialer("smtp.example.com", 587, "user", "123456")
		d.StartTLSPolicy = mail.MandatoryStartTLS

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	},
}

// Message is a simple mail message
func Message(m *mail.Message, from, to, message string) {
	m.SetHeader("From", from)
	m.SetHeader("To", "to")
	m.SetAddressHeader("Cc", "drumac2@gmail.com", "GoWiz")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", message)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := mail.NewDialer("smtp.example.com", 587, "user", "123456")
	d.StartTLSPolicy = mail.MandatoryStartTLS

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
