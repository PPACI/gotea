package libs

import (
	"github.com/go-toast/toast"
)

func Done() {
	notification := toast.Notification{
		AppID: "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe",
		Title: "Tea",
		Message: "Your tea is done !",
	}
	notification.Push()
}
