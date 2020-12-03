package distributedLock

import "testing"

func TestDistributedLock(t *testing.T) {
	_, err := DistributedLock()
	if err != nil {
		t.Fatalf("DistributedLock failed: %s",err )
	}
}