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
		// {
		// 	Slug:  "autumns-beats",
		// 	Name:  "Autumn's Beats",
		// 	Owner: "Lofi for Life",
		// 	URL:   "https://youtu.be/RN0zF_Uh0nY",
		// 	Cards: []string{},
		// 	Links: []socialLink{},
		// },
		{
			Slug:  "chilling-on-the-space",
			Name:  "Chilling on the Space",
			Owner: "Lofi for Life",
			URL:   "https://youtu.be/2tX5RpHxY74",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Celtic Study Music https://www.youtube.com/watch?v=4ICggPlpRQA
		{
			Slug:  "celtic-study-music",
			Name:  "Celtic Study Music",
			Owner: "Visual Melodies",
			URL:   "https://youtu.be/4ICggPlpRQA",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Charlie Brown https://www.youtube.com/watch?v=xj0KQijdGEs
		{
			Slug:  "charlie-brown",
			Name:  "Charlie Brown",
			Owner: "A Worker's Playlist",
			URL:   "https://youtu.be/xj0KQijdGEs",
			Cards: []string{},
			Links: []socialLink{},
		},
		// //Christmas Fireplace - kszQ_pb1lU8 - OCB Relax Music
		// {
		// 	Slug:  "christmas-fireplace",
		// 	Name:  "Christmas Fireplace",
		// 	Owner: "OCB Relax Music",
		// 	URL:   "https://youtu.be/kszQ_pb1lU8",
		// 	Cards: []string{},
		// 	Links: []socialLink{},
		// },
		// //Christmas Instrumental - 0cGLoDhfaWs - Cozy Moments
		// {
		// 	Slug:  "christmas-instrumental",
		// 	Name:  "Christmas Instrumental",
		// 	Owner: "Cozy Moments",
		// 	URL:   "https://youtu.be/0cGLoDhfaWs",
		// 	Cards: []string{},
		// 	Links: []socialLink{},
		// },
		// //Christmas Jazz - B7nut0m6HBQ - Cozy Coffee Shop
		// {
		// 	Slug:  "christmas-jazz",
		// 	Name:  "Christmas Jazz",
		// 	Owner: "Cozy Coffee Shop",
		// 	URL:   "https://youtu.be/B7nut0m6HBQ",
		// 	Cards: []string{},
		// 	Links: []socialLink{},
		// },
		// Cozy winter https://youtu.be/_tV5LEBDs7w?si=IvtsBJ5Jr8SIUZFy
		{
			Slug:  "cozy-winter",
			Name:  "Cozy Winter",
			Owner: "Lofi Girl",
			URL:   "https://youtu.be/_tV5LEBDs7w",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Disney Jazz https://www.youtube.com/watch?v=UQCCwpiFEkI
		{
			Slug:  "disney-jazz",
			Name:  "Disney Jazz",
			Owner: "Massimo Roberti",
			URL:   "https://youtu.be/UQCCwpiFEkI",
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
		// Ocean Lofi https://www.youtube.com/watch?v=YJEnhffr5Vg
		{
			Slug:  "ocean-lofi",
			Name:  "Ocean Lofi",
			Owner: "Chill Village",
			URL:   "https://youtu.be/YJEnhffr5Vg",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Ocean Piano https://www.youtube.com/watch?v=QR3lp0ptpy8
		{
			Slug:  "ocean-piano",
			Name:  "Ocean Piano",
			Owner: "Dream Sounds",
			URL:   "https://youtu.be/QR3lp0ptpy8",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Pirate Study Music https://www.youtube.com/watch?v=EU8Zn5rHIDI
		{
			Slug:  "pirate-study-music",
			Name:  "Pirate Study Music",
			Owner: "Blessed",
			URL:   "https://youtu.be/EU8Zn5rHIDI",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Pirate https://youtu.be/ttZ1zQqMY6U?feature=shared
		{
			Slug:  "pirate",
			Name:  "Pirates of the Caribbean",
			Owner: "Ambient Worlds",
			URL:   "https://youtu.be/ttZ1zQqMY6U",
			Cards: []string{},
			Links: []socialLink{},
		},
		// Relaxing Spring Music https://www.youtube.com/watch?v=u4pQ7O8uQRw
		{
			Slug:  "relaxing-spring-music",
			Name:  "Relaxing Spring Music",
			Owner: "Soothing Relaxation",
			URL:   "https://youtu.be/u4pQ7O8uQRw",
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
		// Spring flowers music https://youtu.be/J4Fk0pujVFA?si=kTDhvdmsiNOZfOvx
		{
			Slug:  "spring-flowers-music",
			Name:  "Spring Flowers Music",
			Owner: "Relaxing Music Films",
			URL:   "https://youtu.be/J4Fk0pujVFA",
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
