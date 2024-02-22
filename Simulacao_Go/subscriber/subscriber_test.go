
package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
	"fmt"
)

func TestMain(t *testing.T) {
    // Canal para sincronização
    var receivedMessages = make(chan struct{})

    // Função para manipular as mensagens recebidas
    var onMessageReceived = func(client MQTT.Client, message MQTT.Message) {
        fmt.Printf("Mensagem recebida no tópico: %s\n", message.Topic())
        fmt.Printf("Payload: %s\n", message.Payload())
        // Envie uma marcação para indicar que uma mensagem foi recebida
        receivedMessages <- struct{}{}
    }

    // Configurações do cliente MQTT
    opts := MQTT.NewClientOptions().AddBroker(broker)
    opts.SetClientID("go-mqtt")
    opts.SetDefaultPublishHandler(onMessageReceived)

    // Criação do cliente MQTT
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    // Inscrição no tópico
    if token := client.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
        panic(token.Error())
        t.Errorf("Erro ao se inscrever no tópico: %s", topic)
    }

    t.Logf("Conectado ao broker MQTT em %s\n", broker)

    // Aguarde a recebimento de três mensagens
    for i := 0; i < 3; i++ {
        select {
        case <-receivedMessages:
            t.Logf("Mensagem %d recebida", i+1)
        case <-time.After(2 * time.Second): // Defina um tempo limite para 2 segundos
            t.Errorf("Tempo limite excedido para receber mensagens")
        }
    }
}