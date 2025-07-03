package cmd

import (
    "github.com/spf13/cobra"
    "github.com/bubu-gigi/wolf/modules/scan"
)

var target string

var scanCmd = &cobra.Command{
    Use:   "scan",
    Short: "Scansione HTTP per testare header bypass 403",
    Run: func(cmd *cobra.Command, args []string) {
        scan.Run(target)
    },
}

func init() {
    scanCmd.Flags().StringVarP(&target, "target", "t", "", "Target URL (es: https://sito.com/admin)")
    scanCmd.MarkFlagRequired("target")
}
