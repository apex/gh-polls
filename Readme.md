# GitHub Polls

User polls for GitHub powered by [Up](https://github.com/apex/up).

[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20A)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20A/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20B)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20B/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20C)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2T00TMSDTZWJ1RP6TQF200/Option%20C/vote)

## About

These polls work by pasting individual markdown SVG images into your issue, each wrapped with a link that tracks a vote. A single vote per IP is allowed for a given poll, which are stored in DynamoDB.

What do they look like? The poll above is live, click one out and give it a try! Please don't abuse it or I'll have to take it down :).

## Installation

Via `go get`:

```
$ go get github.com/tj/gh-polls/cmd/polls
```

## Usage

Create a new poll for who is the best ferret.

```
$ polls new Tobi Loki Jane
```

[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Tobi)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Tobi/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Loki)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Loki/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Jane)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHFZXYKQV9N3HNFXCBH3N/Jane/vote)

Create a new poll for the best species:

```
$ polls new "Cats are cool" "Dogs are better" "Ferrets for the win"
```

[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Cats%20are%20cool)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Cats%20are%20cool/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Dogs%20are%20better)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Dogs%20are%20better/vote)
[![](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Ferrets%20for%20the%20win)](https://m131jyck4m.execute-api.us-west-2.amazonaws.com/prod/poll/01BM2ZHPN7BA19X15SQDGX4D88/Ferrets%20for%20the%20win/vote)

---

[![GoDoc](https://godoc.org/github.com/tj/gh-polls?status.svg)](https://godoc.org/github.com/tj/gh-polls)
![](https://img.shields.io/badge/license-MIT-blue.svg)
![](https://img.shields.io/badge/status-experimental-orange.svg)

<a href="https://apex.sh"><img src="http://tjholowaychuk.com:6000/svg/sponsor"></a>
