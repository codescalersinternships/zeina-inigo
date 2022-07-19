package main

import (
	"reflect"
	"testing"
)

const content = ` ; last modified 1 April 2001 by John Doe
                  [owner]
                  name = John Doe
                  organization = Acme Widgets Inc.

				  ; use IP address in case network name resolution is not working

                [database]   
                server = 192.0.2.62     
                port = 143
                file = payroll.dat

			    `



func TestLoadFromString(t *testing.T){
	
	t.Run( content , func(t *testing.T){
		parser := Parser{}
	    
		dictionary := map[string]map[string]string{
			"owner" :{
				"name":"John Doe",
				"organization":"Acme Widgets Inc.",

			},
			"database" :{
				"server":"192.0.2.62",
				"port":"143",
				"file":"payroll.dat",
			},
		}
        parsed :=parser.LoadFromString(content)
		got := parsed
		want := dictionary
	
		reflect.DeepEqual(got,want)
	})
}
func TestGetSections(t *testing.T){
	t.Run( content , func(t *testing.T){
		parser:= Parser{}
		parser.LoadFromFile("file.txt")
		dictionary := map[string]map[string]string{
			"owner" :{
				"name":"John Doe",
				"organization":"Acme Widgets Inc.",

			},
			"database" :{
				"server":"192.0.2.62",
				"port":"143",
				"file":"payroll.dat",
			},
		}

       
		got := parser.GetSections()
		want := dictionary
	    
		reflect.DeepEqual(got,want)
	})

}
func TestGetSectionNames(t *testing.T){
	t.Run( content , func(t *testing.T){
		parser:= Parser{}
		parser.LoadFromFile("file.txt")
		dictionary :=[] string {"owner","database"}

       
		got := parser.GetSections()
		want := dictionary
	    
		reflect.DeepEqual(got,want)
		
	})

}
func TestGet(t *testing.T){
t.Run( content , func(t *testing.T){
		assertCorrectMessage := func (t testing.TB, got, want string)  {

		t.Helper()
		if got != want{
			t.Errorf("got %q want %q" , got, want)
		}
		
	}
	parser:= Parser{}
	parser.LoadFromFile("file.txt")
        section := "owner"
		key:="name"
		got := parser.Get(section,key)
		want := "John Doe"
	    
		assertCorrectMessage(t,got,want)
	})
}

func TestSet(t *testing.T){
	t.Run( content , func(t *testing.T){
			assertCorrectMessage := func (t testing.TB, got, want string)  {
	
			t.Helper()
			if got != want{
				t.Errorf("got %q want %q" , got, want)
			}
			
		}
		parser:= Parser{}
		parser.LoadFromFile("file.txt")
			section := "owner"
			key:="name"
			parser.Set(section,key,"zeina")
			got := parser.Get(section,key)
			want := "zeina"
			
			assertCorrectMessage(t,got,want)
		})
	}

	

