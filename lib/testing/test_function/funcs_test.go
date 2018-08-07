package ttt

import "testing"

func TestGetData1(t *testing.T) {
	if GetData1(1) != "1" {
		t.Fatal("1!=1")
	} else {
		t.Log("GetData1 OK")
	}
	t.Log(testing.Coverage())
}

func TestGetData2(t *testing.T) {
	if GetData2(2) != 2 {
		t.Fatal("2!=2")
	} else {
		t.Log("GetData2 OK")
	}
	t.Log(testing.Coverage())
}

func TestOther1(t *testing.T) {
	t.Log(t.Failed())
	t.Log(testing.Coverage())
	//success
}

/*
func TestOther2(t *testing.T) {
	t.Error("error ocur")
	t.Log("executr here")
	//success
}

func TestOther3(t *testing.T) {
	t.Fatal("error ocur")
	t.Log("can not executr here") // This is the difference between Error and Fatal
	//success
}

func TestOther4(t *testing.T) {
	t.FailNow()
}
*/

func Test2_1(t *testing.T) {
	//t.Parallel()
	t.Log(testing.Coverage())
}

func Test2_2(t *testing.T) {
	//t.Parallel()
	t.Log(testing.Coverage())
}

func Test2_3(t *testing.T) {
	t.Log(testing.Coverage())
	t.Skip("not testing it now")
}
