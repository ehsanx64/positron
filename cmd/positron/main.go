package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	userHttpHandler "github.com/ehsanx64/positron/internal/infra/delivery/http"
	ui "github.com/ehsanx64/positron/ui"
	"github.com/spf13/viper"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	mux := http.NewServeMux()

	viper.SetConfigType("json")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("failed to load the config file")
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	// Config file found and successfully parsed

	var broker = "elnaz"
	var port = 1883
	var rootTopic = "/positron"
	var clientID = rootTopic + "/core"
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(clientID)
	// opts.SetUsername("emqx")
	// opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Failed to connect to mqtt broker: %+x", token.Error())
	}

	rootHandler := func(c mqtt.Client, msg mqtt.Message) {
		// pp.Print(msg)
	}

	nodemcuHandler := func(c mqtt.Client, msg mqtt.Message) {
	}

	userHttpHandler.NewUserHTTPHandler(mux)
	mux.Handle("/assets/", http.FileServer(http.FS(ui.Assets)))
	mux.Handle("/", http.FileServer(http.FS(ui.Main)))
	log.Println("Starting positron on " + viper.Get("app.port").(string))
	if err := http.ListenAndServe(viper.Get("app.port").(string), mux); err != nil {
		log.Fatal(err)
	}

	// subscribe to subTopic("/a1Zd7n5***/deng/user/get") and request messages to be delivered
	token := client.Subscribe(clientID+"/#", 1, rootHandler)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	token = client.Subscribe(rootTopic+"/nodemcu/#", 1, nodemcuHandler)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(clientID+"/currentTime", 0, false, t.String())
	}
}
