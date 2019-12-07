package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID string
	Location string
}

func FetchUser(id string) (user *User) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/users/" + id, nil)

	client := new(http.Client)
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &user)
	return user
}
