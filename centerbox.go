package main

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func centerBox() gtk.Box {
	boxx := gtk.NewBox(gtk.OrientationVertical, 5)
	boxx.SetVExpand(false)
	boxx.SetHExpand(false)
	boxx.SetVAlign(gtk.AlignStart)
	boxx.SetHAlign(gtk.AlignStart)

	firstLabel := gtk.NewLabel(getDistro() + " Linux")
	secondLabel := gtk.NewLabel("Version " + getKernel())
	thirdLabel := gtk.NewLabel("© Linux. All rights reversed")
	thirdLabel.SetMarginBottom(5)

	fourthLabel := gtk.NewLabel("The"+ getDistro() +" operating system and its user interface are protected by trademark and other pending or existing intellectual property rights in the United States and other countries/regions.")

	labbels(firstLabel)
	labbels(secondLabel)
	labbels(thirdLabel)
	labbels(fourthLabel)

	boxx.Append(firstLabel)
	boxx.Append(secondLabel)
	boxx.Append(thirdLabel)
	boxx.Append(fourthLabel)

	return *boxx
}

func labbels(l *gtk.Label) {
	l.AddCSSClass("medium-text")
	l.SetWrap(true)
	l.SetHAlign(gtk.AlignStart)
}
