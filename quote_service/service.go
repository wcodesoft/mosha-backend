package main

import (
	"context"
	db "github.com/wcodesoft/mosha-backed/common/database"
	qs "github.com/wcodesoft/mosha-backed/common/protos/quoteservice"
	"google.golang.org/protobuf/types/known/emptypb"
)

type quoteService struct {
	qs.UnimplementedQuoteServiceServer

	// database is the connection to the Database.
	database db.Database[qs.Quote]
}

// NewQuoteService creates a new QuoteService.
func NewQuoteService(database db.Database[qs.Quote]) qs.QuoteServiceServer {
	return &quoteService{
		database: database,
	}
}

// CreateQuote creates a new quote.
func (s *quoteService) CreateQuote(ctx context.Context, req *qs.CreateQuoteRequest) (*qs.Quote, error) {
	quote, err := s.database.Create(ctx, req.Quote.Id, req.Quote)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

// GetQuote retrieves a quote.
func (s *quoteService) GetQuote(ctx context.Context, req *qs.GetQuoteRequest) (*qs.Quote, error) {
	quote, err := s.database.Read(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

// UpdateQuote updates a quote.
func (s *quoteService) UpdateQuote(ctx context.Context, req *qs.UpdateQuoteRequest) (*qs.Quote, error) {
	quote, err := s.database.Update(ctx, req.Quote.Id, req.Quote)
	if err != nil {
		return nil, err
	}
	return quote, nil
}

// DeleteQuote deletes a quote.
func (s *quoteService) DeleteQuote(ctx context.Context, req *qs.DeleteQuoteRequest) (*qs.DeleteQuoteResponse, error) {
	err := s.database.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &qs.DeleteQuoteResponse{
		Success: true,
	}, nil
}

// ListQuotes lists all quotes.
func (s *quoteService) ListQuotes(ctx context.Context, _ *emptypb.Empty) (*qs.ListQuotesResponse, error) {
	quotes := s.database.List(ctx)
	return &qs.ListQuotesResponse{
		Quotes: quotes,
	}, nil
}
