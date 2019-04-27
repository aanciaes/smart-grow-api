#!/bin/bash

external_dependencies=(
    "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "github.com/mattn/go-sqlite3"
    "github.com/dgrijalva/jwt-go"
)

for item in ${external_dependencies[*]}
do
    printf "Downloading dependency %s\n" $item
    go get $item
done