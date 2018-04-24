package trans

import "strings"

func Commands(command string, extra ...string) (commands []string) {
	vals := TRANS[command]
	commands = make([]string, 0, len(vals)*2+len(extra))
	for _, val := range vals {
		if val == "" {
			continue
		}
		val = strings.ToLower(strings.TrimLeft(val, "/"))
		for _, c := range commands {
			if c == val { // duplicate
				continue
			}
		}
		commands = append(commands, val)
		if len(val) > 1 && !strings.Contains(val, " ") {
			commands = append(commands, "/"+val)
		}
	}
	for _, val := range extra {
		for _, c := range commands {
			if c == val { // duplicate
				continue
			}
		}
		commands = append(commands, val)
	}
	return
}
