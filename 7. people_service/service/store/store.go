package store

import (
	"context"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	conn *pgx.Conn
}

type People struct {
	ID   int
	Name string
}

// NewStore creates new database connection
func NewStore(connString string) *Store {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New("file://migrations", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	return &Store{
		conn: conn,
	}
}

func (s *Store) ListPeople() ([]People, error) {
	rows, err := s.conn.Query(context.Background(), "SELECT id, name FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var people []People

	err = rows.Err()
	if err != nil {
		log.Fatal()
	}

	for rows.Next() {
		var dude People
		err = rows.Scan(&dude.ID, &dude.Name)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, dude)
	}
	return people, err
}

func (s *Store) GetPeopleByID(id int) (People, error) {
	var name string
	err := s.conn.QueryRow(context.Background(), "SELECT name FROM people where id=$1", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return People{Name: name, ID: id}, err
}
