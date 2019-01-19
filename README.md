# rqmetric
Request Metric: Simple command line program to read data from web app log and returns some insights about the requests (url, request count, min/max/avg response time and response codes).

**Features**

- Parse log files from different format, create custom profile to parse your specific log line format.
- Browse the result directly on the terminal.
- Or view it on the web browser by starting its internal web server.

## Usage

```
$ rqmetric

== [RQ Metric v0.1.0 - https://github.com/ekaputra07/rqmetric] ==

Usage examples:
Import log file  =>	rqmetric --import production.log --profile rails
View the report  =>	rqmetric --view 123456
Serve the report =>	rqmetric --serve 123456 --port 8080
Params help      =>	rqmetric -h
```

## Development

- Clone this repository
- Run `make develop` (this will install the [packer](https://github.com/gobuffalo/packr) utility)
- Hack it!
- Run `make build` or `make install` to use it.
