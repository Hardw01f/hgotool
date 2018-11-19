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
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

- s
- s
- s

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
		fmt.Println("test called")

		if args[0] == "Send" {
			SendTest("Test", "This is Test")

		} else {
			os.Exit(1)
		}

	},
}

type Config struct {
	Detail ServerDetail
}

type ServerDetail struct {
	Slack string
}

var config Config

func SendTest(Name string, Text string) {

	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		os.Exit(1)
	}

	//fmt.Println(config.Detail.Slack)

	channel := "alert"

	jsonStr := `{"channel":"` + channel + `","username":"` + Name + `","text":"` + Text + `"}`

	req, err := http.NewRequest(
		"POST",
		config.Detail.Slack,
		bytes.NewBuffer([]byte(jsonStr)),
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(resp)
	defer resp.Body.Close()
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
