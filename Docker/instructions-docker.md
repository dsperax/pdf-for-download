# **Comandos Docker:**

### **Visualização**

| Comando | Resultado
|----------|:-------------:
| docker ps -a | lista containers
| docker ps -a -q | lista id dos containers
| docker images -a | lista imagens
| docker images -a -q | lista id das imagens

### **Ação**

| Comando | Resultado
|----------|:-------------:
| docker run ...  | criação
| docker stop ... | pausa imagem ou container
| docker start ...| starta imagem ou container

### **Exclusão**

| Comando | Resultado
|----------|:-------------:
| docker rm (nome ou 4 primeiros digitos do id)   | remove container
| docker rmi (nome ou 4 primeiros digitos do id)  | remove imagem
| docker rm ($docker ps -a -q) | exclui todos os containers (funcional apenas executado pelo window powershell)
| docker rmi ($docker images -a -q) | remove todas as imagens (funcional apenas executado pelo window powershell)
    

### **Comandos para docker compose**

* *comandos executados DENTRO DA PASTA ONDE SE LOCALIZA O ARQUIVO*

| Comando | Resultado
|----------|:-------------:
| docker-compose up    | executa comandos dentro do docker-compose.yaml
| docker-compose down  | exclui imagens criadas pelo docker-compose.- yaml
| docker-compose start | starta imagens criadas pelo docker-compose.yaml
| docker-compose stop  | pausa imagens criadas pelo docker-compose.yaml

<hr>

## **Exemplo**

Execução do docker compose para criação de um ambiente com 2 bancos MSSQL e 1 Messageria (kafka).

Execute o windows powershell como administrados e acesse a pasta onde está o arquivo "docker-compose.yaml";

    docker-compose up -d

após esse comando, será feito o download das imagens e os container irão subir;

Nesse ponto, os bancos de dados e o kafka estarão disponíveis para conexão;

Terminei por hoje, e agora?

    docker-compose stop 

Comando para pausar a excução do ambiente;

Com isso, pode voltar amanhã e dar "play" novamente sem perder os dados;

    docker-compose start 

Starta novamente o ambiente;

Terminei os testes de vez, quero zerar o ambiente e excluir tudo da minha maquina;

    docker-compose down

Finaliza tudo e exclui os containers;

    docker rmi ($docker images -a -q)

Exclui todas as imagens;

lembre-se, em caso de teimosia da maquina -f resolve;

<hr>

## **Utilizando os ambientes**

### **MongoDB**

*Via windows powershell*

    docker exec -it mongodb bash
    mongosh
    use admin
    db.auth('root', passwordPrompt())
    root123

-> Retorno 1 se deu certo.

    show dbs
    
-> Nesse ponto aparecem os bancos de dados base. Basta seguir.


*String de conexão em caso de ide*

    string: mongodb://root:root123@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.9.0&authMechanism=DEFAULT


### **MSSQL**

    Host: localhost
    Porta: 1433 ou 1434 (se atentar ao docker-compose.yaml)
    Database/Schema: master
    Authentication: SQL Server Authentication
    Username: sa
    Senha: P@ssword

### **Kafka (via offset explorer)**

    Cluster name: zookeeper-local-docker
    Zookeeper Host: localhost
    Zookeeper Port: 2181
    Advanced → Booststrap servers: localhost:9092

OBS: Mudar o tipo de dado depois de subir o cluster antes de mandar a mensagem.

Fim.