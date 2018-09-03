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
	// "encoding/csv"

	"errors"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"encoding/json"
	"encoding/xml"
	models "qcode/models"

	"github.com/aodin/csv2"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cobra"
)

var f = "db/test.db"

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add collects users from a file and saves to database",
	Long:  `Add is used to add users from a list of files to a sqlite database, files supported are (json,csv,xml)`,
	Run: func(cmd *cobra.Command, args []string) {

		// users := models.Users{
		// 	Store: map[string]*models.User{},
		// }
		// if len(args) == 0 {
		// 	checkError(fmt.Errorf("please specify the nessery command "))
		// }
		// if args[1] == "" {
		// 	checkError(fmt.Errorf("please specfy the directory of the file u want to sync"))
		// }

		// //Opening the file
		// fifo, fileerr := os.OpenFile(args[1], os.O_RDWR, 0755)
		// if fileerr != nil {
		// 	checkError(fmt.Errorf("Error opening the file specified %s", fileerr))
		// }
		// ctype, _ := GetType(fifo)
		// fmt.Println(ctype)
		// // Start session for mongo database
		// // session, err := mgo.Dial("mongodb://127.0.0.1:27017")
		// // if err != nil {
		// // 	panic(err)
		// // }
		// // defer session.Close()
		// os.Remove(f)

		// // Start session for xorm
		// orm, ormerr := SessionSqlite(f)
		// if ormerr != nil {
		// 	checkError(fmt.Errorf("Error: %s", ormerr))
		// }
		// // creating the Tables
		// orm.ShowSQL(true)
		// taberr := Create(orm)
		// if taberr != nil {
		// 	checkError(fmt.Errorf("Error: %s", taberr))
		// }
		// // Decoding the given file format
		// users, err := Decoder(args[0], fifo)
		// if err != nil {
		// 	checkError(fmt.Errorf("Error: %s", err))
		// }
		// for _, user := range users {
		// 	err := users.Adduser(user, orm)
		// 	if err != nil {
		// 		checkError(fmt.Errorf("Error: %s", err))
		// 	}
		// }

	},
}

// SessionSqlite starts the session for the XORM
func SessionSqlite(f string) (*xorm.Engine, error) {
	orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		return nil, err
	}
	return orm, nil
}

// Create Creates the users table
func Create(orm *xorm.Engine) error {
	// orm.ShowSQL(true)
	err := orm.CreateTables(&models.User{})
	if err != nil {
		return err
	}
	return nil
}

// AddUsers Adds users to the Database
// func AddUsers(users models.User, sess *mgo.Session) error {
// 	// sess.SetMode(mgo.Monotonic, true)
// 	C := sess.DB("qcode").C("users")
// 	loc := Mapper(users)
// 	users.Qcode = loc
// 	err := C.Insert(users)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(users)
// 	return nil
// }

// Mapper maps the users to the qrcode
func Mapper(user models.User) string {
	if user.Name == "Name" || user.Name == "" {
		return ""
	}
	js, err := json.Marshal(user)
	if err != nil {
		checkError(err)
	}
	qrCode, _ := qr.Encode(string(js), qr.H, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)
	name := user.Email + ".png"

	wd, _ := os.Getwd()
	path := filepath.Join(wd, "/public/images", name)
	dst, _ := os.Create(path)

	// create the output file
	// file, _ := os.Create(name)
	defer dst.Close()

	// encode the barcode as png
	png.Encode(dst, qrCode)

	return name
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// Decoder decodes the values specfied in any format
func Decoder(ftype string, content *os.File) ([]models.User, error) {
	ctype, err := GetType(content)
	if err != nil {
		return nil, fmt.Errorf("can not get the type of this file %s", err)
	}

	switch ftype {
	case "json":
		if ctype == "json" {
			users, err := JSONDecoder(content)
			if err != nil {
				return nil, err
			}
			return users, nil
		} else {
			return nil, errors.New("This is not a json file")
		}

	case "xml":
		if ctype == "xml" {
			users, err := XMLParser(content)
			if err != nil {
				return nil, err
			}
			return users, nil
		} else {
			return nil, fmt.Errorf("This is not a xml file")
		}

	case "csv":
		if ctype == "csv" {
			users, err := CSVDecoder(content)
			if err != nil {
				return nil, err
			}
			return users, nil
		} else {
			return nil, fmt.Errorf("This is not a csv file")
		}

	default:
		return nil, fmt.Errorf("Unable to find the right data type")
	}
}

// ExtractToken extracts the string out of the doc
func ExtractToken(doc *os.File) (string, error) {
	src, err := os.Open(doc.Name())
	if err != nil {
		return "", err
	}
	bs, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatalln("Unable to Read the document")
	}
	// fmt.Println(len(bs))
	str := string(bs)
	return str, nil
}

// GetType is used to analize what type of file we want to sync
func GetType(file *os.File) (string, error) {
	switch path.Ext(file.Name()) {
	case ".csv":
		return "csv", nil
	case ".json":
		return "json", nil
	case ".xml":
		return "XML", nil
	default:
		log.Fatal("Please use either csv,json or an xml format to sync data")
	}
	return "", fmt.Errorf("Please use either csv,json or an xml format to sync data")
}

// JSONDecoder Parses Json files
func JSONDecoder(byt *os.File) ([]models.User, error) {
	var users []models.User
	data, tocerr := ExtractToken(byt)
	if tocerr != nil {
		return nil, fmt.Errorf("%s", tocerr)
	}

	if err := json.Unmarshal([]byte(data), &users); err != nil {
		return nil, fmt.Errorf("Json unmarshaling error Help! %s", err)
	}
	fmt.Println(users)
	return nil, nil
}

// CSVDecoder Decodes CSV documents
func CSVDecoder(byt *os.File) ([]models.User, error) {
	csvf := csv.NewReader(byt)
	var users []models.User
	if marerr := csvf.Unmarshal(&users); marerr != nil {
		return nil, fmt.Errorf("csrf unmarshaling error Help! %s", marerr)
	}
	fmt.Println(users)
	return users, nil
}

// XMLParser Parses XML
func XMLParser(byt *os.File) ([]models.User, error) {
	data, tocerr := ExtractToken(byt)
	if tocerr != nil {
		return nil, fmt.Errorf("%s", tocerr)
	}
	var users []models.User
	err := xml.Unmarshal([]byte(data), &users)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling XML Help! %s", err)
	}
	fmt.Println(users)
	return users, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
