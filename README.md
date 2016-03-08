# LogBoom

Simple tool for load testing logging pipelines.

## Installation

```console
$ go get -u github.com/ejholmes/logboom
```

## Usage

Log 1000 messages per second indefinitely:

```console
$ echo "This is the message that will be logged" | logboom
```

Log lines from a file randomly for 10 minutes:

```console
$ cat /var/log/syslog | logboom -d 10m
```

## Docker Usage

You can run logboom with Docker as well:

```console
$ echo "This is the message that will be logged" | docker run -i ejholmes/logboom
```

## Full Usage

```console
$ logboom -h
Usage of logboom:
  -d duration
      How long to log for.
  -r int
      The number of lines per second to log. (default 1000)
```
