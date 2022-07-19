package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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


type Parser struct{
	ini map[string]map[string]string
}

func (p *Parser)LoadFromFile(filename string) (str string){

	b, err := os.ReadFile(filename)
	if err != nil {
        fmt.Print(err)
    }
	x := string(b) 

	return x
}
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
  func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
func (p *Parser) LoadFromString(content string)( ) {


	p.ini=loadString(content)

}
func loadString(content string)( x map[string]map[string]string) {
    toRemove := []string{";"}
	res := removeLinesContainingcomments(content,toRemove )
    
     println(res)
	 rest := strings.Split(res," ") 
    fmt.Println(rest)
	 rest = removeEmptyStrings(rest)
	 
	 ress := make(map[string]map[string]string)
	 
	 for i ,s := range rest{
	  fmt.Println(i,s)
	 }

	 
	
	 for i, s := range rest {
		fmt.Println(i, s)
		result := []rune(s)
		var sectionnn string
		firstcharacter := string(result[0:1])

		if firstcharacter == "["{
			ress[s]=make(map[string]string)
			sectionnn = s
			
			println(sectionnn)
		}
		fmt.Println(ress)
		
	
		if rest[i+1] =="="{
			
			
			//rest[i+1] = ress[sectionnn]
			
			
			//ress[sectionnn][s]= rest[i+1]

            println(rest[i-1])
		}
		
		
	}
	 fmt.Println(ress)

	 println("ssss")
	
	
	return ress
}


func (p *Parser)GetSectionNames()(arr []string){

	
	strKeys := Keys(p.ini)

   
  return strKeys
}

func (p *Parser)GetSections()(sections map[string]map[string]string){
	
	sections = p.ini
     return sections

}
func (p *Parser) Get(section string,key string) (value string){
	
    value = p.ini[section][key]
	return value

}

func (p *Parser) Set(section string,key string, value string){
	

	p.ini[section][key] = value
	

}

func (p *Parser) ToString()(final string){

	jsonStr, err := json.Marshal(p.ini)
	
    if err != nil {
        fmt.Printf("Error: %s", err.Error())
    } else {
        fmt.Println(string(jsonStr))
    }
	return string(jsonStr)
}

func (p *Parser) SaveToFile(finalstr string)(err error){
	f, err := os.Create("output.txt")
	if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
	_, err2 := f.WriteString(finalstr)
	if err2 != nil {
        log.Fatal(err2)
    }
	return err
}



func main(){
	
	 fmt.Println(loadString(contentt))
	 ds := Parser{}

	ds.LoadFromFile("file.txt")

	fmt.Println(ds.GetSections())

	ds.GetSectionNames()

}

