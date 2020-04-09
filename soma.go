package main

import "fmt"

func Soma(a, b int) (result int){
	result = a + b
	return result
}

func main() {
        fmt.Println("Vamos testar Bruno - Desafio CI")
        result := Soma(5,5)
        fmt.Println(result)
    }
