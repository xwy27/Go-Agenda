package model

type Participator struct {
	Username string
}

type Meeting struct {
	Title         string
	Sponsor       string
	Participators []Participator
	StartTime     int
	EndTime       int
}

type meetingsJSON struct {
	Meetings []Meeting
}

type meetingsType struct {
	storage    Storage
	dictionary map[string]*Meeting
}

var meetings meetingsType
var meetingsDB meetingsJSON

func InitMeetings() error {
	return nil
}

func AddMeeting(meeting *Meeting) error {
	return nil
}

func DeleteMeeting(title string) error {
	return nil
}

func FindMeetingsBy(filter func(*Meeting) bool) []Meeting {
	return []Meeting{}
}

func FindMeetingByTitle(title string) *Meeting {
	return nil
}

func DeleteParticipator(title, username string) error {
	return nil
}

func AddParticipator(title, username string) error {
	return nil
}
