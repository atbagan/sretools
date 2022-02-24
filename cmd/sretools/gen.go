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
	"github.com/spf13/cobra"
)

// genCmd represents the cfn command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate various useful things for sretools",
	Long:  `Generate documentation, CLI completions, and IAM policies`,
}

func init() {
	rootCmd.AddCommand(genCmd)
}
