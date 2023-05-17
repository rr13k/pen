package main

import (
	"fmt"
	"testing"

	"github.com/rr13k/pen/cmd"
	"github.com/rr13k/pen/common"
)

func Test_main(t *testing.T) {
	fmt.Println(EmbedTempsContent)
	common.EmbedTempsContent = EmbedTempsContent
	cmd.GenerationModel("/Users/zhouyuan11/work/pen/pen-test/internal/app/models/user.go")
}
