/* Package dvdfile implements functions for the manipulation of file. It provides several functions to transform the file content of []byte into [][]byte with line concept or [][][]byte with parameter concept. */
/* dvdfile包实现了一些操纵文件的函数。这些函数可以把[]byte类型的内容分割为行概念的[][]byte，或行和段落概念的[][][]byte类型。*/
package dvdfile

var eol = byte('\n')

//TransToLines parse text of type []byte into lines of type [][]byte. Each member in lines is a line which ends at '\n'. Allocate no new memory and the modification made to lines is visible to text.
func TransToLines(text []byte) (lines [][]byte) {
    start, end := 0, 0
    lines = make([][]byte, 0)
    for i, v := range text {
        if v == eol {
            end = i+1
            lines = append(lines, text[start:end])
            start = i+1
        }
    }
    if end != len(text) { //On Linux, every line contains a '\n'. But on windows, the last line contains no eol.
        lines = append(lines, text[end:])
    }
    return
}

//TransToParas transforms text into paras with line concept. i.e. paras[0] indicate first paragraph in text; paras[0][0] indicate first line in first paragraph in text. Paras are separated by empty line. empty line is the line that contains only a '\n'. Each para also contains either one or several empty line appears below it. As for last para, it contains no empty line if no empty para appears below it.
func TransToParas(text []byte) (paras [][][]byte) {
    lines := TransToLines(text)
    var emptyLine = []byte("\r\n")
    for i:=0; i<len(lines)-1; i++ { //don't consider lastline
        line := lines[i]
        if line[len(line)-2] != '\r' {
            emptyLine = []byte("\n")
            break
        }
    }
    start, end := 0, 0
    paras = make([][][]byte, 0)
    for i:=0; i<len(lines); i++ {
        if sliceEq(lines[i], emptyLine) {
            for i<len(lines)-1 {
                i += 1
                if !sliceEq(lines[i], emptyLine) {
                    i -= 1
                    break
                }
            }
            end = i+1
            paras = append(paras, lines[start:end])
            start = i+1
            i += 1 //Skip next "if" statement check, because it has been checked.
        }
    }
    if end != len(lines) {
        paras = append(paras, lines[end:])
    }
    return
}

//TransToParasNoLine transforms text into paras without line concept. i.e. paras[0] indicates the first paragraph. This function is provided as a easier way to access character from paragraph without access line previously.  Allocate no new memory and the modification made to lines is visible to text. This function is designed for regex match in over multiple lines.
func TransToParasNoLine(text []byte) (paras [][]byte) {
    paras3d := TransToParas(text)
    paras = make([][]byte, len(paras3d))
    for i, para := range paras3d {
        for j:=1; j<len(para); j++ {
            para[0] = append(para[0], para[j]...)
        }
        paras[i] = para[0]
    }
    return
}

//sliceEq checks whether slice s1 equals to slice s2.
func sliceEq(s1, s2 []byte) bool {
    if len(s1) != len(s2) {
        return false
    }
    for i, v := range s1 {
        if v != s2[i] {
            return false
        }
    }
    return true
}
