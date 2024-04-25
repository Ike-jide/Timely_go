package auth

// Login model
type Login struct {
	Password string `json:"password"  bson:"password"`
	Email    string `json:"email" bson:"email"`
}
