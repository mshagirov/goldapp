package ldap

import (
	"log"

	"github.com/go-ldap/ldap/v3"
)

func Search(url, adminDN, adminPassword, baseDN, filter string) (*ldap.SearchResult, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		log.Printf("Encountered %v\n", err)
		return nil, err
	}
	defer l.Close()

	if err := l.Bind(adminDN, adminPassword); err != nil {
		log.Printf("Encountered %v\n", err)
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree,
		0, 0, 0, false,
		filter,
		[]string{},
		nil,
	)

	res, err := l.Search(searchRequest)
	if err != nil {
		log.Printf("Encountered %v\n", err)
		return nil, err
	}

	return res, err
}
