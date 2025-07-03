package scan

import (
    "fmt"
    "github.com/fatih/color"
)

func Run(target string) {
    color.Green("[SCAN] Target selezionato: %s", target)
    fmt.Println("🚧 Qui andrà la logica di scan, header fuzzing, 403 bypass...")
}
