package swagger

type InMemoryClientRepository struct {
	clients map[int64]*Client
	id      int64
}

func NewInMemoryRepo() *InMemoryClientRepository {
	repository := InMemoryClientRepository{make(map[int64]*Client), 0}
	return &repository
}

func (repository *InMemoryClientRepository) Create(employee *Client) {
	repository.id++
	currentEmployee := employee
	currentEmployee.Id = repository.id

	repository.clients[repository.id] = employee
}

func (repository *InMemoryClientRepository) Update(employee *Client) {
	repository.clients[employee.Id] = employee
}

func (repository *InMemoryClientRepository) Delete(id int64) {
	delete(repository.clients, id)
}

func (repository *InMemoryClientRepository) FindById(id int64) *Client {
	return repository.clients[id]
}

func (repository *InMemoryClientRepository) FindAll() []*Client {
	result := make([]*Client, 0)
	var i int64 = 0
	for ; i <= repository.id; i++ {
		element, ok := repository.clients[i]
		if ok {
			result = append(result, element)
		}
	}
	return result
}
