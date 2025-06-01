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

## 🔍 General Questions

### What is the difference between a slice and an array in Go?
An **array** has a fixed length and cannot be resized. A **slice** is a dynamically-sized, flexible view into an array. Slices are more commonly used because they are more versatile.

### What is a map in Go? What type of map is it?
A **map** is a built-in data type in Go used to associate keys with values. It is implemented as a **hash table**.

### What is the time complexity of an insertion/read in a map?
On average, both insertion and read operations in a map are **O(1)**.

### Define generics. What’s its use?
**Generics** allow writing functions and data structures with type parameters, enabling code reuse for different types while maintaining type safety.

\`\`\`go
func Sum[T int | float64](a, b T) T {
    return a + b
}
\`\`\`

### Define struct.
A **struct** is a composite data type that groups together variables under a single name. Each variable inside a struct is called a **field**.

### Define interface.
An **interface** defines a set of method signatures. A type implements an interface by defining all the methods declared by that interface.

### Can interfaces be embedded in other interfaces?
Yes. Interfaces can **embed other interfaces**, effectively inheriting their method sets.

### Can interfaces be embedded in structs? How does this affect struct implementations?
Yes. When an interface is embedded in a struct, the struct must implement the methods of the embedded interface (or embed a type that does). This enables **composition** and **polymorphism**.

### Explain the interface implementation in Go.
Go uses **implicit implementation**: a type satisfies an interface simply by implementing its methods — no explicit declaration is required.

### Name a good use case for interfaces.
Interfaces are ideal for **dependency injection**, allowing for decoupled and testable code. Example: using an `io.Writer` instead of a concrete `os.File`.

### Can a struct implement more than one interface in Go?
Yes. A struct can implement **multiple interfaces** by satisfying all required methods.

### What is inheritance? How to implement that in Go?
**Inheritance** is the reuse of behavior from another type. Go doesn’t support classical inheritance, but supports **composition** through embedding other structs or interfaces.

### Is it possible to override a struct method in Go?
No. Go does **not support method overriding** in the classical OOP sense.

---

## ⚙️ Concurrency and Execution

### Explain multithreading approach in Go.
Go uses **goroutines**, which are lightweight threads managed by the Go runtime, and **channels** for communication between them, allowing for efficient concurrency.

### What does the `defer` statement do?
`defer` schedules a function call to run **after the current function returns**. If multiple defers exist, they are executed in **LIFO (Last In, First Out)** order.

### What are struct field offsets?
Struct field offsets are the **memory positions** of each field in a struct, used to calculate field locations during access. This impacts memory alignment and performance.

### What is the size of a pointer in memory?
On a 64-bit system, a pointer typically occupies **8 bytes**.

### Why do we have zero-based index in all languages?
Zero-based indexing simplifies **pointer arithmetic** and **memory access**. It's also more efficient in terms of computation.

---

## 🧱 Architecture and Systems

### What is the difference between a monolithic application and a microservices application?

| Aspect        | Monolithic                        | Microservices                                         |
|---------------|-----------------------------------|-------------------------------------------------------|
| Structure     | Single application                | Multiple independent services                         |
| Deployment    | Easier initially                  | Complex setup, individual deployment                  |
| Scalability   | Difficult to scale parts          | Easy to scale services independently                  |
| Maintenance   | Becomes hard as app grows         | Enables team autonomy and modularization              |

### What are the programmatic concerns on each model?

- **Monoliths**: Tight coupling, deployment risks, difficult to scale parts independently.
- **Microservices**: Requires network communication, service discovery, data consistency, distributed tracing.

### What is the difference between Stateless and Stateful applications?

- **Stateless**: Doesn’t store session data on the server.  
  _Example_: REST APIs.

- **Stateful**: Stores user session/state.  
  _Example_: Multiplayer games, session-based web servers.

### What’s NGINX?
NGINX is a high-performance **web server**, **reverse proxy**, and **load balancer** used to serve static files and route HTTP traffic.

### What’s an entry point in Docker?
The `ENTRYPOINT` in a Dockerfile defines the **default executable** to run when a container starts.

---

## 🔐 Communication & Security

### What are synchronous requests? Example?
A synchronous request **waits** for the response before proceeding.  
_Example_: HTTP call waiting for a response.

### What are asynchronous requests? Example?
Asynchronous requests **allow other operations to continue** before the response returns.  
_Example_: Sending a message to a Kafka queue.

### What are wildcards in asynchronous communication?
Wildcards allow **pattern matching** in topics/routes, commonly used in Pub/Sub systems like MQTT or AMQP.

### What is the difference between hashing and encryption?

| Feature       | Hashing                         | Encryption                        |
|---------------|----------------------------------|-----------------------------------|
| Reversibility | One-way                         | Two-way (reversible)              |
| Use Case      | Data integrity (e.g., passwords) | Data confidentiality (e.g., TLS) |

### What is the difference between Queues and Topics?

- **Queues**: One message goes to **one consumer** (point-to-point).
- **Topics**: Messages are **broadcast** to all subscribers (publish-subscribe).

---

## 🧑‍💻 Go Practical Use

### How would you initialize anything at application start in Go?
Use the `init()` function — it is **automatically executed** before `main()`.

### How would you treat a timeout routine in Go?
Use `context.WithTimeout()` or `time.After()` to **handle timeouts** in goroutines or HTTP clients.

### What is the difference between API and SDK?

- **API**: Interface to interact with a system or component.
- **SDK**: Toolkit with APIs, libraries, and tools for building applications.

### What is MVC pattern? Example?

- **Model**: Data (e.g., `User` struct)  
- **View**: Presentation (e.g., HTML templates)  
- **Controller**: Logic (e.g., HTTP handlers in Go)

### What is the difference between relational and non-relational databases?

| Type         | Relational                          | Non-relational                      |
|--------------|-------------------------------------|-------------------------------------|
| Structure    | Tables, fixed schema                | Key-value, document, or graph       |
| Examples     | PostgreSQL, MySQL                   | MongoDB, Redis                      |
| Best For     | Structured data, transactions       | Flexible, high-scale applications   |

### What is the advantage of stored procedures in databases?
Stored procedures **improve performance and security** by precompiling logic and restricting direct access to tables.


