package dictionary

type Dictionary map[string]string

// our errors will behave like sentinel values. best to have theem constant in practice, 
// so that their behavior is predictable and the messages meaningful
const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word that does not exist")
)
type DictionaryErr string

func (d Dictionary) Search(word string) (string,error){
	definition, ok := d[word] //returns value or boolean, false if not found

	if !ok{
		return "", ErrNotFound
	}

	return definition, nil
}

//Adds a word to our dictionary. Note that the receiver is not a pointer
// this is because the "map" being passed in is a pointer. Maps are not reference variables, they are pointers
// so they share the same semantics as a pointer would.
// always initialize them to an empty map or use the 'make' keyword
func (d Dictionary) Add(word, definition string ) error{
	_, err := d.Search(word)
	// a good safety net in case we encounter other errors.
	switch err {
		case ErrNotFound:
			d[word] = definition
		case nil: 
			return ErrWordExists
		default: return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrNotFound: 
		return ErrWordDoesNotExist
	case nil: 
		d[word] = definition
	default: 
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil: 
		delete(d, word)
	default:
		return err
	}
	return nil
}

// this method allows our custom dictionary error to implement the Error interface
// so it can be returned like an error and behaves like an error.
func (e DictionaryErr) Error() string{
	return string(e)
}