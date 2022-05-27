package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type StorageUnit struct {
	Dir         string `yaml:"directory"`
	Url         string `yaml:"url"`
	RefreshRate string `yaml:"refresh-rate"`
}

type Config struct {
	Storage []StorageUnit `yaml:"storage"`
}

func readConf() {
	filename, _ := filepath.Abs("./conf.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Value: %#v\n", config.Storage)
}

func main() {
	// readConf()
	fs := http.FileServer(http.Dir("./shared"))
	http.Handle("/", fs)

	log.Print("[-] sglfs Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
