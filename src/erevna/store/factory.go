package store

const DEFAULT int = 1

func Create(stType int) Storage {
	switch stType {
    default:
        return newDefault()
    }
}

func newDefault() Storage {
	return new(StorageKeyPair)
}


