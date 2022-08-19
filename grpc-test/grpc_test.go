package test

import (
	pb "books/grpc-books"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateNewAuthor(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	client := pb.NewStorageManagementClient(conn)

	_, err = client.CreateNewAuthor(ctx, &pb.NewAuthor{Name: "Anton Chaplygin", BookId: 9})
	if err != nil {
		t.Fatalf("CreateNewAuthor failed: %v", err)
		return
	}
}

func TestCreateNewBook(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
		return
	}

	defer conn.Close()
	client := pb.NewStorageManagementClient(conn)

	_, err = client.CreateNewBook(ctx, &pb.NewBook{BookName: "Oscar Wilde Dorian Grey", AuthorId: 9})
	if err != nil {
		t.Fatalf("CreateNewBook failed: %v", err)
		return
	}
}

func TestGetAuthors(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
		return

	}

	defer conn.Close()
	client := pb.NewStorageManagementClient(conn)
	_, err = client.GetAuthors(ctx, &pb.AuthorRequest{BookId: 1})
	if err != nil {
		t.Fatalf("GetUsers failed: %v", err)
		return
	}
}

func TestGetBooks(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
		return

	}

	defer conn.Close()
	client := pb.NewStorageManagementClient(conn)
	_, err = client.GetBooks(ctx, &pb.BooksRequest{AuthorId: 9})
	if err != nil {
		t.Fatalf("GetBooks failed: %v", err)
		return
	}
}
