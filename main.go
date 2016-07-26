package main

import (
	"fmt"
	"git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"os"
	"strconv"
	"time"
)

const (
	clientID = "" // your relayr client-id
	username = "" // your relayr username
	password = "" // your relayr password
)

func main() {
	fmt.Println("==> relayr virtual-device golang example")

	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.relayr.io:1883")
	opts.SetClientID(clientID).SetUsername(username).SetPassword(password)
	c := mqtt.NewClient(opts)

	fmt.Print("==> try to connect... ")
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection Error", token.Error())
		os.Exit(1)
	}
	fmt.Println("successful connected")

	//
	// subscribe
	//
	msgRcvd := func(client *mqtt.Client, message mqtt.Message) {
		fmt.Printf("<-- Received message on topic: %s \t Message: %s\n", message.Topic(), message.Payload())
	}
	//
	if token := c.Subscribe("/v1/"+username+"/cmd", 0, msgRcvd); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	//
	// publish
	//
	counter := 0
	for {
		fmt.Printf("--> publish data %v\n", counter)
		if token := c.Publish("/v1/"+username+"/data", 1, false, publishJSON(counter)); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
		counter++
		time.Sleep(1 * time.Second)
	}
}

func publishJSON(val int) string {
	return `{"meaning":"someMeaning", "value":"` + strconv.Itoa(val) + `"}`
}
