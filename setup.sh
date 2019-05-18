#!/bin/bash

external_dependencies=(
    "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

for item in ${external_dependencies[*]}
do
    printf "Downloading dependency %s\n" $item
    go get -u -v $item
done