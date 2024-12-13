package uuid7_test

import (
	"testing"
	"time"

	"github.com/oleg-vasiliev/uuid7"

	"github.com/stretchr/testify/assert"
)

func TestTimestamps(t *testing.T) {
	tests := []struct {
		name      string
		timestamp time.Time
	}{
		{
			name:      "valid time in and out",
			timestamp: time.Now(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, err := uuid7.NewWithTime(tt.timestamp)
			assert.NoError(t, err)
			assert.Equal(t, tt.timestamp.UnixMilli(), ts.Time().UnixMilli())
		})
	}
}
