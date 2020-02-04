reflex -r '\.go$|\.mod$|\.env$' -s -- sh -c 'go build -o go-webserver && du -h ./go-webserver && ./go-webserver'
