package main

import (
	_ "app"
	"base64"
	"controller"
	"crypto/aes"
	_ "crypto/cipher"
	_ "database/sql"
	_ "encoding/base64"
	"encrypt"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_"gopkg.in/yaml.v2"
	"io/ioutil"
	_ "log"
	"net/http"
	"path/filepath"
	"yaml2"
)

type Page struct {
	Title string
	Body  []byte
}

type Config struct {
	Src Options
}
type Options struct {
	Iv  string
	Key string
}

var (
	id   int
	name string
)

func main() {
	//key := []byte("H\xDE\xD9\xFOy7z3Q4JSWvzxGGjOkNGyw==F\x1D\xD37\xC8\xF1N\x00y\xF6\xE9\xCE\xF3")
	//plaintext := []byte("123456")
	file, _ := filepath.Abs("./config/FedbookCategory.xlsx")
	//var a Config\
	defer fmt.Println("defer sta")
	a := yaml2.Yaml("./config/text.yml")
	fmt.Printf("Value: %#v\n", a.Src)
	yamlFile, _ := ioutil.ReadFile(file)

	fmt.Println(string(yamlFile))
	//var config Config
	//err := yaml.Unmarshal(yamlFile, &config)
	//if err != nil {
	//	panic(err)
	//}
	//text := []byte("askhgasgh")
	//fmt.Printf("Value: %#v\n", config.Src)
	//err2 := ioutil.WriteFile("/home/rony/Documents/go/hello/tmp/dat1", yamlFile, 0644)
	//if err2 != nil {
	//	fmt.Printf("error")
	//}
	// fmt.Println("ashdhfajsdlhasdfhsssssssssssssssssssssssssssssssssssssssss")
	// db, err := sql.Open("mysql",
	// 	"root:123456@tcp(127.0.0.1:3306)/idbi_dev_dump")
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("ss")
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Println("error is db")
	// }
	// stmt, err := db.Prepare("select id, msisdn from customers where id = ?")
	// if err != nil {
	// 	log.Println("error in row6575")
	// }
	// defer stmt.Close()

	// rows, err := stmt.Query(3)
	// if err != nil {
	// 	fmt.Println("+++++++")
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// fmt.Println(rows)

	// err2 := rows.Scan(&id, &name)
	// if err2 != nil {
	// 	fmt.Println("hi")
	// }
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		fmt.Println("error in row1")
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	fmt.Println("error in row2")
	// }

	// fmt.Println(id, name)
	input := []byte("this is a test")
	iv1, _ := base64.Decode(a.Src.Iv)
	iv := iv1[:aes.BlockSize]
	key, _ := base64.Decode(a.Src.Key)
	encrypted := make([]byte, len(input))
	encrypt.EncryptAES(encrypted, input, key, iv)
	fmt.Println(string(encrypted))
	encrypt.DecryptAES(encrypted, input, key, iv)
	fmt.Println(string(encrypted))
	//app.Run
	http.HandleFunc("/view/", controller.ViewHandler)
	http.HandleFunc("/edit/", controller.EditHandler)
	http.HandleFunc("/save/", controller.SaveHandler)
	http.ListenAndServe(":8085", nil)
}
