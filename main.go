package main

import (
	"fmt"
	"os"

	"github.com/mshagirov/goldap/internal/config"
	"github.com/mshagirov/goldap/internal/ldap"
	"github.com/mshagirov/goldap/internal/tui"
)

func main() {
	ldapConfig := config.Read()
	if ldapConfig.LdapUrl == "" {
		fmt.Printf("%v", config.ExampleJson())
		os.Exit(1)
	}

	fmt.Printf("%#v\n", ldapConfig)

	filter := "(objectClass=*)" // all classes
	// filter := fmt.Sprintf("(uid=%s)", "jbourne") // find user

	sr, err := ldap.Search(
		ldapConfig.LdapUrl,
		ldapConfig.LdapAdminDn, "admin123",
		ldapConfig.LdapBaseDn,
		filter,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	choices := []string{}
	for _, entry := range sr.Entries {
		fmt.Printf("\n%s\n", entry.DN)
		choices = append(choices, entry.DN)
		for _, attr := range entry.Attributes {
			fmt.Printf(" |-- %v: %v\n", attr.Name, attr.Values)
		}
	}

	p := tui.NewInitialModel(choices)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
