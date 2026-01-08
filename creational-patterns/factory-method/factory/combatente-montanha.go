package factorymethod

import "fmt"

type CombatenteMontanha struct{}

func (m *CombatenteMontanha) OperarNoBioma() {
	fmt.Println("Escalada, rapel e resgate em encostas de difícil acesso.")
}

func (m *CombatenteMontanha) GetEquipamento() string {
	return "Cordas, Mosquetões e Equipamento de Escalada"
}

func (m *CombatenteMontanha) GetGritoDeGuerra() string { return "Para a frente e para o alto!" }
