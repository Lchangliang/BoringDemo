package main

import (
	"awesomeProject/auth"
	"awesomeProject/token"
	"strconv"
	"time"
	"fmt"
)

func main() {
	var apiAuthenticator auth.ApiAuthenticator
	apiAuthenticator = auth.New()
	params := make(map[string]string)
	params["appId"] = "liuchangliang"
	params["password"] = "0621"
	timestamp := time.Now().UnixNano()/1e6
	authToken := token.New("www.baidu.com?id=123", timestamp, params)
	time.Sleep(90*time.Second)
	error := apiAuthenticator.AuthByUrl("www.baidu.com?id=123&liuchangliang&"+authToken.GetToken()+"&"+strconv.FormatInt(timestamp, 10))
	if error != nil {
		fmt.Println(error.Error())
	}
}
