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

**We can modify the original variable using the pointers**
```go
var number int = 10
*ptr = *ptr * 3
fmt.Println(number)
```
Output :
```
30
```
- This technically means `value at ptr` = `value at ptr` * `3`