package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
    "github.com/gorilla/mux"
	"gopkg.in/olivere/elastic.v3"
)


type Product struct{
	
	_id 
	
}


	func main() {
		
		log.Println("=============== Starting APIMOBILE ===============")
		createEndPoint()
		createConnectElasticSerach()
	}
	
	func createConnectElasticSerach(){
		
		client, err := elastic.NewClient()	
		
		if err != nil {
	    	log.Fatal("Falha ao conectar com elasticSearch	")
		}else{

		_, err = client.CreateIndex("product").Do()
		
		if err != nil {
		    panic(err)
		}

	}


	func createEndPoint(){
	    router := mux.NewRouter().StrictSlash(true)
	    
		router.HandleFunc("/apimobile", Index)
		log.Println("created endpoint '/apimobile'")
		
		router.HandleFunc("/apimobile/createProduct", createProduct).Methods("POST")
		log.Println("created endpoint '/apimobile/createProduct'")
		
		router.HandleFunc("/apimobile/listProductPublic", Index)
		log.Println("created endpoint '/apimobile/listProductPublic'")
		
	    log.Fatal(http.ListenAndServe(":8888", router))
	}
	
	

	func Index(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}
		
	func createProduct(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "createProduct %q", html.EscapeString(r.URL.Path))
	}
	
	func listProductPublic(w http.ResponseWriter, r *http.Request) {
	    fmt.Fprintf(w, "listProductPublic %q", html.EscapeString(r.URL.Path))
	}

