package factorymethod

import "fmt"

// --- LOGÍSTICA DE ACESSO ---

// ObterCentroDeRecrutamento: Função auxiliar para entregar a FÁBRICA correta.
// No Factory Method, o cliente escolhe a fábrica, e a fábrica escolhe o produto.
func ObterCentroDeRecrutamento(bioma string) (ICentroDeRecrutamento, error) {
	switch bioma {
	case "amazonia":
		return &BaseSelva{}, nil
	case "caatinga":
		return &BaseCaatinga{}, nil
	case "pantanal":
		return &BasePantanal{}, nil
	case "montanha":
		return &BaseMontanha{}, nil
	default:
		return nil, fmt.Errorf("bioma %s não possui centro de treinamento", bioma)
	}
}
