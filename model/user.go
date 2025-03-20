package model

import "encoding/json"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
