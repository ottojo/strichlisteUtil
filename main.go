package strichlisteUtil

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

type StrichlisteApi struct {
	Url string
}

func (s *StrichlisteApi) GetUser(id int) (User, error) {
	var u User

	resp, err := http.Get(s.Url + "/api/user/" + strconv.Itoa(id))
	if err != nil {
		return User{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *StrichlisteApi) GetUsers() ([]User, error) {
	var u usersResponse
	resp, err := http.Get(s.Url + "/api/user")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		return nil, err
	}
	return u.Entries, err
}

func (s *StrichlisteApi) GetTransactions(u *User) ([]Transaction, error) {
	var t transactionsResponse
	resp, err := http.Get(s.Url + "/api/user/" + strconv.Itoa(u.ID) + "/transaction")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}
	return t.Entries, err
}
