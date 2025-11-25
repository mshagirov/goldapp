# goldap : go LDAP application

> LDAP TUI in go

## TUI

```
Ldap info config etc.

[tab_i]:tab/shift+tab{

[table]:j/k/up/down{
content
}
}

tab_i \in {user, group, ou}

```

## LDAP

- [LDAP tutorial](https://www.zytrax.com/books/ldap/)
- [LDAP in Go](https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/)
- go-ldap:
  - [repo](https://github.com/go-ldap/ldap)
  - [Docs](https://pkg.go.dev/github.com/go-ldap/ldap)

## SQLite DB Notes

> Not needed yet

### Turso SQL DB

- [Quickstart guide](https://docs.turso.tech/tursodb/quickstart)
- Installation:

```sh
curl --proto '=https' \
--tlsv1.2 -LsSf \
https://github.com/tursodatabase/turso/releases/latest/download/turso_cli-installer.sh | sh
```

### goose

- [goose documentation](https://pressly.github.io/goose/documentation/annotations/)

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```
