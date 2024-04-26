package users

import (
	"context"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
	dbs "timely/config/db"
	"timely/config/responses"
	httplib "timely/libs/http"
	crypt "timely/util/crypto"

	usersmodel "timely/models/users"

	"github.com/spf13/viper"
)

var (
	env          = viper.GetString("env")
	dbName       = "timely"
	dbCollection = "users"
)

// RegisterUser controller
func RegisterUser(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	c := httplib.C{Req: req, Res: res}

	db, err := dbs.ConnectMongodbURL()

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
		return
	}

	defer db.Disconnect(ctx)

	coll := db.Database(dbName).Collection(dbCollection)

	var data usersmodel.Users

	c.BindJSON(&data)

	data.ID = bson.NewObjectId()
	data.Date = time.Now()
	hash := crypt.HashText(data.Password)
	data.Password = hash
	_, err = coll.InsertOne(ctx, data)
	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
		return
	}

	token := crypt.Jwt(data.ID)

	//mask password
	data.Password = ""
	details := struct {
		User  interface{} `json:"user"`
		Token string      `json:"token"`
	}{User: data, Token: token}

	resp := responses.GeneralResponse{Success: true, Data: details, Message: "user created"}
	httplib.Response(res, resp)
}

// GetUserDetailsByEmail controller
func GetUserDetailsByEmail(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	c := httplib.C{Req: req, Res: res}

	db, err := dbs.ConnectMongodbURL()

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
		return
	}

	defer db.Disconnect(ctx)

	coll := db.Database(dbName).Collection(dbCollection)

	email := c.Params("email")

	var user usersmodel.Users

	err = coll.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error getting user"}
		httplib.Response400(res, resp)
		return
	}

	resp := responses.GeneralResponse{Success: true, Data: user, Message: "user details"}
	httplib.Response(res, resp)
}

// DeleteUser controller
func DeleteUser(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	c := httplib.C{Req: req, Res: res}

	db, err := dbs.ConnectMongodbURL()

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
		return
	}

	defer db.Disconnect(ctx)

	coll := db.Database(dbName).Collection(dbCollection)
	email := c.Params("email")

	_, err = coll.DeleteOne(ctx, bson.M{"email": email})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error deleting user"}
		httplib.Response400(res, resp)
		return
	}

	resp := responses.GeneralResponse{Success: true, Data: email, Message: "user deleted"}
	httplib.Response(res, resp)
}

// UpdateUserDetails controller
func UpdateUserDetails(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	c := httplib.C{Req: req, Res: res}

	db, err := dbs.ConnectMongodbURL()

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error creating user"}
		httplib.Response400(res, resp)
		return
	}

	defer db.Disconnect(ctx)

	coll := db.Database(dbName).Collection(dbCollection)

	var updates bson.M

	c.BindJSON(updates)
	userID := c.Params("id")
	_, err = coll.UpdateByID(ctx, bson.M{"_id": bson.ObjectIdHex(userID)}, bson.M{"$set": updates})

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating user"}
		httplib.Response(res, resp)
		return
	}

	var user usersmodel.Users

	err = coll.FindOne(ctx, bson.M{"_id": bson.ObjectIdHex(userID)}).Decode(&user)
	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: err.Error(), Message: "error updating user"}
		httplib.Response(res, resp)
		return
	}

	resp := responses.GeneralResponse{Success: true, Data: user, Message: "user updated"}
	httplib.Response(res, resp)
}
