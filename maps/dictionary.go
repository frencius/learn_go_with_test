package dictionary

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) string {
	return word
}
