package factorymethod

import "fmt"

type GuerreiroCaatinga struct{}

func (c *GuerreiroCaatinga) OperarNoBioma() {
	fmt.Println("Táticas de emboscada e movimento silencioso no semiárido.")
}
func (c *GuerreiroCaatinga) GetEquipamento() string {
	return "Uniforme de Couro Característico"
}
func (c *GuerreiroCaatinga) GetGritoDeGuerra() string {
	return "Caatinga!"
}
