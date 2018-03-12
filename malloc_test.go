package offheap_test

import (
	"testing"

	"github.com/genez/offheap"
)

func TestMalloc(t *testing.T) {

	writeme := []byte("hello memory mapped WORLD\n")

	mm2 := offheap.Malloc(10*1024, "")
	copy(mm2.Mem[0:26], writeme)
	mm2.Free()

}
