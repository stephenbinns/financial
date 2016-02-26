package financial

import (
	"testing"
)

func TestCalculatesPmt(t *testing.T) {
	pmt, err := Pmt(0.4, 12, -1000, 0, false)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if pmt != 407.1821135011055 {
		t.Error("Expected 407.1821135011055, got ", pmt)
	}
}

func TestCalculatesPmtWithZeroRate(t *testing.T) {
	pmt, err := Pmt(0, 10, -1000, 0, false)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if pmt != 100 {
		t.Error("Expected 100, got ", pmt)
	}
}

func TestPmtWithInvalidPeriodReturnsError(t *testing.T) {
	pmt, err := Pmt(0.4, 0, -1000, 0, false)

	if err == nil {
		t.Error("Did not receive expected error")
	}
	if pmt != 0 {
		t.Error("Expected 0, got ", pmt)
	}
}

func TestCalculatingPPmt(t *testing.T) {
	ppmt, err := PPmt(0.4, 1, 12, -1000)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if ppmt != 7.182113501105505 {
		t.Error("Expected 7.182113501105505, got ", ppmt)
	}
}

func TestCalculatingPPmtWithZeroRate(t *testing.T) {
	ppmt, err := PPmt(0, 1, 10, -1000)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if ppmt != 100 {
		t.Error("Expected 100, got ", ppmt)
	}
}

func TestCalculatingIPmt(t *testing.T) {
	ipmt, err := IPmt(0.4, 1, 12, -1000)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if ipmt != 400 {
		t.Error("Expected 400, got ", ipmt)
	}
}

func TestCalculatingIPmtWithZeroRate(t *testing.T) {
	ipmt, err := IPmt(0, 1, 12, -1000)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if ipmt != 0 {
		t.Error("Expected 0, got ", ipmt)
	}
}

func TestCalculatingIPmtInvalidCurrentPeriod(t *testing.T) {
	ipmt, err := IPmt(0, 12, 1, -1000)

	if err == nil {
		t.Error("Did not receive expected error")
	}
	if ipmt != 0 {
		t.Error("Expected 0, got ", ipmt)
	}
}

func TestCalculatingIPmtWithNegativePeriods(t *testing.T) {
	ipmt, err := IPmt(0, -1, -1, -1000)

	if err == nil {
		t.Error("Did not receive expected error")
	}
	if ipmt != 0 {
		t.Error("Expected 0, got ", ipmt)
	}
}

func TestCalculatingPVForEndOfPeriod(t *testing.T) {
	pmt, _ := Pmt(0.4, 12, 1000, 0, false)
	pv, err := Pv(0.4, 2, pmt, 0, true)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if pv != 698.0264802876093 {
		t.Error("Expected 698.0264802876093, got ", pv)
	}
}

func TestCalculatingPVForStartOfPeriod(t *testing.T) {
	pmt, _ := Pmt(0.4, 12, 1000, 0, false)
	pv, err := Pv(0.4, 3, pmt, 0, false)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if pv != 646.9803261169167 {
		t.Error("Expected 646.9803261169167, got ", pv)
	}
}

func TestPVWithZeroRate(t *testing.T) {
	pmt, _ := Pmt(0, 12, 1000, 0, false)
	pv, err := Pv(0, 3, pmt, 0, false)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if pv != 250 {
		t.Error("Expected 250, got ", pv)
	}
}

func TestFVWithNegativePeriods(t *testing.T) {
	fv, err := Fv(0.4, -1, 0, 1000)

	if err == nil {
		t.Error("Did not receive expected error")
	}
	if fv != 0 {
		t.Error("Expected 0, got ", fv)
	}
}

func TestCalculatesFV(t *testing.T) {
	pmt, _ := Pmt(0.4, 12, 1000, 0, false)
	fv, err := Fv(0.4, 12, pmt, 1000)

	if err != nil {
		t.Error("Did not expect an error to be returned")
	}
	if fv != 0 {
		t.Error("Expected 0, got ", fv)
	}
}
