package basedates

import (
	"reflect"
	"testing"
)

func TestGetBaseDates(t *testing.T) {
	var baseDatesMap = GenerateIDsForDates([]string{"BAR", "FOO", "ZEDD", "YIELD", "FOO", "BAR"})
	expectedBaseDatesMap := map[string]string{"BAR": "dt1", "FOO": "dt2", "ZEDD": "dt3", "YIELD": "dt4"}
	if !reflect.DeepEqual(baseDatesMap, expectedBaseDatesMap) {
		t.Fatalf(`getBaseDates(testingFiles) = %s, expected %s`, baseDatesMap, expectedBaseDatesMap)
	}
}
