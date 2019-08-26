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
	"github.com/spf13/cobra"
	"github.com/wejustdostuff/carnot/pkg/crawler"
)

// moveCmd represents the crawl command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Move all files from the source directory to the target directory.",
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

		force, _ := cmd.Flags().GetBool("force")

		// Retrieve files
		files, err := crawler.GetFiles(source)
		if err != nil {
			exit(cmd, "error: could not retrieve files: %s\n", err.Error())
		}

		// Iterate files
		for _, file := range files {
			if !force && file.Exists(target) {
				printWarning("%s -> %s, already exists", file.Path, file.GetPath(target))
				continue
			}

			if err := file.Move(target); err != nil {
				printWarning("warning: could not move file %s (%s)", file.Path, err.Error())
			} else {
				printInfo("%s -> %s", file.Path, file.GetPath(target))
			}
		}
		printSuccess("Done")
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)

	moveCmd.Flags().StringP("source", "s", "", "Specify the source directory to move the files from")
	moveCmd.Flags().StringP("target", "t", "", "Specify the target directory to move the files to")

	moveCmd.Flags().Bool("force", false, "Force files to be moved ignoring existing files")
}
