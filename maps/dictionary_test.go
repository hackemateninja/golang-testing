package maps

import "testing"

const (
	TEST_DEF  = "this is just a test"
	TEST_WORD = "test"
)

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, TEST_WORD)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{TEST_WORD: TEST_DEF}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search(TEST_WORD)
		want := TEST_DEF

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add(TEST_WORD, TEST_DEF)

		assertError(t, err, nil)

		assertDefinition(t, dictionary, TEST_WORD, TEST_DEF)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{TEST_WORD: TEST_DEF}

		err := dictionary.Add(TEST_WORD, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, TEST_WORD, TEST_DEF)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{TEST_WORD: TEST_DEF}

		newDef := "new Definition"

		err := dictionary.Update(TEST_WORD, newDef)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, TEST_WORD, newDef)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update(TEST_WORD, TEST_DEF)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{TEST_WORD: "test definition"}

	dictionary.Delete(TEST_WORD)

	_, err := dictionary.Search(TEST_WORD)

	assertError(t, err, ErrNotFound)
}
