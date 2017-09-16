# GitHub Polls

User polls for GitHub powered by [Up](https://github.com/apex/up).

[![](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20A)](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20A/vote)
[![](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20B)](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20B/vote)
[![](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20C)](https://api.gh-polls.com/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20C/vote)

## About

These polls work by pasting individual markdown SVG images into your issue, each wrapped with a link that tracks a vote. A single vote per IP is allowed for a given poll, which are stored in DynamoDB.

What do they look like? The poll above is live, click one out and give it a try! You can create your own [online](https://app.gh-polls.com/) or via the `polls` CLI.

## Installation

Via `go get`:

```
$ go get github.com/apex/gh-polls/cmd/polls
```

Then create your own poll with `polls new`, the markdown is copied to your clipboard.

## Usage

Create a new poll for who is the best ferret.

```
$ polls new Tobi Loki Jane
```

[![](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Tobi)](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Tobi/vote)
[![](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Loki)](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Loki/vote)
[![](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Jane)](https://api.gh-polls.com/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Jane/vote)

Create a new poll for the best species:

```
$ polls new "Cats are cool" "Dogs are better" "Ferrets for the win"
```

[![](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Cats%20are%20cool)](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Cats%20are%20cool/vote)
[![](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Dogs%20are%20better)](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Dogs%20are%20better/vote)
[![](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Ferrets%20for%20the%20win)](https://api.gh-polls.com/poll/01BM2ZHPN7BA19X15SQDGX4D88/Ferrets%20for%20the%20win/vote)

## Statistics

GitHub poll API statistics powered by [Up](https://github.com/apex/up).

![](https://q3qxtefzqa.execute-api.us-west-2.amazonaws.com/production/timeseries/start:-1M/title:Requests/max-points:750)

![](https://q3qxtefzqa.execute-api.us-west-2.amazonaws.com/production/timeseries/title:Latency/start:-1M/metric:Latency/stat:Average/x-suffix:%20ms/max-points:750)

---

[![GoDoc](https://godoc.org/github.com/apex/gh-polls?status.svg)](https://godoc.org/github.com/apex/gh-polls)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-experimental-orange.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
