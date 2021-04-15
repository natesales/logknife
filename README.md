# logknife

[![Go Report](https://goreportcard.com/badge/github.com/natesales/logknife?style=for-the-badge)](https://goreportcard.com/report/github.com/natesales/logknife)
[![License](https://img.shields.io/github/license/natesales/logknife?style=for-the-badge)](https://raw.githubusercontent.com/natesales/logknife/main/LICENSE)
[![Release](https://img.shields.io/github/v/release/natesales/logknife?style=for-the-badge)](https://github.com/natesales/logknife/releases)

logknife removes sensitive information from your logs. It can be used standalone to hide information before sending a log to someone else, or used at service runtime to prevent private data from being sent to syslog (or any other log consumer).

### Usage:
```
Usage:
  logknife [OPTIONS]

Application Options:
  -i, --ips     Match IP addresses
  -4, --ipv4    Match IPv4 addresses
  -6, --ipv6    Match IPv6 addresses
  -a, --action= "replace" to replace with dummy values, "remove" to replace
                with asterisks
  -V, --version  Show version and exit

Help Options:
  -h, --help    Show this help message
```
