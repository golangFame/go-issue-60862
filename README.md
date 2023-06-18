# Try Execute Otelchi on Go 1.20.5

This repo is created for debugging issue reported [here](https://github.com/riandyrn/otelchi/issues/21).

To make sure the platform is same, we use docker to run the application.

To run the application simply use this command:

```
> make run
```

If everything is okay, we will see following output in console:

```
server_1  | 2023/06/18 11:49:05 server started at :3804
server_1  | 2023/06/18 11:49:05 hello world
try-otelchi-go120_client_1 exited with code 0
```