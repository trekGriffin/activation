package main

import (
	"errors"
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
	des = " a tool for activation in the us server"
)

var (
	appVersion = "Unknown"
	appDate    = "Unknown"
)

type Config struct {
	Port     string   `yaml:"port"`
	Token    []string `yaml:"token"`
	Redirect string   `yaml:"redirect"`
	Contains string   `yaml:"gotrek.top"`
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
	if !strings.Contains(r.Host, config.Contains) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, config.Redirect+r.RequestURI, http.StatusSeeOther)
}
func checkConfig() error {
	if strings.Compare(config.Contains, "") == 0 {
		return errors.New("contains is empty")
	}
	if strings.Compare(config.Port, "") == 0 {
		return errors.New("port is empty")
	}
	if len(config.Token) == 0 {
		return errors.New("token is empty")
	}

	return nil
}

func main() {
	showVersion := false
	configFile := ""
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.StringVar(&configFile, "c", "/etc/trek/config.yaml", "specify config")

	flag.Parse()
	if showVersion {
		fmt.Printf("app version is %s appdate is %s", appVersion, appDate)
		return
	}
	fmt.Println(des)
	_, err := os.Stat(configFile)
	if err != nil {
		const default2 = "./config.yaml"
		_, err = os.Stat(default2)
		if err != nil {
			fmt.Println(" config file is not exist:", configFile, " and", default2)
			os.Exit(1)
		}
		fmt.Println(configFile, " not exist, using the default ", default2)
		configFile = default2
	}

	buf, err := os.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal("open file config.yaml failed", err)
	}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal("unmarshal failed", err)
	}
	checkConfig()
	//	http.HandleFunc("/", handler)
	log.Printf("config: port : %s, token: %s", config.Port, config.Token)

	mux := http.NewServeMux()
	mux.HandleFunc("/activate", activateHandler)
	mux.HandleFunc("/ip", ipHandler)
	mux.HandleFunc("/", rootHandler)
	log.Println("servier is listening", config.Port)
	err = http.ListenAndServe(config.Port, cors.Default().Handler(mux))
	if err != nil {
		log.Fatal(err)
	}
}
