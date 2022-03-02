package academics

import "testing"

func TestStudent_CourseOffered(t *testing.T) {
	courseoffered := []struct {
		firstInput Student
		inputFunc  string
		result     string
	}{
		{Student{"Edmund", "Timi", 81, "English"}, "English", "You have been offered this course"},
		{Student{"Edmund", "Timi", 81, "Biology"}, "Maths", "We're sorry you cannot be offered this course "},
	}
	for _, value := range courseoffered {
		functionResult := value.firstInput.CourseOffered(value.inputFunc)
		if functionResult != value.result {
			t.Errorf("function expected %v, but Got %v", value.result, functionResult)
		}
	}

}
