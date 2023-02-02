package actions

import (
	"fmt"
	"git-uwu/utils"
	"git-uwu/uwuify"
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

func Commit(ctx *cli.Context) error {
	_, worktree, err := utils.GetWorktree()
	if err != nil {
		return err
	}
	message := ctx.String("message")
	if message != "" {
		message = uwuify.From(message)
	}

	hash, err := worktree.Commit(message, &git.CommitOptions{All: true})
	if err != nil {
		return err
	}
	fmt.Println(uwuify.From("Committed all the files with hash "), hash.String())
	return nil
}
