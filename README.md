

## serial

```
$ docker-compose up
Recreating contextexample_back2_1
Recreating contextexample_back3_1
Recreating contextexample_back1_1
Recreating contextexample_client_1
Attaching to contextexample_back1_1, contextexample_back3_1, contextexample_back2_1, contextexample_client_1
back3_1   | back3 START
back1_1   | back1 START
back2_1   | back2 START
client_1  | START request to backs...
client_1  | start req to http://back1/
back1_1   | sleeping...
client_1  | back1
client_1  | end req to http://back1/
client_1  | start req to http://back2/
back2_1   | sleeping...
client_1  | back2
client_1  | end req to http://back2/
client_1  | start req to http://back3/
back3_1   | sleeping...
client_1  | back3
client_1  | end req to http://back3/
client_1  | END request to backs...
contextexample_client_1 exited with code 0
```
