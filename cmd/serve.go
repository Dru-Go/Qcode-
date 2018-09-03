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
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	models "qcode/models"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-xorm/xorm"
	"github.com/gobuffalo/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var tpl *template.Template

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the API through a rest connection",
	Long: ` qcode helps in managing other guest sign up externally by providing a restful api and a template so a local client
		 	can manupilate diferent tasks
				qcode serve Response-Type Port 
					-Responce-Type = restfull 
						eg qcode serve rest  
					-Responce-Type = tempate
						eg qcode serve template  
			
				default port is :1988
			`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			checkError(fmt.Errorf("please specify the nessery command "))
		}
		var Resptype = args[0]

		port := ":1988"
		if Resptype == "rest" {
			fmt.Println("Listen at Port " + port)
			api := rest.NewApi()
			api.Use(rest.DefaultDevStack...)
			router, err := rest.MakeRouter(
				rest.Get("/api/users", GetAllUsers),
				rest.Post("/api/adduser", PostUser),
				rest.Put("/api/present/:id", Present),
				rest.Get("/api/checklist", CheckList),
				rest.Get("/api/count", Count),
				rest.Get("/api/getuser/:name", GetuserbyName),
				rest.Get("/api/user/:mail", GetuserbyEmail),
				// rest.Delete("/user/:demail", Deleteuser),
			)
			if err != nil {
				log.Fatal(err)
			}
			api.SetApp(router)
			log.Fatal(http.ListenAndServe(port, api.MakeHandler()))
		} else if Resptype == "template" {
			tpl = template.Must(template.ParseGlob("templates/html/*.html"))

			r := mux.NewRouter()
			r.HandleFunc("/", home).Methods("GET")
			r.HandleFunc("/add", add).Methods("POST")

			srv := &http.Server{
				Handler: r,
				Addr:    "127.0.0.1:2000",
				// Good practice: enforce timeouts for servers you create!
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}

			log.Fatal(srv.ListenAndServe())

		} else {
			checkError(fmt.Errorf("Unable to Serve please input the nessery commands : %s", " "))
		}

	},
}

/*
_____------______-------______-------_____

This is the Template Side of the program
_____--------________---------____________
*/

func home(res http.ResponseWriter, req *http.Request) {
	// res.Write([]byte("Cool"))
	res.Write([]byte(tpl.DefinedTemplates()))
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func add(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "home.html", nil)

	user := models.User{
		Name:    req.FormValue("icon_prefix"),
		Email:   req.FormValue("email"),
		PhoneNo: req.FormValue("phone-no"),
		Role:    req.FormValue("role"),
	}

	js, _ := json.Marshal(user)
	res.Write(js)
	if err := PutUser(res, user); err != nil {
		http.Error(res, err.Error()+" this is fucked up", 500)
		return
	}
	tpl.ExecuteTemplate(res, "home.html", nil)
}

func PutUser(res http.ResponseWriter, user models.User) error {
	users := models.Users{
		Store: map[string]*models.User{},
	}
	id, _ := uuid.NewV4()
	user1 := models.User{
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: user.PhoneNo,
		Role:    user.Role,
	}
	user.Qcode = Mapper(user1)
	user.ID = id.String()
	js, _ := json.Marshal(user)
	res.Write(js)

	orm, err := xorm.NewEngine("sqlite3", f)
	if err != nil {
		return fmt.Errorf("error connecting to the database: %s", err)
	}

	if err = users.Adduser(user, orm); err != nil {
		return fmt.Errorf("Error Adding user to database: %s", err)
	}
	return nil
}

/*
_____------______-------______-------_____

This is the Restfull Side of the program
_____--------________---------____________
*/
// Present ticks every user who arrived at the event
func Present(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := models.User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status, err := Check(id); err != nil {
		w.WriteJson(err.Error() + string(status) + "/n")
		return
	}
	fmt.Println("Didnt resach here")
	Tieck(id)
	w.WriteJson(&user)
	return
}

