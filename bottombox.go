package main

import "github.com/diamondburned/gotk4/pkg/gtk/v4"

func bottomBox(window *gtk.ApplicationWindow) gtk.Box {
	boxx := gtk.NewBox(gtk.OrientationVertical, 5)

	firstlabel := gtk.NewLabel("This product is licensed under  to:")
	firstlabel.SetMarkup("This product is licensed under <a href='https://www.gnu.org/licenses/gpl-3.0.en.html'>GNU General Public License</a> to:")

	secondLabel := gtk.NewLabel(getUser())
	secondLabel.SetMarginStart(30)

	okButton := gtk.NewButtonWithLabel("OK")
	okButton.SetHAlign(gtk.AlignEnd)
	okButton.ConnectClicked(func() {
		window.Application().Quit()
	})

	labbels(firstlabel)
	labbels(secondLabel)

	boxx.Append(firstlabel)
	boxx.Append(secondLabel)
	boxx.Append(okButton)

	return *boxx
}
