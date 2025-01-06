package t_test

import (
	"testing"

	tt "github.com/mychiux413/goutils/types"
	"github.com/stretchr/testify/assert"
)

func TestUniques(t *testing.T) {
	assert := assert.New(t)

	arr1 := tt.IDArray{1, 2, 3, 1, 2, 3}
	assert.Len(arr1.Unique(), 3)
	arr2 := tt.BigIDArray{1, 2, 3, 1, 2, 3}
	assert.Len(arr2.Unique(), 3)
	arr3 := tt.SmallIDArray{1, 2, 3, 1, 2, 3}
	assert.Len(arr3.Unique(), 3)
	arr4 := tt.TinyIDArray{1, 2, 3, 1, 2, 3}
	assert.Len(arr4.Unique(), 3)
	arr5 := tt.IPArray{"127.0.0.1", "127.0.0.1", "127.0.0.2", "127.0.0.2", "127.0.0.3", "127.0.0.3"}
	assert.Len(arr5.Unique(), 3)
	arr6 := tt.TextArray{"127.0.0.1", "127.0.0.1", "127.0.0.2", "127.0.0.2", "127.0.0.3", "127.0.0.3"}
	assert.Len(arr6.Unique(), 3)
}
