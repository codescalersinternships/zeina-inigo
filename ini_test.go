package main

import "testing"

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



func TestINI(t *testing.T){
	assertCorrectMessage := func (t testing.TB, got, want string)  {

		t.Helper()
		if got != want{
			t.Errorf("got %q want %q" , got, want)
		}
		
	}

	t.Run( content , func(t *testing.T){
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

		got := LoadFromString(content)
		want := dictionary

	
		assertCorrectMessage(t,got,want)
	})

	t.Run( content , func(t *testing.T){
		strArr := []string{"owner" ,"database"}

		got := GetSectionNames(content)
		want := strArr
	
		assertCorrectMessage(t,got,want)
	})

	t.Run( content , func(t *testing.T){
		dictionary := map[string]map[string]string{
			"owner" :{
				"name":"John Doe",
				"organization":"Acme Widgets Inc.",

			},
			
		}

        section := "owner"
		got := GetSections(content,section)
		want := dictionary
	    
		assertCorrectMessage(t,got,want)
	})


	t.Run( content , func(t *testing.T){
		
        section := "owner"
		key:="name"
		got := Get(content,section,key)
		want := "John Doe"
	    
		assertCorrectMessage(t,got,want)
	})

	t.Run( content , func(t *testing.T){
		
        section := "owner"
		key:="name"
		value :="zeina"
        Set(content,section,key,value)
		got := Get(content,section,key)
		want := "zeina"
	    
		assertCorrectMessage(t,got,want)
	})

}