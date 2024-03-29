package ini

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
) 




type Parser struct{
	ini map[string]map[string]string
}

func (p *Parser)LoadFromFile(filename string)(err error){

	b, err := os.ReadFile(filename)
	if err != nil {
        fmt.Print(err)
    }
	x := string(b) 

	p.LoadFromString(x)
	return err
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
func (p *Parser) LoadFromString(content string)(err error ) {


	p.ini , err=loadString(content)

	return err
}


func loadString(content string)( x map[string]map[string]string,err error){
	scanner := bufio.NewScanner(strings.NewReader(content))
	finalmap := make(map[string]map[string]string)
	section := ""
    for scanner.Scan(){
		line:=getline(scanner.Text())
        key:= ""
		value := ""
		if line == "section"{
			result := strings.Split(scanner.Text()," ") 
			//fmt.Println(result)
			result = removeEmptyStrings(result)
			
			result[0] =strings.Replace(result[0],"[","",-1)
			result[0] =strings.Replace(result[0],"]","",-1)
			section=result[0]
            finalmap[section]=make(map[string]string)

		}

		if line == "key and value"{
			result := strings.Split(scanner.Text(),"=")
			result = removeEmptyStrings(result)
            key = strings.TrimSpace(result[0])
			value =strings.TrimSpace(result[1])
			finalmap[section][key]=value

          
		}
	}

	return finalmap , err
}
func getline(line string)(name string){

	if strings.Contains(line,"[") && strings.Contains(line,"]"){
		name= "section"
		return name 

	}
    if strings.Contains(line,"=") {
       name="key and value"
	   return name
	}
	if strings.Contains(line,";"){
      name= "comment"
	  return name
	}
 return
}

func (p *Parser)GetSectionNames()(arr []string){

	for key := range p.ini {
		arr = append(arr, key)
	}
	return arr
   
  
}

func (p *Parser)GetSections()(sections map[string]map[string]string){
	
	sections = p.ini
     return sections

}
func (p *Parser) Get(section string,key string) (value string){
	
    values := p.ini[section][key]
	return values

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
    p.SaveToFile(string(jsonStr))
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



