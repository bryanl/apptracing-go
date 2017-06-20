package main

import (
	"context"

	"time"

	"github.com/bryanl/apptracing-go/internal/platform/db"
	atrand "github.com/bryanl/apptracing-go/internal/platform/rand"
	"github.com/pkg/errors"
)

var (
	queryAllPeople = `SELECT id, first_name, last_name, occupation FROM people LIMIT $1 OFFSET $2`
	queryPerson    = `SELECT id, first_name, last_name, occupation FROM people where id = $1`
)

type person struct {
	ID         int    `db:"id,omitempty" json:"id,omitempty"`
	FirstName  string `db:"first_name,omitempty" json:"first_name,omitempty"`
	LastName   string `db:"last_name,omitempty" json:"last_name,omitempty"`
	Occupation string `db:"occupation,omitempty" json:"occupation,omitempty"`
}

func peopleListService(ctx context.Context, ds *db.DB, page, perPage int) ([]person, error) {
	offset := (page - 1) * perPage

	// delay to simulate serving
	serviceDelay := atrand.Between(10, 75) // add 10 to 75 milliseconds
	time.Sleep(time.Duration(serviceDelay) * time.Millisecond)

	return peopleListData(ctx, ds, perPage, offset)
}

func peopleListData(ctx context.Context, ds *db.DB, limit, offset int) ([]person, error) {
	// fail 20% of the time
	if atrand.Between(1, 100) <= 20 {
		return nil, errors.Errorf("db query failed")
	}

	var people []person
	if err := ds.SelectContext(ctx, &people, queryAllPeople, limit, offset); err != nil {
		return nil, err
	}

	return people, nil
}
