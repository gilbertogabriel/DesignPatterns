package factorymethod

import "fmt"

type GuerreiroSelva struct{}

func (s *GuerreiroSelva) OperarNoBioma() {
	fmt.Println("Natação utilitária e orientação em selva. Foco em sobrevivência.")
}
func (s *GuerreiroSelva) GetEquipamento() string {
	return "Brevê da Onça e Embarcações Leves"
}
func (s *GuerreiroSelva) GetGritoDeGuerra() string {
	return "Selva!"
}
