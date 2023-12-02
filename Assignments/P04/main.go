package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i, url := range urls {
		err := downloadImage(url, fmt.Sprintf("Seq_image%d.jpg", i+1), nil)
		if err != nil {
			fmt.Printf("Error downloading image %d: %v\n", i+1, err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup

	for i, url := range urls {
		wg.Add(1)
		go downloadImage(url, fmt.Sprintf("Async_image%d.jpg", i+1), &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// readURLsFromFile reads URLs from a file and returns a slice of strings
func readURLsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := scanner.Text()
		urls = append(urls, url)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}

func main() {
	fmt.Print("Enter a file name containing URLs: ")
	var inputString string
	fmt.Scanln(&inputString)
	filePath := inputString

	urls, err := readURLsFromFile(filePath)
	if err != nil {
		fmt.Printf("Error reading URLs from file: %v\n", err)
		return
	}

	// Add more image URLs

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// downloadImage.. get this... downloads an image
func downloadImage(url, fileName string, wg *sync.WaitGroup) error {
	if wg != nil {
		defer wg.Done()
	}
	// Create the output file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Download the image
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Copy the image data to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded: %s\n", fileName)
	return nil
}
