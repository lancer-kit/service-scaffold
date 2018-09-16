# routines

This is the package for implementing the worker pool pattern.

## Worker

`Worker` is an interface for async workers which launches and manages by the `Chief`.

The main ideas of the `Worker`: 
- this is simple and lightweight worker;
- it can communicate with surroundings through channels, message queues, etc;
- worker must do one or few small jobs; 
- should be able to run the worker as one independent process;

For example microservice for the image storing can have three workers:

1) Rest API - receives and gives out images from user;
2) Image resizer - compresses and resize the image for the thumbnails;
3) Uploader - uploads files to the S3 storage.

All this workers are part of one microservice, in one binary and able to run them as a one process or as a three different. 

#### Method list:

- `Init(context.Context) Worker` - initializes new instance of the `Worker` implementation. 
- `Run()` - starts the `Worker` instance execution. This should be a blocking call, which in normal mode will be executed in `goroutine`.

## Chief

Chief is a head of workers, it must be used to register, initialize and correctly start and stop asynchronous executors of the type `Worker`.

#### Method list:

- `AddWorker(name string, worker Worker` - register a new `Worker` to the `Chief` worker pool.
- `EnableWorkers(names ...string)` - enables all worker from the `names` list. By default, all added workers are enabled.
- `EnableWorker(name string)` - enables the worker with the specified `name`. By default, all added workers are enabled.
- `IsEnabled(name string) bool` - checks is enable worker with passed `name`.
- `InitWorkers(logger *logrus.Entry)` - initializes all registered workers. 
- `Start(parentCtx context.Context)` - runs all registered workers, locks until the `parentCtx` closes, and then gracefully stops all workers.

In `InitWorkers` **Chief** insert in the context his logger (`*logrus.Entry`), so at the worker you can took by key `routines.CtxKeyLog` and use it.

``` go
// .....
func (w *MyWorker) Init(parentCtx context.Context) routines.Worker {
    logger, ok := parentCtx.Value(routines.CtxKeyLog).(*logrus.Entry)
    if !ok {
        // process error
    }
    // add field with worker name.
    w.Logger = logger.WithField("worker", "my-cool-worker")
    
    // ... do other stuff ... //
    
    return w
}
// .....
```

## Usage 

Just define the `routines.Chief` variable, register your worker using the `AddWorker` method. 
Before starting, you must initialize registered workers using the `InitWorkers(*logrus.Entry)` method.

A very simple example:

``` go
package main

import (
    "gitlab.inn4science.com/gophers/service-kit/routines"
    "context"
)

var WorkersChief routines.Chief

func init()  {
    WorkersChief = routines.Chief{}
    WorkersChief.AddWorker("my-awesome-worker", &MyWorker{})
    // `MyWorker` is a type which implement `Worker` interface.
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
