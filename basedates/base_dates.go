package basedates

import (
	"fmt"
)

func GenerateIDsForDates(dates []string) map[string]string {
	baseDatesMap := map[string]string{}
	i := 1

	for _, date := range dates {
		//https://stackoverflow.com/a/2050629/13152732
		if _, ok := baseDatesMap[date]; !ok && date != "" {
			baseDatesMap[date] = fmt.Sprint("dt", i)
			i++
		}
	}

	return baseDatesMap
}
