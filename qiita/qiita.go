package qiita

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient(c *http.Client) Client {
	return Client{c}
}

type User struct {
	ID string
	Location string
}

func (c Client) FetchUser(id string) (user *User) {
	req, _ := http.NewRequest("GET", "https://qiita.com/api/v2/users/" + id, nil)

	resp, _ := c.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &user)
	return user
}
