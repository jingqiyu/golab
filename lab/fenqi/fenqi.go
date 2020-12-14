package fenqi

func CalculateFenQiIncome(price float64, terms int, plus float64) (income float64, realPrice float64) {

	yearRate := 0.035
	monthRate := yearRate / 12

	every := price / float64(terms)
	zhifu := every
	m := price - zhifu

	for i := 0; i < terms-1; i++ {
		income += m * monthRate
		m = m - every
	}
	return income, price - income
}
