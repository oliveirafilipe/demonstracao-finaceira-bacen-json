package datasource

import (
	"reflect"
	"testing"
	"warrenbrasil/demonstracao-finaceira-bacen-json/basedates"
	"warrenbrasil/demonstracao-finaceira-bacen-json/finstm"
)

func TestGetBaseDates(t *testing.T) {
	csv := CSV{
		Path: "../testdata/test_input1.csv",
	}

	dates, err := csv.GetBaseDates()

	if err != nil {
		t.Fatal(err)
	}

	expectedDates := []string{"BAR", "FOO", "ZEDD"}

	if len(expectedDates) != len(dates) {
		t.Fatalf("Expect len = %d, got len %d", len(expectedDates), len(dates))
	}

	for i := 0; i < len(expectedDates); i++ {
		if expectedDates[i] != dates[i] {
			t.Fatalf("Error in index %d. Expected: %s, got %s", i, expectedDates[i], dates[i])
		}
	}

}

func TestShouldGetStatement(t *testing.T) {
	csv := CSV{
		Path: "../testdata/test_input3.csv",
	}

	dates, err := csv.GetBaseDates()

	if err != nil {
		t.Fatal(err)
	}

	expectedStatements := []finstm.Statement{
		{
			Id:              "conta1",
			Description:     "Ativo",
			Level:           "1",
			ParentStatement: "",
			IndividualizedValues: []finstm.IndividualizedValue{
				{
					BaseDate: "dt1",
					Amount:   56884,
				},
				{
					BaseDate: "dt2",
					Amount:   495.80,
				},
			},
		},
		{
			Id:              "conta2",
			Description:     "Circulante",
			Level:           "1.1",
			ParentStatement: "conta1",
			IndividualizedValues: []finstm.IndividualizedValue{
				{
					BaseDate: "dt1",
					Amount:   55201,
				},
				{
					BaseDate: "dt2",
					Amount:   44907,
				},
			},
		},
		{
			Id:              "conta3",
			Description:     "DISPONIBILIDADES",
			Level:           "1.1.1",
			ParentStatement: "conta2",
			IndividualizedValues: []finstm.IndividualizedValue{
				{
					BaseDate: "dt1",
					Amount:   2357,
				},
			},
		},
	}

	receivedStatements, err := csv.GetStatements(basedates.GenerateIDsForDates(dates))

	if err != nil {
		t.Fatal(err)
	}

	if len(receivedStatements) != len(expectedStatements) {
		t.Fatalf("Expect len = %d, got len %d", len(expectedStatements), len(receivedStatements))
	}

	for i := 0; i < len(expectedStatements); i++ {
		if !reflect.DeepEqual(expectedStatements[i], receivedStatements[i]) {
			t.Fatalf("Expect err %d", i)
		}
	}
}
