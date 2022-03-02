package academics

type Principal struct {
	StudentFirstName  string
	Expulsion         string
	Studentattendance int
	Teacherabsence    int
}

func (p *Principal) Suspend(daysPresent int) string {
	if p.Studentattendance < daysPresent {
		return p.StudentFirstName + " is hereby suspended"
	}
	return p.StudentFirstName + " is free to attend classes"
}

func (p *Principal) Correction(teacherName string) string {
	if p.Teacherabsence > 1 {
		return teacherName + " is hereby queried"
	}
	return "Good Job"
}
