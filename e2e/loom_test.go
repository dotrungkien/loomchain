package main

import (
	"fmt"
	"github.com/loomnetwork/loomchain/e2e/common"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestE2eEvm(t *testing.T) {
	tests := []struct {
		testFile string
		n        int
		genFile  string
	}{
		{"loom-test.toml", 4, ""},
	}
	common.LoomPath = "../loom"
	common.ContractDir = "../contracts"

	for _, test := range tests {
		*common.Validators = test.n
		config, err := common.NewConfig("evm", test.testFile, test.genFile)
		if err != nil {
			t.Fatal(err)
		}

		binary, err := exec.LookPath("go")
		if err != nil {
			t.Fatal(err)
		}
		// required binary
		cmd := exec.Cmd{
			Dir:  config.BaseDir,
			Path: binary,
			Args: []string{
				binary,
				"build",
				"-tags",
				"evm",
				"-o",
				"loom",
				"github.com/loomnetwork/loomchain/cmd/loom",
			},
		}
		if err := cmd.Run(); err != nil {
			t.Fatal(fmt.Errorf("fail to execute command: %s\n%v", strings.Join(cmd.Args, " "), err))
		}

		if err := common.DoRun(*config); err != nil {
			t.Fatal(err)
		}

		// pause before running the next test
		time.Sleep(500 * time.Millisecond)
	}
}