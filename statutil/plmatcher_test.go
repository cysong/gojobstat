package statutil

import "testing"

type testpair struct {
	input string
	ouput map[string]int
}

var test = []testpair{
	{"junior java developer", map[string]int{"java": 1}},
	{"someone strongly experienced with PHP", map[string]int{"php": 1}},
	{"Java UI Designer", map[string]int{"java": 1}},
	{"Embedded C Application Developers - POS & Communications Products", map[string]int{"c": 1}},
	{"C# Web UI Developer", map[string]int{"c#": 1}},
	{"This established software house set in Auckland's CBD is seeking a Senior C# Developer to join their innovative and market leading team.", map[string]int{"c#": 1}},
	{".NET Web Developer", map[string]int{".net": 1}},
	{"Delphi / .NET Developer", map[string]int{".net": 1, "delphi": 1}},
	{"Senior C++ Developer", map[string]int{"c++": 1}},

	// test cases want to pass in the future
	// {"Senior C# .NET Developer", map[string]int{"c#": 1}}, // another evil, fuck!!!
	// {"junior .net developer", map[string]int{".net": 1}},  // evil .net, need create a link for all the things like this
	// {"Senior C#.NET Developer Contract", map[string]int{"c#": 1}},
	// {"Senior MS SQL Developer", map[string]int{"ms sql server": 1}},
	// {"Senior ASP.NET Developer", map[string]int{".net": 1}},
	// {"Senior .NET Developer (VB.NET & C#)", map[string]int{"c#": 1, "vb.net": 1}},
	// {"Winforms C#.NET Developer", map[string]int{"c#": 1}},
	// {"Senior.NET Developer", map[string]int{".net": 1}},
}

func TestAnalyse(t *testing.T) {
	for _, v := range test {
		m := Analyse(v.input)
		if len(m) == len(v.ouput) {
			for i, o := range m {
				if v.ouput[i] == -1 || v.ouput[i] != o {
					t.Error(
						"for \"", v.input,
						"\" expect (", i, ":", v.ouput[i],
						") got (", i, ":", o, ")",
					)
				}
			}
		} else {
			t.Error(
				"for", v.input,
				"expect", v.ouput,
				"got", m,
			)
		}
	}
}
