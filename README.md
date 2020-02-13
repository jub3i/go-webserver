# Dev notes

* install reflex:
```bash
$ go get -v github.com/cespare/reflex
```

* setup env vars:
```bash
$ cp .env_example .env
$ vim .env
```

* exec from go-webserver root dir:
```bash
$ ./start_reflex.sh
```

* start the server and run e2e tests
```bash
$ ./start_reflex.sh
$ go test -v
```

# Todo

* docs
* testing
