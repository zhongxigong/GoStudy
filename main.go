package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "this is home")
}

func HelloList(){

}


func HelloGo() string{
	println("hello world")
	return "hello world"
}

func HelloMuti(){

}

func HelloWeb(){

}

func HelloInterface()  {
	
}

func HelloStruct()  {

}

func HelloClourse(){

}


func HelloMap(){

}

func HelloRange()  {

}

func HelloPointer(){

}




//import (
//	"context"
//	"log"
//	"time"
//
//	"github.com/chromedp/chromedp"
//)
//
//func main() {
//	var err error
//
//	// create context
//	ctxt, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// create chrome instance
//	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// run task list
//	err = c.Run(ctxt, click())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// shutdown chrome
//	err = c.Shutdown(ctxt)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// wait for chrome to finish
//	err = c.Wait()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func click() chromedp.Tasks {
//	return chromedp.Tasks{
//		chromedp.Navigate(`https://golang.org/pkg/time/`),
//		chromedp.WaitVisible(`#footer`),
//		chromedp.Click(`#pkg-overview`, chromedp.NodeVisible),
//		chromedp.Sleep(150 * time.Second),
//	}
//}