## 1 - Descrição: Trabalhar com números, manipulação básica de strings e manipulação básica de arrays. Requer uma combinação de 2-3 desses conceitos. Solução em 5-10 linhas de código.

Escreva uma função em Go que receba uma string contendo números separados por vírgulas, converta esses números em um array de inteiros, e retorne a soma desses números.

```
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumNumbers(s string) int {
	numStrs := strings.Split(s, ",")
	sum := 0
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(strings.TrimSpace(numStr))
		sum += num
	}
	return sum
}

func main() {
	input := "1, 2, 3, 4, 5"
	fmt.Println(sumNumbers(input)) // Output: 15
}

func testSumNumbers() {
	fmt.Println(sumNumbers("1,2,3") == 6)
	fmt.Println(sumNumbers("10, 20, 30") == 60)
	fmt.Println(sumNumbers("5,5,5,5") == 20)
}

```

Explicação:

- strings.Split(s, ","): Divide a string em substrings com base na vírgula, resultando em um array de strings.
= strconv.Atoi(): Converte uma string em um inteiro.
- strings.TrimSpace(): Remove espaços em branco das extremidades da string.
- O loop percorre cada substring, converte para inteiro e acumula a soma.

## 2 - Descrição: Implementação um pouco mais complexa, combinando 3-5 conceitos.

Escreva uma função em Go que receba um array de strings representando palavras, e retorne um novo array com as palavras em ordem inversa e em maiúsculas.

```
package main

import (
	"fmt"
	"strings"
)

func reverseAndUppercase(words []string) []string {
	result := make([]string, len(words))
	for i, word := range words {
		result[len(words)-1-i] = strings.ToUpper(word)
	}
	return result
}

func main() {
	words := []string{"golang", "is", "fun"}
	fmt.Println(reverseAndUppercase(words)) // Output: [FUN IS GOLANG]
}

func testReverseAndUppercase() {
	fmt.Println(reverseAndUppercase([]string{"hello", "world"})) // [WORLD HELLO]
	fmt.Println(reverseAndUppercase([]string{"go", "is", "cool"})) // [COOL IS GO]
}
```

Explicação:

- make([]string, len(words)): Cria um novo slice com o mesmo tamanho do original.
- O loop preenche o novo slice com as palavras em ordem inversa e em maiúsculas.
- strings.ToUpper(word): Converte a string para maiúsculas.

## 3 - Envolve dividir a tarefa em subtarefas/funções, manipular arrays bidimensionais e usar hashmaps. Código com 25-40 linhas.

Escreva uma função em Go que receba uma matriz bidimensional de inteiros representando notas de alunos, e retorne um mapa (hashmap) onde a chave é o índice do aluno e o valor é a média das notas desse aluno.

```
package main

import (
	"fmt"
)

func calculateAverage(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func studentAverages(gradesMatrix [][]int) map[int]float64 {
	averages := make(map[int]float64)
	for i, grades := range gradesMatrix {
		averages[i] = calculateAverage(grades)
	}
	return averages
}

func main() {
	gradesMatrix := [][]int{
		{80, 90, 100},
		{70, 85, 90},
		{60, 75, 80},
	}
	fmt.Println(studentAverages(gradesMatrix))
	// Output: map[0:90 1:81.66666666666667 2:71.66666666666667]
}

func testStudentAverages() {
	gradesMatrix := [][]int{
		{100, 90},
		{80, 70},
		{60, 50},
	}
	averages := studentAverages(gradesMatrix)
	fmt.Println(averages[0] == 95)
	fmt.Println(averages[1] == 75)
	fmt.Println(averages[2] == 55)
}
```

Explicação:

- calculateAverage(): Função que calcula a média de um slice de inteiros.
- studentAverages(): Função que cria um mapa onde cada aluno (índice) está associado à sua média.
- Percorre cada linha da matriz (cada aluno) e calcula a média usando a função auxiliar.

# 4- Descrição: Trabalhar com árvores, entender hashmaps, fundamentos de matemática discreta e busca por força bruta. Código com 20-35 linhas.

Escreva uma função em Go que determine se uma árvore binária é um heap máximo. Em um heap máximo, cada nó é maior ou igual aos seus filhos.

