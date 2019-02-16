package dictionary

var (
	ErrNotFound          = DictionaryErr("could not find the word your were looking for")
	ErrWordExists        = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExists = DictionaryErr("cannot update word because it does not exists")
)

type DictionaryErr string

type Dictionary map[string]string

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Update(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		dictionary[word] = definition
	default:
		return err
	}

	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
