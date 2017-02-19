package customer

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
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

// type Detail struct {
// 	Value map[string]string `json:"value"`
// }

//-----------------------------------------
// for accountdetails not needed now
// --------------------------------------------

// type AccountDetail struct {
// 	//Accounts        map[int]map[string]string `json:"accounts"` // not sure about this line
// 	Accounts        [2]map[string]int `json:"accounts"`
// 	CustId          string            `json:"custId"`
// 	Message         string            `json:"message"`
// 	MobileNo        string            `json:"mobileNo"`
// 	SuccesOrFailure string            `json:"successOrFailure"`
// }
// type Account struct {
// 	CustomerDetail AccountDetail `json:"customerDetail"`
// }
type DataOtp struct {
	Status    string `json:"status"`
	AuthToken string `json:"auth_token"`
	ApiKey    string `json:"api_key"`
	//AcctDetails Account `json:"acctdetails"`
}
type ResponseOtp struct {
	StatusCode    string  `json:"status_code"`
	StatusMessage string  `json:"status_message"`
	Request       Method  `json:"request"`
	Response      DataOtp `json:"response"`
	Version       int     `json:"version"`
}
type Response struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Request       Method `json:"request"`
	Response      Data   `json:"response`
	Version       int    `json:"version"`
}
type StdResponse struct {
	StatusCode    string       `json:"status_code"`
	StatusMessage string       `json:"status_message"`
	Request       Method       `json:"request"`
	Version       int          `json:"version"`
	Response      DataResponse `json:"response"`
}
type DataResponse struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobile_number"`
	CustomerType string `json:"customer_type"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) { // url /users/register
	log.Println("this reggisteration")
	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "GET called", 500)
		return
	}

	r.ParseForm() // parse the form
	//r.Form.Encode gives value in string form "name=Ava"
	method := Method{Method: r.URL.Path}
	uuid, err := Uuid()
	if err != nil {
		http.Error(w, "GET called", 500)
		return
	}
	data := Data{Status: "200", Uuid: uuid, UserId: 1}
	todos := Response{StatusCode: "200", StatusMessage: "success", Request: method, Response: data, Version: 1}
	json.NewEncoder(w).Encode(todos)
}
func OtpVerifyHandler(w http.ResponseWriter, r *http.Request) { // url /users/:id/verify

	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "Wrong method", 500)
		return
	}
	// map[int]map[string]string{
	// 		0: map[string]string{
	// 			"name":  "Hydrogen",
	// 			"state": "gas"}}
	r.ParseForm() // parse the form
	//r.Form.Encode gives value in string form "name=Ava"

	//-----------------------------------------------------------
	// aom1 := map[string]int{"one": 1, "two": 4, "three": 9}
	// aom2 := map[string]int{"one": 1, "two": 4, "three": 9}
	// var arr [2]map[string]int
	// arr[0] = aom1 // put map into the array
	// arr[1] = aom2
	//-----------------------------------------------------------

	method := Method{Method: r.URL.Path}

	//-----------------------------------------------------------
	// accounts := AccountDetail{Accounts: arr, CustId: "1234", Message: "qweq", MobileNo: "9061415632", SuccesOrFailure: "S"}
	// account := Account{CustomerDetail: accounts}
	//-----------------------------------------------------------

	dataotp := DataOtp{Status: "200", AuthToken: "111", ApiKey: "1111"}
	todos := ResponseOtp{StatusCode: "200", StatusMessage: "success", Request: method, Response: dataotp, Version: 1}
	json.NewEncoder(w).Encode(todos)
}
func AccountHandler(w http.ResponseWriter, r *http.Request) { // url /users/:id/accounts

	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "Wrong method", 500)
		return
	}
	r.ParseForm() // parse the form
	dataaccount := DataResponse{Name: "Rony", MobileNumber: "9061415632", CustomerType: "VIP"}
	method := Method{Method: r.URL.Path}
	todos := StdResponse{StatusCode: "200", StatusMessage: "success", Request: method, Response: dataaccount, Version: 1}
	json.NewEncoder(w).Encode(todos)
}
func TransHandler(w http.ResponseWriter, r *http.Request) { // url /users/:id/transactions

	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "Wrong method", 500)
		return
	}
	r.ParseForm() // parse the form
	dataaccount := DataResponse{Name: "Rony", MobileNumber: "9061415632", CustomerType: "VIP"}
	method := Method{Method: r.URL.Path}
	todos := StdResponse{StatusCode: "200", StatusMessage: "success", Request: method, Response: dataaccount, Version: 1}
	json.NewEncoder(w).Encode(todos)
}
func InboxHandler(w http.ResponseWriter, r *http.Request) { // url /messages

	// check method is Post or not
	if r.Method != "POST" {
		http.Error(w, "Wrong method", 500)
		return
	}
	r.ParseForm() // parse the form
	dataaccount := DataResponse{Name: "Rony", MobileNumber: "9061415632", CustomerType: "VIP"}
	method := Method{Method: r.URL.Path}
	todos := StdResponse{StatusCode: "200", StatusMessage: "success", Request: method, Response: dataaccount, Version: 1}
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
