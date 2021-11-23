package manifests

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
)

const version = "latest"
const port = "8080"

type Config struct {
	Version string
	Name    string
	Port    string
	Container string
}



type DockerComposer struct {

}


func NewDockerComposer() DockerComposer {
	return DockerComposer{}
}


func (d DockerComposer)InitContainer(domain string) (int,error) {
	manifest, err := d.ParseFile(version, domain, port)
	if err != nil {
		log.Println("error manifest", err)
		return 0,err
	}

	port, err := d.ApplyCompose(manifest, fmt.Sprintf("%s-service", domain))
	if err !=nil {
		log.Println("error applying compose", err)
		return 0, err
	}

	return port, nil
}

// creates a new docker-compose.yaml for the newly created service
func (DockerComposer)ParseFile(version, name, port string) (string, error){
	 config := Config{
		 Version: version,
		 Name:    name,
		 Port:    port,
		 Container: fmt.Sprintf("%s-service", name),
	 }
	t, err := template.ParseFiles("internal/manifests/docker-compose.tmpl")
	if err != nil {
		log.Println("create file: ", err)
		return "", err
	}

	fileName := fmt.Sprintf("internal/manifests/docker-compose-%s.yaml", name)
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


// applies docker compose
// returns the port for nginx mapping
func (DockerComposer)ApplyCompose(fileName, containerName string)(int, error) {
	out, err := exec.Command("docker", "compose", "-f", fileName, "up", "-d").Output()
	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
		return 0, err
	}

	out, err = exec.Command("docker", "port", containerName).Output()
	if err != nil {
		fmt.Printf("%s", err)
		return 0, err
	}
	output := string(out[:])
	port := 0
	if strings.Contains(output, ":") {
		output = strings.ReplaceAll(strings.Split(output,":")[1], "\n", "")
		port, err = strconv.Atoi(output)
		if err !=nil{
			return 0, err
		}
	}
	return port, nil
}
