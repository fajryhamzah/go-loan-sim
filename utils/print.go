package utils

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/fajryhamzah/go-loan-sim/constants"
)

func PrintNormal(a ...any) {
	fmt.Println(a...)
	fmt.Println(constants.RESET)
}

func PrintInlineGreen(a ...any) {
	fmt.Print(constants.GREEN)
	fmt.Print(a...)
	fmt.Print(constants.RESET)
}

func PrintGreen(a ...any) {
	fmt.Print(constants.GREEN)
	fmt.Println(a...)
	fmt.Print(constants.RESET)
}

func PrintRed(a ...any) {
	fmt.Print(constants.RED)
	fmt.Println(a...)
	fmt.Print(constants.RESET)
}

func PrintBlue(a ...any) {
	fmt.Print(constants.BLUE)
	fmt.Println(a...)
	fmt.Print(constants.RESET)
}

func PrintInlineBlue(a ...any) {
	fmt.Print(constants.BLUE)
	fmt.Print(a...)
	fmt.Print(constants.RESET)
}

func FormatRupiah(amount int) string {
	return "Rp " + humanize.Comma(int64(amount))
}
