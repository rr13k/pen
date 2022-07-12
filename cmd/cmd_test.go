package cmd

import (
	"fmt"
	"testing"
)

func Test_App(t *testing.T) {
	cmdApp.Run(nil, []string{"nihao"})
}

func Test_nihao(t *testing.T) {
	fmt.Println("niaho")
}
