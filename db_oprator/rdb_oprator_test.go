package db_oprator

import (
	"context"
	"testing"
)

func TestDel(t *testing.T){
	hosts := []string{"192.168.1.245:3201","192.168.1.245:3202","192.168.1.245:3203"}
	client := NewClient(hosts)
	ctx := context.Background() 
	keys,err := client.Keys(ctx,"specialcolumn::page:*")
	if err != nil {
		t.Error(err)
	}
	for _,key := range keys {
		t.Log(key)
	}		
}