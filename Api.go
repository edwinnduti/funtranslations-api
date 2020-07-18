/*
An API using funtranslations.com
Created on 18 July,2020
Author: Edwin Nduti

Description:
	This is a simple API for translation text given into any of the choosen translations using funtranslations API.
	Feel free to fork it.
*/

package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
)

//types of translations available
const (
	Minion = "minion"
	Mandal = "mandalorian"
	Yoda   = "yoda"
	Doth   = "dothraki"
	Chef   = "chef"
	Val    = "valyrian"
	Emoji  = "emoji"
	Num    = "numbers"
)


func main(){
	start := time.Now()

	//text entered must be of string type
	Text := "456"

	//the uri for the API
	uri := "https://api.funtranslations.com/translate/"+Num+".json?text="+Text

	//GET endpoint
	site,err := http.Get(uri)
	if err != nil{
		fmt.Printf("Gotten Error:%v",err)
	}

	defer site.Body.Close()
	defer fmt.Printf("Time Taken: %v secs",time.Since(start).Seconds())

	//read byte steam
	body , err := ioutil.ReadAll(site.Body)
	if err != nil{
		fmt.Printf("Error:" ,err)
	}

	//convert bytes to strings
	fmt.Println(string(body))

	fmt.Println("\n")
}
