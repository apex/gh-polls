package poll

import (
	"bytes"
	"html/template"

	"github.com/pkg/errors"
)

// font family.
var fontFamily = `-apple-system,BlinkMacSystemFont,Segoe UI,Roboto,Oxygen,Ubuntu,Cantarell,Fira Sans,Droid Sans,Helvetica Neue,sans-serif`

// option svg.
var option = `<svg width="448px" height="58px" viewBox="0 0 448 58" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
    <g id="Page-1" stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
        <g id="poll">
            <g id="Group" transform="translate(17.000000, 10.000000)">
                <rect id="Rectangle" fill="#F1F3F5" x="0" y="19" width="334" height="14" rx="2"></rect>
                <rect id="Rectangle" fill="#7950F2" x="0" y="19" width="0" height="14" rx="2">
                  <animate attributeName="width"
                      begin="0.5s"
                      dur="600ms"
                      from="0"
                      to="{{.Width}}"
                      repeatCount="1"
                      fill="freeze"
                      calcMode="spline"
                      keyTimes="0; 1"
                      keySplines="0.3, 0.61, 0.355, 1.1" />
                </rect>
                <text id="100%" font-family="{{.FontFamily}}" font-size="12" font-weight="normal" letter-spacing="1.857333" fill="#212529">
                    <tspan x="344" y="30">{{.Percent}}%</tspan>
                </text>
                <text id="Option-A" font-family="{{.FontFamily}}" font-size="12" font-weight="bold" letter-spacing="1" fill="#212529">
                    <tspan x="0" y="12">{{.Name}}</tspan>
                </text>
                <text id="150" font-family="{{.FontFamily}}" font-size="12" font-weight="normal" letter-spacing="1" fill="#868E96">
                    {{if .Votes}}
                      <tspan x="386" y="30">{{.Votes}}</tspan>
                    {{end}}
                </text>
            </g>
        </g>
    </g>
</svg>
`

// Option represents a single poll option.
type Option struct {
	Name    string
	Votes   int
	Percent int

	Width      int
	FontFamily string
}

// Render option as svg.
func (o *Option) Render() ([]byte, error) {
	o.FontFamily = fontFamily

	tmpl, err := template.New("poll").Parse(option)
	if err != nil {
		return nil, errors.Wrap(err, "parsing")
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, o); err != nil {
		return nil, errors.Wrap(err, "executing")
	}

	return buf.Bytes(), nil
}
