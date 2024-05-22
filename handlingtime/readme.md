# Handling date and time in Golang

* To create format we use this format
```go
time.Date(year int, month Month, day, hour, min, sec, nsec int, loc *Location) Time
```
```go
createdDate := time.Date(2020, time.August, 13, 23, 23, 0, 0, time.UTC)
fmt.Println(createdDate)

```
**formatting**
```go
presentTime := time.Now()
fmt.Println((presentTime))
fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

```

* Here to format we always have to use this line ```go
fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
```