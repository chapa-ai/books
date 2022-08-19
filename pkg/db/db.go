package db

import (
	pb "books/grpc-books"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/tern/migrate"
	"log"
	"os"
	"path"
)

var db *pgx.Conn

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "123"
//	dbname   = "postgres"
//)

func InitDB() (*pgx.Conn, error) {
	if db == nil {
		conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
		d, err := pgx.Connect(context.Background(), conn)
		if err != nil {
			log.Fatalf("Unable to establish connection: %v", err)
		}
		db = d
	}

	return db, nil
}

func MigrateDB() error {
	conn, err := InitDB()
	if err != nil {
		return fmt.Errorf("InitDB failed: %w", err)
	}

	migrator, err := migrate.NewMigrator(context.Background(), conn, "schema_version")
	if err != nil {
		return fmt.Errorf("Unable to create a migrator: %w\n", err)
	}

	mydir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Getwd failed: %w\n", err)
	}

	dir := path.Dir(mydir)
	join := path.Join(dir, "migrations")

	err = migrator.LoadMigrations(join)
	if err != nil {
		return fmt.Errorf("Unable to load migrations: %w\n", err)
	}

	err = migrator.Migrate(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to migrate: %w\n", err)
	}
	return nil
}

func InsertAuthor(created_author *pb.Author) (*pb.Author, error) {

	var author *pb.Author = &pb.Author{}

	err := db.QueryRow(context.Background(), `insert into authors(authorname, bookid) values($1, $2) RETURNING "id", "authorname", "bookid"`, &created_author.AuthorName, &created_author.BookId).Scan(&author.Id, &author.AuthorName, &author.BookId)

	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %w", err)
	}

	fmt.Printf("author: %v", author)

	return author, nil
}

func InsertBook(created_book *pb.Book) (*pb.Book, error) {

	var book *pb.Book = &pb.Book{}

	err := db.QueryRow(context.Background(), `insert into books(bookname, authorid) values($1, $2) RETURNING "id", "bookname", "authorid"`, &created_book.BookName, &created_book.AuthorId).Scan(&book.Id, &book.BookName, &book.AuthorId)

	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %w", err)
	}

	fmt.Printf("Book: %v", book)

	return book, nil
}
