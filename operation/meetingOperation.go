package operation

import (
	"Go-Agenda/model"
	"errors"
	"time"
)

type participator = model.Participator

// type session = model.Session

const timeFormat = "2006-01-02T15:04:05"

// ValidateMeeting validates meeting time properties and returns, if something wrong, an error
// Valid meeting contains start time and end time and the time interval is right while,
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
	SMeetings, err := model.FindMeetingsBy(func(m *model.Meeting) bool {
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

	if err != nil {
		return err
	}

	for _, m := range SMeetings {
		if !((end.Before(time.Unix(m.StartTime, 0))) ||
			(start.After(time.Unix(m.EndTime, 0)))) {
			return errors.New("You have meeting to attend in the time interval")
		}
	}

	// Validate unique meeting for participator
	var PMeetings []model.Meeting
	for _, participator := range meeting.Participators {
		meetings, err := model.FindMeetingsBy(func(m *model.Meeting) bool {
			if m.Sponsor == participator.Username {
				return true
			}

			for _, p := range m.Participators {
				if p.Username == participator.Username {
					return true
				}
			}
			return false
		})

		if err != nil {
			return err
		}
		PMeetings = append(PMeetings, meetings...)
	}
	if err != nil {
		return err
	}

	for _, m := range PMeetings {
		if !((end.Before(time.Unix((int64)(m.StartTime), 0))) ||
			(start.After(time.Unix((int64)(m.EndTime), 0)))) {
			return errors.New("You have meeting to attend in the time interval")
		}
	}

	return nil
}

// AddMeeting adds a valid meeting to db and returns, if something wrong, error
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

	// Check Title existence
	m, err := model.FindMeetingByTitle(Title)
	if err != nil {
		return err
	}
	if m != nil {
		return errors.New("Meeting: " + Title + " is existed")
	}

	// Validate participator
	if len(Participators) == 0 {
		return errors.New("Meeting Participator is required")
	}
	for _, p := range Participators {
		user := model.FindUserByName(p.Username)
		if user == nil {
			return errors.New("Participator: " + p.Username + " does not exist")
		}
	}

	// Check Sponsor without in participator
	for _, p := range Participators {
		if currentUser == p.Username {
			return errors.New("You could not attend meeting sponsored by you as a participator")
		}
	}

	// Check valid time
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
		Sponsor:       currentUser,
		Participators: Participators,
		StartTime:     start,
		EndTime:       end,
	}

	if err := ValidateMeeting(newMeeting); err != nil {
		return err
	}

	err = model.AddMeeting(newMeeting)
	return err
}

// DeleteMeeting deletes a meeting by meeting name and returns, if something wrong, error
// Meeting must exist and login user must be the sponsor
func DeleteMeeting(title string) error {
	currentUser, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	meeting, err := model.FindMeetingByTitle(title)
	if err != nil {
		return err
	}
	if meeting == nil {
		return errors.New("Meeting: " + title + " does not exist")
	}

	if meeting.Sponsor != currentUser {
		return errors.New("You don't have the authority to delete others' meeting")
	}

	return model.DeleteMeeting(title)
}

// AddParticipator adds the user into a meeting and returns, if something wrong, error
// Meeting must exist and user must exist
func AddParticipator(title, username string) error {
	meeting, err := model.FindMeetingByTitle(title)
	if err != nil {
		return err
	}
	if meeting == nil {
		return errors.New("Meeting: " + title + " does not exist")
	}

	user := model.FindUserByName(username)
	if user == nil {
		return errors.New("User: " + username + " does not exist")
	}

	for _, p := range meeting.Participators {
		if username == p.Username {
			return errors.New(username + " has been the participator of the meeting")
		}
	}

	if meeting.Sponsor == username {
		return errors.New("You are the sponsor of the meeting")
	}

	return model.AddParticipator(title, username)
}

// DeleteParticipator deletes a participator in a meeting and returns, if something wrong, error
// Meeting must exist, login user must be the sponsor and the user must exist and be the participator
func DeleteParticipator(title, username string) error {
	currentUser, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	meeting, err := model.FindMeetingByTitle(title)
	if err != nil {
		return err
	}
	if meeting == nil {
		return errors.New("Meeting: " + title + " does not exist")
	}

	if meeting.Sponsor != currentUser {
		return errors.New("You don't have the authority to delete participator of others' meeting")
	}

	user := model.FindUserByName(username)
	if user == nil {
		return errors.New("User: " + username + " does not exist")
	}

	for _, p := range meeting.Participators {
		if p.Username == username {
			return model.DeleteParticipator(title, username)
		}
	}

	return errors.New(username + " is not a participator of meeting: " + title)
}

// QuitMeeting helps login user quit a meeting and returns, if something wrong, error
// Meeting must exist and login user must be the participator of the meeting
func QuitMeeting(title string) error {
	currentUser, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	meeting, err := model.FindMeetingByTitle(title)
	if err != nil {
		return err
	}
	if meeting != nil {
		return errors.New("Meeting: " + title + " does not exist")
	}

	for _, p := range meeting.Participators {
		if p.Username == currentUser {
			return model.DeleteParticipator(title, currentUser)
		}
	}

	return errors.New("You have not attended the meeting")
}

// QueryMeetings queries the sponsored meetings matching the given time interval and returns,
// if something wrong, error
func QueryMeetings(startTime, endTime string) ([]model.Meeting, error) {
	user, err := model.GetCurrentUserName()
	if err != nil {
		return nil, err
	}

	start, err := time.Parse(timeFormat, startTime)
	if err != nil {
		return nil, errors.New("Invalid Start Time")
	}

	end, err := time.Parse(timeFormat, endTime)
	if err != nil {
		return nil, errors.New("Invalid End Time")
	}

	if start.After(end) {
		return nil, errors.New("End time must be after Start time")
	}

	return model.FindMeetingsBy(func(m *model.Meeting) bool {
		mStart := time.Unix(m.StartTime, 0)
		mEnd := time.Unix(m.EndTime, 0)
		if !(mStart.After(end) || mEnd.Before(start)) && m.Sponsor == user {
			return true
		}
		for _, p := range m.Participators {
			if user == p.Username {
				return true
			}
		}
		return false
	})
}

// ClearMeetings clears all the meetings sponsored by current user and returns,
// if something wrong, error
// Log in required
func ClearMeetings() error {
	currentUser, err := model.GetCurrentUserName()
	if err != nil {
		return err
	}

	meetings, err := model.FindMeetingsBy(func(m *model.Meeting) bool {
		if m.Sponsor == currentUser {
			return true
		}
		return false
	})

	if err != nil {
		return err
	}

	for _, m := range meetings {
		err = model.DeleteMeeting(m.Title)
		if err != nil {
			return err
		}
	}

	return nil
}
