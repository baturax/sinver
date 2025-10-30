package main

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func activate(app *gtk.Application) {

	gtk.StyleContextAddProviderForDisplay(
		gdk.DisplayGetDefault(), loadCSS(styleCSS),
		gtk.STYLE_PROVIDER_PRIORITY_APPLICATION,
	)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("About Linux")

	mBox := gtk.Box(mainBox())
	tBox := gtk.Box(topBox())
	cBox := gtk.Box(centerBox())
	bBox := bottomBox(window)


	// append
	mBox.Append(&tBox)
	mBox.Append(&cBox)
	mBox.Append(&bBox)

	window.SetChild(&mBox)
	window.SetDefaultSize(650, 500)
	window.SetVisible(true)

}
