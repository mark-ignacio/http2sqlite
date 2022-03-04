# http2sqlite

Slap HTTP requests straight into a local SQLite3 database. Fun for debugging!

GPLv3 licensed because if you wanted to pay for something more full-fledged you'd probably be using one of the SaaS services for this, lol.


## Request filtering

To facilitate more targeted request logging, the [Viper config variable](https://github.com/spf13/viper) `filter` supports a [github.com/google/cel-go](https://github.com/google/cel-go) expression that returns a boolean. 


For those not familiar with Viper, this means you can configure the CEL filter via any of:

* execution args: `--filter 'expression here'`
* config file: `filter: 'expression here'`
* env var: `FILTER='expression here'`

Supported CEL variables are:

* `host` - HTTP host header
* `path` - absolute URL path, including the leading slash
* `method` - HTTP method
* `header` - map of HTTP headers
* `source_address` - IP:port as a string

Due to CEL limitations, all n-dashes are converted to underscores. This is particularly noticable when filtering headers.

### Example filters

Only save `PATCH` requests to `localhost:4676`.

```
method == "PATCH" && host == "localhost:4676"
```

Only save requests with a User-Agent of `curl/*`.

```
header.User_Agent.startsWith("curl/")
// alternatively, use RE2 regex
header.User_Agent.matches("^curl/.+")
matches(header.User_Agent, "^curl/.+")
```

Only save requests from a particular IP address.

```
source_address.startsWith("9.9.9.9")
```

For everything else that you can do with CEL, reference the [language grammar](https://github.com/google/cel-spec/blob/master/doc/langdef.md#syntax).