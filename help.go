package cli

import (
	"bytes"
	"fmt"
	"text/template"
)

const greenColor = "\x1b[32m"

const yellowColor = "\x1b[33m"

const resetColor = "\x1b[0m"

const helpTemplate string = `
 __  __         ___
|  \/  |_ _    |   \
| |\/| | '_|   | |) |
|_|  |_|_| (_) |___/

{{yellowFill "NAME:"}}
  {{greenFill .Name | printf "%-25s"}} Version {{.Version}}

{{yellowFill "USAGE:"}}
  {{greenFill .Usage | printf "%-25s"}} The command name [OPTIONS]
{{if .Options}}
{{yellowFill "OPTIONS:"}}{{range .Options}}
  {{renderOption .}}{{end}}{{end}}
{{if .Commands}}
{{yellowFill "Commands:"}}{{range .Commands}}
  {{ .}}{{end}}{{end}}
`

func loadTemplate(temp string, data interface{}) string {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"greenFill":    func(str string) string { return fillColor(str, greenColor) },
		"yellowFill":   func(str string) string { return fillColor(str, yellowColor) },
		"renderOption": renderOptionTemplate,
	})
	template.Must(t.Parse(temp))
	b := bytes.Buffer{}
	if err := t.Execute(&b, data); err != nil {
		panic(err)
	}

	return b.String()
}

func renderOptionTemplate(o Option) string {
	str := ""
	if o.Short != "" {
		str = fmt.Sprintf("-%s, ", o.Short)
	}
	str = fmt.Sprintf("%-4s--%s", str, o.Name)

	return fmt.Sprintf("%-25s %s", fillColor(str, greenColor), o.Usage)
}

func fillColor(text string, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, resetColor)
}

func (c *Cli) printHelp() {
	fmt.Println(loadTemplate(helpTemplate, struct {
		Name     string
		Version  string
		Usage    string
		Options  []Option
		Commands []string
	}{
		c.Name,
		c.Version,
		fmt.Sprintf("%s %s", c.Name, c.Usage),
		c.Options,
		[]string{
			fmt.Sprintf("%-25s Lists commands", fillColor("list", greenColor)),
			fmt.Sprintf("%-25s Display the current framework environment", fillColor("env", greenColor)),
			fmt.Sprintf("%-25s Bring the application out of maintenance mode", fillColor("up", greenColor)),
		},
	}))
}
