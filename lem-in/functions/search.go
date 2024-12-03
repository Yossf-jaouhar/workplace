package functions

func (a * Info) SearchNumberOfAntsAndRoomsAndTunnels(lines []string) string {
	var MessageOfInvalidInput string

	MessageOfInvalidInput, Index := a.SearchNumberOfAnts(lines)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfRooms(lines, Index)

	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfTunnels(lines, Index)
	if MessageOfInvalidInput != "" {
		return MessageOfInvalidInput
	}

	
	return ""
}
