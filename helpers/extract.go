package helpers

import "regexp"

func ExtractInfo(expression string, line string) map[string]string {
	result := make(map[string]string)
	re := regexp.MustCompile(expression)

	matches := re.FindStringSubmatch(line)
	if len(matches) == 0 {
		return result
	}

	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result
}
