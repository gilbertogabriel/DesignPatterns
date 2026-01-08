package factorymethod

import "fmt"

type Pantaneiro struct{}

func (p *Pantaneiro) OperarNoBioma() {
	fmt.Println("Operações anfíbias e sobrevivência em áreas alagadas com fauna agressiva.")
}
func (p *Pantaneiro) GetEquipamento() string {
	return "Embarcações Leves e Equipamento de Mergulho Raso"
}
func (p *Pantaneiro) GetGritoDeGuerra() string { return "Pantanal!" }
