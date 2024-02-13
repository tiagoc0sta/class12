package main

import (
	"_/C_/Users/chris/source/repos/Portfolio/Golang/alura-oo-Bank-Golang/contas"

	"./clientes"
)



func main() {
	/*
	//Tres maneiras de criar uma nova instancia da struct
	contaDoTiago := contas.ContaCorrente { // 1. Create a new instance of the `ContaCorrente` struct.
		Titular: "Tiago", NumeroAgencia: 123456, Saldo: 125.5,
	}

	contaDaBruna := contas.ContaCorrente{ // 2. Create a new instance of the `ContaCorrente` struct.
		Titular: "Bruna", NumeroAgencia: 222, NumeroConta: 111222, Saldo: 200,
	}
	fmt.Println(contaDoTiago)
	fmt.Println(contaDaBruna)

	/*contaDaSilvia := ContaCorrente{} // 3. Create a new instance of the `ContaCorrente` struct.
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	//Testanto função sacar na conta 
	fmt.Println(contaDaSilvia.saldo)
	contaDaSilvia.Sacar(200)
	fmt.Println(contaDaSilvia.saldo)

	// Testanto função depositar na conta 
	fmt.Println(contaDaSilvia.saldo)
	status, valor := contaDaSilvia.Depositar(2000) // salva os retornos da função Depositar em variaveis
	fmt.Println(status, valor)
	
	contaDaSilvia := contas.ContaCorrente{	Titular:"Silvia", Saldo:300	}
	contaDoGustavo := contas.ContaCorrente{	Titular:"Gustavo", Saldo:100}

	status := contaDaSilvia.Transferir(200,&contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)*/
	

	clienteBruno := clientes.Titular {Nome: "Bruno", CPF: "123.123.123.12",Profissao: "Desenvolvedor Go"}
	contaDoBruno := contas.ContaCorrente {clienteBruno}
	println(contaDoBruno)//

}