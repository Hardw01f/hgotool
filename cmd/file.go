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

	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "file subcommand handles file parts",
	Example: `hgotool file detail [FILE_PATH]
hgotool file monitor [FILE_PATH]`,
	Long: `file subcommand handles file parts.
can show target file simple infomation and 
monitor target file whatever changes in target file 
and If has changes, Send notification to User`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println("file called...\n")
		flag.Parse()
		args = flag.Args()

		if args[1] == "detail" {

			Filename := args[2]

			Name, Size, Permission, err := GetFileDetail(Filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Name : %s    Size : %s    Mod : %s\n", Name, Size, Permission)
		} else if args[1] == "monitor" {
			Filename := args[2]
			_, CurrentSize, CurrentPermission, err := GetFileDetail(Filename)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			CurrentSizeInt, _ := strconv.Atoi(CurrentSize)
			CurrentPermissionInt, _ := strconv.Atoi(CurrentPermission)

			for {
				_, Size, Permission, err := GetFileDetail(Filename)
				if err != nil {
					fmt.Println("Changing File Name or Deleteing File !!")
					os.Exit(1)
				}
				SizeInt, _ := strconv.Atoi(Size)
				PermissionInt, _ := strconv.Atoi(Permission)

				if CurrentSizeInt != SizeInt {
					fmt.Println("Changing File Size")
					os.Exit(1)
				} else if CurrentPermissionInt != PermissionInt {
					fmt.Println("Changing File Permission")
					os.Exit(1)
				} else {
					time.Sleep(1 * time.Second)
				}

			}

		} else {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	fileCmd.PersistentFlags().String("detail", "", "show file FileName,FileSize,FilePermissionCode. show Exsample")
	fileCmd.PersistentFlags().String("monitor", "", "monitor Whether there is no arbitrary change on target File")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//fileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func GetFileDetail(Filename string) (string, string, string, error) {
	Res, err := os.Stat(Filename)
	if err != nil {
		//fmt.Println(" [Error] : target file not exist")
		Funcerr := fmt.Errorf("[Error] : %s", "target file not exist")
		return "", "", "", Funcerr
	}
	ResStr := fmt.Sprintf("%+v", Res)
	//fmt.Println(ResStr)
	SplitedRes := strings.Split(ResStr, " ")

	TrimName := strings.Split(SplitedRes[0], ":")
	Name := TrimName[1]

	TrimSize := strings.Split(SplitedRes[1], ":")
	Size := TrimSize[1]

	TrimPermission := strings.Split(SplitedRes[2], ":")
	Permission := TrimPermission[1]
	return Name, Size, Permission, nil
}
