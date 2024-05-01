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
[Vídeo 1](https://youtu.be/QVEsWWBr-l0)

After recording the video demo, I made two additional videos, one about testing some things that do not are in the basic flow and the last about inserting/updating a trimmed (of spaces) value in the name of the product and general formatting.

[Video 2](https://youtu.be/P1CCSaFhsw4)

[Video 3](https://youtu.be/oCYcXkwUZ4Y)

# Postman collection to make the test more easy

You can access Postman [here](http://postman.com). You will need the desktop agent.

Please import the file `postmanCollection.json` that is in the root of this repository.

**Important!**: The value of the variable serverAddress in Postman is "`http://127.0.0.1:8081`". You can also execute tests of the GET routes (/products, /products/id) in this address. Ok, you can also execute a POST too in the browser, but with a specific form that isn't in this project (but easy to develop :)) 