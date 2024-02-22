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
	AddedAt time.Time
}

func (r liveRadio) youtubeID() string {
	return path.Base(r.URL)
}

func getLiveRadios() []liveRadio {
	radios := []liveRadio{
		//Boot Scootin Boogie - Brooks & Dunn - d05tQrhNMkA
		{
			Slug:  "boot-scootin-boogie",
			Name:  "Boot Scootin Boogie",
			Owner: "Brooks & Dunn",
			URL:   "https://youtu.be/d05tQrhNMkA",
		},
		//Tracy Byrd - Watermelon Crawl - BFPdX2TFiIE
		{
			Slug:  "watermelon-crawl",
			Name:  "Watermelon Crawl",
			Owner: "Tracy Byrd",
			URL:   "https://youtu.be/BFPdX2TFiIE",
		},
		//Alan Jackson - Chattahoochee - JW5UEW2kYvc
		{
			Slug:  "chattahoochee",
			Name:  "Chattahoochee",
			Owner: "Alan Jackson",
			URL:   "https://youtu.be/JW5UEW2kYvc",
		},
		//Garth Brooks- Friends In Low Places - mvCgSqPZ4EM
		{
			Slug:  "friends-in-low-places",
			Name:  "Friends In Low Places",
			Owner: "Garth Brooks",
			URL:   "https://youtu.be/mvCgSqPZ4EM",
		},
		//Reba McEntire - Fancy - FUN0U80hELo
		{
			Slug:  "fancy",
			Name:  "Fancy",
			Owner: "Reba McEntire",
			URL:   "https://youtu.be/FUN0U80hELo",
		},
		//Zach Bryan - Hey Driver - o8LTi_zkGuY
		{
			Slug:  "hey-driver",
			Name:  "Hey Driver",
			Owner: "Zach Bryan",
			URL:   "https://youtu.be/o8LTi_zkGuY",
		},
		//Dolly Parton - Jolene - Ixrje2rXLMA
		{
			Slug:  "jolene",
			Name:  "Jolene",
			Owner: "Dolly Parton",
			URL:   "https://youtu.be/Ixrje2rXLMA",
		},
		//Zac Brown Band - Chicken Fried - e4ujS1er1r0
		{
			Slug:  "chicken-fried",
			Name:  "Chicken Fried",
			Owner: "Zac Brown Band",
			URL:   "https://youtu.be/e4ujS1er1r0",
		},
		//Martina McBride - Independence Day - 4VPpAZ9_qAw
		{
			Slug:  "independence-day",
			Name:  "Independence Day",
			Owner: "Martina McBride",
			URL:   "https://youtu.be/4VPpAZ9_qAw",
		},
		//Toby Keith - Should've Been A Cowboy - aIq1LvzSLsk
		{
			Slug:  "shouldve-been-a-cowboy",
			Name:  "Should've Been A Cowboy",
			Owner: "Toby Keith",
			URL:   "https://youtu.be/aIq1LvzSLsk",
		},
		//George Strait - Amarillo By Morning - wtVeDaZxAXo
		{
			Slug:  "amarillo-by-morning",
			Name:  "Amarillo By Morning",
			Owner: "George Strait",
			URL:   "https://youtu.be/wtVeDaZxAXo",
		},
		//Tim McGraw - I Like It, I Love It - aAZVtj_JsCY
		{
			Slug:  "i-like-it-i-love-it",
			Name:  "I Like It, I Love It",
			Owner: "Tim McGraw",
			URL:   "https://youtu.be/aAZVtj_JsCY",
		},
		//Big & Rich - Save A Horse (Ride A Cowboy) - S9ZbuIRPwFg
		{
			Slug:  "save-a-horse-ride-a-cowboy",
			Name:  "Save A Horse (Ride A Cowboy)",
			Owner: "Big & Rich",
			URL:   "https://youtu.be/S9ZbuIRPwFg",
		},
		//Shania Twain - Any Man Of Mine - 8N2k-gv6xNE
		{
			Slug:  "any-man-of-mine",
			Name:  "Any Man Of Mine",
			Owner: "Shania Twain",
			URL:   "https://youtu.be/8N2k-gv6xNE",
		},
		//Nitty Gritty Dirt Band - Fishin' In The Dark - jBRfkxUAyOk
		{
			Slug:  "fishin-in-the-dark",
			Name:  "Fishin' In The Dark",
			Owner: "Nitty Gritty Dirt Band",
			URL:   "https://youtu.be/jBRfkxUAyOk",
		},
		//Wynonna Judd - No One Else On Earth - HfFeNt7a18k
		{
			Slug:  "no-one-else-on-earth",
			Name:  "No One Else On Earth",
			Owner: "Wynonna Judd",
			URL:   "https://youtu.be/HfFeNt7a18k",
		},
		//Blake Shelton - Austin - bb1DTsxBOfE
		{
			Slug:  "austin",
			Name:  "Austin",
			Owner: "Blake Shelton",
			URL:   "https://youtu.be/bb1DTsxBOfE",
		},
		//Little Big Town - Boondocks - skAOb_EUE_M
		{
			Slug:  "boondocks",
			Name:  "Boondocks",
			Owner: "Little Big Town",
			URL:   "https://youtu.be/skAOb_EUE_M",
		},
		//Brooks & Dunn - Neon Moon - f4sdG5frA0Q
		{
			Slug:  "neon-moon",
			Name:  "Neon Moon",
			Owner: "Brooks & Dunn",
			URL:   "https://youtu.be/f4sdG5frA0Q",
		},
		//Travis Tritt - It's a Great Day to be Alive - d4tSE2w53ts
		{
			Slug:  "its-a-great-day-to-be-alive",
			Name:  "It's a Great Day to be Alive",
			Owner: "Travis Tritt",
			URL:   "https://youtu.be/d4tSE2w53ts",
		},
		//Clay Walker - If I Could Make a Living - 0q8-CwJYvfA
		{
			Slug:  "if-i-could-make-a-living",
			Name:  "If I Could Make a Living",
			Owner: "Clay Walker",
			URL:   "https://youtu.be/0q8-CwJYvfA",
		},
		//Conway Twitty - Tight Fittin Jeans - xs1kwVeKcRg
		{
			Slug:  "tight-fittin-jeans",
			Name:  "Tight Fittin Jeans",
			Owner: "Conway Twitty",
			URL:   "https://youtu.be/xs1kwVeKcRg",
		},
		//The Dance - Garth Brooks - 6FHvP_Yv0wE
		{
			Slug:  "the-dance",
			Name:  "The Dance",
			Owner: "Garth Brooks",
			URL:   "https://youtu.be/6FHvP_Yv0wE",
		},
		//Garth Brooks - Callin' Baton Rouge - L88rx2j011A
		{
			Slug:  "callin-baton-rouge",
			Name:  "Callin' Baton Rouge",
			Owner: "Garth Brooks",
			URL:   "https://youtu.be/L88rx2j011A",
		},
		//Shania Twain - Man! I Feel Like A Woman - ZJL4UGSbeFg
		{
			Slug:  "man-i-feel-like-a-woman",
			Name:  "Man! I Feel Like A Woman",
			Owner: "Shania Twain",
			URL:   "https://youtu.be/ZJL4UGSbeFg",
		},
		//Tim McGraw - Don't Take The Girl - FWaW1IjbmAk
		{
			Slug:  "dont-take-the-girl",
			Name:  "Don't Take The Girl",
			Owner: "Tim McGraw",
			URL:   "https://youtu.be/FWaW1IjbmAk",
		},
		//Darius Rucker - Wagon Wheel - hvKyBcCDOB4
		{
			Slug:  "wagon-wheel",
			Name:  "Wagon Wheel",
			Owner: "Darius Rucker",
			URL:   "https://youtu.be/hvKyBcCDOB4",
		},
		//Jamey Johnson - In Color - EYGwxf1gCC4
		{
			Slug:  "in-color",
			Name:  "In Color",
			Owner: "Jamey Johnson",
			URL:   "https://youtu.be/EYGwxf1gCC4",
		},
		//George Strait - All My Ex's Live In Texas - 9qumxVP8PrE
		{
			Slug:  "all-my-exs-live-in-texas",
			Name:  "All My Ex's Live In Texas",
			Owner: "George Strait",
			URL:   "https://youtu.be/9qumxVP8PrE",
		},
		//Luke Dick - Polyester - oJwd6cu-Fi0
		{
			Slug:  "polyester",
			Name:  "Polyester",
			Owner: "Luke Dick",
			URL:   "https://youtu.be/oJwd6cu-Fi0",
		},
		//Miranda Lambert - Somethin' Bad - o4Yzj-m_SBk
		{
			Slug:  "somethin-bad",
			Name:  "Somethin' Bad",
			Owner: "Miranda Lambert",
			URL:   "https://youtu.be/o4Yzj-m_SBk",
		},
		//Miranda Lambert - Gunpowder & Lead - aWQdEDtveB0
		{
			Slug:  "gunpowder-and-lead",
			Name:  "Gunpowder & Lead",
			Owner: "Miranda Lambert",
			URL:   "https://youtu.be/aWQdEDtveB0",
		},
		//Luke Combs - Beer Never Broke My Heart - 7Lb9dq-JZFI
		{
			Slug:  "beer-never-broke-my-heart",
			Name:  "Beer Never Broke My Heart",
			Owner: "Luke Combs",
			URL:   "https://youtu.be/7Lb9dq-JZFI",
		},
		//Luke Combs - Hurricane - BixwVsiDdZM
		{
			Slug:  "hurricane",
			Name:  "Hurricane",
			Owner: "Luke Combs",
			URL:   "https://youtu.be/BixwVsiDdZM",
		},
		//Luke Combs - She Got the Best of Me - a2a9fgPI_PI
		{
			Slug:  "she-got-the-best-of-me",
			Name:  "She Got the Best of Me",
			Owner: "Luke Combs",
			URL:   "https://youtu.be/a2a9fgPI_PI",
		},
		//Cody Johnson - Dirt Cheap - ZKCwYbPrAGg
		{
			Slug:  "dirt-cheap",
			Name:  "Dirt Cheap",
			Owner: "Cody Johnson",
			URL:   "https://youtu.be/ZKCwYbPrAGg",
		},
		//Cody Johnson - The Painter - ur1cb_OztPQ
		{
			Slug:  "the-painter",
			Name:  "The Painter",
			Owner: "Cody Johnson",
			URL:   "https://youtu.be/ur1cb_OztPQ",
		},
		//Jo Dee Messina - Heads Carolina, Tails California - FvLjJE7Bt88
		{
			Slug:  "heads-carolina-tails-california",
			Name:  "Heads Carolina, Tails California",
			Owner: "Jo Dee Messina",
			URL:   "https://youtu.be/FvLjJE7Bt88",
		},
		//Cody Johnson & Reba McEntire - Dear Rodeo - RlwREmok31o
		{
			Slug:  "dear-rodeo",
			Name:  "Dear Rodeo",
			Owner: "Cody Johnson & Reba McEntire",
			URL:   "https://youtu.be/RlwREmok31o",
		},
		//Josh Abbott Band - Oh, Tonight - CTlQzKOOkeU
		{
			Slug:  "oh-tonight",
			Name:  "Oh, Tonight",
			Owner: "Josh Abbott Band",
			URL:   "https://youtu.be/CTlQzKOOkeU",
		},
		//Josh Abbott Band - My Texas (feat. Pat Green) - z9wV-bI8Lwk
		{
			Slug:  "my-texas",
			Name:  "My Texas (feat. Pat Green)",
			Owner: "Josh Abbott Band",
			URL:   "https://youtu.be/z9wV-bI8Lwk",
		},
		//Josh Abbott Band - She's Like Texas - yapf2QvFHfw
		{
			Slug:  "shes-like-texas",
			Name:  "She's Like Texas",
			Owner: "Josh Abbott Band",
			URL:   "https://youtu.be/yapf2QvFHfw",
		},
		//George Strait - Baby Blue - dfOaLpIrC20
		{
			Slug:  "baby-blue",
			Name:  "Baby Blue",
			Owner: "George Strait",
			URL:   "https://youtu.be/dfOaLpIrC20",
		},
		//Luke Combs - Fast Car - fL7O5wW4wC0
		{
			Slug:  "fast-car",
			Name:  "Fast Car",
			Owner: "Luke Combs",
			URL:   "https://youtu.be/fL7O5wW4wC0",
		},
		//Waylon Jennings - Mammas, Don't Let Your Babies Grow Up to Be Cowboys - qZTwjljm5qc
		{
			Slug:  "mammas-dont-let-your-babies-grow-up-to-be-cowboys",
			Name:  "Mammas, Don't Let Your Babies Grow Up to Be Cowboys",
			Owner: "Waylon Jennings",
			URL:   "https://youtu.be/qZTwjljm5qc",
		},
		//Jon Pardi - Dirt On My Boots - gCxbgqyC2Wg
		{
			Slug:  "dirt-on-my-boots",
			Name:  "Dirt On My Boots",
			Owner: "Jon Pardi",
			URL:   "https://youtu.be/gCxbgqyC2Wg",
		},
		//Brothers Osborne - It Ain’t My Fault - E5RDEXpc8OY
		{
			Slug:  "it-aint-my-fault",
			Name:  "It Ain’t My Fault",
			Owner: "Brothers Osborne",
			URL:   "https://youtu.be/E5RDEXpc8OY",
		},
		//Brothers Osborne - Stay A Little Longer - zY6cMMtLCcQ
		{
			Slug:  "stay-a-little-longer",
			Name:  "Stay A Little Longer",
			Owner: "Brothers Osborne",
			URL:   "https://youtu.be/zY6cMMtLCcQ",
		},
		//Hank Williams, Jr. - A Country Boy Can Survive - 3cQNkIrg-Tk
		{
			Slug:  "a-country-boy-can-survive",
			Name:  "A Country Boy Can Survive",
			Owner: "Hank Williams, Jr.",
			URL:   "https://youtu.be/3cQNkIrg-Tk",
		},
		//Alabama - Mountain Music - M6WfM0cXWSQ
		{
			Slug:  "mountain-music",
			Name:  "Mountain Music",
			Owner: "Alabama",
			URL:   "https://youtu.be/M6WfM0cXWSQ",
		},
		//Hank Williams, Jr. - Family Tradition - GdDyEqnhvNI
		{
			Slug:  "family-tradition",
			Name:  "Family Tradition",
			Owner: "Hank Williams, Jr.",
			URL:   "https://youtu.be/GdDyEqnhvNI",
		},
		//Toby Keith - Red Solo Cup - BKZqGJONH68
		{
			Slug:  "red-solo-cup",
			Name:  "Red Solo Cup",
			Owner: "Toby Keith",
			URL:   "https://youtu.be/BKZqGJONH68",
		},
		//Jimmy Buffett - Margaritaville - mrF4nF8VUb4
		{
			Slug:  "margaritaville",
			Name:  "Margaritaville",
			Owner: "Jimmy Buffett",
			URL:   "https://youtu.be/mrF4nF8VUb4",
		},
		//Jimmy Buffett - Cheeseburger In Paradise - dmDBBu3CTOA
		{
			Slug:  "cheeseburger-in-paradise",
			Name:  "Cheeseburger In Paradise",
			Owner: "Jimmy Buffett",
			URL:   "https://youtu.be/dmDBBu3CTOA",
		},
		//Jimmy Buffett - Fins - kxh8YQKruAQ
		{
			Slug:  "fins",
			Name:  "Fins",
			Owner: "Jimmy Buffett",
			URL:   "https://youtu.be/kxh8YQKruAQ",
		},
	}

	sort.Slice(radios, func(a, b int) bool {
		return strings.Compare(radios[a].Name, radios[b].Name) < 0
	})

	return radios
}
