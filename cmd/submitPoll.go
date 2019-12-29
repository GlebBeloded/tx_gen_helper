/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"inside.omertex.com/txgen/codec"
	"inside.omertex.com/txgen/poll"
	"inside.omertex.com/txgen/stdTx"
)

// submitPollCmd represents the submitPoll command
var submitPollCmd = &cobra.Command{
	Use:   "submit-poll [poll-id] [submission_timestamp] [address]",
	Short: "create sumbitPoll tx if given PollId Exists, else return error",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		pollSubmission := poll.NewMsgPollSubmission(args[0], args[1], args[2])
		tx := stdTx.NewTx(pollSubmission)
		bytes, err := codec.Codec.MarshalJSONIndent(tx, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(submitPollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitPollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitPollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
