package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const crtshURL = "https://crt.sh?q=%s"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run enum.go <domain>")
		os.Exit(1)
	}

	targetDomain := os.Args[1]
	subdomains, err := enumSubdomains(targetDomain)
	if err != nil {
		log.Fatal("Error:", err)
	}

	fmt.Printf("Subdomains for %s:\n", targetDomain)
	for _, subdomain := range subdomains {
		fmt.Println(subdomain)
	}
}

func enumSubdomains(targetDomain string) ([]string, error) {
	url := fmt.Sprintf(crtshURL, targetDomain)

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	subdomains := make(map[string]struct{})
	doc.Find("table tr td:nth-child(5)").Each(func(_ int, s *goquery.Selection) {
		subdomain := strings.TrimSpace(s.Text())
		if subdomain != "" {
			subdomains[subdomain] = struct{}{}
		}
	})

	subdomainList := make([]string, 0, len(subdomains))
	for subdomain := range subdomains {
		subdomainList = append(subdomainList, subdomain)
	}

	return subdomainList, nil
}
