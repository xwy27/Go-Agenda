package operation

import (
	"errors"
	"time"

	"github.com/siskonemilia/Go-Agenda/model"
)

type meeting = model.Meeting
type participator = model.Participator
type session = model.Session

const timeFormat = "2006-01-02 15:04:05"

// validateMeeting validates meeting time properties and returns, if something wrong, an error
// an valid meeting contains start time and end time and the time interval is right while,
// Participators only attend this new meeting in the meeting time interval
func validateMeeting(meeting *meeting) error {
	// Validate Start Time
	if meeting.StartTime == 0 {
		return errors.New("Start Time is required")
	}

	// Validate End Time
	if meeting.EndTime == 0 {
		return errors.New("End Time is required")
	}

	// Validate time interval
	start := time.Unix((int64)(meeting.StartTime), 0)
	end := time.Unix((int64)(meeting.EndTime), 0)
	if start.After(end) {
		return errors.New("Start Time must be before End Time")
	}

	// Validate unique meeting for participators
	Meetings := model.FindMeetingsBy(func(m *model.Meeting) bool {
		if m.Sponsor == meeting.Sponsor {
			return true
		}
		for _, participator := range meeting.Participators {
			if participator.Username == m.Sponsor {
				return true
			}
		}
		return false
	})

	for _, participator := range meeting.Participators {

	}
	return nil
}

func addMeeting(Title string, Participators []participator, StartTime int, EndTime int) error {
	// Check log in
	currentUser, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	// Validate Title
	if len(Title) == 0 {
		return errors.New("Meeting Title is required")
	}

	// Check Title existance
	if model.FindMeetingByTitle(Title) != nil {
		return errors.New("Meeting: " + Title + " is existed")
	}

	// Validate participator
	if len(Participators) == 0 {
		return errors.New("Meeting Participator is required")
	}

	// TODO:Check Sponsor without in participator

	newMeeting := &meeting{
		Title:         Title,
		Sponsor:       currentUser, //TODO:Change code for cur User
		Participators: Participators,
		StartTime:     StartTime,
		EndTime:       EndTime,
	}

	if err := validateMeeting(newMeeting); err != nil {
		return err
	}

	model.AddMeeting(newMeeting)
	return nil
}

func deleteMeeting(Title string) error {
	return nil
}
