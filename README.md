# P2-MQTT_Artefatos-M9

# Teste do Cliente MQTT

Este pacote contém testes para verificar o comportamento do cliente MQTT.

## Subscribe e Publisher (Go Test)

O teste `TestMain` verifica se o cliente MQTT pode se conectar ao broker, subscribir a um tópico e receber mensagens no handler padrão. Ele verifica se a mensagem está sendo recebido em um perído de tempo específico (entre 1 segundo). O Existe ainda um segundo test em /Simulacao_Go que testa se os dados simulados do sensor estão dentro do padrão esperado.

### Funcionalidade

1. O cliente MQTT é configurado com as opções padrão e o ID do cliente definido como "go-mqtt".
2. Conecta-se ao broker MQTT.
3. Inscreve-se no tópico especificado.
4. Aguarda a recebimento de três mensagens.
5. Verifica se as mensagens são recebidas dentro do limite de tempo especificado.
6. Exibe mensagens de log para indicar o estado do teste.

### Pré-requisitos

- Um broker MQTT deve estar em execução localmente em "localhost:1883".

### Observações

Este teste pressupõe que um broker MQTT esteja em execução localmente em "localhost:1883".

## Dependências Externas

O teste depende do pacote `github.com/eclipse/paho.mqtt.golang` para interagir com o broker MQTT.

## Como Executar

Para executar o teste, basta executar o comando `go test` no diretório do pacote.

Exemplo:
```bash
go test -v
```

### Vídeo de Funcionamento

link: https://drive.google.com/file/d/17PU42tLDo2XP2a21-bOLrEUHub5mZxnL/view?usp=sharing