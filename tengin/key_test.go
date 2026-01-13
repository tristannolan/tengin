package tengin_test

import (
	"testing"

	"github.com/tristannolan/tengin/tengin"
)

func TestKeyValue(t *testing.T) {
	tests := []struct {
		name string
		key  tengin.Key
		want string
	}{
		{
			name: "string key returns string",
			key:  tengin.NewStringKey("a"),
			want: "a",
		},
		{
			name: "special key returns name",
			key:  tengin.NewSpecialKey(tengin.KeyEnter),
			want: "Enter",
		},
		{
			name: "empty key returns Empty",
			key:  tengin.NewEmptyKey(),
			want: "Empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.key.Value(); got != tt.want {
				t.Fatalf("Value() = %q, want %q", got, tt.want)
			}
		})
	}
}
