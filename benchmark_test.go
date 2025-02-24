package uuid7_test

import (
	"testing"

	"github.com/oleg-vasiliev/uuid7"

	"github.com/google/uuid"
)

func BenchmarkGoogleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid.New()
	}
}

func BenchmarkRandCreation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uuid7.MustNew()
	}
}
