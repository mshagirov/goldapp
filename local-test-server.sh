#!/usr/bin/env bash

APP=docker
CONTAINER_NAME="goldap-test"

LDAP_ADMIN_DN="cn=admin,dc=example,dc=com"
LDAP_ADMIN_PASSWORD="admin123"

LDAP_BASE_DN="dc=example,dc=com"

LDAP_URL="ldap://127.0.0.1:389"

$APP \
  run \
  --rm \
  -d \
  --name $CONTAINER_NAME \
  --hostname "${CONTAINER_NAME}.example.com" \
  -p "127.0.0.1:3389:389" \
	-p "127.0.0.1:3636:636" \
  -e LDAP_ORGANISATION="Example Inc" \
  -e LDAP_DOMAIN="example.com" \
  -e LDAP_ADMIN_PASSWORD="${LDAP_ADMIN_PASSWORD}" \
  -e LDAP_TLS_VERIFY_CLIENT="never" \
  docker.io/osixia/openldap:1.5.0

