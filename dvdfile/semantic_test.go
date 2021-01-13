package dvdfile

import "testing"

func TestTransToLines(t *testing.T) {
    text := []byte(`    para1, line1
    para1, line2

`)
    wants := [][]byte{
        []byte("    para1, line1\n"),
        []byte("    para1, line2\n"),
        []byte("\n"),
    }
    gotLines := TransToLines(text)
    t.Log("Given the need to transform text into lines.")
    {
        for i, gotLine := range gotLines {
            if !sliceEq(wants[i], gotLine) {
                t.Errorf("want:%s get:%s", wants[i], gotLine)
            }
        }
    }
}

func TestTransToParas(t *testing.T) {
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
    gots := TransToParas(text)
    t.Log("Given the need for transform text into parameters")
    {
        for i, gotPara := range gots {
            for j, gotLine := range gotPara {
                if !sliceEq(wants[i][j], gotLine) {
                    t.Errorf("want %s get %s", wants[i][j], gotLine)
                }
            }
        }
    }
}

func TestTransToParasNoLine(t *testing.T) {
    text := []byte(`    para1, line1


    para2, line1

`)
    wants := [][]byte{
        []byte("    para1, line1\n\n\n"),
        []byte("    para2, line1\n\n"),
    }
    gots := TransToParasNoLine(text)
    t.Log("Given the need for transform text into parameters without line concept")
    {
        for i, gotPara := range gots {
            if !sliceEq(wants[i], gotPara) {
                t.Errorf("want %s get %s", wants[i], gotPara)
            }
        }
    }
}
