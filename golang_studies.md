# 🔧 Perguntas Técnicas sobre Microsserviços e Golang

## 📡 1. Como implementar Service Discovery em uma arquitetura de microsserviços?

A descoberta de serviços pode ser feita de duas formas:

- **Estática**: uso de endereços IPs ou DNS pré-configurados.
- **Dinâmica**: com ferramentas como:
  - [Consul](https://www.consul.io/)
  - [Eureka (Netflix OSS)](https://github.com/Netflix/eureka)
  - [etcd](https://etcd.io/)
  - Kubernetes DNS + services

Essas ferramentas permitem que microsserviços se registrem automaticamente e descubram outros serviços dinamicamente via API ou DNS.

---

## 🔁 2. Como garantir tolerância a falhas em microsserviços em Go? Como usar circuit breakers como Hystrix ou Resilience4j em Golang?

### Estratégias de tolerância a falhas:
- Retry com **backoff exponencial**
- **Timeouts** por requisição
- **Circuit Breakers**
- Fallbacks (respostas alternativas)
- Load balancing
- Isolamento de serviços

### Circuit Breakers em Go:
Embora **Hystrix** e **Resilience4j** sejam para Java, em Go usamos:

- [`github.com/sony/gobreaker`](https://github.com/sony/gobreaker)
- [`github.com/slok/goresilience`](https://github.com/slok/goresilience)
- **Istio** e service meshes também oferecem circuit breakers externos à aplicação.

---

## 🧾 3. Como gerenciar transações envolvendo múltiplos serviços? Como o padrão Saga ajuda nisso?

### Problema:
Transações ACID não funcionam bem em microsserviços distribuídos.

### Solução:
**Saga Pattern** – quebra a transação em eventos locais e define ações de compensação em caso de falhas.

### Tipos de Saga:
- **Orquestrada**: um orquestrador comanda o fluxo entre serviços.
- **Coreografada**: cada serviço reage a eventos publicados por outros.

### Exemplo:
1. Serviço A debita saldo
2. Serviço B reserva estoque
3. B falha → A executa ação compensatória para devolver saldo

---

## 💾 4. Como Go gerencia memória? Como otimizar uso de memória em aplicações de alta performance?

### Gerenciamento de memória:
- Go possui **alocação automática** e **coletor de lixo (GC)**.
- O **GC do Go** é:
  - Concorrente
  - Incremental
  - Otimizado para pausas curtas

### Otimizações manuais:
- Evite alocações desnecessárias
- Prefira tipos por valor (`struct`) a ponteiros
- Use `sync.Pool` para reuso de objetos
- Minimize escapes para o heap
- Use `pprof` para profiling de memória

---

## ⚙️ 5. Como Go lida com concorrência? Como evitar race conditions? Channels vs sync.Mutex?

### Concorrência no Go:
- Go usa **goroutines** e **channels** como primitivas leves e nativas
- Alternativa: **sync.Mutex** para bloqueio de regiões críticas

### Evitando condições de corrida:
- Use `go run -race` para detectar
- Use `sync.Mutex` ou `sync.RWMutex`
- Use **channels** para comunicação segura entre goroutines

> **Channels** são ideais para pipelines e comunicação.  
> **Mutex** é mais eficiente para sincronizar leitura/escrita simples em memória compartilhada.

---

## 🧨 6. Como Go trata erros? Como implementar tratamento estruturado?

### Erros em Go:
- Go não tem exceções – usa valores do tipo `error` retornados

### Tratamento estruturado:
- Use `fmt.Errorf("%w", err)` para wrapping
- Use `errors.Is` e `errors.As` para introspecção
- Defina **erros personalizados** com tipos customizados
- Bibliotecas úteis: `pkg/errors`, `multierr`

### Vantagens de erros personalizados:
- Facilitam o debug
- Permitem categorização e logging com contexto

---

## 🧱 7. Como interfaces e composição em Go diferem da herança tradicional? Como implementar injeção de dependência?

### Diferenças principais:
- Go **não tem herança**, usa **composição**
- Interfaces são **implícitas** (não há necessidade de declarar que um tipo implementa uma interface)

### Injeção de dependência com interfaces:

```go
type Notifier interface {
	Send(msg string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(msg string) error {
	fmt.Println("Enviando email:", msg)
	return nil
}

type Service struct {
	notifier Notifier
}

func NewService(n Notifier) *Service {
	return &Service{notifier: n}
}
```
