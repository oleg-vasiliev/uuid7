package uuid7_test

import (
	"testing"

	"github.com/oleg-vasiliev/uuid7"

	"github.com/google/uuid"
)

func BenchmarkGoogleUUID(b *testing.B) {
	for range b.N {
		uuid.New()
	}
}

func BenchmarkRandCreation(b *testing.B) {
	for range b.N {
		uuid7.MustNew()
	}
}
