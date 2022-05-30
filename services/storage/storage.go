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
	Routes   map[string]string `json:"routes"`
	Users    []User            `json:"users"`
}

func NewStorageManager() *StorageManager {

	storageManager := StorageManager{}

	file, _ := ioutil.ReadFile(database)

	_ = json.Unmarshal([]byte(file), &storageManager)

	return &storageManager
}

func (s *StorageManager) save() {

	file, err := json.MarshalIndent(s, "", " ")

	if err != nil {
		return
	}

	_ = ioutil.WriteFile(database, file, 0666)
}

func (s *StorageManager) GetProxyUrl(host string) string {
	if r, ok := s.Routes[host]; ok {
		return r
	}

	return ""
}

func (s *StorageManager) GetActiveUser() map[string]string {
	for _, u := range s.Users {
		if u.Active {
			return u.Headers
		}
	}
	return make(map[string]string)
}

func (s *StorageManager) SetConfig(routes map[string]string) {
	s.Routes = routes
	s.save()
}

func (s *StorageManager) AddUser(user User) User {
	id := randstr.String(20)

	user.Id = id
	user.Active = false
	s.Users = append(s.Users, user)
	s.save()
	return user
}

func (s *StorageManager) EditUser(id string, user User) User {
	var u User
	for i, v := range s.Users {
		if v.Id == id {
			s.Users[i].Headers = user.Headers
			u = s.Users[i]
		}
	}
	s.save()
	return u
}

func (s *StorageManager) ActivateUser(id string) {
	activated := false
	for i, v := range s.Users {
		if v.Id == id {
			activated = !s.Users[i].Active
			s.Users[i].Active = activated
		} 
	}
	if activated {
		for i, v := range s.Users {
			if v.Id != id {
				s.Users[i].Active = false
			} 
		}
	}
	s.save()
}

func (s *StorageManager) DeleteUser(id string) {
	for i, v := range s.Users {
		if v.Id == id {
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
		}
	}
	s.save()
}
