package constant

const (
	SymbolstatusUnopened = iota
	SymbolstatusOpen
	SymbolstatusWhite
	SymbolstatusClosed
	SymbolstatusBlack
)

func GetConstant(constantKey int) int {
	return constantKey
}
