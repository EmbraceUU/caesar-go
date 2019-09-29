package constant

import (
	"fmt"
	"testing"
)

func TestConstantReturn(t *testing.T) {
	fmt.Println(GetConstant(SymbolstatusUnopened))
	fmt.Println(GetConstant(SymbolstatusOpen))
	fmt.Println(GetConstant(SymbolstatusWhite))
	fmt.Println(GetConstant(SymbolstatusClosed))
	fmt.Println(GetConstant(SymbolstatusBlack))
}
