package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestPutToken01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 1, "o")
	if b.get(1, 1) != "o" {
		t.Errorf("Mistake")
	}
}

func TestPutToken02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 2, "x")
	if b.get(1, 2) != "x" {
		t.Errorf("Mistake")
	}
}

// yoko
func TestCheck01(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "o")
	b.put(1, 0, "o")
	b.put(2, 0, "o")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}

// tate
func TestCheck02(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(1, 0, "x")
	b.put(1, 1, "x")
	b.put(1, 2, "x")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}

// nashi
func TestCheck03(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 0, "x")
	b.put(0, 1, "o")
	b.put(0, 2, "x")
	if b.check() != "no" {
		t.Errorf("Mistake")
	}
}

// naname
func TestCheck04(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 2, "x")
	b.put(1, 1, "x")
	b.put(2, 0, "x")
	if b.check() != "yes" {
		t.Errorf("Mistake")
	}
}

func TestPutToken05(t *testing.T) {
	b := &Board{
		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	b.put(0, 2, "x")
	b.put(1, 1, "x")
	b.put(2, 0, "x")
	b.print()
}

func TestMain(t *testing.T) {
	// シミュレーション入力の作成:
	// Player1(0,0) -> Player2(1,0) -> Player1(0,1) -> Player2(1,1) -> Player1(0,2)
	input := "0,0\n1,0\n0,1\n1,1\n0,2\n"
	tempIn, _ := ioutil.TempFile("", "stdin")
	tempIn.Write([]byte(input))
	tempIn.Seek(0, 0)
	os.Stdin = tempIn

	// シミュレーション出力の作成
	tempOut, _ := ioutil.TempFile("", "stdout")
	os.Stdout = tempOut

	main()

	// 出力チェック
	tempOut.Seek(0, 0)
	outputBytes, _ := ioutil.ReadAll(tempOut)
	output := string(outputBytes)
	if !strings.Contains(output, "Player 1 won") {
		t.Errorf("Expected 'Player 1 won', got '%s'", output)
	}

	// クリーンアップ
	os.Remove(tempIn.Name())
	os.Remove(tempOut.Name())
}
