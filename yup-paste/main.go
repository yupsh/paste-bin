package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/paste"
)

const (
	flagDelimiter = "delimiters"
	flagSerial    = "serial"
	flagZero      = "zero-terminated"
)

func main() {
	app := &cli.App{
		Name:  "paste",
		Usage: "merge lines of files",
		UsageText: `paste [OPTIONS] [FILE...]

   Write lines consisting of the sequentially corresponding lines from
   each FILE, separated by TABs, to standard output.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    flagDelimiter,
				Aliases: []string{"d"},
				Usage:   "use characters from LIST instead of TABs",
			},
			&cli.BoolFlag{
				Name:    flagSerial,
				Aliases: []string{"s"},
				Usage:   "paste one file at a time instead of in parallel",
			},
			&cli.BoolFlag{
				Name:    flagZero,
				Aliases: []string{"z"},
				Usage:   "line delimiter is NUL, not newline",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "paste: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// Add file arguments (or none for stdin)
	for i := 0; i < c.NArg(); i++ {
		params = append(params, gloo.File(c.Args().Get(i)))
	}

	// Add flags based on CLI options
	if c.IsSet(flagDelimiter) {
		params = append(params, Delimiter(c.String(flagDelimiter)))
	}
	if c.Bool(flagSerial) {
		params = append(params, Serial)
	}
	if c.Bool(flagZero) {
		params = append(params, Zero)
	}

	// Create and execute the paste command
	cmd := Paste(params...)
	return gloo.Run(cmd)
}
