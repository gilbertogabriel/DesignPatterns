package factorymethod

// --- PRODUTOS CONCRETOS ---
// Cada struct abaixo representa um especialista treinado para um cenário específico.
// Todos eles devem implementar os métodos da interface ICombatente.

// GuerreiroCaatinga: Especialista em operações no semiárido brasileiro.
type GuerreiroCaatinga struct{}

func (c *GuerreiroCaatinga) OperarNoBioma() string {
	return "Táticas de emboscada e movimento silencioso no semiárido."
}
func (c *GuerreiroCaatinga) GetEquipamento() string {
	return "Uniforme de Couro Característico (Proteção contra espinhos)"
}
func (c *GuerreiroCaatinga) GetGritoDeGuerra() string {
	return "Caatinga!"
}

// GuerreiroSelva: Combatente especializado no bioma amazônico.
type GuerreiroSelva struct{}

func (s *GuerreiroSelva) OperarNoBioma() string {
	return "Natação utilitária e orientação em selva. Foco em sobrevivência."
}
func (s *GuerreiroSelva) GetEquipamento() string {
	return "Brevê da Onça e Embarcações Leves"
}
func (s *GuerreiroSelva) GetGritoDeGuerra() string {
	return "Selva!"
}

// GuerreiroPantanal: Especialista em áreas alagadiças e transição.
type GuerreiroPantanal struct{}

func (p *GuerreiroPantanal) OperarNoBioma() string {
	return "Operações anfíbias e sobrevivência em áreas alagadas com fauna agressiva."
}
func (p *GuerreiroPantanal) GetEquipamento() string {
	return "Embarcações Leves e Equipamento de Mergulho Raso"
}
func (p *GuerreiroPantanal) GetGritoDeGuerra() string {
	return "Pantanal!"
}

// CombatenteMontanha: Focado em terrenos de altitude e difícil acesso.
type CombatenteMontanha struct{}

func (m *CombatenteMontanha) OperarNoBioma() string {
	return "Escalada, rapel e resgate em encostas de difícil acesso."
}
func (m *CombatenteMontanha) GetEquipamento() string {
	return "Cordas, Mosquetões e Equipamento de Escalada"
}
func (m *CombatenteMontanha) GetGritoDeGuerra() string {
	return "Para a frente e para o alto!"
}