```
package main

import (
	"fmt"
)

// Definição do nó da árvore
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Função para verificar se a árvore é um heap máximo
func isMaxHeap(root *Node) bool {
	if root == nil {
		return true
	}

	// Verifica o filho esquerdo
	if root.Left != nil {
		if root.Left.Value > root.Value {
			return false
		}
	}

	// Verifica o filho direito
	if root.Right != nil {
		if root.Right.Value > root.Value {
			return false
		}
	}

	// Recursivamente verifica os subárvores
	return isMaxHeap(root.Left) && isMaxHeap(root.Right)
}

func main() {
	// Construção da árvore
	root := &Node{Value: 10}
	root.Left = &Node{Value: 9}
	root.Right = &Node{Value: 8}
	root.Left.Left = &Node{Value: 7}
	root.Left.Right = &Node{Value: 6}
	root.Right.Left = &Node{Value: 5}
	root.Right.Right = &Node{Value: 4}

	fmt.Println(isMaxHeap(root)) // Output: true
}

func testIsMaxHeap() {
	// Exemplo de heap máximo
	root1 := &Node{Value: 10, Left: &Node{Value: 9}, Right: &Node{Value: 8}}
	fmt.Println(isMaxHeap(root1)) // true

	// Exemplo de não heap máximo
	root2 := &Node{Value: 10, Left: &Node{Value: 11}, Right: &Node{Value: 8}}
	fmt.Println(isMaxHeap(root2)) // false
}
```

Explicação:

- Node: Estrutura que representa um nó da árvore binária.
- isMaxHeap(): Função recursiva que verifica se cada nó satisfaz a propriedade de heap máximo.
- A função verifica se o valor do nó atual é maior ou igual aos valores de seus filhos.
- Se todos os nós satisfazem a condição, a função retorna true.

# Extra: Tarefa com Busca Binária (lista ordenada)

Título: Buscar um Número em um Array Ordenado

Descrição: Escreva uma função em Go que receba um array de inteiros ordenado em ordem crescente e um número alvo. A função deve utilizar o algoritmo de busca binária para encontrar o índice do número alvo no array. Se o número não estiver presente no array, a função deve retornar -1.

Exemplo:

- Input: array = [1, 3, 5, 7, 9, 11, 13], target = 7

- Output: 3 (porque 7 está no índice 3)

- Input: array = [2, 4, 6, 8, 10], target = 5

- Output: -1 (porque 5 não está no array)

```
package main

import (
	"fmt"
)

// binarySearch realiza a busca binária em um array ordenado.
// Retorna o índice do target se encontrado, ou -1 caso contrário.
func binarySearch(array []int, target int) int {
	inicio := 0
	fim := len(array) - 1

	for inicio <= fim {
		meio := inicio + (fim-inicio)/2
		if array[meio] == target {
			return meio
		} else if array[meio] < target {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

func main() {
	// Exemplos de uso da função binarySearch
	array1 := []int{1, 3, 5, 7, 9, 11, 13}
	target1 := 7
	fmt.Printf("Índice de %d: %d\n", target1, binarySearch(array1, target1)) // Output: 3

	array2 := []int{2, 4, 6, 8, 10}
	target2 := 5
	fmt.Printf("Índice de %d: %d\n", target2, binarySearch(array2, target2)) // Output: -1

	array3 := []int{10, 20, 30, 40, 50, 60, 70}
	target3 := 60
	fmt.Printf("Índice de %d: %d\n", target3, binarySearch(array3, target3)) // Output: 5

	array4 := []int{5, 15, 25, 35, 45}
	target4 := 5
	fmt.Printf("Índice de %d: %d\n", target4, binarySearch(array4, target4)) // Output: 0

	array5 := []int{5, 15, 25, 35, 45}
	target5 := 45
	fmt.Printf("Índice de %d: %d\n", target5, binarySearch(array5, target5)) // Output: 4

    testBinarySearch() //for testing
}

func testBinarySearch() {
	// Teste 1: Elemento presente no meio
	array1 := []int{1, 3, 5, 7, 9, 11, 13}
	target1 := 7
	result1 := binarySearch(array1, target1)
	fmt.Printf("Teste 1 - Esperado: 3, Obtido: %d\n", result1)

	// Teste 2: Elemento não presente
	array2 := []int{2, 4, 6, 8, 10}
	target2 := 5
	result2 := binarySearch(array2, target2)
	fmt.Printf("Teste 2 - Esperado: -1, Obtido: %d\n", result2)

	// Teste 3: Elemento no final
	array3 := []int{10, 20, 30, 40, 50, 60, 70}
	target3 := 60
	result3 := binarySearch(array3, target3)
	fmt.Printf("Teste 3 - Esperado: 5, Obtido: %d\n", result3)

	// Teste 4: Elemento no início
	array4 := []int{5, 15, 25, 35, 45}
	target4 := 5
	result4 := binarySearch(array4, target4)
	fmt.Printf("Teste 4 - Esperado: 0, Obtido: %d\n", result4)

	// Teste 5: Elemento no final
	array5 := []int{5, 15, 25, 35, 45}
	target5 := 45
	result5 := binarySearch(array5, target5)
	fmt.Printf("Teste 5 - Esperado: 4, Obtido: %d\n", result5)

	// Teste 6: Array vazio
	array6 := []int{}
	target6 := 10
	result6 := binarySearch(array6, target6)
	fmt.Printf("Teste 6 - Esperado: -1, Obtido: %d\n", result6)

	// Teste 7: Array com um elemento (presente)
	array7 := []int{100}
	target7 := 100
	result7 := binarySearch(array7, target7)
	fmt.Printf("Teste 7 - Esperado: 0, Obtido: %d\n", result7)

	// Teste 8: Array com um elemento (não presente)
	array8 := []int{100}
	target8 := 50
	result8 := binarySearch(array8, target8)
	fmt.Printf("Teste 8 - Esperado: -1, Obtido: %d\n", result8)
}	
```

