package pokecache

import (
	"testing"
	"time"
)


func TestCreatedCach(t *testing.T){
    cache := NewCache(5*time.Minute)
    if cache.data ==nil{
        t.Error("cach failed")
    }
    
}

func TestAddCach(t *testing.T){
    cache := NewCache(5*time.Minute)
   cache.Add("key1",[]byte("val1")) 
    ac,ok := cache.Get("key1")
    if !ok{
        t.Error("key not found")
    }
    if string(ac)!= "val1"{
        t.Error("values doesn't match")
    }
}
