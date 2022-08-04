package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/cucumber/common/gherkin/go/v24"
	messages "github.com/cucumber/common/messages/go/v19"
	"github.com/spf13/cobra"
)

type fileContext struct {
	Steps []Step
}

type Step struct {
	Kind string
	Text string
	Vars string
	Body string
}

var rootCmd = &cobra.Command{
	Use:  "gherkinator",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		feat, err := features(args[0])
		if err != nil {
			return err
		}

		tpl, err := template.New("cypress").Funcs(map[string]interface{}{}).ParseFiles("template/cypress")
		if err != nil {
			return err
		}

		for path, doc := range feat {

			scns := []*messages.Scenario{}

			for _, v := range doc.Feature.Children {
				if v.Scenario != nil {
					if v.Scenario.Keyword == "Scenario Outline" {
						scns = append(scns, expandScenarioOutline(v.Scenario)...)
						continue
					}
					scns = append(scns, v.Scenario)
				}
			}

			ctx := mkFileCtx(scns)

			outPath := strings.TrimSuffix(path, "feature") + "ts"
			os.MkdirAll(filepath.Dir(outPath), 0755)

			if err := writeFile(ctx, outPath, tpl); err != nil {
				return err
			}
		}

		return nil
	},
}

func writeFile(ctx *fileContext, path string, tpl *template.Template) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return tpl.ExecuteTemplate(f, "cypress", ctx)
}

var matchString = regexp.MustCompile(`"[^"]+"`)
var matchFloat = regexp.MustCompile(`\d+\.\d+`)
var matchInt = regexp.MustCompile(`[-]??\d+`)
var matchVar = regexp.MustCompile(`\{(.*?)\}`)

var keywordType = map[messages.StepKeywordType]string{
	messages.StepKeywordType_CONTEXT: "Given",
	messages.StepKeywordType_ACTION:  "When",
	messages.StepKeywordType_OUTCOME: "Then",
}

var tsType = map[string]string{
	"string": "string",
	"int":    "number",
	"float":  "number",
}

func toChar(i int) rune {
	return rune('a' + i)
}

func mkFileCtx(scns []*messages.Scenario) *fileContext {
	ctx := &fileContext{}
	m := map[string]map[string]struct {
		vars string
		body string
	}{
		"Given": map[string]struct {
			vars string
			body string
		}{},
		"When": map[string]struct {
			vars string
			body string
		}{},
		"Then": map[string]struct {
			vars string
			body string
		}{},
	}
	for _, s := range scns {
		lastKeyword := messages.StepKeywordType_CONTEXT
		for _, stp := range s.Steps {
			str := stp.Text
			str = matchString.ReplaceAllString(str, "{string}")
			str = matchFloat.ReplaceAllString(str, "{float}")
			str = matchInt.ReplaceAllString(str, "{int}")
			vars := strings.Builder{}
			body := strings.Builder{}
			mtch := matchVar.FindAllStringSubmatch(str, -1)
			for i, v := range mtch {
				if i != 0 {
					vars.WriteString(", ")
				}
				vars.WriteString(fmt.Sprintf("%c: %s", toChar(i), tsType[v[1]]))
				body.WriteString(", ")
				body.WriteString(fmt.Sprintf("%c", toChar(i)))
			}
			if stp.DataTable != nil {
				if mtch != nil {
					vars.WriteString(", ")
				}
				vars.WriteString("dt: DataTable")
				body.WriteString(", ")
				body.WriteString("dt")
			}
			if stp.DocString != nil {
				if mtch != nil {
					vars.WriteString(", ")
				}
				vars.WriteString("ds: string")
				body.WriteString(", ")
				body.WriteString("ds")
			}
			if stp.KeywordType == messages.StepKeywordType_CONJUNCTION {
				stp.KeywordType = lastKeyword
			}
			m[keywordType[stp.KeywordType]][str] = struct {
				vars string
				body string
			}{
				vars: vars.String(),
				body: fmt.Sprintf("console.log('%s'%s)", stp.Text, body.String()),
			}
			lastKeyword = stp.KeywordType
		}
	}
	for k, v := range m {
		for i, j := range v {
			ctx.Steps = append(ctx.Steps, Step{
				Kind: k,
				Text: i,
				Vars: j.vars,
				Body: j.body,
			})
		}
	}
	return ctx
}

func main() {
	rootCmd.Execute()
}

func features(s string) (map[string]*messages.GherkinDocument, error) {

	out := map[string]*messages.GherkinDocument{}
	s = filepath.Clean(s)

	err := filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.HasSuffix(path, ".feature") {
			doc, err := parseFeatureFile(path)
			if err != nil {
				return err
			}
			out[path] = doc
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

func parseFeatureFile(path string) (*messages.GherkinDocument, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	doc, err := gherkin.ParseGherkinDocument(f, (&messages.Incrementing{}).NewId)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func expandScenarioOutline(so *messages.Scenario) []*messages.Scenario {

	ret := []*messages.Scenario{}

	for _, ex := range so.Examples {

		m := examplesMap(ex)

		for i := range ex.TableBody {

			s := &messages.Scenario{
				Name:    fmt.Sprintf("%s - (%d)", so.Name, i),
				Tags:    so.Tags,
				Keyword: "scenario",
			}

			for _, stp := range so.Steps {
				s.Steps = append(s.Steps, &messages.Step{
					Keyword:     stp.Keyword,
					KeywordType: stp.KeywordType,
					DocString:   stp.DocString,
					Text:        replacePlaceholder(stp.Text, m[i]),
					DataTable:   replacePlaceholderDataTable(stp.DataTable, m[i]),
				})
			}
			ret = append(ret, s)
		}
		return ret
	}

	return nil
}

func examplesMap(ex *messages.Examples) map[int]map[string]string {
	m := map[int]map[string]string{}
	cols := map[int]string{}

	for i, h := range ex.TableHeader.Cells {
		cols[i] = h.Value
	}

	for ri, r := range ex.TableBody {
		m[ri] = map[string]string{}
		for ci, c := range r.Cells {
			m[ri][cols[ci]] = c.Value
		}
	}

	return m
}

func replacePlaceholder(s string, m map[string]string) string {
	bits := strings.Split(s, " ")
	for i, v := range bits {
		e := len(v) - 1
		if v[0] == '<' && v[e] == '>' {
			bits[i] = m[v[1:e]]
		}
	}
	return strings.Join(bits, " ")
}

func replacePlaceholderDataTable(dt *messages.DataTable, m map[string]string) *messages.DataTable {
	if dt == nil {
		return nil
	}
	ret := &messages.DataTable{}
	for _, r := range dt.Rows {
		cols := []*messages.TableCell{}
		for _, c := range r.Cells {
			cols = append(cols, &messages.TableCell{
				Value: replacePlaceholder(c.Value, m),
			})
		}
		ret.Rows = append(ret.Rows, &messages.TableRow{
			Cells: cols,
		})
	}
	return ret
}