```
Teste 1 - Esperado: 3, Obtido: 3
Teste 2 - Esperado: -1, Obtido: -1
Teste 3 - Esperado: 5, Obtido: 5
Teste 4 - Esperado: 0, Obtido: 0
Teste 5 - Esperado: 4, Obtido: 4
Teste 6 - Esperado: -1, Obtido: -1
Teste 7 - Esperado: 0, Obtido: 0
Teste 8 - Esperado: -1, Obtido: -1
```

## Explicação da Solução
1 - Inicialização:

- Definimos dois ponteiros, inicio e fim, que representam o início e o fim do intervalo atual de busca dentro do array.

2- Loop de Busca:

- Enquanto inicio for menor ou igual a fim, calculamos o índice meio como a média de inicio e fim.
- Comparação:
    - Se o elemento no índice meio for igual ao target, retornamos meio como o índice encontrado.
    - Se o elemento no índice meio for menor que o target, ajustamos inicio para meio + 1 para buscar na metade direita do array.
    - Caso contrário, ajustamos fim para meio - 1 para buscar na metade esquerda do array.

3- Elemento Não Encontrado:

 - Se o loop terminar sem encontrar o target, retornamos -1 indicando que o elemento não está presente no array.

## Considerações Finais
A busca binária é um algoritmo eficiente para encontrar elementos em arrays ordenados, com complexidade de tempo O(log n). É essencial garantir que o array esteja ordenado antes de aplicar a busca binária.

Dicas para a Prova:

- Entenda o Pré-requisito: A busca binária só funciona em arrays ordenados. Se o array não estiver ordenado, você precisará ordená-lo primeiro (o que pode impactar a complexidade total).

- Cuidado com o Cálculo do Meio: Use meio := inicio + (fim - inicio) / 2 para evitar possíveis overflows que podem ocorrer com (inicio + fim) / 2.

- Casos de Teste: Considere testar arrays com:

    - Elementos únicos
    - Elementos duplicados
    - Elemento no início, meio e fim
    - Arrays vazios
    - Arrays com um único elemento

- Implementação Iterativa vs. Recursiva: Embora a implementação iterativa seja mais comum devido ao menor uso de memória, conhecer a versão recursiva também pode ser útil.

# Extra: Tarefa com Busca Binária (lista não ordenada)

Problema: Dada uma lista de inteiros não ordenada e um número alvo, encontre o índice do alvo na lista. Se o alvo não estiver presente, retorne -1.

Soluções Possíveis:
1. Busca Linear (Linear Search)
2. Ordenar a Lista e Aplicar Busca Binária
3. Utilizar uma Estrutura de Dados Auxiliar (Hash Map)

<hr>

## 1. Busca Linear

Descrição:
A busca linear percorre cada elemento da lista sequencialmente até encontrar o alvo. É simples e não requer que a lista esteja ordenada.

Complexidade:
- Tempo: O(n), onde n é o número de elementos na lista.
- Espaço: O(1), pois não utiliza espaço adicional significativo.

