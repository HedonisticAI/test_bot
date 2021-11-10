package main

import (
	"net"

	"github.com/ipinfo/go/ipinfo"
)

func ipinit() *ipinfo.Client {
	auth := ipinfo.AuthTransport{Token: "3f2c9320ee5df6"}
	httpClient := auth.Client()
	client := ipinfo.NewClient(httpClient)
	return client
}
func parseip(ip string, client *ipinfo.Client) string {
	var res string
	IPAddress := net.ParseIP(ip)
	info, err := client.GetInfo(IPAddress)
	if err != nil {
		return "ip error"
	}
	res = info.Organization + "/n" + info.City + "/" + info.Region + "/n"
	res = res + info.Hostname + "/n" + info.Phone + "/n" + info.Country
	return res
}
