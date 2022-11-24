package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/consolving/autodns.go/autodns"
)

func getZone(domain string) (*autodns.Zone, error) {
	zone, err := autodns.NewZone(domain)
	if err != nil {
		return nil, err
	}
	return zone, nil
}

func RenderDataXML(data interface{}) {
	fmt.Println("\t\nXML:")
	b, err := xml.MarshalIndent(data, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func RenderDataJSON(data interface{}) {
	fmt.Println("\t\nJSON:")
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {

	client, err := autodns.NewClientFromEnv()
	if err != nil {
		log.Fatal(err)
	}

	request, err := client.NewRequest()
	if err != nil {
		log.Fatal(err)
	}

	task, err := autodns.NewTaskWithKey(autodns.ZONE_INFO)
	if err != nil {
		log.Fatal(err)
	}

	zone, err := getZone("ey4.de")
	if err != nil {
		log.Fatal(err)
	}

	task = task.WithZone(zone)
	request = request.WithTask(task)
	cb := func(resp *http.Response, err error) {
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var response *autodns.Response
		fmt.Printf("===== RAW\n%s\n===== RAW\n", string(body))
		err = xml.Unmarshal(body, &response)
		if err != nil {
			panic(err)
		}
		RenderDataXML(response)
		RenderDataJSON(response)
	}
	client.PerformRequest(request, cb)

}
