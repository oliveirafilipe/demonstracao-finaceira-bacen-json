package finstm

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestShouldGenerateJSONStatement(t *testing.T) {
	financialStatements := FinancialStatements{
		Cnpj:             "12345678",
		DocumentCode:     "9011",
		TypeRemittance:   "I",
		ValuesMultiplier: 1000,
		BaseDate:         "012020",
		BaseDatesReferences: []BaseDatesReference{{
			Id:   "dt1",
			Date: "012020",
		}},
		BalancoPatrimonial: BalancoT{
			Statements: []Statement{
				{
					Id:              "conta1",
					Description:     "Value1",
					Level:           "1",
					ParentStatement: "",
					IndividualizedValues: []IndividualizedValue{
						{
							BaseDate: "dt1",
							Amount:   200,
						},
					},
				},
			},
		},
		DRE:   DRET{},
		Caixa: CaixaT{},
		DMPL:  DMPLT{},
		DRA:   DRAT{},
	}

	if err := financialStatements.SaveWithName("../testdata/output.json"); err != nil {
		t.Fatalf("erro while trying to save output test file: %s", err)
	}

	rawExpectedContent, err := ioutil.ReadFile("../testdata/generated_statements.golden.json")
	if err != nil {
		t.Fatalf("Error loading golden file: %s", err)
	}

	rawGeneratedContent, err := ioutil.ReadFile("../testdata/output.json")
	if err != nil {
		t.Fatalf("Error loading generated file: %s", err)
	}

	got := strings.TrimSpace(string(rawGeneratedContent))
	expectedOutput := strings.TrimSpace(string(rawExpectedContent))
	if got != expectedOutput {
		t.Errorf("Want:\n%s\nGot:\n%s", expectedOutput, got)
	}
}
