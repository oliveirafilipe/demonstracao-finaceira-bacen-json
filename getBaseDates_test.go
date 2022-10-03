package main

import (
	"reflect"
	"testing"
)

func TestGetBaseDates(t *testing.T) {
	var baseDatesMap = getBaseDates([]string{"testdata/test_input1.csv", "testdata/test_input2.csv"})
	expectedBaseDatesMap := map[string]string{"BAR": "dt1", "FOO": "dt2", "ZEDD": "dt3"}
	if !reflect.DeepEqual(baseDatesMap, expectedBaseDatesMap) {
		t.Fatalf(`getBaseDates(testingFiles) = %s, expected %s`, baseDatesMap, expectedBaseDatesMap)
	}
}
