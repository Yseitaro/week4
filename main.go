package main

import (
	"fmt"
)

type Board struct {
	tokens []int // tokens[0] -> (0,0), tokens[1] -> (0,1), ...
}

func (b *Board) put(x, y int, u string) {
	if u == "o" {
		b.tokens[x+3*y] = 1
	} else {
		b.tokens[x+3*y] = -1
	}
}
func (b *Board) get(x, y int) string {
	if b.tokens[x+3*y] == 1 {
		return "o"
	}
	if b.tokens[x+3*y] == -1 {
		return "x"
	}
	return "-"
}

// 縦の値の積が０だった場合ー＞その列に空きがあるため勝負は続く
// 積が0ではなくかつ縦の値が3もしくは-3だった場合勝負がついたと知らせる
// ななめは２通りを条件分岐で調べる
func (b *Board) check() string {

	for y := 0; y < 3; y++ {
		//積が0でない場合
		if b.tokens[0+3*y]*b.tokens[1+3*y]*b.tokens[2+3*y] != 0 {
			if b.tokens[0+3*y]+b.tokens[1+3*y]+b.tokens[2+3*y] == 3 || b.tokens[0+3*y]+b.tokens[1+3*y]+b.tokens[2+3*y] == -3 {
				return "yes"
			}
		}
	}
	for x := 0; x < 3; x++ {
		//積が0でない場合
		if b.tokens[x+3*0]*b.tokens[x+3*1]*b.tokens[x+3*2] != 0 {
			if b.tokens[x+3*0]+b.tokens[x+3*1]+b.tokens[x+3*2] == 3 || b.tokens[x+3*0]+b.tokens[x+3*1]+b.tokens[x+3*2] == -3 {
				return "yes"
			}
		}
	}
	if b.tokens[0]+b.tokens[4]+b.tokens[8] == 3 || b.tokens[0]+b.tokens[4]+b.tokens[8] == -3 {
		return "yes"
	}
	if b.tokens[6]+b.tokens[4]+b.tokens[2] == 3 || b.tokens[6]+b.tokens[4]+b.tokens[2] == -3 {
		return "yes"
	}
	return "no"
}

func (b *Board) print() {
	fmt.Printf("%s%s%s\n", b.get(0, 0), b.get(1, 0), b.get(2, 0))
	fmt.Printf("%s%s%s\n", b.get(0, 1), b.get(1, 1), b.get(2, 1))
	fmt.Printf("%s%s%s\n", b.get(0, 2), b.get(1, 2), b.get(2, 2))
}

func (b *Board) scanput() {
	var x, y int
	var s string
	fmt.Scanf("%d %d %s", &x, &y, &s)
	b.put(x, y, s)
}

// 毎ターン終了後、現在操作しているプレイヤーを切り替える
func switchPlayer() {
	if currentPlayer == player1 {
		currentPlayer = player2
	} else {
		currentPlayer = player1
	}
}

func getPlayerNumber(player string) int {
	if player == player1 {
		return 1
	} else {
		return 2
	}
}

// 入力が正当かどうかを判断する
// x、yの入力値は0～2
// この位置が"-"（未入力状態）でない場合、入力は無効である。
func validMove(x, y int) bool {
	if x < 0 || x >= 3 || y < 0 || y >= 3 {
		return false
	}
	if board.get(x, y) != "-" {
		return false
	}
	return true
}

const (
	player1 = "o"
	player2 = "x"
)

var board Board
var currentPlayer string

// ゲーム本体の実装
func main() {
	currentPlayer = player1
	board.tokens = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for {
		fmt.Printf("Player %d: Input (x,y) ", getPlayerNumber(currentPlayer))
		var x, y int
		fmt.Scanf("%d,%d\n", &x, &y)

		if !validMove(x, y) {
			fmt.Println("Invalid move, try again.")
			continue
		}
		board.put(x, y, currentPlayer)

		if board.check() == "yes" {

			fmt.Printf("Player %d won\n", getPlayerNumber(currentPlayer))
			break
		}
		board.print()
		switchPlayer()
	}
}
