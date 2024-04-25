package routes

import (
	"github.com/rs/cors"
	"net/http"
	"timely/config/responses"
	authController "timely/controllers/v1/auth"
	usersController "timely/controllers/v1/users"
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
	//	APPLY MIDDLEWARES
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"*"},
	})
	route.Use(mws.AccessLogToConsole)

	route.Use(c.Handler)

	//*****************
	// AUTH ROUTES
	//*****************
	authRoute := route.PathPrefix("/v1/auth").Subrouter()
	authRoute.HandleFunc("/login", authController.Login).Methods("POST")
	//mwsWithAuth adds authorization token to endpoints
	mwsWithAuth := mws.AuthorizationSingle

	//************************
	// USERS  ROUTES
	//************************
	usersRoute := route.PathPrefix("/v1/users").Subrouter()
	usersRoute.HandleFunc("", usersController.RegisterUser).Methods("POST")
	usersRoute.HandleFunc("/{email}", mwsWithAuth(usersController.GetUserDetailsByEmail)).Methods("GET")
	usersRoute.HandleFunc("/{id}", mwsWithAuth(usersController.UpdateUserDetials)).Methods("PUT")
	usersRoute.HandleFunc("/{email}", mwsWithAuth(usersController.DeleteUser)).Methods("DELETE")
	return route
}
