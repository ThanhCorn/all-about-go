package linq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestFilter_Where(t *testing.T) {
// 	type TestStruct struct {
// 		Name string
// 		Age  int
// 	}

// 	arr := []TestStruct{
// 		{Name: "John", Age: 20},
// 		{Name: "Doe", Age: 30},
// 		{Name: "Jane", Age: 25},
// 	}

// 	result := AsQueryable(arr).Where(func(v TestStruct) bool {
// 		return v.Age > 20
// 	}).ToList()

// 	nameResult := Map(result, func(v TestStruct) *string {
// 		return &v.Name
// 	})
// 	assert.Equal(t, 2, len(result))
// 	assert.Contains(t, nameResult, "Doe")
// 	assert.Contains(t, nameResult, "Jane")
// }

func TestSumBy(t *testing.T) {
	type Hello struct {
		ID     int64
		Amount float64
	}

	arr := []Hello{
		{ID: 1, Amount: 1000},
		{ID: 2, Amount: 1500},
		{ID: 3, Amount: 2000},
		{ID: 4, Amount: 2500},
	}

	sum := SumBy(arr, func(h Hello) float64 {
		return h.Amount
	})

	assert.Equal(t, float64(7000), sum)
}
