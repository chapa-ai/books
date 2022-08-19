package main

import (
	pb "books/grpc-books"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Errorf("Client connection error: %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewStorageManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createdAuthor, err := c.CreateNewAuthor(ctx, &pb.NewAuthor{Name: "Krylov", BookId: 1})
	if err != nil {
		fmt.Errorf("CreateNewAuthor failed: %v", err)
		return
	}

	fmt.Printf("AuthorName: %v, BookId: %v", &createdAuthor.AuthorName, &createdAuthor.BookId)

	createdBook, err := c.CreateNewBook(ctx, &pb.NewBook{BookName: "Basni Krylov", AuthorId: 1})
	if err != nil {
		fmt.Errorf("CreateNewBook failed: %v", err)
		return
	}

	fmt.Printf("BookName: %v, AuthorId: %d", &createdBook.BookName, &createdBook.AuthorId)

	authors, err := c.GetAuthors(ctx, &pb.AuthorRequest{BookId: 1})
	if err != nil {
		fmt.Errorf("GetAuthors failed: %v", err)
		return
	}
	fmt.Printf("Authors response: %v", authors)

	books, err := c.GetBooks(ctx, &pb.BooksRequest{AuthorId: 1})

	fmt.Printf("Books response: %v", books)
	if err != nil {
		fmt.Errorf("GetBooks failed: %v", err)
		return
	}

}
