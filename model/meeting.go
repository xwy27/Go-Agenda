package model

import (
	"errors"
)

// Participator is the type
// which specifics the format
// that a participator exists
// in the JSON file.
type Participator struct {
	Username string
}

// Meeting is the type which
// specifics the format
// that a meeting exists
// in the JSON file.
type Meeting struct {
	Title         string
	Sponsor       string
	Participators []Participator
	StartTime     int64
	EndTime       int64
}

type meetingsJSON struct {
	Meetings []Meeting
}

type meetingsType struct {
	storage    Storage
	dictionary map[string]Meeting
}

var meetings meetingsType
var meetingsDB meetingsJSON
var isMeetingInit = false

func initMeetings() error {
	if isMeetingInit {
		return nil
	}
	isMeetingInit = true
	meetings.storage.filePath = "data/meetings.json"
	meetings.dictionary = make(map[string]Meeting)
	return loadMeetings()
}

// AddMeeting accept a meeting pointer
// as the parameter. It will tries to
// add this meeting to current meeting
// set without any check.
func AddMeeting(meeting *Meeting) error {
	if err := initMeetings(); err != nil {
		return err
	}
	meetings.dictionary[meeting.Title] = Meeting(*meeting)
	return writeMeetings()
}

// DeleteMeeting tries to delete a
// meeting with the given title.
func DeleteMeeting(title string) error {
	if err := initMeetings(); err != nil {
		return err
	}
	delete(meetings.dictionary, title)
	return writeMeetings()
}

// FindMeetingsBy will use the filter
// passed in to filter those meetings
// that meet the requirement and return
// them.
func FindMeetingsBy(filter func(*Meeting) bool) ([]Meeting, error) {
	if err := initMeetings(); err != nil {
		return []Meeting{}, err
	}
	var resultMeetings []Meeting
	for _, meeting := range meetings.dictionary {
		if filter(&meeting) {
			resultMeetings = append(resultMeetings, meeting)
		}
	}
	return resultMeetings, nil
}

// FindMeetingByTitle tries to find
// the meeting with the given title
// and return it to you.
func FindMeetingByTitle(title string) (*Meeting, error) {
	if err := initMeetings(); err != nil {
		return nil, err
	}
	if meeting, ok := meetings.dictionary[title]; ok {
		return &meeting, nil
	}
	return nil, nil
}

// DeleteParticipator tries to delete a
// paticipator with the given username
// from a meeting with given title.
func DeleteParticipator(title, username string) error {
	if err := initMeetings(); err != nil {
		return err
	}
	if meeting, ok := meetings.dictionary[title]; ok {
		for index, participator := range meeting.Participators {
			if participator.Username == username {
				meeting.Participators = append(meeting.Participators[:index], meeting.Participators[index+1:]...)
				meetings.dictionary[title] = meeting
				break
			}
		}

		if len(meeting.Participators) == 0 {
			delete(meetings.dictionary, title)
		}
		return writeMeetings()
	}
	return errors.New("no such meeting")
}

// AddParticipator tries to append
// a participator with the given
// username without any check.
func AddParticipator(title, username string) error {
	if err := initMeetings(); err != nil {
		return err
	}

	if meeting, ok := meetings.dictionary[title]; ok {
		meeting.Participators = append(meeting.Participators, Participator{username})
		meetings.dictionary[title] = meeting
		return writeMeetings()
	}
	return errors.New("no such meeting")
}

func loadMeetings() error {
	err := meetings.storage.load(&meetingsDB)
	if err != nil {
		return err
	}
	for _, meeting := range meetingsDB.Meetings {
		meetings.dictionary[meeting.Title] = Meeting(meeting)
	}
	return nil
}

func writeMeetings() error {
	var newMeetingDB meetingsJSON
	for _, meeting := range meetings.dictionary {
		newMeetingDB.Meetings = append(newMeetingDB.Meetings, meeting)
	}
	return meetings.storage.write(&newMeetingDB)
}
