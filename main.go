package main

import (
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte(`
			<html>
				<head>
					<title>Date with me</title>
				</head>
				<body>
					<h1>Let' have a date and Go to the hotel after </h1>
				</body>
			</html>
		`))
	})
	// start the web server
	if err := http.ListenAndServe(":8080", nil) ; err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}