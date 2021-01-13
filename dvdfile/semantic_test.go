package dvdfile

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestTransToLines(t *testing.T) {
    text := `    para1, line1
    para1, line2

`
    wants := [][]byte{
        []byte("    para1, line1\n"),
        []byte("    para1, line2\n"),
        []byte("\n"),
    }
    gotLines := TransToLines([]byte(text))
    t.Log("Given the need to transform text into lines no mem.")
    {
        for i, gotLine := range gotLines {
            if !sliceEq(wants[i], gotLine) {
                t.Logf("want:%s get:%s", wants[i], gotLine)
            }
        }
    }
}

func TestTransToLinesMem(t *testing.T) {
    text := []byte(`    para1, line1
    para1, line2

`)
    wants := [][]byte{
        []byte("    para1, line1\n"),
        []byte("    para1, line2\n"),
        []byte("\n"),
    }
    gotLines := TransToLinesMem(text)
    t.Log("Given the need to transform text into lines with mem.")
    {
        t.Log("\t\tTest the content is correctly parsed")
        {
            for i, gotLine := range gotLines {
                if !sliceEq(wants[i], gotLine) {
                    t.Logf("want:%s get:%s", wants[i], gotLine)
                }
            }
        }
        t.Log("\t\tTest the memory is correctly allocate")
        {
            gotLinesModify := append(gotLines, []byte("New line"))
            buf := make([]byte, 0)
            for _, line := range gotLinesModify {
                buf = append(buf, line...)
            }
            if sliceEq(buf, text) {
                t.Log("memory allocate incorrectly")
            }
        }
    }
}

func TestTransToParasMem(t *testing.T) {
    text := []byte(`    para1, line1


    para2, line1

`)
    wants := [][][]byte{
        [][]byte{
            []byte("    para1, line1\n"),
            []byte("\n"),
            []byte("\n"),
        },
        [][]byte{
            []byte("    para2, line1\n"),
            []byte("\n"),
        },
    }
    gots := TransToParasMem(text)
    t.Log("Given the need for transform text into parameters")
    {
        for i, gotPara := range gots {
            for j, gotLine := range gotPara {
                if !sliceEq(wants[i][j], gotLine) {
                    t.Logf("want %s get %s", wants[i][j], gotLine)
                }
            }
        }
    }
}
