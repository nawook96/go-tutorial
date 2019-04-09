# Go
---
<p align="center"><img width="300" src="https://cdn0.tnwcdn.com/wp-content/blogs.dir/1/files/2018/07/go.png" />

`go-tutorial` provides tutorial how to use **Golang**. In the tutorial, There are basic syntax of **Golang**.

# Why use go?
---

**Golang** is designed at **google** by Robert Griesemer, Rob Pike, and Ken Thompsom.

**Golang** is programming language has been developed for fast performance, reliability, convenience, and easy.

**Golang**'s feature :

- Statically type
- Compiled language
- Garbage collection
- Concurrency
- Support multi-core environment
- Fast compile

# Table of Contents
---

### 1. basics
- [X] [How to build](https://github.com/nawook96/go-tutorial/tree/master/src/1_basics/1_Build)
- [X] [Basic Syntax](https://github.com/nawook96/go-tutorial/tree/master/src/1_basics/2_Basic_Syntax)
- [X] [Variable](https://github.com/nawook96/go-tutorial/tree/master/src/1_basics/3_Variable)
- [ ] Number Type
- [ ] String
- [ ] Boolean
- [ ] Constant
- [ ] Enumeration
- [ ] Package
- [ ] Conditional Statement
- [ ] Loop Statement
- [ ] Goto
- [ ] Switch Statement
- [ ] Function
- [ ] Deffered Call

### 2. advanced
- [ ] Array
- [ ] Slice
- [ ] Map
- [ ] Closure
- [ ] Panic, Recover
- [ ] Pointer
- [ ] Structure
- [ ] Interface
- [ ] Goroutine
- [ ] Channel
- [ ] Synchronization
- [ ] Reflection

# Getting Started
---
## Linux
```bash
$ wget https://storage.googleapis.com/golang/go<version>.linux-<architecture>.tar.gz
$ tar vxzf go<version>.linux-<architecture>.tar.gz
$ sudo mv go /usr/local/
```
And set environment variable
```bash
$ echo "export PATH=$PATH:/usr/local/go/bin" >> .bashrc
$ sourch .bashrc
```
## Ubuntu
```bash
$ sudo apt-get update
$ sudo apt-get install golang
```
## CentOS
```bash
$ sudo yum install golang
```
## Mac OS X
You can download pkg file in [here](http://golang.org/dl).
## Windows
You can download install file in [here](http://golang.org/dl).

And you set environment variable. (ex: PATH=C:\Go\bin, GOPATH=C:\myGo\)
And make pkg,bin,src folder.

# Version
---
- Go 1.12.1

# Author
---
- Dongwook, Shin [@nawook96](https://github.com/nawook96)
- Homepage : https://nawook96.github.io

# Reference
---
http://pyrasis.com/go.html