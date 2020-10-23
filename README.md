# Go Architectures Examples

Code examples from our talks about how to structure Go code.

# How to run

All these examples are for study purpose, but all the code can be compiled, so you can test it:

### no-architecture
```sh
$ go run no-architecture/*.go
```

### package-architecture
```sh
$ go run package-architecture/main.go
```

### hexagonal-architecture
```sh
$  go run hexagonal-architecture/cmd/counters-api/main.go 
```

### contexts-architecture

**counters-api**
```sh
$  go run contexts-architecture/counters/cmd/counters-api/main.go 
```

**users-api**
```sh
$  go run contexts-architecture/users/cmd/users-api/main.go 
```

## Links to the videos and slides

* 2019-08-23 GopherCon UK [video](https://www.youtube.com/watch?v=KEUmOomnEqc) | slides pending...
* 2019-10-06 Software Crafter Bcn 2019 [video](https://www.youtube.com/watch?v=_rpDmUzP_ZI) | [slides](https://bit.ly/bcn-crafters-fogo-talk)
