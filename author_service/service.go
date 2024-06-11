package main

import (
	"context"
	db "github.com/wcodesoft/mosha-backed/common/database"
	as "github.com/wcodesoft/mosha-backed/common/protos/authorservice"
)

type authorService struct {
	as.UnimplementedAuthorServiceServer

	// database is the connection to the Database.
	database db.Database[as.Author]
}

func NewAuthorService(database db.Database[as.Author]) as.AuthorServiceServer {
	return &authorService{
		database: database,
	}
}

// CreateAuthor creates a new author.
func (s *authorService) CreateAuthor(ctx context.Context, req *as.CreateAuthorRequest) (*as.Author, error) {
	author, err := s.database.Create(ctx, req.Author.Id, req.Author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

// GetAuthor retrieves an author.
func (s *authorService) GetAuthor(ctx context.Context, req *as.GetAuthorRequest) (*as.Author, error) {
	author, err := s.database.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *authorService) UpdateAuthor(ctx context.Context, req *as.UpdateAuthorRequest) (*as.Author, error) {
	author, err := s.database.Update(ctx, req.Author.Id, req.Author)
	if err != nil {
		return nil, err
	}
	return author, nil
}

func (s *authorService) DeleteAuthor(ctx context.Context, req *as.DeleteAuthorRequest) (*as.DeleteAuthorResponse, error) {
	err := s.database.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &as.DeleteAuthorResponse{
		Success: true,
	}, nil
}
