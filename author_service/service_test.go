package main

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	db "github.com/wcodesoft/mosha-backed/common/database"
	as "github.com/wcodesoft/mosha-backed/common/protos/authorservice"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestAuthorService(t *testing.T) {
	Convey("Given a new author service", t, func() {
		ctx := context.Background()
		database := db.NewMemoryDatabase[as.Author]()
		s := NewAuthorService(database)

		Convey("When a new author is created", func() {
			req := &as.CreateAuthorRequest{
				Author: &as.Author{
					Id:     "1",
					Name:   "John Doe",
					PicUrl: "https://example.com/johndoe.jpg",
				},
			}
			r, err := s.CreateAuthor(ctx, req)
			So(err, ShouldBeNil)
			So(r, ShouldResemble, req.Author)

			Convey("When the same author is created again", func() {
				r, err := s.CreateAuthor(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})
		})

		Convey("With an existing author", func() {
			carr := &as.CreateAuthorRequest{
				Author: &as.Author{
					Id:     "1",
					Name:   "John Doe",
					PicUrl: "https://example.com/johndoe.jpg",
				},
			}
			_, err := s.CreateAuthor(ctx, carr)
			So(err, ShouldBeNil)

			Convey("When the author is retrieved", func() {
				req := &as.GetAuthorRequest{Id: "1"}
				r, err := s.GetAuthor(ctx, req)
				So(err, ShouldBeNil)
				So(r, ShouldResemble, carr.Author)
			})

			Convey("When a non-existing author is retrieved", func() {
				req := &as.GetAuthorRequest{Id: "2"}
				r, err := s.GetAuthor(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})

			Convey("When the author is updated", func() {
				ur := &as.UpdateAuthorRequest{
					Author: &as.Author{
						Id:     "1",
						Name:   "Jane Doe",
						PicUrl: "https://example.com/janedoe.jpg",
					},
				}
				r, err := s.UpdateAuthor(ctx, ur)
				So(err, ShouldBeNil)
				So(r, ShouldResemble, ur.Author)
			})

			Convey("When a non-existing author is updated", func() {
				ur := &as.UpdateAuthorRequest{
					Author: &as.Author{
						Id:     "2",
						Name:   "Jane Doe",
						PicUrl: "https://example.com/janedoe.jpg",
					},
				}
				r, err := s.UpdateAuthor(ctx, ur)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})

			Convey("When the author is deleted", func() {
				req := &as.DeleteAuthorRequest{Id: "1"}
				r, err := s.DeleteAuthor(ctx, req)
				So(err, ShouldBeNil)
				So(r, ShouldResemble, &as.DeleteAuthorResponse{Success: true})

				Convey("When the same author is deleted again", func() {
					r, err := s.DeleteAuthor(ctx, req)
					So(err, ShouldNotBeNil)
					So(r, ShouldBeNil)
				})
			})

			Convey("When a non-existing author is deleted", func() {
				req := &as.DeleteAuthorRequest{Id: "2"}
				r, err := s.DeleteAuthor(ctx, req)
				So(err, ShouldNotBeNil)
				So(r, ShouldBeNil)
			})
		})

		Convey("With multiple authors", func() {
			a1 := &as.CreateAuthorRequest{
				Author: &as.Author{
					Id:     "1",
					Name:   "John Doe",
					PicUrl: "https://example.com/johndoe.jpg",
				},
			}
			a2 := &as.CreateAuthorRequest{
				Author: &as.Author{
					Id:     "2",
					Name:   "Jane Doe",
					PicUrl: "https://example.com/janedoe.jpg",
				},
			}
			_, err := s.CreateAuthor(ctx, a1)
			So(err, ShouldBeNil)
			_, err = s.CreateAuthor(ctx, a2)
			So(err, ShouldBeNil)

			Convey("When all authors are listed", func() {
				req := &emptypb.Empty{}
				r, err := s.ListAuthors(ctx, req)
				So(err, ShouldBeNil)
				So(r.Authors, ShouldHaveLength, 2)
				So(r.Authors[0], ShouldResemble, a1.Author)
				So(r.Authors[1], ShouldResemble, a2.Author)
			})
		})
	})
}
