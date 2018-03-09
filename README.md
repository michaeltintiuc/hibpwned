# hibpwned [![Go Report Card](https://goreportcard.com/badge/github.com/michaeltintiuc/hibpwned)](https://goreportcard.com/report/github.com/michaeltintiuc/hibpwned)

Wrapper around the [HaveIBeenPwned.com API](https://haveibeenpwned.com/API/v2) written in Go

This is a work in progress

## Installation

`go get github.com/michaeltintiuc/hibpwned/src/hibpwned`

## Usage

Passing a password as an argument

```
> hibpwned -p qwerty
You've been PWNED!
```

Manually enetering the password

```
> hibpwned
Input password (hidden):
You're secure, for now...
```
