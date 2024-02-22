package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	// Configuração do cliente MQTT
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("Conexão Zona Leste - SP")
	client := MQTT.NewClient(opts)

	// Conecte ao broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	// Captura de sinal para lidar com o encerramento
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	// Loop para publicar mensagens continuamente
	go func() {
		for {
			zl1, zl2 := simularLeituraSensor()
			message := fmt.Sprintf("Leitura do sensor: PM2.5 %.2f (µg/m³), PM10 %.2f (µg/m³), Penha - Zona Leste - SP", zl1, zl2)
			token := client.Publish("test/topic", 0, false, message)
			token.Wait()
			fmt.Printf("Publicado: %s\n", message)
			time.Sleep(1 * time.Second)
		}
	}()

	// Aguarde um sinal de interrupção
	<-signalChannel
	fmt.Println("Publicação encerrada")
}

// simularLeituraSensor simula a leitura do sensor
func simularLeituraSensor() (float64, float64) {
	// Simulação de leitura de sensor
	zl1 := rand.Float64() * 100
	zl2 := rand.Float64() * 100
	return zl1, zl2
}
