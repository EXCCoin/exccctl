// Copyright (c) 2026 The EXCCoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import "testing"

func TestParseSemVer(t *testing.T) {
	tests := []struct {
		version                   string
		major, minor, patch       uint32
		preRelease, buildMetadata string
		wantErr                   bool
	}{
		{"0.0.0", 0, 0, 0, "", "", false},
		{"2.8.0-pre.1+release.local", 2, 8, 0, "pre.1", "release.local", false},
		{"02.8.0", 0, 0, 0, "", "", true},
		{"2.8", 0, 0, 0, "", "", true},
		{"2.8.0+bad_metadata", 0, 0, 0, "", "", true},
	}
	for _, test := range tests {
		t.Run(test.version, func(t *testing.T) {
			major, minor, patch, preRelease, buildMetadata, err := parseSemVer(test.version)
			if (err != nil) != test.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if test.wantErr {
				return
			}
			if major != test.major || minor != test.minor || patch != test.patch ||
				preRelease != test.preRelease || buildMetadata != test.buildMetadata {
				t.Fatalf("got %d.%d.%d-%s+%s", major, minor, patch, preRelease, buildMetadata)
			}
		})
	}
}
