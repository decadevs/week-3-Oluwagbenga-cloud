package academics

type Student struct {
	StudentName     string
	StudentLastName string
	StudentGrade    int
	SubjectOffered  string
}

func (s *Student) CourseOffered(subject string) string {
	if subject == s.SubjectOffered {
		return "You have been offered this course"
	}
	return "We're sorry you cannot be offered this course "
}
