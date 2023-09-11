package models

type User struct {
	Username string `bson:"username"`
	Salt     []byte `bson:"salt"`
	Password []byte `bson:"password"`
}