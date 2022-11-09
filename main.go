package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Box struct {
	goalX int
	goalY int
	x     int
	y     int
}

func (bx *Box) move(x int, y int) {
	bx.x += x
	bx.y += y
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
					if x == bx.x && y == bx.y {
						// 블럭
						if bx.goalX == bx.x && bx.goalY == bx.y {
							content += "✅"
						} else {
							content += getColor(color)
						}
					} else if x == bx.goalX && y == bx.goalY {
						content += "❌"
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

func checkWin(box []Box) bool {
	var stack int
	for _, bx := range box {
		if bx.x == bx.goalX && bx.y == bx.goalY {
			stack++
		}
	}

	return len(box) == stack
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

	box = append(box, Box{goalX: 2, goalY: 4, x: 1, y: 3})
	// box = append(box, Box{goalX: 1, goalY: 2, x: 3, y: 2})

	rand.Seed(time.Now().UnixNano())
	color := rand.Intn(7)

	gameover := false

	player := "😀"
	playerPos := map[string]int{
		"x": 2, "y": 3,
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
			if box[idx].x == playerPos["x"] && box[idx].y == playerPos["y"] {
				box[idx].move(directionX, directionY)

				if box[idx].x < 0 {
					playerPos["x"] -= directionX
					box[idx].x = 0
				}

				if box[idx].x > 7 {
					playerPos["x"] += directionX
					box[idx].x = 7
				}

				if box[idx].y < 0 {
					playerPos["y"] -= directionY
					box[idx].y = 0
				}

				if box[idx].y > 4 {
					playerPos["y"] -= directionY
					box[idx].y = 4
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

		fmt.Println("-----------------------------")
		gameover = checkWin(box)
	}

	fmt.Println(getBoard(board, box, color, player, playerPos))
	fmt.Println("You win!")
}
