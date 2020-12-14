# backoff
go backoff lib,support jitter,avoid 'thundering herd' effect

## install

`go get github.com/mmaxiaolei/backoff`

## usage

b := backoff.NewExponentialBackoff()
d := b.Next()
