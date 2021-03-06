package parser

import (
	"bytes"
	"encoding/json"
	"github.com/lyraproj/issue/issue"
	"testing"
)

func TestManifest(t *testing.T) {
	expectJSON(t, issue.Unindent(`
      file { '/tmp/foo':
        mode => '0640',
        ensure => present;
      '/tmp/bar':
        mode => '0640',
        ensure => present;
      }

      $rootgroup = $facts['os']['family'] ? 'Solaris' => 'wheel'

      function foo(Integer[2,3] $in, String $n = 'vi') >> Float[0.0] {
          notice("show the ${n}")
        $in * 3.14
      }`),
		`{"^":["block",{"^":["resource",{"#":["type",{"^":["qn","file"]},"bodies",[{"#":["title","/tmp/foo","ops",[{"^":["=>","mode","0640"]},{"^":["=>","ensure",{"^":["qn","present"]}]}]]},{"#":["title","/tmp/bar","ops",[{"^":["=>","mode","0640"]},{"^":["=>","ensure",{"^":["qn","present"]}]}]]}]]}]},{"^":["=",{"^":["var","rootgroup"]},{"^":["?",{"^":["access",{"^":["access",{"^":["var","facts"]},"os"]},"family"]},[{"^":["=>","Solaris","wheel"]}]]}]},{"^":["function",{"#":["name","foo","params",{"#":["in",{"#":["type",{"^":["access",{"^":["qr","Integer"]},2,3]}]},"n",{"#":["type",{"^":["qr","String"]},"value","vi"]}]},"body",[{"^":["invoke",{"#":["functor",{"^":["qn","notice"]},"args",[{"^":["concat","show the ",{"^":["str",{"^":["var","n"]}]}]}]]}]},{"^":["*",{"^":["var","in"]},3.14]}],"returns",{"^":["access",{"^":["qr","Float"]},0]}]}]}]}`)
}

func toJSON(e Expression) string {
	result := bytes.NewBufferString(``)
	enc := json.NewEncoder(result)
	enc.SetEscapeHTML(false)
	enc.Encode(e.ToPN().ToData())
	result.Truncate(result.Len() - 1)
	return result.String()
}

func expectJSON(t *testing.T, source string, expected string) {
	expr, err := CreateParser().Parse(``, source, false)
	if err != nil {
		t.Errorf(err.Error())
	} else {
		actual := toJSON(expr)
		if expected != actual {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	}
}
