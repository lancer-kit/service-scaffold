# log

`log` is a  simple wrapper for [logrus](https://github.com/sirupsen/logrus). 

## Usage 

To start use the `log` package add import:

``` go
...
import (
  "gitlab.inn4science.com/gophers/service-kit/log" // imports as package "log"
)
...
```

- Fill config structure:

| Field | Type | Description |
| ----- | ---- | ----------- |
| AppName | string | identifier of the app |
| Level | string | level of the logging output |
| Sentry | string | dsn string for the sentry hook |
| AddTrace | bool | enable the inclusion of the file name and line number in the log |
| JSON | bool | enable json formatted output |

- Call init function: 

``` go
logger, err := log.Init(config)
```

A once-initialized `logrus.Entry` can then be used anywhere. To access it, use the getter:

``` go
logger := log.Get()
```

#### Example

``` go
package main

import (
    "fmt"
    "gitlab.inn4science.com/gophers/service-kit/log" // imports as package "log"
)

var config = log.Config{
    AppName: "my-app",
    Level: "warn",
    Sentry: "http://....",
    AddTrace: true,  
}

logger, err := log.Init(config)
if err != nil {
    fmt.Println(err)
    panic("unable to init log Entry")
}

logger.Info("log initialized")
logger.Error("log test error")

logger2 := log.Get()

fmt.Println(logger == logger2) // true
```

