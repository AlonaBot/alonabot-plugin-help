// Package help 提供可用命令列表和当前用户身份查询。
//
// New 登记 /help 与 /whoami；命令准入与列表筛选由 Core 按当前会话执行（CONTRACTS §15.2）。
package help

import (
	"strings"

	"github.com/AlonaBot/AlonaBot/sdk"
	"github.com/AlonaBot/AlonaBot/sdk/command"
)

// Register 登记组件的业务路由、存储和生命周期。
func Register(p *sdk.Plugin) {
	p.Command("help").Description("List available commands").Handle(func(c *sdk.Context) error {
		entries, err := command.List(c)
		if err != nil {
			return err
		}
		if len(entries) == 0 {
			return c.Reply(c.T("help.empty"))
		}
		lines := []string{c.T("help.title")}
		for _, entry := range entries {
			names := []string{strings.Join(entry.Path, " ")}
			for _, alias := range entry.Aliases {
				names = append(names, strings.Join(alias, " "))
			}
			text := strings.Join(names, ", ")
			if entry.Description == "" {
				lines = append(lines, c.T("help.line", "names", text))
			} else {
				lines = append(lines, c.T("help.line.description", "names", text, "description", entry.Description))
			}
		}
		return c.Reply(strings.Join(lines, "\n"))
	})
	p.Command("whoami").Description("Show your user ID and instance").Handle(func(c *sdk.Context) error {
		if c.User() == "" {
			return c.Reply(c.T("help.whoami.no_actor"))
		}
		return c.Reply(c.T("help.whoami", "user", c.User(), "instance", c.Instance()))
	})
}
