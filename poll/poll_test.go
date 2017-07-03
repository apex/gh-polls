package poll

import (
	"testing"

	"github.com/tj/assert"
)

// TODO: show count in svg
// TODO: omit voters on load

func TestPoll_Vote(t *testing.T) {
	p := Poll{
		ID: "foo",
	}

	t.Run("has not voted", func(t *testing.T) {
		assert.NoError(t, p.Remove(), "remove")
		assert.NoError(t, p.Create([]string{"Something"}), "create")

		assert.NoError(t, p.Vote("tobi", "Something"), "vote")
		assert.NoError(t, p.Vote("loki", "Something"), "vote")

		assert.NoError(t, p.Load(), "load")
		assert.Equal(t, "foo", p.ID)
		assert.Equal(t, 2, p.Votes, "votes")
		assert.Equal(t, []string{"loki", "tobi"}, p.Voters)
	})

	t.Run("has voted", func(t *testing.T) {
		assert.NoError(t, p.Remove(), "remove")
		assert.NoError(t, p.Create([]string{"Something", "Another"}), "create")

		assert.NoError(t, p.Vote("tobi", "Something"), "vote")
		assert.Equal(t, ErrAlreadyVoted, p.Vote("tobi", "Something"))
		assert.Equal(t, ErrAlreadyVoted, p.Vote("tobi", "Another"))

		assert.NoError(t, p.Vote("loki", "Another"), "vote")
		assert.Equal(t, ErrAlreadyVoted, p.Vote("loki", "Something"))

		assert.NoError(t, p.Vote("jane", "Something"), "vote")

		assert.NoError(t, p.Load(), "load")
		assert.Equal(t, "foo", p.ID)
		assert.Equal(t, 3, p.Votes, "votes")
		assert.Equal(t, []string{"jane", "loki", "tobi"}, p.Voters)
		assert.Equal(t, map[string]int{"Another": 1, "Something": 2}, p.Options)
	})
}
