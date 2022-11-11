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

type Step struct {
	_type string
	idx   int
	x     int
	y     int
	goal  bool
}

const (
	WIDTH  = 15
	HEIGHT = 12
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

func checkWin(box []Box, goal []Goal, steps *[]Step) bool {
	var stack int
	for idx := 0; idx < len(box); idx++ {
		for _, ga := range goal {
			if ga.x == box[idx].x && ga.y == box[idx].y {
				stack++
				if len(*steps) > 0 {
					*steps = append(*steps, Step{_type: "box-goal", idx: idx, goal: box[idx].goal})
				}
				box[idx].setGoal(true)
			}
		}
	}

	return len(goal) == stack
}

func cancelGoal(box []Box, goal []Goal) {
	for idx := 0; idx < len(box); idx++ {
		for _, ga := range goal {
			if box[idx].goal && !(ga.x == box[idx].x && ga.y == box[idx].y) {
				box[idx].setGoal(false)
			}
		}
	}
}

func BoxFilter(vs []Box, f func(Box) bool) []Box {
	vsf := make([]Box, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func contains(s []string, x string) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}

	return false
}

func reset(box *[]Box, goal *[]Goal) {
	*box = []Box{}
	*goal = []Goal{}

	*box = append(*box, Box{goal: false, x: 1, y: 3})
	*box = append(*box, Box{goal: false, x: 3, y: 2})
	*box = append(*box, Box{goal: false, x: 7, y: 4})
	*box = append(*box, Box{goal: false, x: 4, y: 3})
	*box = append(*box, Box{goal: false, x: 3, y: 1})
	*box = append(*box, Box{goal: false, x: 9, y: 5})
	*box = append(*box, Box{goal: false, x: 12, y: 10})
	*box = append(*box, Box{goal: false, x: 3, y: 11})
	*box = append(*box, Box{goal: false, x: 14, y: 4})
	*box = append(*box, Box{goal: false, x: 3, y: 3})
	*box = append(*box, Box{goal: false, x: 1, y: 10})
	*box = append(*box, Box{goal: false, x: 10, y: 2})
	*box = append(*box, Box{goal: false, x: 10, y: 6})
	*box = append(*box, Box{goal: false, x: 7, y: 10})

	*goal = append(*goal, Goal{x: 2, y: 4})
	*goal = append(*goal, Goal{x: 1, y: 2})
	*goal = append(*goal, Goal{x: 9, y: 4})
	*goal = append(*goal, Goal{x: 7, y: 3})
	*goal = append(*goal, Goal{x: 2, y: 7})
	*goal = append(*goal, Goal{x: 1, y: 5})
	*goal = append(*goal, Goal{x: 14, y: 2})
	*goal = append(*goal, Goal{x: 4, y: 11})
	*goal = append(*goal, Goal{x: 2, y: 9})
	*goal = append(*goal, Goal{x: 5, y: 3})
	*goal = append(*goal, Goal{x: 14, y: 1})
	*goal = append(*goal, Goal{x: 10, y: 3})
	*goal = append(*goal, Goal{x: 10, y: 11})
	*goal = append(*goal, Goal{x: 14, y: 11})
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

	var steps [][]Step
	var tempStep []Step

	reset(&box, &goal)

	rand.Seed(time.Now().UnixNano())
	color := rand.Intn(7)

	gameover := false

	player := "😀"
	playerPos := map[string]int{
		"x": 2, "y": 3,
	}

	for !gameover {
		fmt.Println(getBoard(board, box, goal, color, player, playerPos))
		fmt.Printf("> 이동할려면 a, s, d, w, undo, reset 중 입력해주세요 : ")

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

		passCommand := []string{"undo", "u", "reset", "r"}

		if !contains(passCommand, input) {
			tempStep = append(tempStep, Step{_type: "player-move", x: playerPos["x"], y: playerPos["y"]})
		}

		if (input == "undo" || input == "u") && len(steps) > 0 {
			step := steps[len(steps)-1]

			for _, st := range step {
				switch st._type {
				case "player-move":
					playerPos["x"] = st.x
					playerPos["y"] = st.y
				case "box-move":
					box[st.idx].x = st.x
					box[st.idx].y = st.y
				case "box-goal":
					box[st.idx].setGoal(st.goal)
				}
			}

			steps = steps[:len(steps)-1]
		}

		if input == "reset" || input == "r" {
			playerPos = map[string]int{
				"x": 2, "y": 3,
			}
			reset(&box, &goal)
			continue
		}

		playerPos["x"] += directionX
		playerPos["y"] += directionY

		for idx := 0; idx < len(box); idx++ {
			newbox := BoxFilter(box, func(b Box) bool {
				return b.x == box[idx].x+directionX && b.y == box[idx].y+directionY
			})

			if box[idx].x == playerPos["x"] && box[idx].y == playerPos["y"] && len(newbox) == 0 {
				tempStep = append(tempStep, Step{_type: "box-move", idx: idx, x: box[idx].x, y: box[idx].y})
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

		colidbox := BoxFilter(box, func(b Box) bool {
			return b.x == playerPos["x"] && b.y == playerPos["y"]
		})
		if len(colidbox) != 0 {
			playerPos["x"] -= directionX
			playerPos["y"] -= directionY
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
		cancelGoal(box, goal)
		gameover = checkWin(box, goal, &tempStep)

		if len(tempStep) > 0 {
			steps = append(steps, tempStep)
		}
		tempStep = []Step{}
	}

	fmt.Println(getBoard(board, box, goal, color, player, playerPos))
	fmt.Println("You win!")
}
