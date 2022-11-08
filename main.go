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
		fmt.Printf("> 이동할려면 a,s,d,w 중 입력해주세요 : ")

		var input string
		fmt.Scanln(&input)

		if input == "a" {
			playerPos["x"] -= 1
		}
		if input == "d" {
			playerPos["x"] += 1
		}
		if input == "w" {
			playerPos["y"] -= 1
		}
		if input == "s" {
			playerPos["y"] += 1
		}

		if board[playerPos["y"]][playerPos["x"]] == 1 {
			if input == "a" {
				playerPos["x"] -= 1
			}
			if input == "d" {
				playerPos["x"] += 1
			}
			if input == "w" {
				playerPos["y"] -= 1
			}
			if input == "s" {
				playerPos["y"] += 1
			}
		}

		if playerPos["x"] < 0 {
			playerPos["x"] = 0
		}
		if playerPos["x"] > 7 {
			playerPos["x"] = 7
		}
		if playerPos["y"] < 0 {
			playerPos["y"] = 0
		}
		if playerPos["y"] > 4 {
			playerPos["y"] = 4
		}

		fmt.Println(input)
		fmt.Println("-----------------------------")
	}
}
