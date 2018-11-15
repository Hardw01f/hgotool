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

	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "show process",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		args = flag.Args()
		//fmt.Println(args[1], args[2])

		if args[1] == "show" {
			processes, err := ps.Processes()
			if err != nil {
				os.Exit(1)
			}

			for _, process := range processes {
				fmt.Printf("%T type  :  ", process)
				fmt.Printf("%v\n", process)
				//fmt.Println(process)
				//str := fmt.Sprintf("%T", process.Pid)
				//fmt.Println(str)
			}
		} else if args[1] == "search" {
			pid, _ := strconv.Atoi(args[2])
			targetprocess, err := ps.FindProcess(pid)
			if err != nil {
				os.Exit(1)
			}

			if targetprocess != nil {
				fmt.Printf("%v", targetprocess)
			} else {
				fmt.Println("the process is not exist")
			}
		} else if args[1] == "monitor" {
			pid, _ := strconv.Atoi(args[2])
			for {
				MonitorTarget, err := ps.FindProcess(pid)
				if err != nil {
					os.Exit(1)
				}

				if MonitorTarget != nil {
					time.Sleep(5 * time.Second)
				} else if MonitorTarget == nil {
					for i := 0; i < 2; i++ {
						fmt.Println("Send")
					}
					os.Exit(1)
				} else {
					fmt.Println("error")
					os.Exit(1)
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(psCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// psCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
