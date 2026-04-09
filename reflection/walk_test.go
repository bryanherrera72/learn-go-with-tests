package reflection

import (
	"reflect"
	"testing"
)

type Person struct{
	Name string
	Profile Profile
}

type Profile struct{
	Age int
	City string
}

func TestWalk(t *testing.T){

	// test converted into a "table test"
	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",// Name
			struct{Name string}{"Chris"},// Input
			[]string{"Chris"}, // Expected Calls
		},
		{
			"struct with two string fields",
			struct{
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{
			"struct with non string field",
			struct{
				Name string
				Age int
			}{"Chris", 33},
			[]string {"Chris"},
		},
		{
			"nested fields",
			Person{"Chris", Profile{33, "London"}},
			[]string{"Chris", "London"},
		},
		{
			"pointers to things",
			&Person{ //"pointer to a Person"
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},
		{
			"arrays",
			[2]Profile {
				{33, "London"},
				{34, "Reykjavik"},
			},
			[]string{"London", "Reykjavik"},
		},//side note: always need a trailing comma
	}

	t.Run("with maps", func(t *testing.T) { // we've moved maps to their own tests, because they are likely unordered.
		aMap := map[string] string{
			"Cow": "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string){
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "baa")
	})

	t.Run("with channels", func(t *testing.T){
		aChannel := make(chan Profile)

		go func(){
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func()(Profile, Profile){
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string){
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want){
			t.Errorf("got %v, want %v", got, want)
		}
	})

	for _, test := range cases{
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string){
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls){
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}


}

func assertContains(t testing.TB, haystack []string, needle string){
	t.Helper()
	contains := false
	for _, x := range haystack{
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but id didn't", haystack, needle)
	}
}

func getValue(x interface{}) reflect.Value{
	val := reflect.ValueOf(x)

	if val.Kind()== reflect.Pointer{ // this is how we type check when we have a pointer
		val = val.Elem() // extracts the actual value the pointer points to
	}

	return val
}


func walk(x interface{}, fn func(input string)){
	val := getValue(x)

	walkValue := func(value reflect.Value){
		walk(value.Interface(), fn)
	}

	switch val.Kind(){
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i:=0; i < val.NumField(); i++{
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++{
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan: 
		for {
			if v, ok := val.Recv(); ok{// iterate through values of a channel with Recv. Disregards type of channel messages
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult{
			walkValue(res)
		}
	}
	
	//PRE REFACTOR. HAD TO CHANGE THIS WHEN WE INTRODUCED ARRAYS
	// numberOfValues := 0
	// var getField func(int) reflect.Value
	// switch val.Kind(){
	// case reflect.String:
	// 	fn(val.String())
	// case reflect.Struct:// get the number of fields in a struct, and their values
	// 	numberOfValues = val.NumField()
	// 	getField = val.Field // kind of confusing, but this method will return the ith element of a struct val
	// case reflect.Slice, reflect.Array:// get the number of elements in a slice. val.Index provides values per index of our slice / array
	// 	numberOfValues = val.Len()
	// 	getField = val.Index // kind of confusing, but this method will return the ith element of a slice val
	// case reflect.Map:
	// 	for _, key := range val.MapKeys() {
	// 		walk(val.MapIndex(key).Interface(), fn)
	// 	}
	// }


	// for i:=0; i < numberOfValues; i++{
	// 	//getField will return the ith element of "this" which in this case is "val"
	// 	walk(getField(i).Interface(), fn)
	// }
}

