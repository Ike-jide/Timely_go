package auth

import (
	"net/http"
	"strings"
	"timely/config/responses"
	httplib "timely/libs/http"
	loginModel "timely/models/auth"
)

// Login controller
func Login(res http.ResponseWriter, req *http.Request) {
	c := httplib.C{Res: res, Req: req}

	var data *loginModel.Login

	c.BindJSON(&data)

	data.Email = strings.ToLower(data.Email)

	userAccount, token, err := CheckHashAndUpdate(data, data.Password)

	if err != nil {
		resp := responses.GeneralResponse{Success: false, Error: "Error authenticating account", Message: err.Error()}
		httplib.Response400(res, resp)
		return
	}

	datum := struct {
		User  interface{} `json:"user"`
		Token string      `json:"token"`
	}{User: userAccount, Token: token}

	resp := responses.GeneralResponse{Success: true, Data: datum, Message: "Authentication successful"}
	httplib.Response(res, resp)
}

func CheckHashAndUpdate(data *loginModel.Login, password string) (interface{}, string, error) {
	return nil, "", nil
}
