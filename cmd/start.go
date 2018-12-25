// Copyright © 2018 Asaf Ohayon <asaf@sysbind.co.il>
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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts a postgres server",
	Long: `Consults pgpool (using pcp commands) for initial role before launching postgres
possibly synchronizing data before that. (e.g: from backup, from master)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start command is empty")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
