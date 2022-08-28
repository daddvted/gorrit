package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"

	"golang.org/x/crypto/ssh"
)

func Connect2Gerrit() error {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println("shit")
	}
	fmt.Println(currentUser.Username)
	fmt.Println(currentUser.HomeDir)

	key, err := ioutil.ReadFile(fmt.Sprintf("%s/.ssh/id_rsa", currentUser.HomeDir))
	if err != nil {
		log.Fatal("unable to read private key: %v", err)
	}

	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: currentUser.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(singer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "10.0.0.111:29418", config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatal("Failed to create session: ", err)
	}
	defer session.Close()

	var b bytes.Buffer

	session.Stdout = &b

	if err := session.Run("gerrit ls-projects --format json"); err != nil {
		log.Fatal("Failed to run: " + err.Error())
	}

	fmt.Println(b.String())

	return nil

}
