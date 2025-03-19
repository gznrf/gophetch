package system_info

import "github.com/JustinTimperio/osinfo"

func getOS() string {
	osName := osinfo.GetVersion().Name
	return osName
}
