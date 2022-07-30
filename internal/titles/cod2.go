package titles

import (
	"github.com/cetteup/joinme.click-launcher/internal/domain"
	"github.com/cetteup/joinme.click-launcher/internal/titles/internal"
	"github.com/cetteup/joinme.click-launcher/pkg/game_launcher"
	"github.com/cetteup/joinme.click-launcher/pkg/software_finder"
)

var Cod2 = domain.GameTitle{
	ProtocolScheme: "cod2",
	GameLabel:      "Call of Duty 2",
	FinderConfigs: []software_finder.Config{
		{
			ForType:           software_finder.RegistryFinder,
			RegistryPath:      "SOFTWARE\\WOW6432Node\\Activision\\Call of Duty 2",
			RegistryValueName: "InstallPath",
		},
	},
	LauncherConfig: game_launcher.Config{
		ExecutableName:    "CoD2MP_s.exe",
		CloseBeforeLaunch: true,
	},
	URLValidator: internal.IPPortURLValidator,
	CmdBuilder:   internal.PlusConnectCmdBuilder,
}