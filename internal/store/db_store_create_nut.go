package store

import (
	"crypto/rand"
	"errors"
	"strconv"
)

// ErrConflict ошибка, которую используем для сигнала о нарушении целостности данных.
var ErrConflict = errors.New("data conflict")

func generateRandom(size int) ([]byte, error) {
	// генерируем криптостойкие случайные байты в b
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func packBytesToString(b []byte) string {
	var result string
	for i := 0; i < len(b); i++ {
		intVal := int(b[i])
		strVal := strconv.Itoa(intVal)
		result += strVal + "!"
	}
	return result
}
