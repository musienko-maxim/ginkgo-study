package spec_context

import (
	"ginkgo/books"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Books", func() {
	It("can extract the author's last name", func() {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		Expect(book.AuthorLastName()).To(Equal("Hugo"))
	})

	It("can fetch a summary of the book from the library service", func(ctx SpecContext) {
		book = &books.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		summary, err := library.FetchSummary(ctx, book)
		Expect(err).NotTo(HaveOccurred())
		Expect(summary).To(ContainSubstring("Jean Valjean"))
	}, SpecTimeout(time.Second))
})
