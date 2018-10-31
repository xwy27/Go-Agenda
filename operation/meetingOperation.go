package operation

import (
	"GO-AGENDA/entity"
	"errors"
	"time"
)

type meeting = entity.Meeting

const timeFormat = "2006-01-02 15:04:05"

// validateMeeting validates meeting time properties and returns, if something wrong, an error
// an valid meeting contains start time and end time and the time interval is right while,
// participators only attend this new meeting in the meeting time interval
func validateMeeting(meeting *meeting) error {
	// Validate Start Time
	if meeting.startTime == 0 {
		return errors.New("Start Time is required")
	}

	// Validate End Time
	if meeting.endTime == 0 {
		return errors.New("End Time is required")
	}

	// Validate time interval
	start := time.Unix(meeting.startTime, 0)
	end := time.Unix(meeting.endTime, 0)
	if start.After(end) {
		return errors.New("Start Time must be before End Time")
	}

	// Validate unique meeting for participators
	meetings := entity.meetings.FindMeetingsBy(func(m *Meeting) bool {
		if m.sponsor == meeting.sponsor {
			return true
		}
		for _, participator := range meeting.participators {
			if participator == m.sponsor {
				return true
			}
		}
	})

	for _, participator := range meeting.participators {

	}
}

func addMeeting(title string, participators []string, startTime int, endTime int) error {
	// TODO:Check log in

	// Validate title
	if len(title) == 0 {
		return errors.New("Meeting Title is required")
	}

	// Check title existance
	if entity.meetings.FindByTitle(title).size != 0 {
		return errors.New("Meeting: " + title + " is existed")
	}

	// Validate participator
	if len(participators) == 0 {
		return errors.New("Meeting Participator is required")
	}

	// TODO:Check sponsor without in participator

	newMeeting := &meeting{
		title:         title,
		sponsor:       currentUser, //TODO:Change code for cur User
		participators: participators,
		startTime:     startTime,
		endTime:       endTime,
	}

	if err := validateMeeting(newMeeting); err != nil {
		return err
	}

	entity.meetings.AddMeeting(newMeeting)
	return nil
}

func deleteMeeting(title string) error {

}
