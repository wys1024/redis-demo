package model

import "encoding/json"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// MarshalBinary 将 User 结构体序列化为 JSON 格式的字节切片。
// 该方法实现了 encoding.BinaryMarshaler 接口，用于将 User 对象转换为二进制数据。
func (u *User) MarshalBinary() (data []byte, err error) {
	// 使用 json.Marshal 将 User 结构体序列化为 JSON 格式的字节切片
	return json.Marshal(u)
}
