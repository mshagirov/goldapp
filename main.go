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

	// filters := []struct {
	// 	name   string
	// 	filter string
	// }{
	// 	{name: "Users", filter: "PosixAccount"},     // all PosixGroups
	// 	{name: "Groups", filter: "PosixGroup"},      // all PosixGroups
	// 	{name: "OUs", filter: "OrganizationalUnit"}, // ou's
	// }
	// "(objectClass=*)" // all classes
	// "(uid=*)" // all ldap users
	// "(cn=*)" // all ldap users
	// fmt.Sprintf("(uid=%s)", "jbourne") // find user

	var (
		tabnames []string
		contents []table.Model
		dn       [][]string
	)

	w, h := tabs.GetTableDimensions()

	users := ldap.Users()

	contents = append(contents,
		table.New(table.WithColumns(users[0].Cols),
			table.WithRows(users[0].Rows),
			table.WithFocused(true),
			table.WithHeight(h),
			table.WithWidth(w),
			table.WithStyles(tabs.GetTableStyle()),
		),
	)
	tabnames = append(tabnames, "Users")
	dn = append(dn, users[0].DN)

	contents = append(contents,
		table.New(table.WithColumns(users[1].Cols),
			table.WithRows(users[1].Rows),
			table.WithFocused(true),
			table.WithHeight(h),
			table.WithWidth(w),
			table.WithStyles(tabs.GetTableStyle()),
		),
	)
	tabnames = append(tabnames, "Groups")
	dn = append(dn, users[1].DN)

	orgUnits := ldap.OUs()
	contents = append(contents,
		table.New(table.WithColumns(orgUnits.Cols),
			table.WithRows(orgUnits.Rows),
			table.WithFocused(true),
			table.WithHeight(h),
			table.WithWidth(w),
			table.WithStyles(tabs.GetTableStyle()),
		),
	)
	tabnames = append(tabnames, "OrgUnits")
	dn = append(dn, orgUnits.DN)

	tabs.Run(tabnames, contents, dn)
}
