package main

import (
	pb "books/grpc-books"
	"books/pkg/db"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type StorageManagementServer struct {
	conn *pgx.Conn
	pb.UnimplementedStorageManagementServer
}

func NewStorageManagementServer() *StorageManagementServer {
	return &StorageManagementServer{}
}

func (server *StorageManagementServer) CreateNewAuthor(ctx context.Context, in *pb.NewAuthor) (*pb.Author, error) {
	created_author := &pb.Author{AuthorName: in.Name, BookId: in.BookId}

	insertedAuthor, err := db.InsertAuthor(created_author)
	if err != nil {
		return nil, fmt.Errorf("insertAuthor failed: %w", err)
	}

	return insertedAuthor, nil
}

func (server *StorageManagementServer) CreateNewBook(ctx context.Context, in *pb.NewBook) (*pb.Book, error) {
	created_book := &pb.Book{BookName: in.BookName, AuthorId: in.AuthorId}

	insertedBook, err := db.InsertBook(created_book)
	if err != nil {
		return nil, fmt.Errorf("insertBook failed: %w", err)
	}

	return insertedBook, nil
}

func (server *StorageManagementServer) GetAuthors(ctx context.Context, in *pb.AuthorRequest) (*pb.AuthorList, error) {
	var authors_list *pb.AuthorList = &pb.AuthorList{}
	rows, err := server.conn.Query(context.Background(), "select id, authorName, bookid from authors where bookid = $1", in.BookId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		author := pb.AuthorResponse{}
		err = rows.Scan(&author.Id, &author.AuthorName, &author.BookId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed: %w", err)
		}

		authors_list.Authors = append(authors_list.Authors, &author)
	}
	fmt.Printf("author_list: %v", authors_list)

	return authors_list, nil
}

func (server *StorageManagementServer) GetBooks(ctx context.Context, in *pb.BooksRequest) (*pb.BooksList, error) {
	var books_list *pb.BooksList = &pb.BooksList{}
	rows, err := server.conn.Query(context.Background(), "select id, bookName, authorid from books where authorid = $1", in.AuthorId)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		author := pb.BooksResponse{}
		err = rows.Scan(&author.Id, &author.BookName, &author.AuthorId)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan failed: %w", err)
		}

		books_list.Books = append(books_list.Books, &author)
	}
	fmt.Printf("books_list: %v", books_list)

	return books_list, nil
}

func (server *StorageManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterStorageManagementServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func main() {
	conn, err := db.InitDB()
	if err != nil {
		fmt.Errorf("InitDB failed: %w", err)
		return
	}

	err = db.MigrateDB()
	if err != nil {
		fmt.Errorf("MigrateDB failed: %w", err)
		return
	}

	defer conn.Close(context.Background())
	var user_mgmt_server *StorageManagementServer = NewStorageManagementServer()
	user_mgmt_server.conn = conn
	if err := user_mgmt_server.Run(); err != nil {
		fmt.Errorf("failed to server: %w", err)
		return
	}
}