```
package main

import (
	"fmt"
)

// linearSearch percorre a lista sequencialmente para encontrar o target.
func linearSearch(array []int, target int) int {
	for index, value := range array {
		if value == target {
			return index
		}
	}
	return -1
}

func main() {
	// Exemplos de uso da função linearSearch
	lista1 := []int{10, 22, 35, 70, 55, 35}
	target1 := 70
	fmt.Printf("Índice de %d: %d\n", target1, linearSearch(lista1, target1)) // Output: 3

	lista2 := []int{5, 15, 25, 35, 45}
	target2 := 20
	fmt.Printf("Índice de %d: %d\n", target2, linearSearch(lista2, target2)) // Output: -1

    testLinearSearch() //test cases
}

func testLinearSearch() {
	// Teste 1: Elemento presente no meio
	lista1 := []int{10, 22, 35, 70, 55, 35}
	target1 := 70
	result1 := linearSearch(lista1, target1)
	fmt.Printf("Teste 1 - Esperado: 3, Obtido: %d\n", result1)

	// Teste 2: Elemento não presente
	lista2 := []int{5, 15, 25, 35, 45}
	target2 := 20
	result2 := linearSearch(lista2, target2)
	fmt.Printf("Teste 2 - Esperado: -1, Obtido: %d\n", result2)

	// Teste 3: Elemento no início
	lista3 := []int{100, 200, 300}
	target3 := 100
	result3 := linearSearch(lista3, target3)
	fmt.Printf("Teste 3 - Esperado: 0, Obtido: %d\n", result3)

	// Teste 4: Elemento no final
	lista4 := []int{1, 2, 3, 4, 5}
	target4 := 5
	result4 := linearSearch(lista4, target4)
	fmt.Printf("Teste 4 - Esperado: 4, Obtido: %d\n", result4)

	// Teste 5: Lista vazia
	lista5 := []int{}
	target5 := 10
	result5 := linearSearch(lista5, target5)
	fmt.Printf("Teste 5 - Esperado: -1, Obtido: %d\n", result5)

	// Teste 6: Lista com um elemento (presente)
	lista6 := []int{42}
	target6 := 42
	result6 := linearSearch(lista6, target6)
	fmt.Printf("Teste 6 - Esperado: 0, Obtido: %d\n", result6)

	// Teste 7: Lista com um elemento (não presente)
	lista7 := []int{42}
	target7 := 24
	result7 := linearSearch(lista7, target7)
	fmt.Printf("Teste 7 - Esperado: -1, Obtido: %d\n", result7)
}
```

```
Teste 1 - Esperado: 3, Obtido: 3
Teste 2 - Esperado: -1, Obtido: -1
Teste 3 - Esperado: 0, Obtido: 0
Teste 4 - Esperado: 4, Obtido: 4
Teste 5 - Esperado: -1, Obtido: -1
Teste 6 - Esperado: 0, Obtido: 0
Teste 7 - Esperado: -1, Obtido: -1
```
Vantagens:
- implicidade na implementação.
- Não requer que a lista esteja ordenada.
- Funciona bem para listas pequenas ou quando a lista é percorrida apenas uma vez.

Desvantagens:

- Ineficiente para listas grandes devido à complexidade linear.


<hr>

## 2. Ordenar a Lista e Aplicar Busca Binária

Descrição:
Primeiro, ordenamos a lista e depois aplicamos a busca binária. Isso permite aproveitar a eficiência da busca binária (O(log n)) após a ordenação (O(n log n)).

Complexidade:
- Tempo: O(n log n) para ordenar + O(log n) para buscar = O(n log n) no total.
- Espaço: Depende do algoritmo de ordenação utilizado. Em Go, `sort.Ints` ordena in-place, portanto, O(1) espaço adicional.

```
package main

import (
	"fmt"
	"sort"
)

// binarySearch realiza a busca binária em um array ordenado.
func binarySearch(array []int, target int) int {
	inicio := 0
	fim := len(array) - 1

	for inicio <= fim {
		meio := inicio + (fim-inicio)/2
		if array[meio] == target {
			return meio
		} else if array[meio] < target {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

// searchInUnsortedList ordena a lista e realiza a busca binária.
// Retorna o índice na lista ordenada. Se precisar do índice na lista original,
// será necessário mapear antes ou após a ordenação.
func searchInUnsortedList(array []int, target int) int {
	// Copiar a lista para não modificar a original
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	// Ordenar a cópia
	sort.Ints(sortedArray)

	// Realizar a busca binária
	return binarySearch(sortedArray, target)
}

func main() {
	// Exemplos de uso
	lista1 := []int{10, 22, 35, 70, 55, 35}
	target1 := 70
	fmt.Printf("Índice de %d na lista ordenada: %d\n", target1, searchInUnsortedList(lista1, target1)) // Output: 4

	lista2 := []int{5, 15, 25, 35, 45}
	target2 := 20
	fmt.Printf("Índice de %d na lista ordenada: %d\n", target2, searchInUnsortedList(lista2, target2)) // Output: -1
}
```

