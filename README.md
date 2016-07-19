# Few simple middleware stuff for github.com/emicklei/go-restful library

Some middleware components for the github.com/emicklei/go-restful library

# Ideas
**access_logger** - log basic access stats from the clients in std. apache format
**api_metrics** - logging execution time, methods, etc per endpoint
**wrapper** - add your own middleware
**hook** - only hijack the subset of the middleware features (pre/post/etc)
**timeout** - time boxed execution of the API function(s)
**negroni** - using negroni with emicklei?