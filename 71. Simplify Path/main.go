package main

func simplifyPath(path string) string {
	parts := make([]string, 0)
	var part string
	idx := 0
	for idx < len(path) {
		if path[idx] == '/' {
			if len(part) > 0 && part != "." {
				if part == ".." {
					if len(parts) > 0 {
						parts = parts[0:len(parts)-1]
					}
				} else {
					parts = append(parts, part)
				}
			}

			part = ""
			idx = idx + 1
		} else {
			part = part + string(path[idx])
			idx = idx + 1
		}
	}

	if len(part) > 0 && part != "." {
		if part == ".." {
			if len(parts) > 0 {
				parts = parts[0:len(parts)-1]
			}
		} else {
			parts = append(parts, part)
		}
	}

	if len(parts) > 0 {
		result := ""
		for _, seg := range parts {
			result = result + "/" + seg
		}

		return result
	}

	return "/"
}