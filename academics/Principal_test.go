package academics

import "testing"

func TestPrincipal_Suspend(t *testing.T) {
	suspendStudent := []struct {
		firstInput Principal
		inputFunc  int
		result     string
	}{
		{Principal{"Edmund", "Nil", 6, 0}, 7, "Edmund is hereby suspended"},
		{Principal{"Anne", "Nil", 6, 0}, 7, "Anne is hereby suspended"},
		{Principal{"Lisa", "Nil", 11, 0}, 7, "Lisa is free to attend classes"},
		{Principal{"Veronica", "Nil", 11, 0}, 7, "Veronica is free to attend classes"},
	}
	for _, value := range suspendStudent {
		functionResult := value.firstInput.Suspend(value.inputFunc)
		if functionResult != value.result {
			t.Errorf("function expected %v, but Got %v", value.result, functionResult)
		}
	}

}
func TestPrincipal_Correction(t *testing.T) {
	correctTeacher := []struct {
		firstInput Principal
		inputFunc  string
		result     string
	}{
		{Principal{"Edmund", "Nil", 6, 3}, "joseph", "joseph is hereby queried"},
		{Principal{"Edmund", "Nil", 6, 3}, "Olu", "olu is hereby queried"},
		{Principal{"Veronica", "Nil", 11, 0}, "magret", "Good Job"},
		{Principal{"Veronica", "Nil", 11, 0}, "victor", "Good Job"},
	}
	for _, value := range correctTeacher {
		functionResult := value.firstInput.Correction(value.inputFunc)
		if functionResult != value.result {
			t.Errorf("function expected %v, but Got %v", value.result, functionResult)
		}
	}

}
