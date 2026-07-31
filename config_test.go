// Copyright (c) 2026 The EXCCoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestNormalizeAddress(t *testing.T) {
	tests := []struct {
		name                    string
		addr                    string
		testnet, simnet, wallet bool
		want                    string
	}{
		{"mainnet node", "localhost", false, false, false, "localhost:9109"},
		{"mainnet wallet", "localhost", false, false, true, "localhost:9110"},
		{"testnet node", "localhost", true, false, false, "localhost:19109"},
		{"testnet wallet", "localhost", true, false, true, "localhost:19110"},
		{"simnet node", "localhost", false, true, false, "localhost:19556"},
		{"simnet wallet", "localhost", false, true, true, "localhost:19557"},
		{"explicit port", "localhost:1234", false, false, false, "localhost:1234"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := normalizeAddress(test.addr, test.testnet, test.simnet, test.wallet)
			if got != test.want {
				t.Fatalf("got %q, want %q", got, test.want)
			}
		})
	}
}

func TestCreateDefaultConfigFile(t *testing.T) {
	testDir := t.TempDir()
	originalHome := dcrdHomeDir
	dcrdHomeDir = testDir
	t.Cleanup(func() { dcrdHomeDir = originalHome })

	source := filepath.Join(testDir, "exccd.conf")
	err := os.WriteFile(source, []byte("rpcuser=user\nrpcpass=pass\n"), 0600)
	if err != nil {
		t.Fatal(err)
	}
	destination := filepath.Join(testDir, "ctl", "exccctl.conf")
	if err := createDefaultConfigFile(destination); err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(destination)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(content), "rpcuser=user\nrpcpass=pass\n"; got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
	if runtime.GOOS != "windows" {
		info, err := os.Stat(destination)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := info.Mode().Perm(), os.FileMode(0600); got != want {
			t.Fatalf("mode is %o, want %o", got, want)
		}
	}
	if err := createDefaultConfigFile(destination); err == nil {
		t.Fatal("existing credential file was overwritten")
	}
}