// Check the exitsance of the specified ID
func Check(id string) (int, error) {
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	// check if the user exists
	var users []models.User
	orm.ShowSQL(true)

	len, err := orm.Where("i_d = ?", id).FindAndCount(&users)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("Error Searching user :%s", err)
	}
	if len == 0 {
		return 404, fmt.Errorf("There is no user with this id :%s", id)
	}
	// check if the user is checked in

	for _, user := range users {
		if user.Check {
			return http.StatusNotAcceptable, fmt.Errorf("This user Has already been scanned :%s", user.Email)
		}
	}
	return 0, nil
}

// Tieck is used for checking list
func Tieck(ID string) {
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	orm.ShowSQL(true)
	_, err = orm.Where("i_d = ?", ID).Cols("presencecheck").Update(&models.User{Check: true})
	if err != nil {
		fmt.Println(err)
		return
	}
}

//  Count counts the users who are present at the event
func Count(w rest.ResponseWriter, req *rest.Request) {
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	var users []models.User
	orm.ShowSQL(true)

	len, err := orm.Where("presencecheck = ?", true).FindAndCount(&users)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.WriteJson(&users)
	w.WriteJson(len)
}

// CheckList shows the users in a checklist
func CheckList(w rest.ResponseWriter, req *rest.Request) {

	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	orm.ShowSQL(true)

	users := new([]models.User)
	err = orm.Where("presencecheck = ?", true).Find(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteJson(&users)
}

// GetAllUsers Serves all the users from the api
func GetAllUsers(res rest.ResponseWriter, req *rest.Request) {
	fmt.Println("list called")
	users := models.Users{
		Store: map[string]*models.User{},
	}
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	user, err := users.Getuser(orm)
	if err != nil {
		fmt.Println(err)
		return
	}
	orm.ShowSQL(true)

	res.WriteJson(&user)
}

// PostUser adds the user captured to the database
func PostUser(w rest.ResponseWriter, req *rest.Request) {
	users := models.Users{
		Store: map[string]*models.User{},
	}
	var user models.User
	err := req.DecodeJsonPayload(&user)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Email == "" {
		rest.Error(w, "user email required", 400)
		return
	}
	if user.Name == "" {
		rest.Error(w, "user name required", 400)
		return
	}
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}

	if err = users.Adduser(user, orm); err != nil {
		rest.Error(w, err.Error(), 400)
		return
	}
	orm.ShowSQL(true)

	w.WriteJson(&user)
}

// GetuserbyName serves the users by matching names
func GetuserbyName(w rest.ResponseWriter, req *rest.Request) {
	name := req.PathParam("name")
	var users = new([]models.User)
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	orm.ShowSQL(true)
	if err := orm.Where("name = ?", name).Find(users); err != nil {
		return
	}
	if users == nil {
		rest.NotFound(w, req)
		return
	}
	w.WriteJson(users)
}

// GetuserbyEmail serves the users by matching emails
func GetuserbyEmail(w rest.ResponseWriter, req *rest.Request) {
	email := req.PathParam("mail")
	var users = new([]models.User)
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	orm.ShowSQL(true)

	if err := orm.Where("email = ?", email).Find(users); err != nil {
		return
	}
	if users == nil {
		rest.NotFound(w, req)
		return
	}
	w.WriteJson(users)
}

// Deleteuser deletes the user with the specified name
func Deleteuser(w rest.ResponseWriter, req *rest.Request) {
	name := req.PathParam("demail")
	var users = new(models.User)
	orm, err := SessionSqlite(f)
	if err != nil {
		log.Printf("Error Creating Database Session: %s", err)
	}
	orm.ShowSQL(true)
	if _, err := orm.Where("name = ?", name).Delete(users); err != nil {
		w.WriteJson(err)
		return
	}
	if users == nil {
		rest.NotFound(w, req)
		return
	}
	w.WriteJson(users)
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// curl -i -H 'Content-Type: application/json' \
//     -d '{"name": "Drake",
//     "phone_no": "2519876543",
//     "email": "Drakeman@gmail.com",
//     "role": "Guest",
//     "presencecheck": false,
//     "qcode": "./images/asdasdasd.png",
//     "Created": "2018-08-22T14:02:46+03:00",
//     "Updated": "2018-08-22T14:02:46+03:00"}' http://127.0.0.1:1988/adduser
