package user

type User struct {
	ID           string `json:"id" bson:"_id, omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password_hash"`
	Email        string `json:"email" bson:"email"`
}

type CreateUserDto struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}
