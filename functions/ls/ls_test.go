package goterminal

import (
	"reflect"
	"testing"
)

type MockListFiles struct{}

func (m MockListFiles) LS() []string {

	ret_value := []string{
		"a.txt",
		"b.txt",
	}

	return ret_value

}

func Test_LS_Mock(t *testing.T) {

	mockLister := MockListFiles{}

	got := ListFiles(mockLister)
	want := []string{"a.txt", "b.txt"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot  %s \nwat %s", got, want)
	}

}
