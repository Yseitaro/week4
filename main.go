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
