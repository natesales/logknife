<div align="center">
<h1>logknife</h1>

Remove sensitive information from your logs.

[![Go Report](https://goreportcard.com/badge/github.com/natesales/logknife?style=for-the-badge)](https://goreportcard.com/report/github.com/natesales/logknife)
[![License](https://img.shields.io/github/license/natesales/logknife?style=for-the-badge)](https://raw.githubusercontent.com/natesales/logknife/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/natesales/logknife?style=for-the-badge)](https://github.com/natesales/logknife/releases)

![screenshot](carbon.svg)
</div>

### Examples

```bash
$ echo "12345678-1234-4321-b321-7e3fe946f781 1.1.1.1 2606:4700:4700::1111" | logknife -
8ac8d0fc-4f7e-4969-8d5b-6f5f4e116b37 10.39.213.116 2001:db8:bb14:9f27::1111

$ logknife --redact test.txt                                                                                                                                             130 ↵ nate@altair
logknife is a log redactor.
It uses regex to substitute sensitive data with an innocuous replacement or a redaction message.
IP addresses: ******** and ********
UUID: ********
```

### Installation

`logkinfe` is available in binary form from:

* `go install github.com/natesales/logknife@latest`
* [My public package repositories](https://github.com/natesales/repo)
* [GitHub releases](https://github.com/natesales/logknife/releases)

### Usage

```
Remove sensitive information from your logs.

Usage:
  logknife [flags]

Flags:
  -t, --entropy-threshold int      Minimum entropy threshold (default 50)
  -h, --help                       help for logknife
  -i, --no-ips                     Don't match IP addresses
  -u, --no-uuids                   Don't match UUIDs
  -r, --redact                     Replace matches with redaction pattern instead of innocuous substitutes
      --redaction-pattern string   Redaction pattern (default "********")
  -v, --verbose                    Enable verbose logging
  -V, --version                    Show version
```
