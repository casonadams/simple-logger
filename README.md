# simple-logger
golang simpler logger

## Env var options
- LOG_LEVEL `[ debug, trace, info, warn, error, fatal ]` sets logging level
- LOG_LEVEL `[ 6, 5, 4, 3, 2, 1 ]` can use numbers instead
- LOG_DATE `[ false, 0 ]` remove date line from logs
- LOG_COLOR `[ false, 0 ]` remove color from logs
- LOG_FUNC `[ false, 0 ]` remove function from logs
- LOG_UTC `[ false, 0 ]` use local time instead of UTC from logs

## Example

```bash
go get -u github.com/casonadams/simple-logger/v2
```

```go
package main

import logger "github.com/casonadams/simple-logger/v2"

func main() {
        log := logger.NewLogger("test")
        log.Info("Hello World")
        log.Infof("Hello Earth %v", 5)
        log.Warn("Warn message")
}
```

## Run

Should only print warn message
```bash
LOG_LEVEL=debug go run .
```

![output](https://github.com/casonadams/simple-logger/examples/output.png)