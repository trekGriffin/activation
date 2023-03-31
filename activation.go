package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	"gopkg.in/yaml.v3"
)

const (
	des = " a tool for activation"
)

var (
	appVersion = "Unknown"
	appDate    = "Unknown"
)

type Config struct {
	Port  string   `yaml:"port"`
	Token []string `yaml:"token"`
}

var config Config

func checkToken(query string) bool {
	for _, ele := range config.Token {
		if strings.Compare(ele, query) == 0 {
			return true
		}
	}
	return false
}
func ipHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got a new request", "/")
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println("split source ip port err", err)
		return
	}
	fmt.Fprint(w, ip)
}
func activateHandler(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query().Get("token")
	log.Println("got request token from IP token:", query, " remote ip:", req.RemoteAddr)
	if checkToken(query) {
		fmt.Println("check token ok", query)
		res.WriteHeader(http.StatusOK)
		return
	}
	fmt.Println("check token failed", query)
	res.WriteHeader(http.StatusUnauthorized)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func main() {
	showVersion := false
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("app version is %s appdate is %s", appVersion, appDate)
		return
	}
	fmt.Println(des)

	buf, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("open file config.yaml failed", err)
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal("unmarshal failed", err)
	}

	//	http.HandleFunc("/", handler)
	log.Printf("config: port : %s, token: %s", config.Port, config.Token)

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/activate", activateHandler)
	mux.HandleFunc("/ip", ipHandler)
	log.Println("servier is listening", config.Port)
	err = http.ListenAndServe(config.Port, cors.Default().Handler(mux))
	if err != nil {
		log.Fatal(err)
	}
}
