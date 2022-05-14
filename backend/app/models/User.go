package models

import (
	"encoding/json"
	"strings"
)

type User struct {
	// Id       int64  `json:"id"`
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Created  string `json:"created"`
}

// hset转码接口，无用
func (user User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(user)
}
func (user User) UnmarshalBinary(data []byte) error {
	dec := json.NewDecoder(strings.NewReader(string(data)))
	return dec.Decode(&user)
}

// jwt接口方法
func (user User) GetUid() string {
	// return strconv.Itoa(int(user.Id))
	return user.Id
}
