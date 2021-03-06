package page

import (
	"github.com/TobiasYin/go_web_ui/node"
	"github.com/TobiasYin/go_web_ui/tool/template/component"
)

func GetMainPage() *node.Page {
	return node.NewPage("Main", func(this *node.Context) node.Widget {
		return node.Column{
			Alignment: node.Center,
			Children: []node.Widget{
				node.Link{
					Child: node.Text{Content: "Image page"},
					Href: "/image?title=hello!",
				},
				node.Text{Content: "Hello"},
				component.StatefulDemo{
					Key: node.GenerateKeyWithCallLine(),
					Value: " Hello ",
					Size: 22,
				},
				component.StatelessDemo{
					Value: "World",
				},
			},
		}
	})
}
