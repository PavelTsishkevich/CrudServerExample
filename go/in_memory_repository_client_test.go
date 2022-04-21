package swagger

import (
	"testing"
)

type addTest struct {
	id       int64
	expected *Client
}

var addTests = []addTest{
	{1, &Client{1, "Micky", "Mik", "First", "email", "123"}},
	{2, &Client{2, "Rrr", "aaa", "First", "email", "123"}},
	{3, nil},
}

func getFilledRepository() *InMemoryClientRepository {
	repository := NewInMemoryRepo()
	repository.Create(addTests[0].expected)
	repository.Create(addTests[1].expected)
	return repository
}

func TestCreate(t *testing.T) {
	repository := NewInMemoryRepo()

	if len(repository.FindAll()) != 0 {
		t.Errorf("Output should be empty but shouldn't")
	}

	repository.Create(addTests[0].expected)
	if len(repository.FindAll()) != 1 {
		t.Errorf("Output should contain one element but shouldn't")
	}

	if output := repository.FindAll()[0]; output != addTests[0].expected {
		t.Errorf("Output contains %v but it should contain %v", output, addTests[0].expected)
	}
}

func TestUpdate(t *testing.T) {
	repository := getFilledRepository()

	updatedClient := Client{2, "Rrr2", "aaa2", "First2", "email2", "1232"}
	repository.Update(&updatedClient)

	if output := repository.FindById(2); output != &updatedClient {
		t.Errorf("Output %v not equal to expected %v", output, updatedClient)
	}

	if output := repository.FindById(1); output == &updatedClient {
		t.Errorf("Output %v equal to expected %v but should not", output, updatedClient)
	}
}

func TestDelete(t *testing.T) {
	repository := getFilledRepository()

	repository.Delete(1)

	if output := repository.FindById(1); output != nil {
		t.Errorf("Output %v is found but shouldn't", output)
	}

	if output := repository.FindById(2); output == nil {
		t.Errorf("Output %v is not found but should be", output)
	}
}

func TestFindById(t *testing.T) {
	repository := getFilledRepository()

	for _, test := range addTests {
		if output := repository.FindById(test.id); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}

func TestFindAll(t *testing.T) {
	repository := NewInMemoryRepo()

	if len(repository.FindAll()) != 0 {
		t.Errorf("Output should be empty but shouldn't")
	}

	repository.Create(addTests[0].expected)
	if len(repository.FindAll()) != 1 {
		t.Errorf("Output should contain one element but shouldn't")
	}

	repository.Create(addTests[1].expected)
	if len(repository.FindAll()) != 2 {
		t.Errorf("Output should contain two elements but shouldn't")
	}
}
