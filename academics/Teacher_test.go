package academics

import "testing"

func TestTeacher_StudentGrade(t *testing.T) {
	var TestStudentGrade = []struct {
		FirstInput  Teacher
		SecondInput int

		Output string
	}{
		{Teacher{300, []int{10, 20, 12, 3}, "Wayne"}, 280, "Student passes the subject"},
		{Teacher{250, []int{10, 20, 12, 3}, "Wayne"}, 280, "Student Fails the subject"},
	}
	for _, value := range TestStudentGrade {
		functionResult := value.FirstInput.StudentGrade(value.SecondInput)
		if value.Output != functionResult {
			t.Errorf("expected %v, got %v", value.Output, functionResult)
		}
	}
}
