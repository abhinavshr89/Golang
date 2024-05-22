# Maps in Golang


### creating a map in golang

```go
language := make(map[string]string)

	language["en"] = "English"
	language["fr"] = "French"
	language["es"] = "Spanish"
	language["de"] = "German"
	language["it"] = "Italian"
	language["ru"] = "Russian"
	language["pt"] = "Portuguese"

```

**syntax for creating a map in golang**
```
mapName := make(map[keyType]valueType)

```
---------------------------------------------
## Different ways to print maps using loops 


```go
for key, value := range language {
    fmt.Println(key, value)
}
```
Output:
```
pt Portuguese
en English
fr French
es Spanish
de German
it Italian
ru Russian
```


```go
for key, value := range language {
    fmt.Printf("key is %v and value is %v \n",key, value)
}
```
Output:
```
key is en and value is English
key is fr and value is French
key is es and value is Spanish
key is de and value is German
key is it and value is Italian
key is ru and value is Russian
key is pt and value is Portuguese
```

```go
for key, value := range language {
    fmt.Println("Key : [",key,"]  value : [",value,"]")
}
```

Output:
```
Key : [ ru ]  value : [ Russian ]
Key : [ pt ]  value : [ Portuguese ]
Key : [ en ]  value : [ English ]
Key : [ fr ]  value : [ French ]
Key : [ es ]  value : [ Spanish ]
Key : [ de ]  value : [ German ]
Key : [ it ]  value : [ Italian ]
```

```go
for _, value := range language {
    fmt.Println("value: ",value)
}

```

Output:
```
value:  German
value:  Italian
value:  Russian
value:  Portuguese
value:  English
value:  French
value:  Spanish
```


```go
for value := range language {
    fmt.Println("key: ",value)
}
```
Output:
```
key:  en
key:  fr
key:  es
key:  de
key:  it
key:  ru
key:  pt
```