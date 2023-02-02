package actions

import (
	"fmt"
	"git-uwu/uwuify"
	"github.com/urfave/cli/v2"
	"os"
)

func Magic(ctx *cli.Context) error {
	file := ctx.Path("file")
	message := ctx.String("message")

	if message != "" {
		fmt.Println(uwuify.From(message))
	}

	if file != "" {
		bts, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		contents := uwuify.From(string(bts))
		err = os.WriteFile(file+".uwu", []byte(contents), 0666)
		if err != nil {
			return err
		}
		fmt.Println("created file", file+".uwu")
	}
	return nil
}
