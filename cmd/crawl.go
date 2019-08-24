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
	"os"

	"github.com/spf13/cobra"
	"github.com/wejustdostuff/carnot/pkg/crawler"
)

// crawlCmd represents the crawl command
var crawlCmd = &cobra.Command{
	Use:   "crawl",
	Short: "Crawl the source directory and organize all files to the target directory.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		if source == "" {
			fmt.Fprintf(os.Stderr, "error: source cannot be empty\n")
			cmd.Usage()
			os.Exit(1)
		}

		target, _ := cmd.Flags().GetString("target")
		if target == "" {
			fmt.Fprintf(os.Stderr, "error: target cannot be empty\n")
			cmd.Usage()
			os.Exit(1)
		}

		dry, _ := cmd.Flags().GetBool("dry")
		move, _ := cmd.Flags().GetBool("move")
		force, _ := cmd.Flags().GetBool("force")

		if err := crawler.Run(source, target, dry, move, force); err != nil {
			exit(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(crawlCmd)

	crawlCmd.Flags().StringP("source", "s", "", "Specify the source directory to move the files from")
	crawlCmd.Flags().StringP("target", "t", "", "Specify the target directory to move the files to")

	crawlCmd.Flags().Bool("dry", false, "Run in dry mode")
	crawlCmd.Flags().Bool("move", false, "Copy or move files")
	crawlCmd.Flags().Bool("force", false, "Force files to be moved ignoring existing files")
}
