package config

import (
	"encoding/json"
	"fmt"

	"github.com/shotastage/GFileable"
)

type UserConfig struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func InitialConfig() {

	conf := UserConfig{}

	conf.Name = ""
	conf.Email = ""
	conf.Username = ""

	outputJson, err := json.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	GFileable.Touch("UserConfig.json")

	file := GFileable.Path("UserConfig.json")

	file.WriteString(fmt.Sprintf("%s", outputJson))

}
