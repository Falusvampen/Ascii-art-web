package asciiart

import "strings"

func DnaDiff() string {
	nice := "╭━━━━-╮<br>╰┃ ┣▇━▇)<br> ┃ ┃  ╰━▅╮<br> ╰┳╯ ╰━━┳╯ DNA DIFF<br>  ╰╮ ┳━━╯ GG EZ<br> ▕▔▋ ╰╮╭━╮NICE TUTORIAL<br>╱▔╲▋╰━┻┻╮╲╱▔▔▔╲<br>▏  ▔▔▔▔▔▔▔  O O┃<br>╲╱▔╲▂▂▂▂╱▔╲▂▂▂╱<br> ▏╳▕▇▇▕ ▏╳▕▇▇▕<br> ╲▂╱╲▂╱ ╲▂╱╲▂╱"
	nice = strings.ReplaceAll(nice, " ", "&nbsp;")
	nice = strings.ReplaceAll(nice, "U+2003", "&nbsp;")

	return nice
}
