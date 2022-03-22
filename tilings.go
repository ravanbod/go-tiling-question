package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter n: ")
	fmt.Scan(&n)
	tiles := make([]rune, n)
	tile(tiles, 0)
}

func tile(tiles []rune, x int) {
	if x >= len(tiles)-1 {
		if x == len(tiles)-1 {
			tiles[x] = '|'
		}
		showTiles(tiles)
		return
	}
	tiles[x] = '|'
	tile(tiles, x+1)
	tiles[x] = '-'
	tiles[x+1] = '-'
	tile(tiles, x+2)
}

func showTiles(tiles []rune) {
	for i := 0; i < 2; i++ {
		for j := 0; j < len(tiles); j++ {
			fmt.Print(string(tiles[j]))
		}
		fmt.Println()
	}
	fmt.Println()
}
