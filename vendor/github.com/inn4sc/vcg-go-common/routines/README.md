# routines

This is the package for implementing the worker pool pattern.

## Usage 

Just define the `routines.Chief` variable, register your worker using the `AddWorkman` method. 
Before starting, you must initialize registered workers using the `InitWorkers(*logrus.Entry)` method.

A very simple example:

```go
package main

import (
	"github.com/inn4sc/vcg-go-common/routines"
	"context"
)

var WorkersChief routines.Chief

func init()  {
	WorkersChief = routines.Chief{}
	WorkersChief.AddWorkman("my-awesome-worker", &MyWorkman{})
	// `MyWorkman` is a type which implement `Workman` interface.
}

func main () {
	WorkersChief.InitWorkers(nil)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		WorkersChief.Start(ctx)
	}()
	
	defer func() {
		cancel()
	}()
}
```
