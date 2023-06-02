package learn

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define pair as a struct with two fields, ints named x and y.
type pair struct {
	x, y int
}

// A single function from package http starts a web server.
func LearnWebProgramming() {

	// First parameter of ListenAndServe is TCP address to listen to.
	// Second parameter is an interface, specifically http.Handler.
	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

// Make pair an http.Handler by implementing its only method, ServeHTTP.
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Serve data with a method of http.ResponseWriter.
	w.Write([]byte("You learned Go in Y minutes!"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}

/////////////

func HttpServerV2() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello, World!")

		// Iterate over all headers
		for headerName, headerVal := range req.Header {
			// headerVal is a array, _ is index,
			for _, value := range headerVal {
				fmt.Println(headerName+":", value)
			}
		}
	}

	// Register the handler function with the default ServeMux
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
