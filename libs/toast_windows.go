package libs

import (
	"github.com/go-toast/toast"
)

func Done() {
	notification := toast.Notification{
		AppID: "gotea",
		Title: "Tea",
		Message: "Your tea is done !",
	}
	notification.Push()
}
