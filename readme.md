# Learning Go


## how to run?
- build
  ```
  go build
  ./hello
  ```
- run
  ```
  go run main.go
  ```


### test
- go to test file, just use IDE hint to run


## reference

- [Learn Go in ~5mins](https://gist.github.com/prologic/5f6afe9c1b98016ca278f4d507e65510?utm_source=hackernewsletter&utm_medium=email&utm_term=code)
- [Learn Go in Y Minutes](https://learnxinyminutes.com/docs/go/)
- [A Tour of Go](https://go.dev/tour/concurrency/11)
- [A curated list of awesome Go frameworks, libraries and software - Awesome Go / Golang](https://awesome-go.com/)
- [Standard library - Go Packages](https://pkg.go.dev/std)
- [public-apis/public-apis: A collective list of free APIs](https://github.com/public-apis/public-apis)
- [Scripting with Go — Bitfield Consulting](https://bitfieldconsulting.com/golang/scripting)
- [11 Solutions to Exercises in GoLang Tour | by NMTechBytes | Medium](https://medium.com/@anumsarmadmalik/11-solutions-togolang-tours-exercises-7ee61b7b94f5)
### Go scripting


## todo
### awesome go
- [A curated list of awesome Go frameworks, libraries and software - Awesome Go / Golang](https://awesome-go.com/)

###
  - [google/wire: Compile-time Dependency Injection for Go](https://github.com/google/wire)
### seems a good article
  - [Scripting with Go — Bitfield Consulting](https://bitfieldconsulting.com/golang/scripting)
  - [bitfield/script: Making it easy to write shell-like scripts in Go](https://github.com/bitfield/script)
### more tutorial
  - [A Tour of Go](https://go.dev/tour/concurrency/11)
### db related
  - postgres(like typeorm in ts)
    - [go-gorm/gorm: The fantastic ORM library for Golang, aims to be developer friendly](https://github.com/go-gorm/gorm?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego)
  - redis (like ioredis in ts)

### make it fast!!
  - [Codewalk: Share Memory By Communicating - The Go Programming Language](https://go.dev/doc/codewalk/sharemem/)
  - [Google I/O 2013 - Advanced Go Concurrency Patterns - YouTube](https://www.youtube.com/watch?v=QDDwwePbDtw)
  - [Google I/O 2012 - Go Concurrency Patterns - YouTube](https://www.youtube.com/watch?v=f6kdp27TYZs)
### publish a go lib 
  - see above Learn go in ~5 min
### a cli tool
  - [Welcome - urfave/cli](https://cli.urfave.org/)
### a web server app using a framework
  - [gin-gonic/gin: Gin is a HTTP web framework written in Go (Golang). It features a Martini-like API with much better performance -- up to 40 times faster. If you need smashing performance, get yourself some Gin.](https://github.com/gin-gonic/gin?utm_campaign=awesomego&utm_medium=referral&utm_source=awesomego)
  - [gofiber/fiber: ⚡️ Express inspired web framework written in Go](https://github.com/gofiber/fiber)
### vid
  - Golang University 101 introduces fundamental Go concepts and shows you how to use the Go tools to create and manage Go code: https://www.youtube.com/playlist?list=PLEcwzBXTPUE9V1o8mZdC9tNnRZaTgI-1P
  - Golang University 201 steps it up a notch, explaining important techniques like testing, web services, and APIs: https://www.youtube.com/playlist?list=PLEcwzBXTPUE_5m_JaMXmGEFgduH8EsuTs
  - Golang University 301 dives into more advanced topics like the Go scheduler, implementation of maps and channels, and optimisation techniques: https://www.youtube.com/playlist?list=PLEcwzBXTPUE8KvXRFmmfPEUmKoy9LfmAf
### hardcore
  - [The Go Programming Language Specification - The Go Programming Language](https://go.dev/ref/spec)


## Note: something about note different from JS/TS
- statically typed (type is fixed in compiled type)
  - safer
  - clear intention
- less parentheses in syntax
- better way of OOP
  - implement method on type
  - interface to enforce contract
- explicit value/reference sematic
- error handing
  - no try-catch, return error instead
  - this design make which method/fn can product error very obvious
  - compared to TS, which require dev to know if this method could throw error to handle it
- native co-current/parallelism
  - goroutine (lightweight thread)
  - to leverage multi-core modern system for parallelism (not only co-current)
  - channel to allow different goroutine to communicate
  - mutex lock, unlock
- easy to read thr source code
    - in TS ( TS -> JS -> V8(c++ binding to native))
    - transpilation step just make it complicated
    - try to read the node.js http implementation, you need to:
      - git clone node.js
      - find it by hand
    - in Go, 2-3 click away to see the source code
- easy to get executable
  - go build -> binary ( go src -> native machine code)



## Note: how to setup this repo
```bash

mkdir hello
cd hello

## init a go module called hello
go mod init hello

# write code
# .... skip ...

# compile
go build

# run it
./hello

## install external lib
go get -u github.com/bozd4g/go-http-client/

```


