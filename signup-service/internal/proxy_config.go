package internal

import (
	"fmt"
	"github.com/juju/fslock"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

const HostName = "localhost"


type Config struct {
	DOMAIN string
	HOST   string
	PORT   string
}



type ProxyConfig struct {
}

func (p ProxyConfig) ConfigureProxyPass(port, domain string) error {

	fileName, err := createNewProxyPassConfig(domain, port, HostName)
	if err != nil {
		return err
	}

	proxy := fmt.Sprintf("include %s;", fileName)

	lock := fslock.New("./internal/assets/nginx1.conf")
	defer lock.Unlock()
	lockErr := lock.LockWithTimeout(5*time.Second)
	if lockErr != nil {
		fmt.Println("failed to acquire lock > " + lockErr.Error())
		return err
	}

	// write at the beginning.
	contents,_ := ioutil.ReadFile("./internal/assets/nginx1.conf")
	contentsToWrite := fmt.Sprintf("%s\n %s", proxy, contents)

	ioutil.WriteFile("./internal/assets/nginx1.conf",[]byte(contentsToWrite) , 0644)

	// sudo nginx -s reload
	_, err = exec.Command("sudo", "nginx", "-s", "reload").Output()
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}

	//symbolic link for nginx
	// ln /opt/homebrew/etc/nginx/nginx.conf internal/assets/nginx.conf
	return nil
}


func createNewProxyPassConfig(domain, port, host string)(string, error) {
	config := Config{
		DOMAIN: domain,
		HOST:   host,
		PORT:   port,
	}

	t, err := template.ParseFiles("./internal/assets/location.tmpl")
	if err != nil {
		log.Println("create file: ", err)
		return "", err
	}

	fileName := fmt.Sprintf("./internal/assets/location-%s.conf", domain)
	f, err := os.Create(fileName)
	if err != nil {
		log.Println("create file: ", err)
		return "", err
	}

	defer f.Close()

	err = t.Execute(f, config)
	if err != nil {
		log.Println("create template file: ", err)
		return "", err
	}
	return fileName, nil
}


