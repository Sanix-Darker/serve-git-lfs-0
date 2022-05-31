package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

var MASTERBRANCH = "master"
var SHAREDDIR = "./shared"

func readConf() Config {
	filename, _ := filepath.Abs("./conf.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func execCommand(bin string, args ...string) {
	cmd := exec.Command(bin, args...)
	// because we execute all our commands inside the shared directory
	cmd.Dir = SHAREDDIR

	stdout, err := cmd.Output()

	if err != nil {
		log.Panic(err)
	}

	// Print the output
	log.Print(string(stdout))
}

func gitClone(path string, url string) {
	// We want to clone only the subdirectory not the whole project
	// we parse to get the repo name
	urlParts := strings.Split(url, "/")
	repoName := urlParts[len(urlParts)-1]

	execCommand("git", "init", repoName)
	os.Chdir(repoName)
	commands := [...]string{
		"git remote add origin " + url,
		"git config core.sparsecheckout true",

		"echo " + path + "/* .git/info/sparse-checkout",
		"git pull origin " + MASTERBRANCH,
	}
	for _, cmdd := range commands {
		binn := strings.Split(cmdd, " ")[0]
		args := strings.Split(cmdd, " ")[1:]
		// strArgs := strings.Join(args, " ")
		log.Print(binn)
		log.Print(args)
		execCommand(binn, args...)
	}
}

func gitPull(path string, url string) {

}

func main() {
	execCommand("git-lfs", "--version")
	execCommand("git", "--version")
	gitClone("audio-files", "https://github.com/osscameroon/podcasts")
	fs := http.FileServer(http.Dir(SHAREDDIR))
	http.Handle("/", fs)

	log.Print("[-] sglfs Listening on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
