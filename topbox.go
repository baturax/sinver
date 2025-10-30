package main

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func topBox() gtk.Box {
	boxx :=  gtk.NewBox(gtk.OrientationHorizontal, 10)
	boxx.SetHExpand(true)
	boxx.SetVExpand(false)
	boxx.SetVAlign(gtk.AlignStart)
	boxx.SetHAlign(gtk.AlignCenter)

	logo := gtk.NewImageFromFile("./tux.svg")
	logo.SetPixelSize(100)

	label := gtk.NewLabel("About Linux")
	label.AddCSSClass("big-text")

	boxx.Append(logo)
	boxx.Append(label)

	return *boxx
}
