package libs

import "github.com/ctcpip/notifize"

func Done() {
	notifize.Display("Tea", "Your tea is done !", true, "")
}
