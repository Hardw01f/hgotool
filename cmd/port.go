// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/anvie/port-scanner"
	"github.com/spf13/cobra"
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:     "port",
	Short:   "port subcommand handle port parts",
	Example: "hgotool port scan [Scan_StartPort] [Scan_EndPort]",
	Long: `[help] : The port subcommand handle port parts , 
can do Scan and detect Using target Port Service`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		flag.Parse()
		args = flag.Args()
		fmt.Println("port called")

		StartPort, _ := strconv.Atoi(args[2])
		EndPort, _ := strconv.Atoi(args[3])

		if args[1] == "scan" {
			ps := portscanner.NewPortScanner("localhost", 3*time.Second, 10)

			fmt.Printf("scanning port %s-%s...\n", args[2], args[3])

			openedPorts := ps.GetOpenedPort(StartPort, EndPort)

			for i := 0; i < len(openedPorts); i++ {
				port := openedPorts[i]
				fmt.Print(" ", port, " [open]")
				fmt.Println("  -->  ", ps.DescribePort(port))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
