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
}

var rootCmd = &cobra.Command{
  Use: "gherkinator",
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

    outDir := "gen"

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


    }

    return nil
  },
}

var matchString = regexp.MustCompile(`"[^"]+"`)
var matchFloat = regexp.MustCompile(`\d+\.\d+`)
var matchInt = regexp.MustCompile(`[-]??\d+`)


func mkFileCtx(scns []*messages.Scenario) *fileContext {
  ctx := fileContext{}
  for _, s := range scns {

  }
  return &fileContext{}
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
        Name: fmt.Sprintf("%s - (%d)", so.Name, i),
        Tags: so.Tags,
        Keyword: "scenario",
      }

      for _, stp := range so.Steps {
        s.Steps = append(s.Steps, &messages.Step{
         Keyword: stp.Keyword,
         KeywordType: stp.KeywordType,
         DocString: stp.DocString,
         Text: replacePlaceholder(stp.Text, m[i]), 
         DataTable: replacePlaceholderDataTable(stp.DataTable, m[i]),
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
    e := len(v)-1
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