Considerações Importantes:
- Indices Diferentes: Após ordenar a lista, os índices dos elementos podem mudar. Se for necessário manter a relação com os índices originais, essa abordagem pode não ser adequada sem um mapeamento adicional.
- Modificação da Lista Original: Se for permitido modificar a lista original, você pode ordenar in-place usando sort.Ints(array) antes de aplicar a busca binária.

Vantagens:
- Aproveita a eficiência da busca binária após a ordenação.
- Útil quando você precisa realizar múltiplas buscas na mesma lista.

Desvantagens:
- A ordenação adiciona uma sobrecarga de tempo.
- A necessidade de manter a relação com os índices originais pode complicar a implementação.

<hr>

## 3. Utilizar uma Estrutura de Dados Auxiliar (Hash Map)

Descrição:
Crie um mapa (hash map) onde as chaves são os elementos da lista e os valores são seus respectivos índices. Isso permite buscas em tempo constante O(1) após a construção do mapa.

Complexidade:
- Tempo: O(n) para construir o mapa + O(1) para cada busca.
- Espaço: O(n) para armazenar o mapa.

```
package main

import (
	"fmt"
)

// buildHashMap constrói um mapa onde a chave é o elemento e o valor é o índice.
func buildHashMap(array []int) map[int]int {
	hashMap := make(map[int]int)
	for index, value := range array {
		// Em caso de elementos duplicados, armazenamos o primeiro índice encontrado.
		if _, exists := hashMap[value]; !exists {
			hashMap[value] = index
		}
	}
	return hashMap
}

// searchUsingHashMap utiliza o hash map para encontrar o índice do target.
func searchUsingHashMap(array []int, target int) int {
	hashMap := buildHashMap(array)
	if index, exists := hashMap[target]; exists {
		return index
	}
	return -1
}

func main() {
	// Exemplos de uso
	lista1 := []int{10, 22, 35, 70, 55, 35}
	target1 := 70
	fmt.Printf("Índice de %d: %d\n", target1, searchUsingHashMap(lista1, target1)) // Output: 3

	lista2 := []int{5, 15, 25, 35, 45}
	target2 := 20
	fmt.Printf("Índice de %d: %d\n", target2, searchUsingHashMap(lista2, target2)) // Output: -1

	lista3 := []int{1, 2, 3, 4, 5, 3}
	target3 := 3
	fmt.Printf("Índice de %d: %d\n", target3, searchUsingHashMap(lista3, target3)) // Output: 2
}
```
Considerações Importantes:
- Elementos Duplicados: Nesta implementação, armazenamos o primeiro índice encontrado para cada elemento. Se precisar de todos os índices de elementos duplicados, você pode armazenar uma lista de índices como valor no mapa.
- Espaço Adicional: Requer espaço adicional para armazenar o mapa.

Vantagens:
- Busca muito rápida (O(1)) após a construção do mapa.
- Útil quando múltiplas buscas são realizadas na mesma lista.

Desvantagens:
- Utiliza espaço adicional O(n).
- Requer pré-processamento para construir o mapa.

<hr>

Resumo das Soluções:

| Método                   | Tempo de Busca | Pré-processamento        | Espaço Adicional      | Notas                                                                  |
|--------------------------|----------------|--------------------------|-----------------------|------------------------------------------------------------------------|
| **Busca Linear**         | O(n)           | Nenhum                   | O(1)                  | Simples e direto, sem requisitos especiais.                            |
| **Ordenar + Busca Binária** | O(n log n)     | Ordenar (O(n log n))     | O(n) se copiar        | Eficiente para múltiplas buscas, índices alterados após a ordenação.   |
| **Hash Map**             | O(1)           | Construção do mapa (O(n))| O(n)                  | Extremamente rápido para buscas, mas utiliza mais memória.             |


<hr>

# Exemplo Completo com Todas as Soluções e Testes

Para facilitar o estudo, abaixo está um exemplo completo em Golang que implementa as três soluções discutidas, juntamente com testes para cada uma.

