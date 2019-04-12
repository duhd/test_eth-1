package utils

import (
   "testing"
   "fmt"
)

func TestConfig(*testing.T) {
   config_file := "../config.yaml"
   cfg = LoadConfig(config_file)
   fmt.Print(cfg)
   // if total != 10 {
   //    t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
   // }
}
