package main

import "fmt"

func main() {
    
	language := make(map[string]string)

	language["en"] = "English"
	language["fr"] = "French"
	language["es"] = "Spanish"
	language["de"] = "German"
	language["it"] = "Italian"
	language["ru"] = "Russian"
	language["pt"] = "Portuguese"

	for key, value := range language {
        fmt.Println(key, value)
    }
	for key, value := range language {
        fmt.Printf("key is %v and value is %v \n",key, value)
    }
	for key, value := range language {
        fmt.Println("Key : [",key,"]  value : [",value,"]")
    }
	for _, value := range language {
        fmt.Println("value: ",value)
    }

	for value := range language {
        fmt.Println("key: ",value)
    }
}