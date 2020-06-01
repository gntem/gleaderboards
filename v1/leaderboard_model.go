package leaderboard

// Leaderboard struct holds entries
type Leaderboard struct {
	entries []Score
}

// NewLeaderboard func create new leaderboard
func NewLeaderboard() *Leaderboard {
	l := &Leaderboard{}
	return l
}

// HasScore method check if leaderboard has score for entityID
func (l *Leaderboard) HasScore(sc Score) *Score {
	for _, entry := range l.entries {
		if entry.entityID == sc.entityID {
			return &entry
		}
	}
	return nil
}

// UpsertScore func updates or creates score entry
func (l *Leaderboard) UpsertScore(sc Score) {
	e := l.HasScore(sc)
	if e == nil {
		l.entries = append(l.entries, sc)
		return
	}
	e.value = sc.value
	return
}
