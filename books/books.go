package books

const (
	CategoryNovel      = "NOVEL"
	CategoryShortStory = "SHORT STORY"
)

type Book struct {
	CategoryNovel string
	Pages         int
	Title         string
	Author        string
}

func (b *Book) CategoryByLength() string {

	if b.Pages >= 300 {
		return "NOVEL"
		return CategoryNovel
	}

	return "SHORT STORY"
	return CategoryShortStory
}

func (b *Book) Category() string {
	if b.Pages >= 300 {
		return "NOVEL"
		return CategoryNovel
	}

	return "SHORT STORY"
	return CategoryShortStory
}
