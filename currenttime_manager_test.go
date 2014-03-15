package trade

import (
	"testing"
	"time"
)

func TestCurrentTimeManager(t *testing.T) {
	engine, err := NewEngine()

	if err != nil {
		t.Fatalf("cannot connect engine: %s", err)
	}
	defer engine.Stop()

	ctm, err := NewCurrentTimeManager(engine)
	defer ctm.Close()

	var m Manager = ctm
	SinkManagerTest(t, &m, 5*time.Second, 1)

	ctmTime := ctm.Time()
	t.Logf("got time: %s\n", ctmTime.String())

	if ctmTime.IsZero() {
		t.Fatal("Expected time to have been updated")
	}

	if ctmTime.Before(engine.serverTime) {
		t.Fatal("Expected time to be later than serverTime of %s, got: %s", engine.serverTime.String(), ctmTime.String())
	}

	if b, ok := <-ctm.Refresh(); ok {
		t.Fatal("Expected the refresh channel to be closed, but got %v", b)
	}

}
