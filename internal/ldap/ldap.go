package ldap

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func Search(url, adminDN, adminPassword, baseDN, filter string) (*ldap.SearchResult, error) {
	l, err := ldap.DialURL(url)
	if err != nil {
		return nil, fmt.Errorf("ldap.Search DialURL; %v", err)
	}
	defer l.Close()

	if err := l.Bind(adminDN, adminPassword); err != nil {
		return nil, fmt.Errorf("ldap.Search Bind; %v", err)
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
		return nil, fmt.Errorf("ldap.Search Search; %v", err)
	}

	return res, err
}
