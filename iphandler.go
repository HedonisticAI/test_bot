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
func parseip(ip string, id int, client *ipinfo.Client) string {
	var res string
	IPAddress := net.ParseIP(ip)
	info, err := client.GetInfo(IPAddress)
	if err != nil {
		return "ip error"
	}
	res = "org " + info.Organization + "\n" + "city " + info.City + "\n" + "region " + info.Region + "\n"
	res = res + "host " + info.Hostname + "\n" + "phone  " + info.Phone + "\n" + "country " + info.Country
	addtodb(res, id)
	return res
}
