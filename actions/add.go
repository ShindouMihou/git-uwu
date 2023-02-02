package actions

import (
	"git-uwu/utils"
	"git-uwu/uwuify"
	"github.com/urfave/cli/v2"
	"log"
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
	log.Println(uwuify.From("added all files in " + path))
	return nil
}
