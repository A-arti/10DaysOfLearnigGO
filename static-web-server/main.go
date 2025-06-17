package main
import {
	"fmt"
	"net/http"
	"log"
}


func main(){
	myserver:= http.FileServer(http.dir("/static"))
	

}