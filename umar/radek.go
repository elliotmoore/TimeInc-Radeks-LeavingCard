package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kyokomi/emoji"
	"github.com/urfave/cli"
)

type Being struct {
	name    string
	species string
}

type Radek struct {
	Being     Being
	Something int
}

var radek = &Radek{Being{"Radek", "Robot"}, 123}

func (r *Radek) sayARadek() (string, error) {

	emojis := []string{
		emoji.Sprint(":tada: :elephant: :tada: :+1: :+1: :+1:"),
		emoji.Sprint("This now looks :ok_hand:"),
		emoji.Sprint("That's nice :+1:"),
		emoji.Sprint("LGTM :ship:"),
		emoji.Sprint("win-win :tada:"),
	}

	rand.Seed(time.Now().Unix())

	return emojis[rand.Int()%len(emojis)], nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Radek's Goodbye"
	app.Version = "v0.1"

	app.Commands = []cli.Command{

		{
			Name: "say-a-radek",
			Action: func(c *cli.Context) error {
				output, err := radek.sayARadek()
				if err != nil {
					return err
				}

				fmt.Printf("Things Radek has actually said on pull requests: %s\n", output)

				return nil
			},
		},
		{
			Name: "do-a-radek",
		},
		{
			Name: "eat-a-radek",
		},
		{
			Name: "drink-a-radek",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "You wrote dirty code and made an error"+err.Error())
	}
}
