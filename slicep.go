package sliceP

import "github.com/clipperhouse/typewriter"

var sliceP = &typewriter.Template{
	Name: "sliceP",
	Text: `// {{.SliceName}} is a slice of type {{.Type}}. Use it where you would use []*{{.Type}}.
type {{.SliceName}} []*{{.Type}}
`,
}
