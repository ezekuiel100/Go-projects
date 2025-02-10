package main

import (
	"fmt"  
    "strconv"
)


var expression  = "666 + 25"
var operators = [4]string{"+", "-" , "/" , "*"}
var op string

var num1 int 
var num2 int

var result int 
var str1 string
var str2 string

func main() {
 

	for _, el := range expression  {
		char := string(el)
		if char == " " {
			continue
		}

		for _ , o := range operators {			
		  if char == o {
			op = char	
			break	
		  }
		}

		if(op == ""){
		    str1 += char 
			num1, _ = strconv.Atoi(str1)
		}else if char != op {
			str2 += char 
			num2, _  = strconv.Atoi(str2)
		}
		
	}
	

	switch op {
	case "+":
     result =  num1 + num2
	case "-":		
	 result = num1 - num2		
	case "*":
	 result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: Divisão por zero")
			return
		}
		
	 result = num1 / num2
	 
	 default: 
	 fmt.Println("Erro: Operador inválido")
	 return
	}
	
	fmt.Printf("num1: %d\nop: %s\nnum2: %d\nResultado: %d\n", num1, op, num2, result)



}
