package functions

func (a * Info) SearchNumberOfAntsAndRoomsAndTunnels(lines []string) string {
	var MessageOfInvalidInput string

	MessageOfInvalidInput = a.SearchNumberOfAnts(lines)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfRooms(lines)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfTunnels(lines)
	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	
	return ""
}
