package main

import (
	"fmt"
	"math/rand"
)

func getColor(color int) string {
	switch color {
	case 0:
		return "🟥"
	case 1:
		return "🟧"
	case 2:
		return "🟨"
	case 3:
		return "🟩"
	case 4:
		return "🟦"
	case 5:
		return "🟪"
	case 6:
		return "🟫"
	}

	return ""
}

func getBlock(tile int, color int) string {
	result := "⬛"

	switch tile {
	case 0:
		result = "⬛"
		break
	case 1:
		result = getColor(color)
		break
	}

	return result
}

func getBoard(board [5][8]int, color int, player string, playerPos map[string]int) string {
	content := ""

	for y := -1; y < 6; y++ {
		for x := -1; x < 9; x++ {
			if x == -1 || y == -1 || x == 8 || y == 5 {
				// 테두리
				content += getColor(color)
			} else if x == playerPos["x"] && y == playerPos["y"] {
				content += player
			} else {
				// 안
				content += getBlock(board[y][x], color)
			}
		}
		content += "\n"
	}

	return content
}

func main() {
	board := [5][8]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
	color := rand.Intn(7)

	gameover := false

	player := "😀"
	playerPos := map[string]int{
		"x": rand.Intn(8), "y": rand.Intn(5),
	}

	for !gameover {
		fmt.Println(getBoard(board, color, player, playerPos))
		fmt.Println("> 이동할려면 a,s,d,w 중 입력해주세요 : ")

		value, _ := fmt.Scanln()
		fmt.Println(value)
	}
}
