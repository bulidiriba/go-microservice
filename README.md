# go-microservice

## To run the restful product api

First run the server in one terminal


```go run main.go```


Then in another terminal or browser run the client

1. To get list of the products -- GET request

    ```curl localhost:9090```

2. To Add new product -- POST Request

    ```curl -v localhost:9090 -d '{"id":1, "name":"tea", "description":"a nice cup of tea"}'```


3. To Update existing product -- PUT request

    ```curl -v localhost:9090/1 -XPUT -d '{"name":"tea", "description":"a nice cup of tea"}'```


