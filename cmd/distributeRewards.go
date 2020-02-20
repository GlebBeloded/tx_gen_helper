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
	sdk "inside.omertex.com/bitbucket/scm/mf/blockchain_mediafm.git/types"
	"inside.omertex.com/txgen/codec"
	"inside.omertex.com/txgen/session"
	"inside.omertex.com/txgen/stdTx"
)

// distributeRewardsCmd represents the distributeRewards command
var distributeRewardsCmd = &cobra.Command{
	Use:   "distribute-rewards [payout] [integration_id...]",
	Short: "create distribute-rewards tx as well as generate and store ad_bytes",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		gas_val, _ := cmd.Flags().GetInt("gas")
		stdTx.GasValue = int(gas_val)
		ads := args[1:]
		payout, err := sdk.ParseCoin(args[0])
		if err != nil {
			panic(err.Error())
		}
		msg := session.NewMsgDistributeRewards(payout, ads)

		if len(msg.Ads) == 0 {
			msg.Ads = append(msg.Ads, session.MsgIntegrationData{IntegrationID: "plug", AdBytes: ""})
		}

		tx := stdTx.NewTx(msg)
		bytes, err := codec.Codec.MarshalJSONIndent(tx, "", "\t")
		if err != nil {
			panic(err)
		}
		fmt.Print(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(distributeRewardsCmd)
	distributeRewardsCmd.PersistentFlags().Int("gas", 200000, "set custom gas value for tx")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// distributeRewardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// distributeRewardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
