package main

import (
	"encoding/json"
	"fmt"

	"strings"
) 



const contentt = ` ; last modified 1 April 2001 by John Doe

                [owner]
                name = John Doe
                organization = Acme Widgets Inc.

				  ; use IP address in case network name resolution is not working

                [database]   
                server = 192.0.2.62     
                port = 143
                file = payroll.dat

			    `




// func LoadFromFile(){

// }
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}
func removeLinesContainingcomments(input string, toRemove []string) string {
	if !strings.HasSuffix(input, "\n") {
	  input += "\n"
	}
  
	lines := strings.Split(input, "\n")
  
	for i, line := range lines {
	  for _, rm := range toRemove {
		if strings.Contains(line, rm) {
		  lines = append(lines[:i], lines[i+1:]...)
		}
	  }
	}
  
	input = strings.Join(lines, "\n")
	input = strings.TrimSpace(input)
	input += "\n"
  
	return input
  }

func LoadFromString(content string)( x map[string]map[string]string) {
    toRemove := []string{";"}
	res := removeLinesContainingcomments(content,toRemove )
    
     println(res)
	 ress := make(map[string]map[string]string)
	 z := json.Unmarshal([]byte(res), &ress)
	 println("ssss")
	 fmt.Println(z)
	  fmt.Println(ress)
	
	return ress
}


func GetSectionNames(content string)(arr []string){

	mapp := LoadFromString(content)
	strKeys := Keys(mapp)

   
  return strKeys
}

func GetSections(content string,name string)(sections map[string]string){
	mapp := LoadFromString(content)
	sections = mapp[name]
     return sections

}
func Get(content string,section string,key string) (value string){
	mapp := LoadFromString(content)
    value = mapp[section][key]
	return value

}

func Set(content string,section string,key string, value string){
	mapp := LoadFromString(content)

	mapp[section][key] = value
	SaveToFile()

}

// func ToString(){

// }
func SaveToFile(){

}



func main(){
	
	 fmt.Println(LoadFromString(contentt))
}