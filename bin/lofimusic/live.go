package main

import (
	"path"
	"sort"
	"strings"
	"time"
)

type liveRadio struct {
	Slug    string
	Name    string
	Owner   string
	URL     string
	Cards   []string
	Links   []socialLink
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		{
			Slug:  "autumns-beats",
			Name:  "Autumn's Beats",
			Owner: "Lofi for Life",
			URL:   "https://youtu.be/RN0zF_Uh0nY",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "chilling-on-the-space",
			Name:  "Chilling on the Space",
			Owner: "Lofi for Life",
			URL:   "https://youtu.be/2tX5RpHxY74",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "deep-focus-for-studying",
			Name:  "Deep Focus For Studying",
			Owner: "MITON.W",
			URL:   "https://youtu.be/c9UK50nrN8c",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "disney-lofi",
			Name:  "Disney Lofi",
			Owner: "LlamaLoops",
			URL:   "https://youtu.be/upY2z9fMfVw",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "disney-relaxing-piano",
			Name:  "Disney Relaxing Piano",
			Owner: "kno Piano Music",
			URL:   "https://youtu.be/g8NVwN0_mks",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Harry Potter Lofi - cuDzX5jIcAE - Chill Astronaut
		{
			Slug:  "harry-potter-lofi",
			Name:  "Harry Potter Lofi",
			Owner: "Chill Astronaut",
			URL:   "https://youtu.be/cuDzX5jIcAE",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Lofi Hip Hop Mix - wzBn3gKbhxk - Lofi Cat
		{
			Slug:  "lofi-hip-hop-mix",
			Name:  "Lofi Hip Hop Mix",
			Owner: "Lofi Cat",
			URL:   "https://youtu.be/wzBn3gKbhxk",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Nintendo Lofi - 7JMvn0wfABQ - Helynt
		{
			Slug:  "nintendo-lofi",
			Name:  "Nintendo Lofi",
			Owner: "Helynt",
			URL:   "https://youtu.be/7JMvn0wfABQ",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "rainy-day-coffee-shop",
			Name:  "Rainy Day Coffee Shop",
			Owner: "Relaxing Jazz Piano",
			URL:   "https://youtu.be/0L38Z9hIi5s",
			Cards: []string{},
			Links: []socialLink{},
		},
		{
			Slug:  "relaxing-guitar-music",
			Name:  "Relaxing Guitar Music",
			Owner: "Yellow Brick Cinema",
			URL:   "https://youtu.be/iIjSS_MbCGg",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Relaxing Piano for Sleep - mo14InsSnIc - Tazad Beat
		{
			Slug:  "relaxing-piano-for-sleep",
			Name:  "Relaxing Piano for Sleep",
			Owner: "Tazad Beat",
			URL:   "https://youtu.be/mo14InsSnIc",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Star Wars Lofi - wv38obj0D_k - Chill Astronaut
		{
			Slug:  "star-wars-lofi",
			Name:  "Star Wars Lofi",
			Owner: "Chill Astronaut",
			URL:   "https://youtu.be/wv38obj0D_k",
			Cards: []string{},
			Links: []socialLink{},
		},
		//Taylor Swift Piano - teX3yyB9PQc - Juan L. Otaiza
		{
			Slug:  "taylor-swift-piano",
			Name:  "Taylor Swift Piano",
			Owner: "Juan L. Otaiza",
			URL:   "https://youtu.be/teX3yyB9PQc",
			Cards: []string{},
			Links: []socialLink{},
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	for _, r := range radios {
		sort.Slice(r.Links, func(a, b int) bool {
			return strings.Compare(r.Links[a].Slug, r.Links[b].Slug) < 0
		})
	}

	return radios
}

type socialLink struct {
	Slug string
	URL  string
}

func socialIcon(slug string) string {
	switch slug {
	case "youtube":
		return youtubeSVG

	case "reddit":
		return redditSVG

	case "facebook":
		return facebookSVG

	case "instagram":
		return instagramSVG

	case "twitter":
		return twitterSVG

	case "spotify":
		return spotifySVG

	case "discord":
		return discordSVG

	case "website":
		return websiteSVG

	default:
		return linkSVG
	}
}
