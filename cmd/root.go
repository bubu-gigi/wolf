package cmd

import (
    "github.com/spf13/cobra"
    "github.com/bubu-gigi/wolf/core"
)

var rootCmd = &cobra.Command{
    Use:   "wolf",
    Short: "WOLF - Offensive Security Framework",
    Long:  `Wolf è un framework modulare per offensive security in Go.`,
}

// Entry point CLI
func Execute() {
    core.Banner()
    cobra.CheckErr(rootCmd.Execute())
}

func init() {
    rootCmd.AddCommand(scanCmd)
}
