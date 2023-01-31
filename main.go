package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const html = `<!DOCTYPE html>
<html>
  <head>
    <title>Hello world from Cartographer</title>
    <style type="text/css">
      body {
        background-color: #ffffff;
        font-family: Arial, sans-serif;
        margin: 0;
        padding: 0;
      }

      h1 {
        text-align: center;
        color: #333;
        margin-top: 50px;
      }

      img {
        display: block;
        margin: 0 auto;
        max-width: 80%;
      }

      footer {
        text-align: center;
        margin-top: 20px;
        padding: 10px;
      }
	  </style>
  </head>
  <body>
    <main>
      <h1>Hello World from... ¡¡Cartographer!!</h1>
      <img src="https://supply-chain-pictures.s3.eu-south-2.amazonaws.com/supply-chain.png">
    </main>
    <footer>
      <p>Copyright © 2023 IPM, a Ricoh Company</p>
    </footer>
  </body>
</html>`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, html)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
