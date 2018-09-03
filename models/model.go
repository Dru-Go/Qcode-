package model

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/nu7hatch/gouuid"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/go-xorm/xorm"
)

var f = "db/test.db"

// User hold information about every user in using the system
type User struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	PhoneNo string    `json:"phone_no,omitempty"`
	Email   string    `json:"email,omitempty"`
	Role    string    `json:"role,omitempty"`
	Qcode   string    `json:"qcode,omitempty" xorm:"qlocation"`
	Check   bool      `json:"check,omitempty" xorm:"presencecheck"`
	Created time.Time `xorm:"created" json:"created,omitempty"`
	Updated time.Time `xorm:"updated" json:"updated,omitempty"`
}

type Users struct {
	sync.RWMutex
	Store map[string]*User
}

var users = Users{
	Store: map[string]*User{},
}

// Adduser Add users to the database
func (u Users) Adduser(users User, orm *xorm.Engine) error {
	u.Lock()
	id, _ := uuid.NewV4()
	users.ID = id.String()
	orm.ShowSQL(true)
	users.Check = false
	users.Qcode = Mapper(users)
	_, err := orm.Insert(&users)
	if err != nil {
		return err
	}
	u.Unlock()
	return nil
}

// Mapper maps the users to the qrcode
func Mapper(user User) string {
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

// Getuser Gets the users info from the database
func (u Users) Getuser(orm *xorm.Engine) ([]User, error) {
	u.Lock()
	// orm.ShowSQL(true)
	var info []User
	err := orm.Find(&info)
	if err != nil {
		return nil, err
	}
	u.Unlock()
	return info, nil

}

// UpuserID updates the user by id
func UpuserID(id uuid.UUID, users User, orm *xorm.Engine) error {
	orm.ShowSQL(true)
	_, err := orm.ID(id).Update(&users)
	if err != nil {
		return err
	}
	return nil
}

// UpuserID updates the user by id
func GetID(id string, orm *xorm.Engine) (*User, error) {
	var user = new(User)
	orm.ShowSQL(true)
	err := orm.ID(id).Find(users)
	if err != nil {
		log.Println(err)
	}
	return user, nil
}

// UpuserName updates the user by id
func UpuserName(id string, users User, orm *xorm.Engine) error {
	orm.ShowSQL(true)
	_, err := orm.Update(&users)
	if err != nil {
		return err
	}
	return nil
}

// Delusers deletes the user from the database
func Delusers(id int, users User, orm *xorm.Engine) error {
	orm.ShowSQL(true)
	_, err := orm.ID(id).Delete(&users)
	if err != nil {
		return err
	}
	return nil
}

// Parse is responsible for converting some barcode to json
func Parse(bar barcode.Barcode) ([]byte, error) {
	img, err := json.Marshal(bar)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// SaveImage saves the specified barcode to image and store it in images folder
func SaveImage(bar barcode.Barcode, name string) (string, error) {
	// Scale the barcode to 200x200 pixels
	qrCode, _ := barcode.Scale(bar, 200, 200)
	name = "./images" + name
	// create the output file
	file, _ := os.Create(name)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
	return name, nil
}
