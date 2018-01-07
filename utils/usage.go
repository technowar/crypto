package utils

import "fmt"

func Usage() {
	fmt.Println("Commands available:")
	fmt.Println("  help   Display help.")
	fmt.Println("  watch  Watch top 10 coins.")
	fmt.Println("  coin")
	fmt.Println("    [<symbol>|<name>]  Display details of specific coin. Defaults to Bitcoin.")
	fmt.Println("\nAll data are from https://coinmarketcap.com/api/")
}
