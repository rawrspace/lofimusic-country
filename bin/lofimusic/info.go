package main

import (
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/maxence-charriere/go-app/v9/pkg/ui"
)

const (
	infoLinkIconSize    = 18
	cardVisibleDuration = time.Second * 10
	cardHiddenDuration  = time.Second * 5
)

type info struct {
	app.Compo

	Iclass   string
	Iradio   liveRadio
	Iplaying bool

	currentCard   int
	isCardVisible bool
}

func newInfo() *info {
	return &info{}
}

func (i *info) Class(v string) *info {
	if v == "" {
		return i
	}
	if i.Iclass != "" {
		i.Iclass += " "
	}
	i.Iclass += v
	return i
}

func (i *info) Radio(v liveRadio) *info {
	i.Iradio = v
	return i
}

func (i *info) Playing(v bool) *info {
	i.Iplaying = v
	return i
}

func (i *info) OnMount(ctx app.Context) {
	i.currentCard = -1

	ticker := time.NewTicker(cardHiddenDuration)
	ctx.Async(func() {
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				ctx.Dispatch(func(ctx app.Context) {
					if i.isCardVisible {
						ticker.Reset(cardHiddenDuration)
						i.isCardVisible = false
					} else {
						ticker.Reset(cardVisibleDuration)
						i.isCardVisible = true
					}
				})
			}
		}
	})
}

func (i *info) OnUpdate(ctx app.Context) {
}

func (i *info) Render() app.UI {
	titleVisibility := ""
	if i.Iplaying {
		titleVisibility = "info-title-show"
	}

	return app.Article().
		Class("info").
		Class("fill").
		Class("no-overflow").
		Body(
			app.Header().
				Class("info-title").
				Class(titleVisibility).
				Class("center").
				Class("fit").
				Body(
					app.H1().
						Class("h1").
						Class("glow").
						Text(i.Iradio.Name),
					app.Div().Class("info-title-separator"),
					ui.Stack().
						Class("info-links").
						Center().
						Middle().
						Content(
							app.H2().
								Class("h2").
								Class("glow").
								Text(i.Iradio.Owner),
						),
				),
		)
}
