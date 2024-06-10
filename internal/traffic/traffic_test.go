package traffic

import "testing"

func TestLightStateString(t *testing.T) {
	testcases := []struct {
		input LightState
		want  string
	}{
		{
			input: LightRed,
			want:  "0 RED",
		},
		{
			input: LightRedAmber,
			want:  "1 RED AMBER",
		},
		{
			input: LightGreen,
			want:  "2 GREEN",
		},
		{
			input: LightAmber,
			want:  "3 AMBER",
		},
	}

	for i, tc := range testcases {
		got := tc.input.String()
		if tc.want != got {
			t.Fatalf("Case: %d Got: %v Want: %v", i, got, tc.want)
		}
	}
}

func TestLightTransition(t *testing.T) {
	testcases := []struct {
		input LightState
		want  LightState
	}{
		{
			input: LightRed,
			want:  LightRedAmber,
		},
		{
			input: LightRedAmber,
			want:  LightGreen,
		},
		{
			input: LightGreen,
			want:  LightAmber,
		},
		{
			input: LightAmber,
			want:  LightRed,
		},
	}

	for i, tc := range testcases {
		l := Light{
			state: tc.input,
		}
		l.Transition()
		got := l.State()
		if tc.want != got {
			t.Fatalf("Case: %d Got: %v Want: %v", i, got, tc.want)
		}
	}
}

func TestLightString(t *testing.T) {
	testcases := []struct {
		input LightState
		want  string
	}{
		{
			input: LightRed,
			want:  "0 RED",
		},
		{
			input: LightRedAmber,
			want:  "1 RED AMBER",
		},
		{
			input: LightGreen,
			want:  "2 GREEN",
		},
		{
			input: LightAmber,
			want:  "3 AMBER",
		},
	}

	for i, tc := range testcases {
		l := Light{
			state: tc.input,
		}
		got := l.String()
		if tc.want != got {
			t.Fatalf("Case: %d Got: %v Want: %v", i, got, tc.want)
		}
	}
}
