package functions

import (
	"errors"
	"strconv"
)

func (a * Info) SearchNumberOfAntsAndRoomsAndTunnels(lines []string) error {
	var MessageOfInvalidInput error

	var err error

	a.NumberOfAnts , err = strconv.Atoi(lines[0])

	if err != nil {
		return errors.New("invalid nm ants")
	}
	if a.NumberOfAnts <= 0 {
		return errors.New("invalid nm ants")
	}

	MessageOfInvalidInput = a.SearchOfRooms(lines)

	if MessageOfInvalidInput != nil {
		return MessageOfInvalidInput
	}

	MessageOfInvalidInput = a.SearchOfTunnels(lines)
	if MessageOfInvalidInput != nil {
		return MessageOfInvalidInput
	}

	
	return nil
}
