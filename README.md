# hibpwned

> Have I been pwned?

[![Go Report Card](https://goreportcard.com/badge/github.com/michaeltintiuc/hibpwned)](https://goreportcard.com/report/github.com/michaeltintiuc/hibpwned)
[![Build Status](https://travis-ci.org/michaeltintiuc/hibpwned.svg?branch=master)](https://travis-ci.org/michaeltintiuc/hibpwned)
[![codecov](https://codecov.io/gh/michaeltintiuc/hibpwned/branch/master/graph/badge.svg)](https://codecov.io/gh/michaeltintiuc/hibpwned)

Wrapper around the [HaveIBeenPwned.com API](https://haveibeenpwned.com/API/v2) written in Go

> This is a work in progress

## Installation

Core:

`go get github.com/michaeltintiuc/hibpwned`

Separate packages:

`go get github.com/michaeltintiuc/hibpwned/pkg/pwd`

## Usage

### Password range

Passing a plain-text password as an argument:

```
> hibpwned -p qwerty
Checking plain-text password
Your password was pwned 3599486 times
```

Passing a SHA-1 hash as an argument:

```
> hibpwned -h B1B3773A05C0ED0176787A4F1574FF0075F7521E
Checking SHA-1 password hash
Your password was pwned 3599486 times
```

Nerding out with same as above:

```
> hibpwned -h $(sha1sum<<(printf "%s" "qwerty"))
Checking SHA-1 password hash
Your password was pwned 3599486 times
```

### Account breaches

> work in progress, dumping out JSON is not very user-friendly, is it?

Passing an email address:

```
>hibpwned -e test@example.com
[{"Title":"000webhost","Name":"000webhost","Domain":"000webhost.com","BreachDate":"2015-03-01","AddedDate":"2015-10-26T23:35:45Z","ModifiedDate":"2017-12-10T21:44:27Z","PwnCount":14936670,"Description":"In approximately March 2015, the free web hosting provider <a href=\"http://www.troyhunt.com/2015/10/breaches-traders-plain-text-passwords.html\" target=\"_blank\" rel=\"noopener\">000webhost suffered a major data breach</a> that exposed almost 15 million customer records. The data was sold and traded before 000webhost was alerted in October. The breach included names, email addresses and plain text passwords.","DataClasses":["Email addresses","IP addresses","Names","Passwords"],"IsVerified":true,"IsFabricated":false,"IsSensitive":false,"IsActive":true,"IsRetired":false,"IsSpamList":false,"LogoType":"png"}, ...]
```

Passing additional parameters:
- domain `-d`
- truncated `-t`
- unverified `-u`

```
>hibpwned -e test@example.com -d example.com -t -u
```

For more info see: https://haveibeenpwned.com/API/v2#BreachesForAccount
