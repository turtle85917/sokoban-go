package main

import (
	"fmt"
	"math/rand"
)

type Box struct {
	Goal bool
	X    int
	Y    int
}

func (bx *Box) move(x int, y int) {
	bx.X += x
	bx.Y += y
}

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

func getBoard(board [5][8]int, box []Box, color int, player string, playerPos map[string]int) string {
	content := ""

	for y := -1; y < 6; y++ {
		for x := -1; x < 9; x++ {
			if x == -1 || y == -1 || x == 8 || y == 5 {
				// 테두리
				content += getColor(color)
			} else if x == playerPos["x"] && y == playerPos["y"] {
				// 플레이어
				content += player
			} else {
				for _, bx := range box {
					if x == bx.X && y == bx.Y {
						// 블럭
						if bx.Goal {
							content += "✅"
						} else {
							content += getColor(color)
						}
					} else {
						// 안
						content += "⬛"
					}
				}
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
	var box []Box

	box = append(box, Box{Goal: false, X: 4, Y: 3})

	color := rand.Intn(7)

	gameover := false

	player := "😀"
	playerPos := map[string]int{
		"x": rand.Intn(8), "y": rand.Intn(5),
	}

	for !gameover {
		fmt.Println(getBoard(board, box, color, player, playerPos))
		fmt.Printf("> 이동할려면 a,s,d,w 중 입력해주세요 : ")

		var input string
		fmt.Scanln(&input)

		var directionX int
		var directionY int

		if input == "a" {
			directionX = -1
		}
		if input == "d" {
			directionX = 1
		}
		if input == "w" {
			directionY = -1
		}
		if input == "s" {
			directionY = 1
		}

		playerPos["x"] += directionX
		playerPos["y"] += directionY

		for idx := 0; idx < len(box); idx++ {
			if box[idx].X == playerPos["x"] && box[idx].Y == playerPos["y"] {
				box[idx].move(directionX, directionY)

				if box[idx].X < 0 {
					playerPos["x"] -= directionX
					box[idx].X = 0
				}

				if box[idx].X > 7 {
					playerPos["x"] += directionX
					box[idx].X = 7
				}

				if box[idx].Y < 0 {
					playerPos["y"] -= directionY
					box[idx].Y = 0
				}

				if box[idx].Y > 4 {
					playerPos["y"] -= directionY
					box[idx].Y = 4
				}
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
