package core

import "github.com/fatih/color"

func Banner() {
    cyan := color.New(color.FgCyan).Add(color.Bold)
    cyan.Println("██╗    ██╗ ██████╗ ██╗     ███████╗")
    cyan.Println("██║    ██║██╔═══██╗██║     ██╔════╝")
    cyan.Println("██║ █╗ ██║██║   ██║██║     █████╗  ")
    cyan.Println("██║███╗██║██║   ██║██║     ██╔══╝  ")
    cyan.Println("╚███╔███╔╝╚██████╔╝███████╗███████╗")
    cyan.Println(" ╚══╝╚══╝  ╚═════╝ ╚══════╝╚══════╝")
    color.Yellow("WOLF v0.1 - Offensive Security Framework")
}
