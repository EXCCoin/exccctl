exccctl
======

[![Build Status](https://github.com/EXCCoin/exccctl/workflows/Build%20and%20Test/badge.svg)](https://github.com/EXCCoin/exccctl/actions)
[![ISC License](https://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![Go Report Card](https://goreportcard.com/badge/github.com/EXCCoin/exccctl)](https://goreportcard.com/report/github.com/EXCCoin/exccctl)

exccctl is a command-line client for interacting with the JSON-RPC servers of
[exccd](https://github.com/EXCCoin/exccd) and
[exccwallet](https://github.com/EXCCoin/exccwallet).

## Usage

In its default configuration, exccctl connects to exccd's mainnet RPC port on
localhost.  The `--wallet` and `--testnet` flags change these defaults to the
exccwallet RPC port and/or testnet ports.  The `--rpcserver/-s` flag can be used
to specify other hostnames or IP addresses of the server, and can also be used
to override the port defaults.

exccctl will attempt to read exccd and exccwallet config files for the
user/password authentication.  If these fields cannot be read, exccctl must be
manually configured with the correct authentication.  Permanent configuration
changes may be written to a config file in a platform-specific location:

* macOS: `~/Library/Application Support/exccctl/exccctl.conf`
* Windows: `%LOCALAPPDATA%\exccctl\exccctl.conf`
* Linux and other Unix: `~/.exccctl/exccctl.conf`

## Build and installation

Development build can be performed by running `go install` in a
locally checked-out repository.

Appending `@master` to this will perform a release build using the latest code
from the master branch.  This may be useful to execute RPC methods not yet found
in the latest Exchangecoin release.

## Developing

When developing either the exccd or exccwallet RPC servers and making
modifications to the RPC methods, you may want to build a development version of
exccctl supporting these changes.  Due to exccctl being built around
package-global method registrations and reflection, supporting these changes
only requires building with the updated packages.  To perform this, module
replacements may be utilized to point to development versions of exccd and
exccwallet in your build environment:

```
$ go mod edit -replace=github.com/EXCCoin/exccd/rpc/jsonrpc/types/v2=../exccd/rpc/jsonrpc/types
$ go mod edit -replace=github.com/EXCCoin/exccwallet/v2=../exccwallet
```

These replaces should be removed prior to committing any updated module
requires:

```
$ go mod edit -dropreplace=github.com/EXCCoin/exccd/rpc/jsonrpc/types/v2
$ go mod edit -dropreplace=github.com/EXCCoin/exccd/v2
```

## Contact

If you have any further questions you can find us at:

https://excc.co

## Issue Tracker

The [integrated github issue tracker](https://github.com/EXCCoin/exccctl/issues)
is used for this project.

## License

exccctl is licensed under the [copyfree](http://copyfree.org) ISC License.
