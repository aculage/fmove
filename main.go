package main

import (
	"fmove/internal/get"
	"fmove/internal/send"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/go-yaml/yaml"
)

type ProgramInfo struct {
	Program  string                       `yaml:"program"`
	Version  string                       `yaml:"version"`
	Author   string                       `yaml:"author"`
	Commands map[string]map[string]string `yaml:"commands"`
}

func LoadProgramInfo() ProgramInfo {
	f, err := ioutil.ReadFile("meta.yaml")
	if err != nil {
		log.Fatal("Config file missing -- consider fetching by fmove fetch")
	}
	pinfo := ProgramInfo{}
	yaml.Unmarshal(f, &pinfo)
	return pinfo
}

func main() {
	args := os.Args[1:]

	pinfo := LoadProgramInfo()

	if len(args) == 0 {
		args = append(args, "help")
	}
	switch args[0] {
	case "get":
		address, err := net.ResolveTCPAddr("tcp", "localhost:9000")
		log.Printf("Awaiting for connections at %v:%v\n", address.IP, address.Port)
		if err != nil {
			log.Fatal(err.Error())
		}
		server, err := net.ListenTCP("tcp", address)
		if err != nil {
			log.Fatal(err.Error())
		}
		for {
			connection, err := server.Accept()
			if err != nil {
				log.Print(err.Error())
			}
			log.Println("Established Connection")
			log.Println(connection.RemoteAddr())
			get.GetFile(connection, "./recv/")

		}
	case "send":
		connection, err := net.Dial("tcp", "localhost:9000")
		if err != nil {
			log.Fatal(err.Error())
		}
		send.SendFile(connection, "./textfile")
		os.Exit(1)
	default:
		fmt.Printf("%v by %v\nversion %v\nupdate status %v\n", pinfo.Program, pinfo.Author, pinfo.Version, "NOT IMPLEMENTED YA MORON")

	}

}
