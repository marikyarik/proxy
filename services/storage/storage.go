package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/thanhpk/randstr"
)

const database = "database/db.json"

type User struct {
	Id      string            `json:"id"`
	Active  bool              `json:"active"`
	Headers map[string]string `json:"headers"`
}

type StorageManager struct {
	ProxyUrl string `json:"proxy_url"`
	Users    []User `json:"users"`
}

func NewStorageManager() *StorageManager {

	storageManager := StorageManager{}

	file, _ := ioutil.ReadFile(database)

	_ = json.Unmarshal([]byte(file), &storageManager)

	return &storageManager
}

func (s *StorageManager) Save() {

	file, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		return
	}

	_ = ioutil.WriteFile(database, file, 0666)
}

func (s *StorageManager) GetActiveUser() map[string]string {
	for _, u := range s.Users {
		if u.Active {
			return u.Headers
		}
	}
	return make(map[string]string)
}

func (s *StorageManager) AddUser(user User) string {
	id := randstr.String(20)

	user.Id = id;
	user.Active = false;
	s.Users = append(s.Users, user)
	s.Save()
	return id
}

func (s *StorageManager) EditUser(id string, user User) {
	for i, v := range s.Users {
        if v.Id == id {
            s.Users[i] = user
        }
    }
	s.Save()
}

func (s *StorageManager) ActivateUser(id string) {
	for i, v := range s.Users {
        if v.Id == id {
            s.Users[i].Active = !s.Users[i].Active
        }
    }
	s.Save()
}

func (s *StorageManager) DeleteUser(id string) {
	for i, v := range s.Users {
        if v.Id == id {
            s.Users = append(s.Users[:i], s.Users[i+1:]...)
        }
    }
	s.Save()
}
