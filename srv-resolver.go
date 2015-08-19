package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type record struct {
	Service string `json:"service"`
	Host    string `json:"host"`
	IP      string `json:"ip"`
	Port    string `json:"port"`
}

func main() {
	server := flag.String("server", "localhost", "Define the IP of the Mesos-DNS Host")
	port := flag.Int("port", 8123, "Define the port for the Mesos-DNS REST API")
	service := flag.String("service", "none", "Define the service you want to get the srv entry")
	protocol := flag.String("protocol", "tcp", "Define the protocl the service is using")
	framework := flag.String("framework", "marathon", "Define the framework which runs the service")
	domain := flag.String("domain", "mesos", "Define the domain of Mesos-DNS")
	flag.Parse()

	var address = fmt.Sprintf("http://%s:%s/v1/services/_%s._%s.%s.%s", *server, strconv.Itoa(*port), *service, *protocol, *framework, *domain)
	response, err := http.Get(address)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	var records []record
	err = json.Unmarshal(contents, &records)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if len(records) < 1 {
		fmt.Printf("No entries for this service")
		os.Exit(1)
	}

	fmt.Printf("%s %s", records[0].IP, records[0].Port)
	os.Exit(0)
}
