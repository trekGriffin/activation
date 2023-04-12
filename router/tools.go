package router

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

var (
	CountryContent *geoip2.Reader
)

func CheckIp(ip string) (string, error) {

	defer CountryContent.Close()
	record, err := CountryContent.City(net.ParseIP(ip))
	if err != nil {
		return "", err
	}

	return record.Country.IsoCode, nil
}
func IpDetai(w http.ResponseWriter, r *http.Request) {
	ip := strings.Replace(r.URL.Path, "/ip/", "", 1)
	log.Println(" got ip detail request", ip)

	if ip == "" {
		fmt.Fprintf(w, "empty ip")
		return
	}
	isocode, err := CheckIp(ip)
	if err != nil {
		fmt.Fprintf(w, "eroror:%s", err)
		return
	}
	fmt.Fprintf(w, "%s", isocode)
}
