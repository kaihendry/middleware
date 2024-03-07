# Usage

```go
slog.Info("starting", "port", port)

if err := http.ListenAndServe(":"+port, middleware.LogStatus(mux)); err != nil {
    slog.Error("error listening", err)
}
```
