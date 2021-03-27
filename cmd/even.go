package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "A brief description of add sub command even",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using add sub command even`,
	Run: func(cmd *cobra.Command, args []string) {
		var sum int
		for _, arg := range args {
			val, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(err)
			}
			if val%2 == 0 {
				sum += val
			}
		}
		fmt.Printf("The even addition of %s is %d\n", args, sum)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
