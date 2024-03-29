package finstm

import (
	"encoding/json"
	"os"
)

type IndividualizedValue struct {
	BaseDate string  `json:"@dtBase"`
	Amount   float64 `json:"@valor"`
}

type Statement struct {
	Id                   string                `json:"@id"`
	Description          string                `json:"@descricao"`
	Level                string                `json:"@nivel"`
	ParentStatement      string                `json:"@contaPai"`
	IndividualizedValues []IndividualizedValue `json:"valoresIndividualizados"`
}

// BalancoPatrimonial
type BalancoT struct {
	Statements []Statement `json:"contas"`
}

// DemonstracaoDoResultado
type DRET struct {
	Statements []Statement `json:"contas"`
}

// DemonstracaoDosFluxosDeCaixa
type CaixaT struct {
	Statements []Statement `json:"contas"`
}

// DemonstracaoDasMutacoesDoPatrimonioLiquido
type DMPLT struct {
	Statements []Statement `json:"contas"`
}

// DemonstracaoDoResultadoAbrangente
type DRAT struct {
	Statements []Statement `json:"contas"`
}

type BaseDatesReference struct {
	Id   string `json:"@id"`
	Date string `json:"@data"`
}

type FinancialStatements struct {
	Cnpj                string               `json:"@cnpj"`
	DocumentCode        string               `json:"@codigoDocumento"`
	TypeRemittance      string               `json:"@tipoRemessa"`
	ValuesMultiplier    int                  `json:"@unidadeMedida"`
	BaseDate            string               `json:"@dataBase"`
	BaseDatesReferences []BaseDatesReference `json:"datasBaseReferencia"`
	BalancoPatrimonial  BalancoT             `json:"BalancoPatrimonial"`
	DRE                 DRET                 `json:"DemonstracaoDoResultado"`
	Caixa               CaixaT               `json:"DemonstracaoDosFluxosDeCaixa"`
	DMPL                DMPLT                `json:"DemonstracaoDasMutacoesDoPatrimonioLiquido"`
	DRA                 DRAT                 `json:"DemonstracaoDoResultadoAbrangente"`
}

func (fst *FinancialStatements) Save() error {
	return fst.SaveWithName("resultado.json")
}

func (fst *FinancialStatements) SaveWithName(fileName string) error {
	file, _ := json.MarshalIndent(fst, "", "  ")

	err := os.WriteFile(fileName, file, 0644)

	return err
}
