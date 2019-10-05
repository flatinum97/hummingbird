# hummingbird
Tiny HTTP server written in Go. Mainly for debbuging.

## Installation

```
go get -u github.com/supernetnavi/hummingbird/cmd/hb
```

## Usage
Start up a server. Default port is set to 8080.
```
hb
```
If you want to use another port, you can specify it using `-p`
```
hb -p 3000
```

hummingbird provides some endpoints.

Root path `/` shows you fully request details.
And others are as bellow.

|path|description|
|:---|:---|
|/|Fully request details|
|/method|HTTP method|
|/headers|HTTP request headers|
|/body|HTTP request body|
|/remote_addr|Remote IP immediately before. It's not neccessary client IP|
