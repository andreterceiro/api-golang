# Install

Please run the following commands:

```
go mod init module
go get .
```


# Build, if necessary after a change

```
go build main.go
```


# Run
```
./main
```


# Demo
[Vídeo](https://youtu.be/QVEsWWBr-l0)


# Postman collection to make the test more easy

You can access Postman [here](http://postman.com). You will need the desktop agent.

**Important!**: The value of the variable serverAddress in Postman is "`http://127.0.0.1:8081`". You can also execute tests of the GET routes (/products, /products/id) in this address. Ok, you can also execute a POST too in the browser, but with a specific form that isn't in this project (but easy to develop :)) 