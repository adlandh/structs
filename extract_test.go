package structs_test

import (
	"structs"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

type ID struct {
	Number int
	Issuer string
}

type Person struct {
	ID
	Age  int
	Name string
}

type Address string

type Employee struct {
	Person
	Address
	Company string
}

type Company struct {
	Name string
}

func TestExtractEmbedValue(t *testing.T) {
	id := ID{
		Number: int(gofakeit.Uint16()),
		Issuer: gofakeit.Name(),
	}
	person := Person{
		ID:   id,
		Age:  gofakeit.IntRange(18, 100),
		Name: gofakeit.Name(),
	}

	address := Address(gofakeit.Address().Address)

	employee := Employee{
		Person:  person,
		Address: address,
		Company: gofakeit.Company(),
	}

	t.Run("extract value of embed type Person from Employee", func(t *testing.T) {
		val, ok := structs.ExtractEmbedValue[Person](employee)
		require.True(t, ok)
		require.Equal(t, person, val)
	})

	t.Run("extract value of embed type ID from Employee", func(t *testing.T) {
		val, ok := structs.ExtractEmbedValue[ID](employee)
		require.True(t, ok)
		require.Equal(t, id, val)
	})

	t.Run("extract value of embed type Address from Employee", func(t *testing.T) {
		val, ok := structs.ExtractEmbedValue[Address](employee)
		require.True(t, ok)
		require.Equal(t, address, val)
	})

	t.Run("extract value of not embed type Company from Employee", func(t *testing.T) {
		val, ok := structs.ExtractEmbedValue[Company](employee)
		require.False(t, ok)
		require.Empty(t, val)
	})

	t.Run("extract value of string from non struct", func(t *testing.T) {
		val, ok := structs.ExtractEmbedValue[string]([]string{"test"})
		require.False(t, ok)
		require.Empty(t, val)
	})
}

func BenchmarkExtractEmbedValue(b *testing.B) {
	id := ID{
		Number: int(gofakeit.Uint16()),
		Issuer: gofakeit.Name(),
	}
	person := Person{
		ID:   id,
		Age:  gofakeit.IntRange(18, 100),
		Name: gofakeit.Name(),
	}

	address := Address(gofakeit.Address().Address)

	employee := Employee{
		Person:  person,
		Address: address,
		Company: gofakeit.Company(),
	}
	b.ResetTimer()
	b.Run("extract value of embed type Person from Employee", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = structs.ExtractEmbedValue[Person](employee)
		}
	})

	b.Run("extract value of embed type ID from Employee", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = structs.ExtractEmbedValue[ID](employee)
		}
	})
}
