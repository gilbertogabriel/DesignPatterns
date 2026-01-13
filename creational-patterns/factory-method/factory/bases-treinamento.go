package factorymethod

// --- FÁBRICAS CONCRETAS (Bases de Treinamento) ---
// Cada base abaixo é responsável por instanciar o soldado correto.
// Note que o retorno é sempre a interface ICombatente, ocultando a struct concreta.

// BaseCaatinga: Sabe exatamente como treinar e equipar um combatente do sertão.
type BaseCaatinga struct{}

func (b *BaseCaatinga) Recrutar() ICombatente {
	return &GuerreiroCaatinga{}
}

// BaseSelva: Especializada em preparar soldados para a selva úmida.
type BaseSelva struct{}

func (b *BaseSelva) Recrutar() ICombatente {
	return &GuerreiroSelva{}
}

// BasePantanal: Focada em operações anfíbias e áreas alagadas.
type BasePantanal struct{}

func (b *BasePantanal) Recrutar() ICombatente {
	return &GuerreiroPantanal{}
}

// BaseMontanha: Prepara especialistas para grandes altitudes e escalada.
type BaseMontanha struct{}

func (b *BaseMontanha) Recrutar() ICombatente {
	return &CombatenteMontanha{}
}
