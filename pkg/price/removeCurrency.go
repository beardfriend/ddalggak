package price

import (
	"regexp"
	"strconv"
	"strings"
)

func RemoveKoreanCurrencySymbol(price string) int {
	pattern1 := `\([^)]*\)`
	re1 := regexp.MustCompile(pattern1)
	result := re1.ReplaceAllString(price, "")

	pattern := "[,~원]"
	re := regexp.MustCompile(pattern)
	result = re.ReplaceAllString(result, "")
	result = strings.TrimSpace(result)
	priceInt, _ := strconv.Atoi(result)
	return priceInt
}
