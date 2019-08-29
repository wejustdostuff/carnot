/*
Copyright © 2019 We Just Do Stuff <hello@wejustdostuff.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/table"
	"github.com/wejustdostuff/carnot/pkg/crawler"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all files in the given source directory.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Get flags
		source, _ := cmd.Flags().GetString("source")
		if source == "" {
			exit(cmd, "error: source cannot be empty\n")
		}

		target, _ := cmd.Flags().GetString("target")
		if target == "" {
			exit(cmd, "error: target cannot be empty\n")
		}

		// Retrieve files
		files, err := crawler.GetFiles(source)
		if err != nil {
			exit(cmd, "error: could not retrieve files: %s\n", err.Error())
		}

		// Initialization
		t := table.NewWriter()
		t.Style().Options.DrawBorder = false
		t.Style().Options.SeparateColumns = false
		t.Style().Options.SeparateFooter = false
		t.Style().Options.SeparateHeader = false
		t.Style().Options.SeparateRows = false

		t.AppendHeader(table.Row{"Source", "", "Target", "Date Field", "Date", "Time"})

		for _, file := range files {
			t.AppendRow(table.Row{file.Path, "->", file.GetPath(target), file.DateField, file.Date.Format("2006.01.02"), file.Date.Format("15:04:05")})
		}

		fmt.Println(t.Render())
		fmt.Println()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("source", "s", "", "Specify the source directory to move the files from")
	listCmd.Flags().StringP("target", "t", "", "Specify the target directory to move the files to")
}
