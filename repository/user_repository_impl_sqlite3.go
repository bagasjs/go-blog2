package repository

import (
	"database/sql"
	"log"

	"github.com/bagasjs/go-blog/entity"
)

type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users (email, name, password) VALUES (?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(user.Email, user.Name, user.Password)
    if err != nil {
        return err
    }
    log.Print("Creating user")
    return nil
}

func (repo *userRepositoryImplSQLite3) FindAll() (users []entity.User, err error) {
    rows, err := repo.db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        user := entity.User{}
        err = rows.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (repo *userRepositoryImplSQLite3) DeleteAll() error {
    return nil
}

func NewUserSQLite3Repository(db *sql.DB) UserRepository {
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email VARCHAR(255) NOT NULL UNIQUE,
        name VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL
    )`)

    if err != nil {
        log.Fatal(err)
    }

    return &userRepositoryImplSQLite3{
        db: db,
    }
}
