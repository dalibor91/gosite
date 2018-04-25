package user

import m "app/model"

type UserModel struct {
	m.DefaultModel

	Id string
	Username string
	Password string
	Email string
}

func (u * UserModel) GetTableName() string {
	return "users"
}

func (u * UserModel) Update() bool {
	return true
}

func (u * UserModel) Save() bool {
	return true
}

func (u * UserModel) Delete() bool {
	return true
}


