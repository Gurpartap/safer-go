package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"go/format"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

type generator struct {
	packageName    string
	importPackages []string
	wrapperName    string
	wrappedName    string
}

func (g *generator) generate() ([]byte, error) {
	filename := "./template/optional.go.tmpl"

	funcMap := template.FuncMap{"title": strings.Title}

	t := template.Must(template.New("optional.go.tmpl").Funcs(funcMap).ParseFiles(filename))

	data := struct {
		Timestamp      time.Time
		PackageName    string
		ImportPackages []string
		WrapperName    string
		WrappedName    string
	}{
		time.Now().UTC(),
		g.packageName,
		g.importPackages,
		g.wrapperName,
		g.wrappedName,
	}

	var buf bytes.Buffer

	err := t.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	src, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return src, nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("optional: ")

	wrapperName := flag.String("wrapper", "", "wrapper type and file name; default [o|O]ptional<type> and srcdir/optional_<type>.go")
	wrappedName := flag.String("wrapped", "", "wrapped type name; must be set")

	flag.Parse()

	if len(*wrappedName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	pkg, err := build.Default.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}

	var (
		filename string
		g        generator
	)

	g.wrappedName = *wrappedName
	g.packageName = pkg.Name

	if len(*wrapperName) == 0 {
		// no output specified, use default optional_<type>

		// TODO: may not be the most reliable method
		exported := strings.Title(g.wrappedName) == g.wrappedName

		if exported {
			g.wrapperName = "Optional" + strings.Title(g.wrappedName)
		} else {
			g.wrapperName = "optional" + strings.Title(g.wrappedName)
		}
		filename = fmt.Sprintf("%s.go", to_snake_case(g.wrapperName))
	} else {
		g.wrapperName = *wrapperName
		filename = to_snake_case(g.wrapperName) + ".go"
	}

	filename = strings.Replace(filename, "u_int", "uint", -1)

	if strings.Contains(g.wrappedName, "timestamp.Timestamp") {
		g.importPackages = append(g.importPackages, "github.com/golang/protobuf/ptypes/timestamp")
	}

	src, err := g.generate()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filename, src, 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

// to_snake_case convert the given string to snake case following the Golang format:
// acronyms are converted to lower-case and preceded by an underscore.
func to_snake_case(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		out = append(out, unicode.ToLower(runes[i]))
	}

	return string(out)
}
