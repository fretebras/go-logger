# Go Logger SDK


This is a package to facilitate the creation and standardization of structured logs in systems developed in Golang.



## Screenshots
![App Screenshot](doc/output_example.png)

```text
{"level":"info","app":"system-example","version":"1.10.0","environment":"production","trace_id":"01GDEC07RTG85ZWEM1BVMBSC6J","service":"sample-service","context":["632a11dc34ad9aff2b447acf","2022-09-20 18:13:58 -0300 -03","example],"time":1663708438311234,"message":"Success message"}
{"level":"error","app":"system-example","version":"1.10.0","environment":"production","trace_id":"01GDEC5V567XFBWX7WXGFVWWKM","service":"sample-service","context":["Message Error","632a2d2c34ad9aff2b447ad7","confirmed"],"trace":"/go/src/app/services/example/service.go:130","time":1663708621993235,"message":"message_error"}
{"level":"info","app":"system-example","version":"1.10.0","environment":"production","trace_id":"01GDEC5V567XFBWX7WXGFVWWKM","service":"sample-service":["632a2d2c34ad9aff2b447ad7","2022-09-20 18:17:01 -0300 -03","example"],"time":1663708621993326,"message":"Info Message"}
```


## Getting Started

```bash
  go get github.com/fretebras/go-logger
```

### Simple Logging Example

For simple logging, import the global logger package **github.com/fretebras/go-logger**

```go
package main

import (
    log "github.com/fretebras/go-logger"
)

func main() {
	logger:= log.New()
	logger.Debug("hello world")
}

// Output: {"level":"debug","app":"my-application","version":"0.0.1","environment":"local","service":"","context":[],"time":1682459856974376,"message":"hello world"}
```

### Info Logging

```go
package main

import (
	log "github.com/fretebras/go-logger"
	"errors"
)

func main() {
	logger := log.New()
	logger.Info("Info logging")
}

// Output: {"level":"info","app":"my-application","version":"0.0.1","environment":"local","service":"","context":[],"time":1682460626356282,"message":"Info logging"}

```


### Warning Logging

```go
package main

import (
	log "github.com/fretebras/go-logger"
	"errors"
)

func main() {
	logger := log.New()
	logger.Info("Info logging")
}

// Output: {"level":"warn","app":"my-application","version":"0.0.1","environment":"local","service":"","context":[],"time":1682460693606309,"message":"Warning logging"}


```

### Error Logging

```go
package main

import (
	log "github.com/fretebras/go-logger"
	"errors"
)

func main() {
	logger := log.New()
	err := errors.New("error")
	logger.Error(err, "Error log")
}

// Output: {"level":"error","app":"my-application","version":"0.0.1","environment":"local","service":"","context":["Error log"],"trace":"/home/felipeteixeira/Projetos/test/main.go:11","time":1682460148031756,"message":"error"}
```

### Logging Critical Messages

```go
package main

import (
    "errors"

    log "github.com/fretebras/go-logger"
    "github.com/rs/zerolog/log"
)

func main() {
	logger := log.New()
	err := errors.New("A repo man spends his life getting into tense situations")
	logger.Critical(err, "Cannot start service")
}

// Output: {"level":"fatal","app":"my-application","version":"0.0.1","environment":"local","service":"","context":["Cannot start service"],"trace":"/home/felipeteixeira/Projetos/test/main.go:11","time":1682460350651457,"message":"A repo man spends his life getting into tense situations"}
exit status 1
```

### Set Name Log Service

```go
package main

import (
	log "github.com/fretebras/go-logger"
	"errors"
)

func main() {
	logger := log.New()
	logger.SetLogService("my-simple-service")
	logger.Info("Info logging from Service")
}

// Output: {"level":"info","app":"my-application","version":"0.0.1","environment":"local","service":"my-simple-service","context":[],"time":1682460860632831,"message":"Info logging from Service"}

