# Pointers in Golang

```go
var number int = 34234

var ptr *int = &number

fmt.Println(number)
fmt.Println(ptr)
fmt.Println(*ptr)
```

- Here `ptr` holds the address of the number variable
- `*ptr` will print the value present at that address , `*` means value at 

**What if poiner is not intialized**
```go
var ptr2 *int ;
fmt.Println(ptr2)
```
Output :
```
<nil>
```

```go
*ptr = *ptr * 3
fmt.Println(number)
```
- This technically means `value at ptr` = `value at ptr` * `3`