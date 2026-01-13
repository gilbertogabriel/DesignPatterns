package main

import (
	"fmt"

	factorymethod "github.com/gilbertogabriel/design-patterns/factory-method/factory"
)

func main() {
	fmt.Println("O País está em guerra! Mobilizando as forças especiais!")
	biomasDoBrasil := []string{"amazonia", "caatinga", "pantanal"}

	fmt.Println("=== MOBILIZAÇÃO DAS FORÇAS ESPECIAIS BRASILEIRAS 2026 ===")

	for _, bioma := range biomasDoBrasil {
		// 1. O Comandante pede o CENTRO (Fábrica), não o soldado.
		centro, err := factorymethod.ObterCentroDeRecrutamento(bioma)
		if err != nil {
			fmt.Printf("Erro: %s\n", err)
			continue
		}

		// 2. O Centro (Fábrica) entrega o soldado.
		// Desacoplamento total: a main não sabe qual tipo de soldado é esse.
		soldado := centro.Recrutar()

		fmt.Printf("\n[Bioma: %s]\n", bioma)
		fmt.Printf("Equipamento: %s\n", soldado.GetEquipamento())
		fmt.Printf("Grito de Guerra: %s\n", soldado.GetGritoDeGuerra())
		fmt.Printf("Operação: %s\n", soldado.OperarNoBioma())
	}
}
