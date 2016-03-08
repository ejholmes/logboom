# LogBoom

Simple tool for load testing logging pipelines.

## Installation

```console
$ go get -u github.com/ejholmes/logboom
```

## Usage

Log 1000 messages per second for 10 minutes:

```console
$ echo "This is the message that will be logged" | logboom -d 10m -r 1000
```

Log lines from a file randomly:

```console
$ cat /var/log/syslog | logboom -d 10m -r 1000
```
