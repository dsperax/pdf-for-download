
## DQL com SQS

Configure o Redrive Policy: Ajuste o ‘maxReceiveCount’ para um número adequado de tentativas antes de mover mensagens para a DLQ, permitindo retries suficientes​;
[Documentação AWS](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)

Defina o Retention Period: Garanta que o período de retenção da DLQ seja maior que o da fila original para gerenciar adequadamente a expiração das mensagens

Monitore e Diagnostique: Use alarmes, logs e análises de conteúdo para identificar e resolver problemas com mensagens na DLQ

Utilize Redrive para Reciclagem de Mensagens: Reenvie mensagens da DLQ de volta para a fila de origem, após a correção dos problemas, para processá-las novamente​

[Diagrama]()

How to do:

Para configurar Dead Letter Queues (DLQs) no Amazon SQS na prática, siga estes passos:

Crie a DLQ: Vá ao console do SQS e crie uma nova fila, que será a sua DLQ.
Associe a DLQ à Fila Principal: No console do SQS, selecione a fila principal e configure a DLQ associando-a e definindo o maxReceiveCount (número de tentativas antes de enviar para a DLQ).
Configure os Consumidores: Ajuste o código do consumidor para processar mensagens da fila principal e tratar exceções adequadamente.
Monitore e Ajuste: Use o CloudWatch para monitorar a fila e ajustar as configurações conforme necessário.

Não há código específico envolvido neste processo, pois é realizado principalmente através da configuração no console do AWS SQS.

Exemplo:

```
package main

import (
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
    // Inicialize a sessão do AWS SDK
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))

    // Crie um novo cliente SQS usando a sessão acima
    svc := sqs.New(sess)

    // URL da fila principal
    queueURL := "url_da_sua_fila_principal_aqui"

    // Recebendo mensagens
    result, _ := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
        QueueUrl:            &queueURL,
        MaxNumberOfMessages: aws.Int64(10),
        WaitTimeSeconds:     aws.Int64(20),
    })

    // Itere sobre as mensagens recebidas
    for _, message := range result.Messages {
        fmt.Printf("Mensagem recebida: %s\n", *message.Body)

        // Processamento da mensagem aqui
        // ...

        // Deletar a mensagem após o processamento
        _, _ = svc.DeleteMessage(&sqs.DeleteMessageInput{
            QueueUrl:      &queueURL,
            ReceiptHandle: message.ReceiptHandle,
        })
    }
}
```
