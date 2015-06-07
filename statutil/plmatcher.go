package statutil

import (
	"log"
	"strings"
)

func init() {
	List_Of_Programming_Languages_A = []string{
		"a# .net", "a# (axiom)", "a-0 system", "a+", "a++", "abap", "abc", "abc algol", "able",
		"abset", "absys", "acc", "accent", "ace dasl", "acl2", "act-iii", "action!", "actionscript",
		"ada", "adenine", "agda", "agilent vee", "agora", "aimms", "alef", "alf", "algol 58", "algol 60",
		"algol 68", "algol w", "alice", "alma-0", "ambienttalk", "amiga e", "amos", "ampl",
		"apl", "app inventor for android's visual block language", "applescript", "arc", "arexx", "argus", "aspectj",
		"assembly language", "ats", "ateji px", "autohotkey", "autocoder", "autoit", "autolisp / visual lisp", "averest",
		"awk", "axum", "android",
	}

	List_Of_Programming_Languages_B = []string{
		"b", "babbage", "bash", "basic", "bc",
		"bcpl", "beanshell", "batch (windows/dos)", "bertrand", "beta", "bigwig",
		"bistro", "bitc", "bliss", "blue", "bon", "boo", "boomerang",
		"bourne shell (including bash and ksh)", "brew", "bpel",
	}

	List_Of_Programming_Languages_C = []string{
		"c", "c--", "c++", "c# - iso/iec 23270",
		"c/al", "caché objectscript", "c shell", "caml", "candle", "cayenne", "cduce",
		"cecil", "cel", "cesil", "ceylon", "cfengine", "cfml", "cg", "ch", "chapel", "chain", "charity", "charm",
		"chef", "chill", "chip-8", "chomski", "chuck", "cics", "cilk", "cl (ibm)", "claire", "clarion", "clean", "clipper",
		"clist", "clojure", "clu", "cms-2", "cobol - iso/iec 1989", "cobra", "code",
		"coffeescript", "cola", "coldc", "coldfusion", "comal", "combined programming language (cpl)", "comit",
		"common intermediate language (cil)", "common lisp (also known as cl)", "compass", "component pascal", "constraint handling rules (chr)",
		"converge", "cool", "coq", "coral 66", "corn", "corvision", "cowsel", "cpl", "csh", "csp", "csound",
		"cuda", "curl", "curry", "cyclone", "cython", "c#",
	}
	List_Of_Programming_Languages_D = []string{
		"d", "dasl (datapoint's advanced systems language)", "dasl (distributed application specification language)",
		"dart", "dataflex", "datalog", "datatrieve", "dbase", "dc", "dcl", "deesel (formerly g)", "delphi", "dinkc",
		"dibol", "dog", "draco", "drakon", "dylan", "dynamo",
	}
	List_Of_Programming_Languages_E = []string{
		"e", "e#", "ease", "easy pl/i", "easy programming language",
		"easytrieve plus", "ecmascript", "edinburgh imp", "egl", "eiffel", "elan", "elixir", "elm", "emacs lisp", "emerald", "epigram", "epl",
		"erlang", "es", "escapade", "escher", "espol", "esterel", "etoys", "euclid",
		"euler", "euphoria", "euslisp robot programming language", "ecms exec", "exec 2", "executable uml",
	}

	List_Of_Programming_Languages_F_G_H = []string{
		"f", "f#", "factor", "falcon", "fancy", "fantom", "faust", "felix", "ferite", "ffp", "fjölnir", "fl", "flavors",
		"flex", "flow-matic", "focal", "focus", "foil", "formac", "@formula", "forth", "fortran - iso/iec 1539", "fortress", "foxbase", "foxpro", "fp",
		"fpr", "franz lisp", "frege", "f-script", "fsprog", "g", "google apps script", "game maker language", "gamemonkey script",
		"gams", "gap", "g-code", "genie", "gdl", "gibiane", "gj", "george", "glsl", "gnu e", "gm", "go", "go!", "goal",
		"gödel", "godiva", "gom (good old mad)", "goo", "gosu", "gotran", "gpss", "graphtalk", "grass", "groovy", "hack (programming language)", "hal/s",
		"hamilton c shell", "harbour", "hartmann pipelines", "haskell", "haxe", "high level assembly", "hlsl", "hop", "hope",
		"hugo", "hume", "hypertalk",
	}

	List_Of_Programming_Languages_I_J_K_L = []string{
		"ibm basic assembly language", "ibm hascript", "ibm informix-4gl", "ibm rpg", "ici", "icon", "id",
		"idl", "idris", "imp", "inform", "io", "ioke", "ipl", "iptscrae", "islisp", "ispf", "iswim", "j", "j#", "j++",
		"jade", "jako", "jal", "janus", "jass", "java", "javascript", "jcl", "jean", "join java", "joss", "joule",
		"jovial", "joy", "jscript", "jscript .net", "javafx script", "julia", "jython", "k", "kaleidoscope", "karel",
		"karel++", "kee", "kixtart", "kif", "kojo", "kotlin", "krc", "krl", "krl (kuka robot language)", "krypton", "ksh",
		"l", "l# .net", "labview", "ladder", "lagoona", "lansa", "lasso", "latex", "lava", "lc-3", "leda", "legoscript",
		"lil", "lilypond", "limbo", "limnor", "linc", "lingo", "linoleum", "lis", "lisa",
		"lisaac", "lisp - iso/iec 13816", "lite-c", "lithe", "little b", "logo", "logtalk", "lpc",
		"lse", "lsl", "livecode", "livescript", "lua", "lucid", "lustre", "lyapas", "lynx",
	}
	List_Of_Programming_Languages_M = []string{
		"m2001", "m4", "machine code", "mad (michigan algorithm decoder)", "mad/i", "magik", "magma", "make", "maple", "mapper now part of bis",
		"mark-iv now vision:builder", "mary", "masm microsoft assembly x86", "mathematica", "matlab", "maxima (see also macsyma)",
		"max (max msp - graphical programming environment)", "maxscript internal language 3d studio max", "maya (mel)", "mdl", "mercury",
		"mesa", "metacard", "metafont", "metal", "microcode", "microscript", "miis", "millscript", "mimic", "mirah", "miranda",
		"miva script", "ml", "moby", "model 204", "modelica", "modula", "modula-2", "modula-3", "mohol", "moo", "mortran", "mouse", "mpd",
		"msil - deprecated name for cil", "msl", "mumps",
	}

	List_Of_Programming_Languages_N_O = []string{
		"nasm", "natural", "napier88", "neko", "nemerle", "nesc",
		"nesl", "net.data", "netlogo", "netrexx", "newlisp", "newp", "newspeak", "newtonscript", "ngl", "nial", "nice",
		"nickle", "nim", "npl", "not exactly c (nxc)", "not quite c (nqc)", "nsis", "nu", "nwscript", "nxt-g",
		"o:xml", "oak", "oberon", "obix", "obj2", "object lisp", "objectlogo", "object rexx", "object pascal", "objective-c",
		"objective-j", "obliq", "obol", "ocaml", "occam", "occam-π", "octave", "omnimark", "onyx", "opa", "opal", "opencl",
		"openedge abl", "opl", "ops5", "optimj", "orc", "orca/modula-2", "oriel", "orwell", "oxygene", "oz",
	}

	List_Of_Programming_Languages_P_Q_R = []string{
		"p#", "parasail", "pari/gp", "pascal - iso 7185", "pawn", "pcastl", "pcf", "pearl", "peoplecode",
		"perl", "pdl", "php", "phrogram", "pico", "picolisp", "pict", "pike", "pikt", "pilot", "pipelines", "pizza", "pl-11", "pl/0",
		"pl/b", "pl/c", "pl/i - iso 6160", "pl/m", "pl/p", "pl/sql", "pl360", "planc", "plankalkül", "planner", "plex", "plexil", "plus",
		"pop-11", "postscript", "portable", "powerhouse", "powerbuilder - 4gl gui appl. generator from sybase", "powershell", "ppl",
		"processing", "processing.js", "prograph", "proiv", "prolog", "promal", "promela", "prose modeling language", "protel", "providex",
		"pro*c", "pure", "python", "q (equational programming language)", "q (programming language from kx systems)", "qalb", "qtscript", "quakec", "qpl", "r",
		"r++", "racket", "rapid", "rapira", "ratfiv", "ratfor", "rc", "rebol", "red", "redcode", "refal", "reia", "revolution", "rex",
		"rexx", "rlab", "robotc", "roop", "rpg", "rpl", "rsl", "rtl/2", "ruby", "runescript", "rust",
	}

	List_Of_Programming_Languages_S_T_U_V_W_X_Y_Z = []string{
		"s", "s2", "s3", "s-lang", "s-plus",
		"sa-c", "sabretalk", "sail", "salsa", "sam76", "sas", "sasl", "sather", "sawzall", "sbl", "scala", "scheme", "scilab",
		"scratch", "script.net", "sed", "seed7", "self", "sensetalk", "sequencel", "setl", "shift script", "simpol", "signal", "simple", "simscript", "simula", "simulink",
		"sisal", "slip", "small", "smalltalk", "small basic", "sml", "snap!", "snobol(spitbol)", "snowball", "sol", "span", "spark",
		"speedcode", "spin", "sp/k", "sps", "squeak", "squirrel", "sr", "s/sl", "stackless python", "starlogo", "strand", "stata", "stateflow", "subtext",
		"supercollider", "supertalk", "swift (apple programming language)", "swift (parallel scripting language)", "sympl", "synccharts",
		"systemverilog", "t", "tacl", "tacpol", "tads", "tal", "tcl", "tea", "teco", "telcomp", "tex", "tex", "tie", "timber", "tmg, compiler-compiler", "tom",
		"tom", "topspeed", "tpu", "trac", "ttm", "t-sql", "ttcn", "turing", "tutor", "txl", "typescript", "turbo c++", "ubercode",
		"ucsd pascal", "umple", "unicon", "uniface", "unity", "unix shell", "unrealscript", "vala", "vba", "vbscript", "verilog",
		"vhdl", "visual basic", "visual basic .net", "visual dataflex", "visual dialogscript", "visual fortran", "visual foxpro", "visual j++", "visual j#",
		"visual objects", "visual prolog", "vsxu", "vvvv", "watfiv, watfor", "webdna", "webql", "windows powershell", "winbatch",
		"wolfram", "wyvern", "x++", "x#", "x10", "xbl", "xc (exploits xmos architecture)", "xharbour", "xl", "xojo", "xotcl", "xpl",
		"xpl0", "xquery", "xsb", "xslt - see xpath", "xtend", "yorick", "yql", "z notation", "zeno",
	}

	Spicial_Programming_Languages = []string{
		".net",
	}

}

