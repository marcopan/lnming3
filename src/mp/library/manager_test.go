package library

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("New Manager Music failed.")
	}

	if mm.Len() != 0 {
		t.Error("New Manager failed, not empty.")
	}

	m0 := &MusicEntry{"1", "my heart",
		"pop", "source", "mp3"}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("Music manager Add failed.")
	}

	m := mm.Find(m0.Name)
	if mm == nil {
		t.Error("Music Manager .Find() failed.")
	}
	if m.Id != m0.Id || m.Artist != m0.Artist {
		t.Error("Music Manager find failed.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("Music Manager Get failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("Music Manager Remove failed.", err)
	}
}
