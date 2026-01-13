package factorymethod

// ICombatente: O "Contrato" (Produto Abstrato)
// Define o que todo soldado, independente do bioma, deve ser capaz de fazer.
// O código cliente (Comandante) interagirá apenas com estes métodos.
type ICombatente interface {
	OperarNoBioma() string
	GetEquipamento() string
	GetGritoDeGuerra() string
}
