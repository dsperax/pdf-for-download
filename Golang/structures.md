## For structure

```
func main() {
    fruits := []string{"apple", "banana", "cherry"}

    // Usando _ para ignorar o índice
    for _, fruit := range fruits {
        fmt.Println(fruit)
    }
}
```

```
func main() {
    fruits := []string{"apple", "banana", "cherry"}

    // Usando _ para ignorar o valor
    for index, _ := range fruits {
        fmt.Println("Index:", index)
    }
}
```

```
func main() {
    // Loop de 5 iterações
    for _ = 0; _ < 5; _++ {
        fmt.Println("This will not compile") // Isso não é válido, pois `_` não pode ser atualizado.
    }
}
```

## Order list (sort)

```
func main() {
	numbers := []int{5, 2, 9, 1, 6}
	sort.Ints(numbers) // Ordena os numeros em ordem crescente
	fmt.Println("Sorted integers:", numbers)
}
```

```
func main() {
	numbers := []int{5, 2, 9, 1, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(numbers))) // Ordem decrescente
	fmt.Println("Sorted integers (descending):", numbers)
}
```

```
func main() {
	strings := []string{"banana", "apple", "cherry", "date"}
	sort.Strings(strings) // Ordena em ordem alfabética
	fmt.Println("Sorted strings:", strings)
}
```

```
type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	// Ordena por idade em ordem crescente
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by age:", people)
}
```

```
func main() {
	numbers := []int{1, 2, 3, 4, 5}
	isSorted := sort.IntsAreSorted(numbers)
	fmt.Println("Is sorted:", isSorted)
}
```