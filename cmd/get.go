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
	"github.com/sysbind/pgtrunk/pkg/pgpool"
)

// getCmd represents the get command
var getCmd = &cobra.Command {
	Use:   "get",
	Short: "Gets a postgres master",
	Long: `Consults pgpool (using pcp commands) for the master`,
	Run: func(cmd *cobra.Command, args []string) {
		pcp := pgpool.PCPConnection("127.0.0.1", 9898, "root", "Password1")
		primary := pgpool.GetPrimaryNode(pcp)
		fmt.Println("Primary: ", primary.Host, primary.Port)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
