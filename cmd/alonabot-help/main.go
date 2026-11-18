// alonabot-help 命令运行官方 help 插件。
package main

import (
	"log/slog"
	"os"

	"github.com/AlonaBot/AlonaBot/plugins/help"
	"github.com/AlonaBot/AlonaBot/sdk"
)

func main() {
	if err := sdk.Run(help.New()); err != nil {
		slog.Error("component stopped", "error", err)
		os.Exit(1)
	}
}
