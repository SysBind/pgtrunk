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
	"os"
	"log"	
	"path"
	"strconv"
	"golang.org/x/sys/unix"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sysbind/pgtrunk/pkg/pgpool"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pgtrunk",
	Short: "A Postgres / Pgpool-II cluster wire-up",
	Long: `A Postgres / Pgpool-II cluster wire-up,`,
	Run: func(cmd *cobra.Command, args []string) {
		executable := viper.GetString("executable")
		datadir := viper.GetString("datadir")
		host := viper.GetString("hostname")
		port := viper.GetInt("port")
		initdb := viper.GetBool("initdb")

		primary := pgpool.GetPrimaryNode()

		if primary.Port == port && primary.Host == host {			
			fmt.Println("primary is me")
			if (initdb) {
				pgpool.InitDB(datadir)
			}
		} else {
			fmt.Println("synching from primary")
			pgpool.Sync(primary, datadir)
		}

		fmt.Println("all done, launching ", executable)
		argv := append([]string{path.Base(executable)},
			"-D", datadir,
			"-p", strconv.Itoa(port),
			"-h", host)

		// use exec to completly replace the current process with postgres:
		argv = append(argv, args...)
		err := unix.Exec(executable, argv, os.Environ())
		log.Fatal(err)	
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pgtrunk.yaml)")

	rootCmd.Flags().String("executable", "/usr/bin/postgres" ,"executable or script for launching postgres")
	rootCmd.Flags().StringP("datadir", "D", "/var/lib/postgres/data", "data directory, passed on to postgres")
	rootCmd.Flags().IntP("port", "p", 5432, "port as refferd to by pgpool")
	rootCmd.Flags().StringP("hostname", "H", "localhost", "hostname as reffered to by pgpool")
	rootCmd.Flags().Bool("initdb", false, "run initdb if datadir empty")
	viper.BindPFlag("executable", rootCmd.Flags().Lookup("executable"))
	viper.BindPFlag("datadir", rootCmd.Flags().Lookup("datadir"))
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("hostname", rootCmd.Flags().Lookup("hostname"))
	viper.BindPFlag("initdb", rootCmd.Flags().Lookup("initdb"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".pgtrunk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pgtrunk")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
