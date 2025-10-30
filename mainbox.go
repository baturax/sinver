package main

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func mainBox() gtk.Box {
	boxx := gtk.NewBox(gtk.OrientationVertical, 40)
	boxx.SetMarginStart(50)
	boxx.SetMarginEnd(20)
	boxx.SetMarginTop(20)
	boxx.SetMarginBottom(20)

	return *boxx
}
