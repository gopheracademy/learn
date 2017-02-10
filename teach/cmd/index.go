package cmd

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/velvet"
)

func IndexHandler(res http.ResponseWriter, req *http.Request) {
	t, err := velvet.Render(indexTemplate, velvet.NewContextWith(map[string]interface{}{
		"modules": modules,
	}))
	if err != nil {
		res.WriteHeader(500)
		fmt.Fprint(res, err)
		return
	}
	res.WriteHeader(200)
	fmt.Fprint(res, t)
}

const indexTemplate = `
<html>
<head>
	<title>GopherAcademy</title>
</head>
<body>
	<ul>
	{{#each modules as |m|}}
		<li><a href="/module/{{m.Slug}}">{{m.Title}}</a></li>
	{{/each}}
	</ul>
</body>
</html>
`
