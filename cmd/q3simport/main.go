// The MIT License (MIT)

// Copyright (c) 2016 Maciej Borzecki

// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:

// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package main

import (
	"fmt"
	"github.com/bboozzoo/q3stats/version"
	"github.com/codegangsta/cli"
	"os"
)

func doImport(c *cli.Context) error {
	if c.Args().Present() == false {
		cli.ShowAppHelp(c)
		os.Exit(1)
	}

	srcfile := c.Args().First()
	mh, err := runImport(srcfile, c.String("host"))
	if err != nil {
		return err
	}

	fmt.Println(mh)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Usage = "Q3 Match Statistics import tool"
	app.Version = version.GetVersion()
	app.ArgsUsage = "<match-file>"
	app.HideHelp = true
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host, t",
			Usage: "Q3 Stats server host:port address",
			Value: "localhost:9090",
		},
	}
	app.Action = doImport
	app.Commands = []cli.Command{}

	app.Run(os.Args)
}
