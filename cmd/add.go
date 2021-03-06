/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of add sub command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using the add sub command`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the flag value, its default value is false
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus { // if status is true, call addFloat
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
}

func addInt(args []string) {
	var sum int

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}

	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func addFloat(args []string) {
	var sum float64

	for _, arg := range args {
		num, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum += num
	}

	fmt.Printf("Addition of floating numbers %s is %f\n", args, sum)
}
