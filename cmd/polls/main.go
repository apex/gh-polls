package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tj/kingpin"

	"github.com/apex/gh-polls/internal/cli"
	"github.com/apex/gh-polls/internal/client"
)

// Config.
var (
	version  = "master"
	endpoint = "https://api.gh-polls.com"
)

func main() {
	app := kingpin.New("polls", "GitHub polls.")
	app.Version(version)

	create := app.Command("new", "Create a new poll.")
	create.Example(`polls new Tobi Loki Jane`, "Create a new poll for who is the best ferret.")
	create.Example(`polls new "Cats are better" "Ferrets are better"`, "Create a new poll with larger options.")
	options := create.Arg("options", "Poll options.").Required().Strings()

	versionCommand := app.Command("version", "Output program version.")
	versionCommand.Example(`polls version`, "Show the version :).")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case versionCommand.FullCommand():
		fmt.Println(version)
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

		if err := cli.CopyToClipboard(out.ID, *options); err != nil {
			log.Fatalf("error copying to clipboard: %s", err)
		}

		fmt.Printf("Copied markdown for poll %s to the clipboard!\n", out.ID)
	}
}
