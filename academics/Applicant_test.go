package academics

import "testing"

func TestApplicant_GainAdmission(t *testing.T) {
	admitStudent := []struct {
		FirstInput Applicant
		inputFunc  int
		output     string
	}{
		{Applicant{12, 300, "Jonathan"}, 250, "Hello Jonathan You have been offered admission into FGC"},
		{Applicant{12, 300, "Jonathan"}, 350, "We are so sorry for you have not met the requirements for admission into FGC"},
	}
	for _, value := range admitStudent {
		functionOutput := value.FirstInput.GainAdmission(value.inputFunc)
		if functionOutput != value.output {
			t.Errorf("function expected %v, but got %v", value.output, functionOutput)
		}
	}

}
