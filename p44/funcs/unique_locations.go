package funcs

func Unique(locations LocationsDATA) map[string]interface{} {
	MyMap := make(map[string]interface{})
	for _, TheStructs := range locations.Index {
		for _, location := range TheStructs.Locations {
			MyMap[location] = nil
		}
	}

	return MyMap
}
