# Middleware modules for github.com/emicklei/go-restful library

Middleware components for the github.com/emicklei/go-restful library

# In progress

**metrics** - logging execution time, methods, etc per endpoint

# List of components

## Logger
Log access stats from the clients in standard apache combined access log format

# Ideas / TODO

**wrapper** - add your own middleware

**hook** - only hijack the subset of the middleware features (pre/post/etc)

**timeout|timebox** - time boxed execution of the API function(s)

**negroni** - using negroni with emicklei?