package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/bmizerany/pat"
	"github.com/gohttp/response"
	"github.com/segmentio/go-env"

	"github.com/tj/gh-polls/internal/poll"
)

func main() {
	app := pat.New()
	app.Post("/poll", http.HandlerFunc(addPoll))
	app.Get("/poll/:id/:option", http.HandlerFunc(getPollOption))
	app.Get("/poll/:id/:option/vote", http.HandlerFunc(getPollOptionVote))
	addr := env.MustGet("UP_ADDR")
	if err := http.ListenAndServe(addr, app); err != nil {
		log.WithError(err).Fatal("binding")
	}
}

// addPoll creates a poll, responds with .id.
func addPoll(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-Real-IP")

	var body struct {
		Options []string `json:"options"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.WithError(err).Error("parsing body")
		response.BadRequest(w, "Malformed request body.")
		return
	}

	p := poll.New(user, body.Options)

	if err := p.Create(); err != nil {
		log.WithError(err).Error("creating poll")
		response.InternalServerError(w)
		return
	}

	response.OK(w, map[string]string{
		"id": p.ID,
	})
}

// getPollOptionVote performs a vote.
func getPollOptionVote(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	option := r.URL.Query().Get(":option")
	user := r.Header.Get("X-Real-IP")

	ctx := log.WithFields(log.Fields{
		"id":     id,
		"option": option,
		"user":   user,
	})

	p := poll.Poll{
		ID: id,
	}

	err := p.Vote(user, option)

	if err == poll.ErrAlreadyVoted {
		ctx.WithError(err).Warn("already voted")
		response.BadRequest(w, "Cheater!")
		return
	}

	if err != nil {
		ctx.WithError(err).Error("voting")
		response.InternalServerError(w, "Error voting.")
		return
	}

	http.ServeFile(w, r, "static/voted.html")
}

// getPollOption responds with a poll option svg.
func getPollOption(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	option := r.URL.Query().Get(":option")

	ctx := log.WithFields(log.Fields{
		"id":     id,
		"option": option,
	})

	p := poll.Poll{
		ID: id,
	}

	if err := p.Load(); err != nil {
		ctx.WithError(err).Error("loading poll")
		response.InternalServerError(w, "Error loading poll.")
		return
	}

	votes, ok := p.Options[option]
	if !ok {
		ctx.Warn("option does not exist")
		response.NotFound(w, "Option does not exist.")
		return
	}

	barWidth := 334
	percent := 0
	width := 0

	if p.Votes > 0 {
		percent = int(float64(votes) / float64(p.Votes) * 100)
		width = int(float64(barWidth) * (float64(votes) / float64(p.Votes)))
	}

	opt := poll.Option{
		Name:    option,
		Votes:   votes,
		Percent: percent,
		Width:   width,
	}

	b, err := opt.Render()
	if err != nil {
		http.Error(w, "Error rendering poll option.", http.StatusInternalServerError)
		return
	}

	setETag(w, b)
	setCacheControl(w)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(b)
}

func setCacheControl(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "private")
}

func setETag(w http.ResponseWriter, body []byte) {
	hash := md5.New()
	hash.Write(body)
	etag := hex.EncodeToString(hash.Sum(nil))
	w.Header().Set("ETag", `w/"`+etag+`"`)
}
