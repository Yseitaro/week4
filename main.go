package main

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
	return "*"
}

func (b *Board) check() string {
	var token string
	token = b.get(0, 0)
	if token == b.get(1, 0) && token == b.get(2, 0) {
		return "yes"
	}
	return "no"
}
