# hibpwned

[![Go Report Card](https://goreportcard.com/badge/github.com/michaeltintiuc/hibpwned)](https://goreportcard.com/report/github.com/michaeltintiuc/hibpwned)
[![Build Status](https://travis-ci.org/michaeltintiuc/hibpwned.svg?branch=master)](https://travis-ci.org/michaeltintiuc/hibpwned)
[![codecov](https://codecov.io/gh/michaeltintiuc/hibpwned/branch/master/graph/badge.svg)](https://codecov.io/gh/michaeltintiuc/hibpwned)
[![Circle CI](https://circleci.com/gh/michaeltintiuc/hibpwned.png?circle-token=baa346fa811747f79bb0faec3184133e07465a1e)](https://circleci.com/gh/michaeltintiuc/hibpwned.png?circle-token=baa346fa811747f79bb0faec3184133e07465a1e)


Wrapper around the [HaveIBeenPwned.com API](https://haveibeenpwned.com/API/v2) written in Go

This is a work in progress

## Installation

`go get github.com/michaeltintiuc/hibpwned/src/hibpwned`

## Usage

Passing a plain-text password as an argument

```
> hibpwned -p qwerty
Checking plain-text password
Your password was pwned 3599486 times
```

Passing a SHA-1 hash as an argument

```
> hibpwned -h B1B3773A05C0ED0176787A4F1574FF0075F7521E
Checking SHA-1 password hash
Your password was pwned 3599486 times
```

Nerding out with same as above

```
> hibpwned -h $(sha1sum<<(printf "%s" "qwerty"))
Checking SHA-1 password hash
Your password was pwned 3599486 times
```
