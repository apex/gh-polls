package poll

import (
	"testing"

	"github.com/tj/assert"
)

// TODO: show count in svg
// TODO: omit voters on load

func TestPoll_Vote(t *testing.T) {
	p := New("tobi", []string{"Option A", "Option B"})
	defer p.Remove()

	assert.NotEmpty(t, p.ID, "id missing")

	t.Run("has not voted", func(t *testing.T) {
		assert.NoError(t, p.Remove(), "remove")
		assert.NoError(t, p.Create(), "create")

		assert.NoError(t, p.Vote("tobi", "Option A"), "vote")
		assert.NoError(t, p.Vote("loki", "Option A"), "vote")

		assert.NoError(t, p.Load(), "load")
		assert.Equal(t, 2, p.Votes, "votes")
		assert.Equal(t, []string{"loki", "tobi"}, p.Voters)
	})

	t.Run("has voted", func(t *testing.T) {
		assert.NoError(t, p.Remove(), "remove")
		assert.NoError(t, p.Create(), "create")

		// tobi tries three times! abuse!
		assert.NoError(t, p.Vote("tobi", "Option A"), "vote")
		assert.Equal(t, ErrAlreadyVoted, p.Vote("tobi", "Option A"))
		assert.Equal(t, ErrAlreadyVoted, p.Vote("tobi", "Option B"))

		// loki tries twice
		assert.NoError(t, p.Vote("loki", "Option B"), "vote")
		assert.Equal(t, ErrAlreadyVoted, p.Vote("loki", "Option A"))

		// jane is cool
		assert.NoError(t, p.Vote("jane", "Option B"), "vote")

		assert.NoError(t, p.Load(), "load")
		assert.Equal(t, 3, p.Votes, "votes")
		assert.Equal(t, []string{"jane", "loki", "tobi"}, p.Voters)
		assert.Equal(t, map[string]int{"Option A": 1, "Option B": 2}, p.Options)
	})
}
