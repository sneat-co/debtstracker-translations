package trans

func Commands(command string, extra... string) (commands []string) {
	vals := TRANS[command]
	commands = make([]string, len(vals)+len(extra))
	var i int
	for _, val := range vals {
		if val == "" {
			continue
		}
		commands[i] = "/" + val
		i += 1
	}
	for _, val := range extra {
		commands[i] = val
		i += 1
	}
	return
}
