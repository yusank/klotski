/*
 * MIT License
 *
 * Copyright (c) 2017 Yusan Kurban.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/08/30        Yusan Kurban
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func init() {
	chessBoard = make([][]character, 5)
	for i := range chessBoard {
		column := make([]character, 4)
		chessBoard[i] = column
	}

	initCharacter()
}

// clear chess board
func clearBoard() {
	for y, chess := range chessBoard {
		for x := range chess {
			chessBoard[y][x] = empty
		}
	}
}

// put character to default position
func setBoard() {
	for _, c := range all {
		positionX := c.Position.X
		positionY := c.Position.Y
		for y, t := range c.Cover {
			for x := range t {
				chessBoard[positionY+y][positionX+x] = *c
			}
		}
	}
	printBoard()
}

// print current chess board
func printBoard() {
	defaultColor.Println("+----------------------------------------------+")
	for _, col := range chessBoard {
		for _, c := range col {
			if c.Control == "" {
				fmt.Printf("|          |")
			} else {
				c.Color.Printf("| %s(%s) |", c.Name, c.Control)
			}
		}

		fmt.Println()
		defaultColor.Println("+----------------------------------------------+")
	}

}

// check user input
func isValidInt(x int) (ok bool) {
	if x == 1 || x == 0 || x == -1 {
		ok = true
	}

	return
}

// move character. when x = 1 and y = 0 means move c to right
func (c *character) move(x, y int) {
	tx, ty := c.Position.X+x, c.Position.Y+y
	ok := c.isValidMove(tx, ty)
	if !ok {
		warnColor.Println("you cannot move like that")
	} else {
		c.Position.X = tx
		c.Position.Y = ty

	}
}

// if the target pane is empty, it represent it`s valid movement, return true
func (c *character) isValidMove(tx, ty int) bool {
	var ok = true

	for y, cov := range c.Cover {
		if !ok {
			break
		}

		for x := range cov {
			if ty+y > 4 || tx+x > 3 {
				ok = false
				break
			}

			if chessBoard[ty+y][tx+x].Name == c.Name {
				continue
			}

			if chessBoard[ty+y][tx+x].Name == "" {
				ok = true
			} else {
				ok = false
			}
		}
	}

	return ok
}

// check is character name is validated, if input wrong name return nil string
func queryName(input string) (s *character, ok bool) {
	for _, c := range all {
		if c.Control == input {
			s = c
			ok = true
			break
		}
	}

	return
}

var inputReader *bufio.Reader

// read user input
func readInput() (op string) {
	fmt.Print("Command:")
	inputReader = bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		errorColor.Println("read failed")
		return op
	}

	op = input[:len(input)-1]
	return
}

// check if it`s win
func checkWin() (win bool) {
	if caoCao.Position.Y == 3 && caoCao.Position.X == 1 {
		win = true
	}

	return
}

// start to waiting user input, after input handle what user input
func start() {
	byeColor.Println("Please input character name and move direction and divide with space")
	for {
		clearBoard()
		if checkWin() {
			sucColor.Printf("\n Yaay,you win !!!! \n")
			initCharacter()
		}
		setBoard()

		operation := readInput()
		if operation == "" {
			continue
		}

		if operation == "q" {
			byeColor.Println("Thank you for playing.")
			os.Exit(0)
		}

		if operation == "save" {
			fileName, err := marshalAndSave()
			if err != nil {
				errorColor.Printf("Save to file failed,please try again.\n")
			}
			sucColor.Printf("Save to file %s \n", fileName)
			continue
		}

		sli := strings.Split(operation, " ")
		if len(sli) != 3 {
			errorColor.Println("!!! wrong input")
			continue
		}

		char, ok := queryName(sli[0])
		if !ok {
			errorColor.Println("!!! wrong character name")
			continue
		}

		x, err := strconv.Atoi(sli[1])
		if err != nil || !isValidInt(x) {
			errorColor.Println("!!! wrong input move")
			continue
		}

		y, err := strconv.Atoi(sli[2])
		if err != nil || !isValidInt(y) {
			errorColor.Println("!!! wrong input move")
			continue
		}

		char.move(x, y)

	}
}

// print some tips for player
func printTips() {
	tips := "Welcome to klotski game!! \n" +
		"Here is some tips for how to play: \n" +
		"1.There is some character in game and the same name in different pane represent one character who cover multi pane." +
		"As you can see caocao(cc) cover 4 pane. \n" +
		"2.You can only move one character one step at one times to empty pane. \n" +
		"3.It`s easy to find that there only two pane is empty. \n" +
		"4.If you want move `zu` to right, you just input `zu 1 0` and press enter. \n" +
		"5.If you want to move one character to up input `name 0 -1`, to down `name 0 1`, to left `name -1 0`. \n" +
		"6.If you want to save you schedule just input save,then you can exit your game. \n" +
		"7.When you wand load your game file which you saved before,run game with `./klotski filename.kls`. \n" +
		"8.Once you move the Boss `cc` to the bottom of ChessBoard, you win.\n" +
		"9.Input `q` for quit, please insure your game schedule is saved. \n" +
		"Have a fun! \n \n"

	fmt.Printf(tips)
}

func main() {
	if len(os.Args) > 1 {
		postfix := strings.Split(os.Args[1], ".")
		if postfix[1] != validPostfix {
			fmt.Printf("invalid file type %s \n", postfix[1])
			os.Exit(1)
		}

		err := readFile(os.Args[1])
		if err != nil {
			fmt.Println("Read file error with:", err)
		} else {
			fmt.Printf("Read %s， please continue your game. \n\n", os.Args[1])
		}
	}

	//setBoard()
	printTips()
	start()
}
