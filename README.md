# <p align="center"><img src="https://github.com/PiterPentester/goGen/blob/develop/logo.jpeg" width="500" align="center"></p>
# goGen
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/99669174d0db4a9aa75153904bd04a80)](https://app.codacy.com/manual/shtormless/goGen?utm_source=github.com&utm_medium=referral&utm_content=PiterPentester/goGen&utm_campaign=Badge_Grade_Dashboard)
[![CodeFactor](https://www.codefactor.io/repository/github/piterpentester/gogen/badge)](https://www.codefactor.io/repository/github/piterpentester/gogen)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=PiterPentester_goGen&metric=alert_status)](https://sonarcloud.io/dashboard?id=PiterPentester_goGen)

# goGen - a simple password generator on golang.

## It can generate pure random character pass like this ```SMuf462COG1!``` or memorable but super-secure pass like this ```Let#Aside@Connected*6```</p>

## Installation
### From sources:
You need to install golang version > 1.11 

```sh
$ git clone https://github.com/PiterPentester/goGen.git
$ cd goGen
$ go build -o goGen .
```

### Get release: [goGen](https://github.com/PiterPentester/goGen/releases)

## How to use:
```sh
$ ./goGen -abr
```
returning randomised password with default length 16 symbols ```Abracadabra password: yq>$BY23T@lqJ!cS```

```sh
$ ./goGen -mem
```
returning more readable password with default number of words - 3 ```Manta282Grievously438Stank830#```
