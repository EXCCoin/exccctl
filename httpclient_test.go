// Copyright (c) 2026 The EXCCoin developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendPostRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method is %q, want POST", r.Method)
		}
		user, pass, ok := r.BasicAuth()
		if !ok || user != "user" || pass != "pass" {
			t.Errorf("unexpected basic authentication")
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
		}
		if string(body) != `{"jsonrpc":"1.0"}` {
			t.Errorf("body is %q", body)
		}
		_, _ = io.WriteString(w, `{"result":{"ok":true},"error":null,"id":1}`)
	}))
	defer server.Close()

	cfg := &config{
		NoTLS:       true,
		RPCServer:   strings.TrimPrefix(server.URL, "http://"),
		AuthType:    authTypeBasic,
		RPCUser:     "user",
		RPCPassword: "pass",
	}
	result, err := sendPostRequest([]byte(`{"jsonrpc":"1.0"}`), cfg)
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(result), `{"ok":true}`; got != want {
		t.Fatalf("result is %s, want %s", got, want)
	}
}
