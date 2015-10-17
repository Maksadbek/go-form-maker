package form

import (
	"testing"
)

func TestFormCreate(t *testing.T) {
	f := MyForm{}
	str, err := FormCreate(&f)
	if err != nil {
		t.Error(err)
	}
	t.Log(str)
}
