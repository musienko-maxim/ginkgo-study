package books

type Book struct {
	Title         string
	Author        string
	Pages         int
	CategoryNovel string
}

func (b *Book) CategoryByLength() string {

	if b.Pages >= 300 {
		return "NOVEL"
	}

	return "SHORT STORY"
}

func (b *Book) Category() string {

	novel := b.CategoryNovel
	return novel
}
