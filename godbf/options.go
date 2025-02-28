package godbf

type OptionFn func(*DbfTable)

func WithIncreasedDatastore(increasedSize int) OptionFn {
	return func(dt *DbfTable) {
		newStore := make([]byte, 0, cap(dt.dataStore)+increasedSize)
		dt.dataStore = append(newStore, dt.dataStore...)
	}
}

func NoTrimSpace() OptionFn {
	return func(dt *DbfTable) {
		dt.noTrimSpace = true
	}
}
