package models

type UserModel struct {
	ID   string `firestore:"id"`
	Name string `firestore:"name"`
}
