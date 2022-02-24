// Package sretools
// Copyright (c) 2022, Andrew Bagan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package sretools

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// docsCmd represents the cfn command
var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate documentation for sretools",
	Long: `Generate documentation for sretools in Markdown format
This is used for the documentation in the repository, but can be run separately. By default it will generate it in the docs directory from where you run the command, but you can override this with the --directory flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := doc.GenMarkdownTree(rootCmd, docsdir)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var docsdir string

func init() {
	genCmd.AddCommand(docsCmd)
	docsCmd.Flags().StringVarP(&docsdir, "directory", "d", "./docs", "The directory where the documentation will be generated")

}
