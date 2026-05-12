package pokecache

import (
	"testing"
	"time"
)

// declaring a struct
type Human struct {

	// defining struct variables
	Name    string
	Address string
	Age     int
}

func TestCache(t *testing.T) {

	cache := NewCache(time.Second * 5)
	cache.Add("1", []byte(`{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`))
	time.Sleep(time.Second * 2)
	cache.Add("2", []byte(`{ 
        "Name": "Filip",   
        "Address": "Karlsson", 
        "Age": 27 
    }`))
	time.Sleep(time.Second * 2)
	cache.Add("3", []byte(`{ 
        "Name": "Tommy",   
        "Address": "Karlsson", 
        "Age": 61
    }`))
	time.Sleep(time.Second * 2)

	_, ok := cache.Get("1")
	if ok {
		t.Errorf("Key %s should not exist with val %s", "1", `{ 
        "Name": "Deeksha",   
        "Address": "Hyderabad", 
        "Age": 21 
    }`)
	}

}
