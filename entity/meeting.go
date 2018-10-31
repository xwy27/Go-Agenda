package entity

type participator struct {
	username string
}

type Meeting struct {
	title         string
	sponsor       string
	participators []participator
	startTime     int
	endTime       int
}

type meetingsJSON struct {
	meetings []Meeting
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
