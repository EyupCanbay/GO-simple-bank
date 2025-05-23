package utils

// constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	TRY = "TRY"
)

// IsSupportedCurrency returns true is the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, TRY:
		return true
	}
	return false
}
