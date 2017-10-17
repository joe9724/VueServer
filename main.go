package main


import (
_ "github.com/go-sql-driver/mysql"
"log"
"net/http"
"fmt"
"encoding/json"
)
type LoginModel struct {
	Error       bool    `json:"error"`
	User  string `json:"user"`
	Token string `json:"token"`
	Redirect string `json:"redirect"`

}
func main() {
	fmt.Println("server start success...")
	http.HandleFunc("/login", Login)
	http.HandleFunc("/table",Table)
	err := http.ListenAndServe(":88", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	var LoginReturn LoginModel
        LoginReturn.Error = false
	LoginReturn.User = "pf"
	LoginReturn.Token = "token"
	LoginReturn.Redirect = ""

	data, err := json.Marshal(LoginReturn)
	if err!=nil{
		fmt.Println("marshal error")
	}
	defer func() {
		w.Write(data)
	}()
}

//
func Table(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	var TableArr []LoginModel

	for i:=0;i<20;i++{
		var LoginReturn LoginModel
		LoginReturn.Error = false
		LoginReturn.User = "pf"
		LoginReturn.Token = "token"
		LoginReturn.Redirect = ""
		TableArr = append(TableArr,LoginReturn)
	}


	data, err := json.Marshal(TableArr)
	if err!=nil{
		fmt.Println("marshal error")
	}
	defer func() {
		w.Write(data)
	}()
}