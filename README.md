# Middleware modules for github.com/emicklei/go-restful library

Middleware components for the github.com/emicklei/go-restful library

# In progress

**recorder** - record the the last x. call
set buffer size (global, per endpoint)
playback - curl, json,?

# List of components

## Logger
Log access stats from the clients in standard apache combined access log format

## ApiMetrics
Tracking execution time, http methods, status codes per endpoint

# Ideas / TODO

**forwarder** - all requests are forwarded to a remote url

**stats** - global app stats, https://github.com/thoas/stats

**hook** - only hijack the subset of the middleware features (pre/post/etc)

**timeout|timebox** - time boxed execution of the API function(s)

**negroni** - using negroni with emicklei?

**wrapper** - add your own middleware
