package main


type Store interface {
	Put(key string, value string)
	Get(key string) (string, bool)
	Delete(key string)
}

type Entry struct {
	Value string
}

type InMemoryStore struct {
	data map[string]Entry
}

func INewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]Entry),
	}
}

func (kvs *InMemoryStore) Put(key string, value string) {
	kvs.data[key] = Entry{Value: value}
}

func (kvs *InMemoryStore) Get(key string) (string, bool) {
	entry, exists := kvs.data[key]
	if !exists {
		return "", false
	}
	return entry.Value, true
}

func (kvs *InMemoryStore) Delete(key string) {
	delete(kvs.data, key)
}
