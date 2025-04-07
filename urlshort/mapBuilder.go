package urlshort

func buildMap(parsedYAML []map[string]string) map[string]string {
	res := make(map[string]string)
	for _, entry := range parsedYAML {
		res[entry["path"]] = entry["url"]
	}
	return res
}
