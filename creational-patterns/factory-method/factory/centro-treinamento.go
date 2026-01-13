package factorymethod

// ICentroRecrutamento: A "Fábrica Abstrata" (Creator)
// Define o método de fabricação. Qualquer nova base militar deve seguir este padrão.
type ICentroDeRecrutamento interface {
	Recrutar() ICombatente
}
