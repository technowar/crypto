package utils

import "fmt"

func Usage() {
	fmt.Println("Commands available:")
	fmt.Println("  help               Display help.")
	fmt.Println("  watch              Watch top 10 coins.")
	fmt.Println("  coin <coin>        Display details of specific coin. Defaults to Bitcoin.")
	fmt.Println("    args:")
	fmt.Println("      <coin>  [<symbol>|<name>]")
	fmt.Println("  price <from> <to>  Converts price. Defaults to Bitcoin.")
	fmt.Println("    args:")
	fmt.Println("      <from>  [<symbol>]")
	fmt.Println("      <to>    [<symbol>]")
}
