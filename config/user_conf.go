package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mirage-go/core"
	"mirage-go/shared"

	"github.com/shotastage/GFileable"
)

type UserConfig struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func InitialConfig() {

	conf := UserConfig{}

	print("Your full name: ")
	conf.Name = core.StrStdin()

	print("Your email: ")
	conf.Email = core.StrStdin()

	print("Your username: ")
	conf.Username = core.StrStdin()

	outputJson, err := json.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	out := new(bytes.Buffer)
	json.Indent(out, outputJson, "", "    ")

	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	//GFileable.Touch(shared.UserConfigPath + "/UserConfig.json")

	file := GFileable.Join(shared.UserConfigPath, "/UserConfig.json")

	file.WriteString(fmt.Sprintf("%s", out))
}
