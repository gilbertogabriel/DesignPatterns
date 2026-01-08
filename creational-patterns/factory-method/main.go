package main

import (
	"fmt"

	factorymethod "github.com/gilbertogabriel/design-patterns/factory-method/factory"
)

func main() {
	fmt.Println("O País está em guerra! Mobilizando as forças especiais para defender nossos biomas!")
	biomasDoBrasil := []string{"amazonia", "caatinga", "montanha", "pantanal"}

	fmt.Println("=== MOBILIZAÇÃO DAS FORÇAS ESPECIAIS BRASILEIRAS 2026 ===")

	for _, bioma := range biomasDoBrasil {
		soldado, err := factorymethod.RecrutarCombatentes(bioma)
		if err != nil {
			fmt.Printf("Erro: %s\n", err)
			continue
		}

		fmt.Printf("\n[Bioma: %s]\n", bioma)
		fmt.Printf("Equipamento: %s\n", soldado.GetEquipamento())
		fmt.Printf("Grito de Guerra: %s\n", soldado.GetGritoDeGuerra())
		fmt.Print("Operar no bioma: ", soldado.OperarNoBioma())
	}
}
