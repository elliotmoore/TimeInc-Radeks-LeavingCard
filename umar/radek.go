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

func (r *Radek) dmARadek() (string, error) {
	return fmt.Sprintf("You're having problems everyone else is having? Instead of sharing with the channel,I WILL DM YOU BECAUSE I AM A %s", emoji.Sprint(":shit:")), nil
}

func (r *Radek) annoyARadek() (string, error) {

	annoyances := []string{
		"git push -f origin master",
		"why people use forks on day2day basis in projects they have write access to, I really don't know",
		"Was it someone’s evil plan to keep people away from kitchen to reduce the noise?",
		"morning from BFB, where it’s also cold",
		"direct link silently disappeared from the intranet and nobody seems to be updating the BFB kitchen menu anymore",
		"be able to read from their DB directly_ :thinking_face: how about no",
	}

	rand.Seed(time.Now().Unix())

	return annoyances[rand.Int()%len(annoyances)], nil
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
			Name: "dm-a-radek",
			Action: func(c *cli.Context) error {
				output, _ := radek.dmARadek()
				fmt.Printf("%s\n", output)

				return nil
			},
		},
		{
			Name: "annoy-a-radek",
			Action: func(c *cli.Context) error {
				output, _ := radek.annoyARadek()
				fmt.Printf("%s\n", output)

				return nil
			},
		},
		{
			Name: "test-a-radek",
			Action: func(c *cli.Context) error {
				fmt.Println("I always return true because I don't write proper tests.")

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "You wrote dirty code and made an error"+err.Error())
	}
}
