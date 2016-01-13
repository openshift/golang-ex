package models

import "testing"

func TestArchive(t *testing.T) {
	addTestEventsToArchive()
	tests := []struct {
		lastReceivedTimestamp int
		numberOfEvents        int
	}{
		{0, 3},
		{1257894010, 0},
	}
	for _, test := range tests {
		got := len(GetEvents(test.lastReceivedTimestamp))
		if got != test.numberOfEvents {
			t.Errorf("len(GetEvents(%v)) = %v; want %v",
				test.lastReceivedTimestamp, got, test.numberOfEvents)
		}
	}
}

func addTestEventsToArchive() {
	events := []Event{
		{
			Type:      EVENT_JOIN,
			User:      "gopher",
			Timestamp: 1257894000,
		},
		{
			Type:      EVENT_MESSAGE,
			User:      "gopher",
			Timestamp: 1257894001,
			Content:   "Hello test!",
		},
		{
			Type:      EVENT_LEAVE,
			User:      "gopher",
			Timestamp: 1257894002,
		},
	}
	for _, e := range events {
		NewArchive(e)
	}
}
