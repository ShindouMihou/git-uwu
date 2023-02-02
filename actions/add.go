package actions

import (
	"fmt"
	"git-uwu/utils"
	"git-uwu/uwuify"
	"github.com/urfave/cli/v2"
)

func Add(ctx *cli.Context) error {
	_, worktree, err := utils.GetWorktree()
	if err != nil {
		return err
	}
	path := ctx.String("path")
	err = worktree.AddGlob(path)
	if err != nil {
		return err
	}
	fmt.Println(uwuify.From("added all files in " + path))
	return nil
}
