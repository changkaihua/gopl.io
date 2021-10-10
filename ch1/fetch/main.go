// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const testUrl = "http://10.112.158.59:8888/getOneJobDetail?jobId=1"

func main() {
	if len(os.Args) > 1 {
		for _, url := range os.Args[1:] {
			if !strings.HasPrefix(url, "http://") {
				url = "http://" + url
			}
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			// b, err := ioutil.ReadAll(resp.Body)
			// resp.Body.Close()
			// if err != nil {
			// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			// 	os.Exit(1)
			// }
			// fmt.Printf("%s", b)

			fmt.Println("http status:", resp.Status)

			if resp.StatusCode == http.StatusOK {
				len, err := io.Copy(os.Stdout, resp.Body)
				fmt.Printf("\n ====> len: %v\n", len)
				if err != nil {
					fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", testUrl, err)
					os.Exit(1)
				}
			} else {
				fmt.Println("fetch status err", resp.Status)
			}

		}
	} else {
		resp, err := http.Get(testUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", testUrl, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)

		// excerise 1
		fmt.Println("http status:", resp.Status)

		if resp.StatusCode == http.StatusOK {
			len, err := io.Copy(os.Stdout, resp.Body)
			fmt.Printf("\n ====> len: %v\n", len)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", testUrl, err)
				os.Exit(1)
			}
		} else {
			fmt.Println("fetch status err", resp.Status)
		}
	}

}

//!-
