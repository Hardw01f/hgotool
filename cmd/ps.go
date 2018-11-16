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
	"strings"
	"time"

	"github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "ps subcommand handles process parts",
	Example: `hgotool ps show
	hgotool ps search [PID]
	hgotool ps monitor [PID]`,
	Long: `ps subcommand handles process parts
can SHOW current running process list like ps shell command
and SEARCH current runnnig process by process PID
MONITOR whether current running process was killed If had killed , Send notification to User`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		flag.Parse()
		args = flag.Args()
		//fmt.Println(args[1], args[2])

		if args[1] == "show" {
			processes, err := ps.Processes()
			if err != nil {
				os.Exit(1)
			}

			for _, process := range processes {
				resstr := fmt.Sprintf("%v", process)
				trimstr := strings.Split(resstr, "{")
				trimedstr := strings.Split(trimstr[1], "}")
				res := trimedstr[0]
				//fmt.Println(res)

				splitRes := strings.Split(res, " ")
				fmt.Printf("PID : %s   PPID : %s   NAME :        %s \n", splitRes[0], splitRes[1], splitRes[2])

				//fmt.Printf("%T type  :  ", process)
				//fmt.Printf("%v\n", process)
				//test := fmt.Sprintf("%v", process)
				//fmt.Println(test)
				//fmt.Printf("type is : %T", test)
				//fmt.Println(strings.Split(test, "{"))
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
	psCmd.PersistentFlags().String("ps", "", "show current running processes,show Example")
	psCmd.PersistentFlags().String("search", "", "search specific running process,show Example")
	psCmd.PersistentFlags().String("monitor", "", "monitor specific running processes,show Example")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// psCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
