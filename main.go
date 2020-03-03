package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/chzyer/readline"
	"github.com/notnil/chess"
)

type ChessAI struct {
	game *chess.Game
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ai := NewChessAI()
	if err := ai.repl(); err != nil {
		panic(err)
	}
}

func NewChessAI() *ChessAI {
	return &ChessAI{
		game: chess.NewGame(),
	}
}

func (ai *ChessAI) Reset() {
	ai.game = chess.NewGame()
}

func (ai *ChessAI) DoRandomMove() error {
	moves := ai.game.ValidMoves()
	if len(moves) == 0 {
		return fmt.Errorf("no moves")
	}
	return ai.game.Move(moves[rand.Intn(len(moves))])
}

func (ai *ChessAI) DoUserInput(moveStr string) error {
	return ai.game.MoveStr(moveStr)
}

func (ai *ChessAI) repl() error {
	rl, err := readline.New("white to move> ")
	if err != nil {
		return err
	}

	fmt.Println(ai.game.Position().Board().Draw())

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}

		if err := ai.DoUserInput(line); err != nil {
			fmt.Println(err.Error())
			continue
		}

		if ai.game.Outcome() != chess.NoOutcome {
			fmt.Println(ai.game.Outcome())
			break
		}

		fmt.Println(ai.game.Position().Board().Draw())

		if err := ai.DoRandomMove(); err != nil {
			return err
		}

		if ai.game.Outcome() != chess.NoOutcome {
			fmt.Println(ai.game.Outcome())
			break
		}

		fmt.Println(ai.game.Position().Board().Draw())
		fmt.Println(ai.game.Outcome().String())
	}

	return nil
}
