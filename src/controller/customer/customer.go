package customer

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net/http"
)

type Method struct {
	Method string `json:"method"`
}
type Data struct {
	Status string `json:"status"`
	Uuid   string `json:"uuid"`
	UserId int    `json:"user_id"`
}
type Response struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Request       Method `json:"request"`
	Response      Data   `json:"response`
	Version       int    `json:"version"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "GET called", 500)
		return
	}

	r.ParseForm() // parse the form
	//r.Form.Encode gives value in string form "name=Ava"
	method := Method{Method: r.URL.Path}
	uuid, _ := Uuid()
	data := Data{Status: "200", Uuid: uuid, UserId: 1}
	todos := Response{StatusCode: "200", StatusMessage: "success", Request: method, Response: data, Version: 1}
	json.NewEncoder(w).Encode(todos)
}

func RandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Uuid() (string, error) {
	var first, last uint32
	var middle [4]uint16
	randomBytes, err := RandomBytes(16)
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBuffer(randomBytes)
	binary.Read(buffer, binary.BigEndian, &first)
	for i := 0; i < 4; i++ {
		binary.Read(buffer, binary.BigEndian, &middle[i])
	}
	binary.Read(buffer, binary.BigEndian, &last)
	middle[1] = (middle[1] & 0x0fff) | 0x4000
	middle[2] = (middle[2] & 0x3fff) | 0x8000
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%04x%08x",
		first, middle[0], middle[1], middle[2], middle[3], last), nil
}
