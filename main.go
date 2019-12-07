package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID string
	Name string
	Description string
}

func main() {
	user := fetchUser("yyh-gl")
	fmt.Println("============ RESULT ============")
	fmt.Printf("%+v\n", user)
	fmt.Println("============ RESULT ============")
}

func fetchUser(id string) (user *User) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/users/yyh-gl", nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &user)
	return user
}
