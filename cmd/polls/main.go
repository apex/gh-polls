package main

import (
	"log"
	"os"

	"github.com/tj/kingpin"

	"github.com/tj/gh-polls/internal/cli"
	"github.com/tj/gh-polls/internal/client"
)

// Config.
var (
	version  = "master"
	endpoint = "https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod"
)

func main() {
	app := kingpin.New("polls", "GitHub polls.")
	app.Version(version)

	create := app.Command("new", "Create a new poll.")
	options := create.Arg("options", "Poll options.").Required().Strings()
	create.Example(`polls new Tobi Loki Jane`, "Create a new poll for who is the best ferret.")
	create.Example(`polls new "Cats are better" "Ferrets are better"`, "Create a new poll with larger options.")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case create.FullCommand():
		polls := client.Client{
			Endpoint: endpoint,
		}

		out, err := polls.Create(&client.CreateInput{
			Options: *options,
		})

		if err != nil {
			log.Fatalf("error creating poll: %s", err)
		}

		cli.OutputOptions(out.ID, *options)
	}
}
