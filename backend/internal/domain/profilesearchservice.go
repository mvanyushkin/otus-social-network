package domain

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SearchType uint8

const (
	ByFirstNameLastName = SearchType(1)
	ByCity              = SearchType(2)
	ByAge               = SearchType(3)
)

type ProfileSearcher interface {
	FindByFirstNameLastName(ctx context.Context, firstName string, lastName string) ([]Profile, error)
	FindByCity(ctx context.Context, city string) ([]Profile, error)
	FindByAge(ctx context.Context, minAge uint8, maxAge uint8) ([]Profile, error)
}

type ProfileSearcherImplementation struct {
	db *sqlx.DB
}

const selectByFirstNameLastNameQuery = `SELECT Id, Email, FirstName, LastName, Age, Gender, City, Hobby FROM Profile ` +
	`WHERE (FirstName LIKE :firstname AND LastName LIKE :lastname) OR (FirstName LIKE :lastname AND LastName LIKE :firstname)`

func (p ProfileSearcherImplementation) FindByFirstNameLastName(ctx context.Context, firstName string, lastName string) ([]Profile, error) {
	data := struct {
		FirstName string
		LastName  string
	}{}
	data.LastName = fmt.Sprintf("%v%v%v", "%", lastName, "%")
	data.FirstName = fmt.Sprintf("%v%v%v", "%", firstName, "%")

	row, err := p.db.NamedQueryContext(ctx, selectByFirstNameLastNameQuery, data)
	if err != nil {
		return nil, err
	}

	profiles := make([]Profile, 0)
	for row.Next() {
		var profile Profile
		err = row.Scan(&profile.Id, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Age, &profile.Gender, &profile.City, &profile.Hobby)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

const selectWhereByCity = `SELECT Id, Email, FirstName, LastName, Age, Gender, City, Hobby FROM Profile WHERE City LIKE :city`

func (p ProfileSearcherImplementation) FindByCity(ctx context.Context, city string) ([]Profile, error) {
	data := struct {
		City string
	}{}
	data.City = fmt.Sprintf("%v%v%v", "%", city, "%")

	row, err := p.db.NamedQueryContext(ctx, selectWhereByCity, data)
	if err != nil {
		return nil, err
	}

	profiles := make([]Profile, 0)
	for row.Next() {
		var profile Profile
		err = row.Scan(&profile.Id, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Age, &profile.Gender, &profile.City, &profile.Hobby)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

const selectWhereByAge = `SELECT Id, Email, FirstName, LastName, Age, Gender, City, Hobby FROM Profile WHERE Age > :minage AND Age < :maxage`

func (p ProfileSearcherImplementation) FindByAge(ctx context.Context, minAge uint8, maxAge uint8) ([]Profile, error) {
	data := struct {
		MinAge uint8
		MaxAge uint8
	}{
		MinAge: minAge,
		MaxAge: maxAge,
	}

	row, err := p.db.NamedQueryContext(ctx, selectWhereByAge, data)
	if err != nil {
		return nil, err
	}

	profiles := make([]Profile, 0)
	for row.Next() {
		var profile Profile
		err = row.Scan(&profile.Id, &profile.Email, &profile.FirstName, &profile.LastName, &profile.Age, &profile.Gender, &profile.City, &profile.Hobby)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func NewProfileSearcher(db *sqlx.DB) ProfileSearcher {
	return ProfileSearcherImplementation{
		db,
	}
}
