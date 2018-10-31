package operation

import (
	"Go-Agenda/model"
	"errors"
	"time"
)

type participator = model.Participator

// type session = model.Session

const timeFormat = "2006-01-02 15:04:05"

// ValidateMeeting validates meeting time properties and returns, if something wrong, an error
// an valid meeting contains start time and end time and the time interval is right while,
// sponsor and participators only attend this new meeting in the meeting time interval
func ValidateMeeting(meeting *model.Meeting) error {
	// Validate Start Time
	if meeting.StartTime == 0 {
		return errors.New("Start Time is required")
	}

	// Validate End Time
	if meeting.EndTime == 0 {
		return errors.New("End Time is required")
	}

	// Validate time interval
	start := time.Unix(meeting.StartTime, 0)
	end := time.Unix(meeting.EndTime, 0)
	if start.After(end) {
		return errors.New("Start Time must be before End Time")
	}

	// Validate unique meeting for sponsor
	SMeetings := model.FindMeetingsBy(func(m *model.Meeting) bool {
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

	for _, m := range SMeetings {
		if !((end.Before(time.Unix((int64)(m.StartTime), 0))) ||
			(start.After(time.Unix((int64)(m.EndTime), 0)))) {
			return errors.New("You have meeting to attend in the time interval")
		}
	}

	// Validate unique meeting for participator
	var PMeetings []model.Meeting
	for _, participator := range meeting.Participators {
		PMeetings = append(PMeetings, model.FindMeetingsBy(func(m *model.Meeting) bool {
			if m.Sponsor == participator.Username {
				return true
			}

			for _, p := range m.Participators {
				if p.Username == participator.Username {
					return true
				}
			}
			return false
		})...)
	}

	for _, m := range PMeetings {
		if !((end.Before(time.Unix((int64)(m.StartTime), 0))) ||
			(start.After(time.Unix((int64)(m.EndTime), 0)))) {
			return errors.New("You have meeting to attend in the time interval")
		}
	}
	return nil
}

// AddMeeting adds a valid meeting to db
func AddMeeting(Title string, Participators []participator, StartTime string, EndTime string) error {
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

	s, err := time.Parse(timeFormat, StartTime)
	if err != nil {
		return errors.New("Start Time is invalid")
	}
	start := s.Unix()

	e, err := time.Parse(timeFormat, EndTime)
	if err != nil {
		return errors.New("End Time is invalid")
	}
	end := e.Unix()

	newMeeting := &model.Meeting{
		Title:         Title,
		Sponsor:       currentUser, //TODO:Change code for cur User
		Participators: Participators,
		StartTime:     start,
		EndTime:       end,
	}

	if err := ValidateMeeting(newMeeting); err != nil {
		return err
	}

	model.AddMeeting(newMeeting)
	return nil
}

func DeleteMeeting(Title string) error {

}
