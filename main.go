package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Box struct {
	goal bool
	x    int
	y    int
}

type Goal struct {
	x int
	y int
}

const (
	WIDTH  = 10
	HEIGHT = 8
)

func (bx *Box) move(x, y int) {
	bx.x += x
	bx.y += y
}

func (bx *Box) setGoal(goal bool) {
	bx.goal = goal
}

func getBlock(tile int, color int) string {
	switch tile {
	case 0:
		return "⬛"
	case 1:
		return getColor(color)
	case 2:
		return "❌"
	case 3:
		return "✅"
	}

	return "⬛"
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

func getBoard(board [HEIGHT][WIDTH]int, box []Box, goal []Goal, color int, player string, playerPos map[string]int) string {
	content := ""

	for _, ga := range goal {
		board[ga.y][ga.x] = 2
	}

	for _, bx := range box {
		if bx.goal {
			board[bx.y][bx.x] = 3
		} else {
			board[bx.y][bx.x] = 1
		}
	}

	for y := -1; y < HEIGHT+1; y++ {
		for x := -1; x < WIDTH+1; x++ {
			if x == -1 || y == -1 || x == WIDTH || y == HEIGHT {
				// 테두리
				content += getColor(color)
			} else if x == playerPos["x"] && y == playerPos["y"] {
				// 플레이어
				content += player
			} else {
				content += getBlock(board[y][x], color)
			}
		}
		content += "\n"
	}

	return content
}

func checkWin(box []Box, goal []Goal) bool {
	var stack int
	for idx := 0; idx < len(box); idx++ {
		for _, ga := range goal {
			if ga.x == box[idx].x && ga.y == box[idx].y {
				stack++
				box[idx].setGoal(true)
			}
		}
	}

	return len(box) == stack
}

func main() {
	board := [HEIGHT][WIDTH]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	var box []Box
	var goal []Goal

	box = append(box, Box{goal: false, x: 1, y: 3})
	box = append(box, Box{goal: false, x: 3, y: 2})
	box = append(box, Box{goal: false, x: 7, y: 4})
	box = append(box, Box{goal: false, x: 4, y: 3})
	box = append(box, Box{goal: false, x: 3, y: 1})

	goal = append(goal, Goal{x: 2, y: 4})
	goal = append(goal, Goal{x: 1, y: 2})
	goal = append(goal, Goal{x: 9, y: 4})
	goal = append(goal, Goal{x: 1, y: 3})
	goal = append(goal, Goal{x: 2, y: 7})

	rand.Seed(time.Now().UnixNano())
	color := rand.Intn(7)

	gameover := false

	player := "😀"
	playerPos := map[string]int{
		"x": 2, "y": 3,
	}

	for !gameover {
		fmt.Println(getBoard(board, box, goal, color, player, playerPos))
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

				if box[idx].x > WIDTH-1 {
					playerPos["x"] += directionX
					box[idx].x = WIDTH - 1
				}

				if box[idx].y < 0 {
					playerPos["y"] -= directionY
					box[idx].y = 0
				}

				if box[idx].y > HEIGHT-1 {
					playerPos["y"] -= directionY
					box[idx].y = HEIGHT - 1
				}
			}
		}

		if playerPos["x"] < 0 {
			playerPos["x"] = 0
		}
		if playerPos["x"] > WIDTH-1 {
			playerPos["x"] = WIDTH - 1
		}
		if playerPos["y"] < 0 {
			playerPos["y"] = 0
		}
		if playerPos["y"] > HEIGHT-1 {
			playerPos["y"] = HEIGHT - 1
		}

		fmt.Println("-----------------------------")
		gameover = checkWin(box, goal)
	}

	fmt.Println(getBoard(board, box, goal, color, player, playerPos))
	fmt.Println("You win!")
}
