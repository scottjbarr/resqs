package resqs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorker(t *testing.T) {
	c := Config{}
	w := NewWorker(c)

	assert.Equal(t, c, w.Config)
	assert.Equal(t, 36, len(w.ID))

	assert.Nil(t, w.Register())

	m, err := w.Put(`{"id": "42", "data": "towel"}`)

	assert.NotNil(t, m)
	assert.Nil(t, err)

	ch := make(chan Message)
	go func(ch chan Message) {
		if err := w.Get(ch); err != nil {
			panic(err)
		}
	}(ch)

	fmt.Printf("%+v\n", <-ch)
}
