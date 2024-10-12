// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
//     "github.com/PuerkitoBio/goquery"
// )

// // Function to scrape a website
// func scrape(url string) {
//     // Step 1: Send a GET request to the URL
//     res, err := http.Get(url)
//     if err != nil {
//         log.Fatalf("Error fetching the URL: %v", err)
//     }
	
//     defer res.Body.Close()

//     if res.StatusCode != 200 {
//         log.Fatalf("Failed to fetch page, status code: %d", res.StatusCode)
//     }

//     // Step 2: Parse the HTML
//     doc, err := goquery.NewDocumentFromReader(res.Body)
//     fmt.Printf("Page Title: %s\n", doc)
//     if err != nil {
//         log.Fatalf("Error parsing the page: %v", err)
//     }

//     // Step 3: Find and print the page title
//     title := doc.Find("title").Text()
//     fmt.Printf("Page Title: %s\n", title)
// }
// func main() {
//     // Scrape a single website
//     url := "https://en.wikipedia.org/wiki/FIFA_World_Cup"
//     fmt.Println("Scraping:", url)
//     scrape(url)
// }





package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Step 1: Fetch the HTML page
	url := "https://en.wikipedia.org/wiki/FIFA_World_Cup"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status is OK
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Received non-200 response status:", response.StatusCode)
		return
	}

	// Step 2: Read the response body
	htmlContent, err := ioutil.ReadAll(response.Body)



	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}


   



	// Step 3: Write the content to a new file
	ioutil.WriteFile("fifa_world_cup.html", htmlContent, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Content has been saved to fifa_world_cup.html")
}

