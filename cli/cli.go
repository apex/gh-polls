package cli

import (
	"fmt"
	"net/url"
)

// Link returns a poll option link with image.
func Link(id, option string) string {
	return fmt.Sprintf(`[%s](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s/vote)`, Image(id, option), id, url.PathEscape(option))
}

// Image returns a poll option image.
func Image(id, option string) string {
	return fmt.Sprintf(`![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/%s/%s)`, id, url.PathEscape(option))
}
