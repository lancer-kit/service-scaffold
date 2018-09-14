# Natswrap

`natswrap` is a  simple wrapper for [nats.io](nats-io/go-nats) client. 

## Usage 

To start use the `natswrap` package add import:

``` go
...
import (
  "gitlab.inn4science.com/gophers/service-kit/natswrap"
)
...
```

- Fill config structure:

| Field | Type | Required |
| ----- | ---- | ---- |
| Host | string | + |
| Port | int | + |
| User | string |   |
| Password | string |

- Set config: 

``` go
natswrap.SetCfg(cfg)
```

Connect will be initialized at first try to push or subscribe a message
``` go
err := PublishMessage(topic, obj)
```

#### Example

``` go
package main

import (
    "fmt"
    "gitlab.inn4science.com/gophers/service-kit/natswrap"
)

var config = natswrap.Config{
    Host: "127.0.0.1",
    Port: 4222,
    User: "user",
    Password: "password",  
}

err := config.Validate()
if err != nil {
    fmt.Println('invalid nats configuration')
}

natswrap.SetConfig(&config)

testMsg := []string {"1", "2"}
err := natswrap.PublishMessage("Topic", testMsg)
if err != nil {
    log.Default.Error(err)
}

```

