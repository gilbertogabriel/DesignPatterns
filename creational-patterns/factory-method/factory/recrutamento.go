package factorymethod

import "errors"

func RecrutarCombatentes(tipo string) (ICombatente, error) {
	switch tipo {
	case "amazonia":
		return &GuerreiroSelva{}, nil
	case "caatinga":
		return &GuerreiroCaatinga{}, nil
	case "montanha":
		return &CombatenteMontanha{}, nil
	case "pantanal":
		return &Pantaneiro{}, nil
	default:
		return nil, errors.New("unidade especializada não encontrada para este bioma")
	}
}
