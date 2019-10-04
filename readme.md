Get trace headers from a context:

```
headers, found := traceheaders.FromContext(ctx)
```

Inject trace headers into the header of an http request:

```
headers.Inject(req.Header)
```
