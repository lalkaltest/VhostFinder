package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadDomains(wordlist string, domainList []string) ([]string, error) {
	var domains []string
	var dnSet bool = len(domainList) > 0

	f, err := os.Open(wordlist)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		guess := scanner.Text()
		if dnSet {
			for _, domain := range domainList {
				domains = append(domains, fmt.Sprintf("%s.%s", guess, domain))
			}
		} else {
			domains = append(domains, guess)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return domains, nil
}
