package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/atotto/clipboard"

	"github.com/tj/kingpin"
)

// Config.
var (
	version  = "master"
	endpoint = "https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod"
)

// TODO: move all this stuff into a pkg
type input struct {
	Options []string `json:"options"`
}

type output struct {
	ID string `json:"id"`
}

func main() {
	app := kingpin.New("polls", "GitHub polls.")
	app.Version(version)

	create := app.Command("new", "Create a new poll.")
	options := create.Arg("options", "Poll options.").Required().Strings()
	create.Example(`polls new Tobi Loki Jane`, "Create a new poll for who is the best ferret.")
	create.Example(`polls new "Cats are better" "Ferrets are better"`, "Create a new poll with larger options.")

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case create.FullCommand():
		// TODO: move all this stuff into a pkg

		b, err := json.Marshal(input{Options: *options})
		if err != nil {
			log.Fatalf("error marshaling: %s", err)
		}

		res, err := http.Post(endpoint+"/poll", "application/json", bytes.NewReader(b))
		if err != nil {
			log.Fatalf("error requesting: %s", err)
		}
		defer res.Body.Close()

		var out output
		if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
			log.Fatalf("error unmarshaling resonse: %s", err)
		}

		var buf bytes.Buffer

		for _, o := range *options {
			fmt.Fprintf(&buf, link(out.ID, o))
		}

		if err := clipboard.WriteAll(buf.String()); err == nil {
			fmt.Fprintln(os.Stderr, "Copied to clipboard!")
		}

		io.Copy(os.Stdout, &buf)
	}
}

func link(id, option string) string {
	return fmt.Sprintf(`[%s](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s/vote)`, image(id, option), id, url.PathEscape(option))
}

func image(id, option string) string {
	return fmt.Sprintf(`![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s)`, id, url.PathEscape(option))
}
