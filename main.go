package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int 
	numeroConta int
	saldo float64
}

func (c *ContaCorrente) sacar (valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <=c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}else {
		return "Saldo insuficiente"
	}
}


func main() {
	//Tres maneiras de criar uma nova instancia da struct
	contaDoTiago := ContaCorrente { // 1. Create a new instance of the `ContaCorrente` struct.
		titular: "Tiago", numeroAgencia: 123456, saldo: 125.5,
	}

	contaDaBruna := ContaCorrente{ // 2. Create a new instance of the `ContaCorrente` struct.
		"Bruna", 222, 111222, 200,
	}
	fmt.Println(contaDoTiago)
	fmt.Println(contaDaBruna)

	contaDaSilvia := ContaCorrente{} // 3. Create a new instance of the `ContaCorrente` struct.
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.saldo)
	fmt.Println(contaDaSilvia.sacar(-600))
	fmt.Println(contaDaSilvia.saldo)

	

}