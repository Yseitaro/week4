# Week4: Tic Tac Toe Game

This is a simple implementation of the Tic Tac Toe game written in Golang, runnable on the command line.

## How to Run the Game

First, ensure you have Go installed on your system. Then, in your terminal, navigate to the directory containing the `main.go` file and run the following command to start the game:

```shell
go run main.go
```
## Game Rules

This is a two-player game where one player is represented by "o" and the other by "x". Players take turns to place their markers on a 3x3 grid. The player who first succeeds in placing three of their markers in a horizontal, vertical, or diagonal row is the winner.

## Game Interface

Upon running the game, you'll see an empty 3x3 grid. Players input the coordinates (x, y) where they want to place their marker, with x being the row and y being the column, both ranging from 0-2, then press enter to confirm your choice.

For example, if you want to place a marker at the top left corner, you would input 0,0.

## Project File Description
- `main.go`: This is the main file of the project containing the logic for starting and running the game.

- `main_test.go`: This file contains unit tests  and main function‘s test.


## Testing

You can test this program by running the following command:
```shell
go test -v
```
This will run all the tests defined in the `main_test.go` file.