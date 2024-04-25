package routes

import (
	"net/http"
	"timely/config/responses"
	userscontoller "timely/controllers/v1/users"
	httplib "timely/libs/http"
	mws "timely/middlewares"

	"github.com/gorilla/mux"
)

// Router for all routes
func Router() *mux.Router {
	route := mux.NewRouter()

	//BASE ROUTE
	route.HandleFunc("/v1", func(res http.ResponseWriter, req *http.Request) {
		resp := responses.GeneralResponse{Success: true, Message: "timely  server running....", Data: "vsm SERVER v1.0"}
		httplib.Response(res, resp)
	})

	route.Use(mws.AccessLogToConsole)

	//************************
	// USERS  ROUTES
	//************************
	usersRoute := route.PathPrefix("/v1/users").Subrouter()
	usersRoute.HandleFunc("", userscontoller.RegisterUser).Methods("POST")
	usersRoute.HandleFunc("/{email}", userscontoller.GetUserDetailsByEmail).Methods("GET")
	usersRoute.HandleFunc("/{id}", userscontoller.UpdateUserDetials).Methods("PUT")
	usersRoute.HandleFunc("/{email}", userscontoller.DeleteUser).Methods("DELETE")
	return route
}
