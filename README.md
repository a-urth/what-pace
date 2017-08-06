#### What Pace?

Small and simple tool to help You to determine pace, which You need be running with to finish given distance in given time.


##### Installation

```go
go get github.com/a-urth/what-pace
```

##### Usage
```
whatpace 42 3:30
Pace to finish 42 km in 3:30 - 5m 00s per km

whatpace 10 37
Pace to finish 10 km in 37 - 3m 42s per km

whatpace 5 18
Pace to finish 5 km in 18 - 3m 36s per km

whatpace 21 1:35
Pace to finish 21 km in 1:35 - 4m 31s per km
```

##### Tests
```go
go test
```