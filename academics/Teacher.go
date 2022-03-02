package academics

type Teacher struct {
	gradeStudent   int
	markAttendance []int
	studentName    string
}

func (t *Teacher) StudentGrade(cutoffPoint int) string {
	if t.gradeStudent >= cutoffPoint {
		return "Student passes the subject"
	}
	return "Student Fails the subject"
}
