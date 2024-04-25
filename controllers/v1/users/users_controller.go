package visitorscontoller

import (
	"net/http"
	"time"
	dbs "timely/config/db"
	"timely/config/responses"
	httplib "timely/libs/http"

	"gopkg.in/mgo.v2/bson"

	usersmodel "timely/models/users"

	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

var (
	env          = viper.GetString("env")
	dbName       = "timely"
	dbCollection = "users"
)

// RegisterUser controller
func RegisterUser(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var data usersmodel.Users

	c.BindJSON(&data)

	data.ID = bson.NewObjectId()
	data.Date = time.Now()

	err := coll.Insert(data)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: data, Message: "user created"}
	httplib.Response(res, resp)
}

// GetUserDetailsByEmail controller
func GetUserDetailsByEmail(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	email := c.Params("email")

	var user interface{}

	err := coll.Find(bson.M{"email": email}).One(&user)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting user"}
		httplib.Response400(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: user, Message: "user details"}
	httplib.Response(res, resp)
}

// DeleteUser controller
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}

	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	email := c.Params("email")

	err := coll.Remove(bson.M{"email": email})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error deleting visitor"}
		httplib.Response400(res, resp)
	}

	resp := responses.GeneralResponse{Success: true, Data: email, Message: "visitor deleted"}
	httplib.Response(res, resp)
}

// UpdateUserDetials controller
func UpdateUserDetials(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Req: req, Res: res}
	var db *mgo.Session
	env := viper.GetString("env")

	if env == "prod" {
		db = dbs.ConnectMongodbTLS()
	} else {
		db = dbs.ConnectMongodb()
	}

	defer db.Close()

	coll := db.DB(dbName).C(dbCollection)

	var updates bson.M

	c.BindJSON(updates)
	userID := c.Params("id")

	err := coll.Update(bson.M{"_id": bson.ObjectIdHex(userID)}, bson.M{"$set": updates})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating visitor"}
		httplib.Response(res, resp)
	}
	resp := responses.GeneralResponse{Success: true, Data: updates, Message: "visitor updated"}
	httplib.Response(res, resp)
}
