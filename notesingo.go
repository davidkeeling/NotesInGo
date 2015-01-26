package notesingo

import (
	"html/template"
	"net/http"
    "strings"
)

type ChordEntry struct {
    Root, ChordInfo string
}

var theChords []ChordEntry

var mainPage = template.Must(template.New("song").Parse(
`<html>
    <head>
        <style>
            .entry {display:inline-block; width:150px;}
        </style>
    </head>

    <body>
        <form action="/addchord" method="post">
            <div>
                <input name="content" size="50" placeholder="ex: C Maj7 #11, F 7, Bb Major" />
            </div>
            <div>
                <input type="submit" value="Add chord" />
            </div>
        </form>
        {{range .}}
            <div class="entry">
                <b>{{.Root}}</b>
                <span>{{.ChordInfo}}</span>
            </div>
            {{end}}
    </body>
</html>
`))


func init() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/addchord", addchord)
}

func handle(w http.ResponseWriter, r *http.Request) {
    mainPage.Execute(w, theChords)
}

func addchord(w http.ResponseWriter, r *http.Request) {
    theData := parseInput(r.FormValue("content"))
    if len(theData) == 0 {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    newChord := &ChordEntry{
		Root: theData[0],
		ChordInfo: theData[1],
	}
    theChords = append(theChords, *newChord)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func parseInput (input string) []string {
    if len(input) == 0 {
        return nil
    }
    
    s := strings.Split(input, " ")
    for i := 0; i < len(s); i++ {
        s[i] = strings.Title(s[i])
        if i > 1 {
            s[1] += " " + s[i]
        }
    }
    
    if len(s) == 1 {
        return append(s, " ")
    }
    
    return s
}
