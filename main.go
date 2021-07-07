package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"voice0726/sqlboiler-playground/models"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=55432 user=root password=sekret dbname=sqlboiler sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	users, err := models.Users(
		qm.InnerJoin("\"department\" on \"department\".id = \"user\".department_id"),
	).All(boil.WithDebug(context.Background(), true), db)
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}

	var result []JoinedObj
	err = models.NewQuery(
		qm.Select("\"user\".*, \"department\".*"),
		qm.From("user"),
		qm.InnerJoin("\"department\" on \"department\".id = \"user\".department_id"),
	).Bind(boil.WithDebug(context.Background(), true), db, &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, obj := range result {
		fmt.Println(obj)
	}
}

type JoinedObj struct {
	user       models.User       `boil:",bind"`
	department models.Department `boil:",bind"`
}
