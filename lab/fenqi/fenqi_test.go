package fenqi

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalculateFenQiIncome(t *testing.T) {
	//fmt.Println(CalculateFenQiIncome(11988, 24, 0))
	s := "invte_code_2BC3LK"

	split := strings.Split(s, "invte_code_")[1]
	fmt.Println(split)
}
