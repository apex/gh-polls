package cli

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/atotto/clipboard"
)

// OutputOptions outputs markdown options and copies to the clipboard.
func OutputOptions(id string, options []string) {
	var buf bytes.Buffer

	for _, o := range options {
		fmt.Fprintln(&buf, Link(id, o))
	}

	if err := clipboard.WriteAll(buf.String()); err == nil {
		fmt.Fprintln(os.Stderr, "Copied to clipboard!")
	}

	io.Copy(os.Stdout, &buf)
}

// Link returns a poll option link with image.
func Link(id, option string) string {
	return fmt.Sprintf(`[%s](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s/vote)`, Image(id, option), id, url.PathEscape(option))
}

// Image returns a poll option image.
func Image(id, option string) string {
	return fmt.Sprintf(`![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s)`, id, url.PathEscape(option))
}
