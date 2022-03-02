package academics

type Applicant struct {
	ApplicantAge   int
	ApplicantGrade int
	ApplicantName  string
}

func (a *Applicant) ScoreApplicant() int {
	return a.ApplicantGrade
}

func (a *Applicant) GainAdmission(score int) string {
	if score <= a.ApplicantGrade {
		return "Hello " + a.ApplicantName + " You have been offered admission into FGC"
	}
	return "We are so sorry for you have not met the requirements for admission into FGC"
}
