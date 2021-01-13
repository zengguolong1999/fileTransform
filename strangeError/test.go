package main

import "fmt"
import "MrLiu_filetransform/dvdfile"

var testTxt = `Hello,
    this is zgl.
    How are you?


    I'm fine.
    Thank you.`

//底层存储结构相同，由编程语言来做类型检查，于是可以分为强类型，弱类型语言
func main() {
      parasNoLine := dvdfile.TransToParasMemNoLine([]byte(testTxt))
      fmt.Println("parasNoLine:")
      for _, v := range parasNoLine {
          fmt.Printf("%s", v)
      }
//para test:
//    paras := dvdfile.TransToParas([]byte(testTxt))
//    parasMem := dvdfile.TransToParasMem([]byte(testTxt))
//    fmt.Println("paras:")
//    for _, v := range paras {
//        for _, line := range v {
//            fmt.Printf("%s%s", "% ", line)
//        }
//    }
//    paras[0] = append(paras[0], []byte("MIB"))
//    fmt.Println("paras change:")
//    for _, v := range paras {
//        for _, line := range v {
//            fmt.Printf("%s%s", "% ", line)
//        }
//    }
//    fmt.Println("parasMem:")
//    for _, v := range parasMem {
//        for _, line := range v {
//            fmt.Printf("%s%s", "% ", line)
//        }
//    }
//    parasMem[0] = append(parasMem[0], []byte("MIB"))
//    fmt.Println("parasMem change:")
//    for _, v := range parasMem {
//        for _, line := range v {
//            fmt.Printf("%s%s", "% ", line)
//        }
//    }
//line test:
//    lines := dvdfile.TransToLines([]byte(testTxt))
//    linesMem := dvdfile.TransToLinesMem([]byte(testTxt))
//    fmt.Println("lines:")
//    for _, v := range lines {
//        fmt.Printf("%s", v)
//    }
//    lines[0] = append(lines[0], 'M', 'I', 'B')
//    fmt.Println("lines change:")
//    for _, v := range lines {
//        fmt.Printf("%s", v)
//    }
//    fmt.Println("linesMem:")
//    for _, v := range linesMem {
//        fmt.Printf("%s", v)
//    }
//    linesMem[0] = append(linesMem[0], 'M', 'I', 'B')
//    fmt.Println("linesMem change:")
//    for _, v := range linesMem {
//        fmt.Printf("%s", v)
//    }
}

//process:
//get length
//copy based on length
func mycopy(des []byte, src []byte) int {
    var l int
    if len(des) > len(src) {
        l = len(src)
    } else {
        l = len(des)
    }
    for i:=0; i<l; i++ {
        des[i] = src[i]
    }
    return l
}

func myappend(src []byte, ele ...byte) []byte {
    l := len(src)
    if l + len(ele) > cap(src) {
        var newSrc []byte
        if cap(src) < 1000 {
            newSrc = make([]byte, (l+len(ele))*2)
        } else {
            newSrc = make([]byte, l+len(ele) + 100)
        }
        copy(newSrc, src)
        src = newSrc[:l+len(ele)]
    }
    copy(src[l:], ele)
    return src
}

//slice
//map
//Print
//append
