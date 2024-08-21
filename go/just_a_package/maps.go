package justapackage

import "fmt"

func ImplementMaps() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites["Google"])
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
}
