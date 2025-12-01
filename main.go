package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	"github.com/mshagirov/goldap/internal/config"
	"github.com/mshagirov/goldap/internal/login"
	"github.com/mshagirov/goldap/internal/tabs"
)

func main() {
	// need to have config file
	ldapConfig := config.Read()
	if ldapConfig.LdapUrl == "" {
		fmt.Printf("%v", config.ExampleJson())
		os.Exit(1)
	}

	// enter ldap admin password
	secret, err := login.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ldap := config.LdapApi{
		Config: &ldapConfig,
		Secret: secret,
	}

	filters := []struct {
		name   string
		filter string
	}{
		{name: "Users", filter: "PosixAccount"},     // all PosixGroups
		{name: "Groups", filter: "PosixGroup"},      // all PosixGroups
		{name: "OUs", filter: "OrganizationalUnit"}, // ou's
	}
	// "(objectClass=*)" // all classes
	// "(uid=*)" // all ldap users
	// "(cn=*)" // all ldap users
	// fmt.Sprintf("(uid=%s)", "jbourne") // find user

	var (
		tabnames []string
		contents []table.Model
	)

	w, h := tabs.GetTabledDimensions()

	for _, f := range filters {
		sr, err := ldap.Search(fmt.Sprintf("(objectClass=%v)", f.filter))
		tabnames = append(tabnames, f.name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		col_names_map := make(map[string]struct{})
		for _, entry := range sr.Entries {
			for _, attr := range entry.Attributes {
				col_names_map[attr.Name] = struct{}{}
			}
		}

		col_names := make([]string, 0, len(col_names_map))
		for n := range col_names_map {
			col_names = append(col_names, n)
		}

		rows := []table.Row{}
		for _, entry := range sr.Entries {
			row_i := make([]string, len(col_names))
			for _, attr := range entry.Attributes {
				for n_i, n := range col_names {
					if n == attr.Name {
						row_i[n_i] = fmt.Sprintf("%v", attr.Values)
					}
				}
			}
			row_i = append(row_i, entry.DN)
			rows = append(rows, row_i)
		}
		cols := []table.Column{}
		for _, n := range col_names {
			cols = append(cols, table.Column{Title: n, Width: len(n) + 2})
		}
		cols = append(cols, table.Column{Title: "DN", Width: 4})

		contents = append(contents,
			table.New(table.WithColumns(cols),
				table.WithRows(rows),
				table.WithFocused(true),
				table.WithHeight(h),
				table.WithWidth(w),
				table.WithStyles(tabs.GetTableStyle()),
			),
		)

	}

	tabs.Run(tabnames, contents)
}
