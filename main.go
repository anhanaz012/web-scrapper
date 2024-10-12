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





// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	// Step 1: Fetch the HTML page
// 	url := "https://en.wikipedia.org/wiki/FIFA_World_Cup"
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error fetching the URL:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Check if the response status is OK
// 	if response.StatusCode != http.StatusOK {
// 		fmt.Println("Error: Received non-200 response status:", response.StatusCode)
// 		return
// 	}

// 	// Step 2: Read the response body
// 	htmlContent, err := ioutil.ReadAll(response.Body)



// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	// Step 3: Write the content to a new file
// 	ioutil.WriteFile("fifa_world_cup.html", htmlContent, 0644)
// 	if err != nil {
// 		fmt.Println("Error writing to file:", err)
// 		return
// 	}

// 	fmt.Println("Content has been saved to fifa_world_cup.html")
// }

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"github.com/jung-kurt/gofpdf"
	"github.com/russross/blackfriday/v2"
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

	// Step 2: Load the HTML document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error loading HTML document:", err)
		return
	}

	// Step 3: Extract table data and convert to Markdown
	var markdownContent strings.Builder

	doc.Find("table").Each(func(i int, table *goquery.Selection) {
		markdownContent.WriteString("\n| ")
		table.Find("tr").Each(func(j int, row *goquery.Selection) {
			row.Find("th, td").Each(func(k int, cell *goquery.Selection) {
				// Append the cell content to Markdown format
				markdownContent.WriteString(cell.Text() + " | ")
			})
			markdownContent.WriteString("\n| --- |" + "\n") // Add Markdown table separator
		})
	})

	// Step 4: Write Markdown to a file
	markdownFileName := "tables_output.md"
	err = os.WriteFile(markdownFileName, []byte(markdownContent.String()), 0644)
	if err != nil {
		fmt.Println("Error writing Markdown file:", err)
		return
	}
	fmt.Println("Markdown saved to:", markdownFileName)

	// Step 5: Convert Markdown to PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetFont("Arial", "B", 16)
	pdf.AddPage()

	// Convert Markdown to HTML
	htmlContent := blackfriday.Run([]byte(markdownContent.String()))

	// Write HTML content to PDF
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 10, string(htmlContent), "", "L", false)

	// Save PDF to file
	pdfFileName := "tables_output.pdf"
	err = pdf.OutputFileAndClose(pdfFileName)
	if err != nil {
		fmt.Println("Error saving PDF file:", err)
		return
	}
	fmt.Println("PDF saved to:", pdfFileName)
}
