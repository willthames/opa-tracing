package spans

import (
	"testing"
	"time"
)

func TestJSONUnmarshal(t *testing.T) {
	b := []byte(`{"id":"hello","name":"bowser","timestamp":1480979203000000,"binary_annotations":[{"key":"hello","value":"world"}],"duration":1000}`)
	span := new(Span)
	err := span.UnmarshalJSON(b)
	if err != nil {
		t.Errorf("Failed to unmarshal span from json data %s: %v", string(b), err)
	}
	if span.Duration != 1000*time.Microsecond {
		t.Errorf("span duration was incorrectly parsed")
	}
	value, ok := span.BinaryAnnotations["hello"]
	if !ok || value != "world" {
		t.Errorf("binary annotations incorrectly parsed")
	}
}