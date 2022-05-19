package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/thanhpk/randstr"
)

const database = "database/db.json"

type StorageManager struct {
	ProxyUrl  string                 `json:"proxy_url"`
	Users     map[string]interface{} `json:"users"`
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

func (s *StorageManager) GetActiveUser() string {
	for _, v := range s.Users {
		u := v.(map[string]interface{})
		if u["active"].(bool) {
			jsonUser, err := json.Marshal(u)
			if err != nil {
				return ""
			}

			return string(jsonUser)
		}
	}
	return ""
}

func (s *StorageManager) AddUser(user interface{}) string {
	hash := randstr.String(20)

	s.Users[hash] = user

	s.Save()
	return hash
}

func (s *StorageManager) EditUser(hash string, user interface{}) {
	s.Users[hash] = user
	s.Save()
}

func (s *StorageManager) DeleteUser(hash string) {
	delete(s.Users, hash)
	s.Save()
}
