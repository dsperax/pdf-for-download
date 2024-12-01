package main

import "fmt"

func fizzBuzz(n int32) {
    var i int32
    for i = 1 ; i <= n; i++ {
        if (i%3==0 && i%5==0){
            fmt.Println("FizzBuzz")
        }
        if (i%3==0 && i%5!=0){
			fmt.Println("Fizz")
        }
        if (i%3!=0 && i%5==0){
            fmt.Println("Buzz")
        }
        if (i%3!=0 && i%5!=0){
            fmt.Println(i)
        }
    }
}

func main(){
	fmt.Println("iniciando")

	fizzBuzz(15)

	fmt.Println("finalizou!")
}