package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringToFloat(price string) (f float64) {
	if price == "" {
		fmt.Println("empty price")
		return
	}
	normPrice := normalizeGerman(strings.ReplaceAll(price, "€", ""))
	f, err := strconv.ParseFloat(strings.Trim(normPrice, " "), 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func normalizeGerman(old string) string {
	count := strings.Count(old, ".")
	s := strings.Replace(old, ",", ".", -1)
	return strings.Replace(s, ".", "", count)
}
