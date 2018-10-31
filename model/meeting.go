package model

import (
	"errors"
)

type Participator struct {
	Username string
}

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

func AddMeeting(meeting *Meeting) error {
	if err := initMeetings(); err != nil {
		return err
	}
	meetings.dictionary[meeting.Title] = Meeting(*meeting)
	return writeMeetings()
}

func DeleteMeeting(title string) error {
	if err := initMeetings(); err != nil {
		return err
	}
	delete(meetings.dictionary, title)
	return writeMeetings()
}

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

func FindMeetingByTitle(title string) (*Meeting, error) {
	if err := initMeetings(); err != nil {
		return nil, err
	}
	if meeting, ok := meetings.dictionary[title]; ok {
		return &meeting, nil
	}
	return nil, nil
}

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
