package main

import (
	"fmt"
	"encoding/json"
	"log"
	"io/ioutil"
	"bufio"
	"os"
	"os/exec"
)

func main(){
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("An error occured while trying to open config.json - " + err.Error())
	}
	config, _ := UnmarshalConfig(configFile)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please enter your search query: ")
		var input string
		input, _ = reader.ReadString('\n')
		url := 'https://google.com/search?q=' + input + '&cr=country' + config.Countrycode
		exec.Command("rund1132", "url.dll,FileProtocolHandler", url).Start()
	}
}

func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Config) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Config struct {
	Countrycode string 'json:"countrycode"'
}