var List_Of_Programming_Languages_A []string
var List_Of_Programming_Languages_B []string
var List_Of_Programming_Languages_C []string
var List_Of_Programming_Languages_D []string
var List_Of_Programming_Languages_E []string
var List_Of_Programming_Languages_F_G_H []string
var List_Of_Programming_Languages_I_J_K_L []string
var List_Of_Programming_Languages_M []string
var List_Of_Programming_Languages_N_O []string
var List_Of_Programming_Languages_P_Q_R []string
var List_Of_Programming_Languages_S_T_U_V_W_X_Y_Z []string
var Spicial_Programming_Languages []string

func Analyse(str string) (count map[string]int) {
	count = make(map[string]int)
	str = strings.ToLower(str)
	strs := strings.Split(str, " ")
	if len(strs) > 0 {
		for _, word := range strs {
			arr := getPLList(strings.ToLower(word))
			if len(arr) != 0 && match(word, arr) {
				if count != nil && count[word] != -1 {
					count[word] = count[word] + 1
				} else {
					log.Println("...")
					count[word] = 1
				}
			}
		}
	}
	return
}

func getPLList(w string) []string {
	log.Println("process word: ", w)
	if len(w) == 0 {
		return nil
	}
	switch w[0] {
	case 'a':
		return List_Of_Programming_Languages_A
	case 'b':
		return List_Of_Programming_Languages_B
	case 'c':
		return List_Of_Programming_Languages_C
	case 'd':
		return List_Of_Programming_Languages_D
	case 'e':
		return List_Of_Programming_Languages_E
	case 'f', 'g', 'h':
		return List_Of_Programming_Languages_F_G_H
	case 'i', 'j', 'k', 'l':
		return List_Of_Programming_Languages_I_J_K_L
	case 'm':
		return List_Of_Programming_Languages_M
	case 'n', 'o':
		return List_Of_Programming_Languages_N_O
	case 'p', 'q', 'r':
		return List_Of_Programming_Languages_P_Q_R
	case 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
		return List_Of_Programming_Languages_S_T_U_V_W_X_Y_Z
	default:
		return Spicial_Programming_Languages
	}
}

func match(str string, arr []string) bool {
	for _, n := range arr {
		if str == n {
			return true
		}
	}
	return false
}
