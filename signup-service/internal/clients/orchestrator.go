package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func ScheduleContainer(domain string)(string, error) {
	resp, err := http.Post(fmt.Sprintf("http://localhost:9001/container/%s", domain),"text/plain", nil)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)

	return string(body), nil
}