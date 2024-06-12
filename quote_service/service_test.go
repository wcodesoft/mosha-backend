package main

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	db "github.com/wcodesoft/mosha-backed/common/database"
	qs "github.com/wcodesoft/mosha-backed/common/protos/quoteservice"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestQuoteService(t *testing.T) {

	Convey("Give a new quote service", t, func() {
		ctx := context.Background()
		database := db.NewMemoryDatabase[qs.Quote]()
		s := NewQuoteService(database)

		Convey("When a new quote is created", func() {
			req := &qs.CreateQuoteRequest{
				Quote: &qs.Quote{
					Id:        "1",
					AuthorId:  "1",
					Text:      "Hello, World!",
					Timestamp: 10,
				},
			}

			r, err := s.CreateQuote(ctx, req)
			So(err, ShouldBeNil)
			So(r, ShouldResemble, req.Quote)

			Convey("When the same quote is created again", func() {
				r, err := s.CreateQuote(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})
		})

		Convey("With an existing quote", func() {
			cqr := &qs.CreateQuoteRequest{
				Quote: &qs.Quote{
					Id:        "1",
					AuthorId:  "1",
					Text:      "Hello, World!",
					Timestamp: 10,
				},
			}
			_, err := s.CreateQuote(ctx, cqr)
			So(err, ShouldBeNil)

			Convey("When the quote is retrieved", func() {
				req := &qs.GetQuoteRequest{Id: "1"}
				r, err := s.GetQuote(ctx, req)
				So(err, ShouldBeNil)
				So(r, ShouldResemble, cqr.Quote)
			})

			Convey("When a non-existing quote is retrieved", func() {
				req := &qs.GetQuoteRequest{Id: "2"}
				r, err := s.GetQuote(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})

			Convey("When the quote is updated", func() {
				req := &qs.UpdateQuoteRequest{
					Quote: &qs.Quote{
						Id:        "1",
						AuthorId:  "2",
						Text:      "Hello, Universe!",
						Timestamp: 20,
					},
				}
				r, err := s.UpdateQuote(ctx, req)
				So(err, ShouldBeNil)
				So(r, ShouldResemble, req.Quote)
			})

			Convey("When a non-existing quote is updated", func() {
				req := &qs.UpdateQuoteRequest{
					Quote: &qs.Quote{
						Id:        "2",
						AuthorId:  "2",
						Text:      "Hello, Universe!",
						Timestamp: 20,
					},
				}
				r, err := s.UpdateQuote(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})

			Convey("When the quote is deleted", func() {
				req := &qs.DeleteQuoteRequest{Id: "1"}
				r, err := s.DeleteQuote(ctx, req)
				So(err, ShouldBeNil)
				So(r.Success, ShouldBeTrue)
			})

			Convey("When non-existing quote is deleted", func() {
				req := &qs.DeleteQuoteRequest{Id: "2"}
				r, err := s.DeleteQuote(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})
		})

		Convey("With multiple quotes", func() {
			cqr1 := &qs.CreateQuoteRequest{
				Quote: &qs.Quote{
					Id:        "1",
					AuthorId:  "1",
					Text:      "Hello, World!",
					Timestamp: 10,
				},
			}
			_, err := s.CreateQuote(ctx, cqr1)
			So(err, ShouldBeNil)

			cqr2 := &qs.CreateQuoteRequest{
				Quote: &qs.Quote{
					Id:        "2",
					AuthorId:  "1",
					Text:      "Hello, Universe!",
					Timestamp: 20,
				},
			}
			_, err = s.CreateQuote(ctx, cqr2)
			So(err, ShouldBeNil)

			Convey("When all quotes are listed", func() {
				req := &emptypb.Empty{}
				r, err := s.ListQuotes(ctx, req)
				So(err, ShouldBeNil)
				So(r.Quotes, ShouldHaveLength, 2)
				So(r.Quotes[0], ShouldResemble, cqr1.Quote)
				So(r.Quotes[1], ShouldResemble, cqr2.Quote)
			})
		})
	})
}
