# Info Worker

This worker is used for health check and workers monitoring
via API
- Allows to specify route prefix; mount pprof 
- **Api** 
  - [swagger](./swagger.yaml) - workers api

##How to use:
1. Add following structure into config file: 
    ```
    info_worker:
        host: localhost
        port: 2443
        profiler: true
        prefix: "/test"
    ```
    Where **profiler** property is responsible for mounting net/http/pprof from chi middleware.Profiler()
     
2. Initialize worker:
    ```
    worker := infoworker.GetInfoWorker(wApi, ctx, wInfo)
    ```
    Where:
    * wApi --> **Conf** structure with worker configuration
    * ctx  --> background context with pointer to worker chief
    * wInfo --> **Info** structure with info about current service 
    
    *This structures belong to infoworker packages* 
3. Add worker into worker chief and config struct
    ```
    WorkerChief.AddWorker("info-server", worker)
    ```
###Warning!
**Run this worker only in dev mode**