```
package main

import (
	"fmt"
	"sort"
)

// Busca Linear
func linearSearch(array []int, target int) int {
	for index, value := range array {
		if value == target {
			return index
		}
	}
	return -1
}

// Busca Binária
func binarySearch(array []int, target int) int {
	inicio := 0
	fim := len(array) - 1

	for inicio <= fim {
		meio := inicio + (fim-inicio)/2
		if array[meio] == target {
			return meio
		} else if array[meio] < target {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}

// Ordenar + Busca Binária
func searchInUnsortedList(array []int, target int) int {
	// Copiar a lista para não modificar a original
	sortedArray := make([]int, len(array))
	copy(sortedArray, array)

	// Ordenar a cópia
	sort.Ints(sortedArray)

	// Realizar a busca binária
	return binarySearch(sortedArray, target)
}

// Hash Map
func buildHashMap(array []int) map[int]int {
	hashMap := make(map[int]int)
	for index, value := range array {
		if _, exists := hashMap[value]; !exists {
			hashMap[value] = index
		}
	}
	return hashMap
}

func searchUsingHashMap(array []int, target int) int {
	hashMap := buildHashMap(array)
	if index, exists := hashMap[target]; exists {
		return index
	}
	return -1
}

// Função de Testes
func runTests() {
	listas := [][]int{
		{10, 22, 35, 70, 55, 35},
		{5, 15, 25, 35, 45},
		{1, 2, 3, 4, 5, 3},
		{},
		{100},
	}

	targets := []int{70, 20, 3, 10, 100, 50}

	fmt.Println("=== Testes de Busca Linear ===")
	for i, lista := range listas {
		target := targets[i%len(targets)]
		result := linearSearch(lista, target)
		fmt.Printf("Lista: %v, Target: %d, Índice: %d\n", lista, target, result)
	}

	fmt.Println("\n=== Testes de Ordenar + Busca Binária ===")
	for i, lista := range listas {
		target := targets[i%len(targets)]
		result := searchInUnsortedList(lista, target)
		fmt.Printf("Lista Original: %v, Target: %d, Índice na Lista Ordenada: %d\n", lista, target, result)
	}

	fmt.Println("\n=== Testes de Hash Map ===")
	for i, lista := range listas {
		target := targets[i%len(targets)]
		result := searchUsingHashMap(lista, target)
		fmt.Printf("Lista: %v, Target: %d, Índice: %d\n", lista, target, result)
	}
}

func main() {
	runTests()
}
```
```
=== Testes de Busca Linear ===
Lista: [10 22 35 70 55 35], Target: 70, Índice: 3
Lista: [5 15 25 35 45], Target: 20, Índice: -1
Lista: [1 2 3 4 5 3], Target: 3, Índice: 2
Lista: [], Target: 10, Índice: -1
Lista: [100], Target: 100, Índice: 0

=== Testes de Ordenar + Busca Binária ===
Lista Original: [10 22 35 70 55 35], Target: 70, Índice na Lista Ordenada: 5
Lista Original: [5 15 25 35 45], Target: 20, Índice na Lista Ordenada: -1
Lista Original: [1 2 3 4 5 3], Target: 3, Índice na Lista Ordenada: 2
Lista Original: [], Target: 10, Índice na Lista Ordenada: -1
Lista Original: [100], Target: 100, Índice na Lista Ordenada: 0

=== Testes de Hash Map ===
Lista: [10 22 35 70 55 35], Target: 70, Índice: 3
Lista: [5 15 25 35 45], Target: 20, Índice: -1
Lista: [1 2 3 4 5 3], Target: 3, Índice: 2
Lista: [], Target: 10, Índice: -1
Lista: [100], Target: 100, Índice: 0
```

## Conclusão

Quando a lista não está ordenada, a busca linear é a abordagem mais direta e simples. Se a eficiência nas buscas é crucial e você pode gastar tempo e espaço adicionais no pré-processamento, utilizar uma estrutura de dados auxiliar como um hash map pode ser a melhor escolha. A combinação de ordenação e busca binária é útil quando você precisa realizar múltiplas buscas e pode tolerar a alteração dos índices ou pode mapear índices de forma apropriada.

Dicas para a Prova no CodeSignal:

- Entenda as Restrições: Se a lista não está ordenada e você precisa realizar apenas uma busca, a busca linear é provavelmente a melhor opção devido à sua simplicidade e ausência de pré-processamento.

- Considere a Eficiência: Para múltiplas buscas, pense em construir um hash map para otimizar o tempo de busca, mesmo que consuma mais memória.

- Pratique Implementações: Certifique-se de saber implementar tanto a busca linear quanto a busca binária e entender quando cada uma é apropriada.

- Gerencie Casos de Teste: Considere listas vazias, listas com um único elemento, listas com elementos duplicados e alvos no início, meio ou fim da lista.