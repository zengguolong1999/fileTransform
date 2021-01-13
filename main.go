package main

import (
    "os"
    "io"
    "regexp"
    "fmt"
    "bytes"
    "log"
    "strings"
    "strconv"
    "flag"

    "MrLiu_filetransform/dvdfile"
)

func main() {
    var iFilename string
    flag.StringVar(&iFilename, "f", "file", "input file name")
    var oFilename string
    flag.StringVar(&oFilename, "o", "a.out", "output file name")
    flag.Parse()
    file, err := os.Open(iFilename)
    defer file.Close()
    if err != nil {
        fmt.Println(err)
    }
    data := make([]byte, 0)
    buf := make([]byte, 1000)
    for {
        count, err := file.Read(buf)
        data = append(data, buf[:count]...)
        if err != nil {
            if err == io.EOF {
                break
            }
            log.Fatal(err)
        }
    }
    reComment := regexp.MustCompile(`^;|\[[a-zA-Z ]*\]`)
    reDeclare := regexp.MustCompile(`\[([a-zA-Z]+ ?)*\]`)
    re2BTo1W := regexp.MustCompile(`0x[0-9A-Z]{2}  0x[0-9A-Z]{2}`)
    paras := dvdfile.TransToParas(data)
    for _, para := range paras {
        for j, line := range para {
            if reComment.Match(line) {
                para[j] = insSlashCmt(line) //add comment
            }
        }
    }
    cplexParas := paras[1:len(paras)-1]
    for i, para := range cplexParas {
        var reserveLineLen int
        if i == 0 {
            reserveLineLen = 6
        } else {
            reserveLineLen = 3
        }
        lineCount2B := len(para) - reserveLineLen
        for j, line := range para {
            if reDeclare.Match(line) { //declare line
                varName := string(line[3:len(line)-3]) //pattern: "//[Frequency Bank]\r\n"
                varNameEle := strings.Fields(varName)
                varName = ""
                for _, v := range varNameEle {
                    varName += v
                }
                declareLine := "unsigned int " + varName + "[" + strconv.Itoa(lineCount2B) + "]" + " = {\r\n"
                para[j+1] = []byte(declareLine)
            } else if re2BTo1W.Match(line) { //pattern: "0xAF  0x23\r\n"
                newLine := make([]byte, 9)
                copy(newLine[:4], line[:4])
                copy(newLine[4:6], line[8:10])
                copy(newLine[6:], []byte(",\r\n"))
                para[j] = newLine
            }
        }
        newPara := make([][]byte, len(cplexParas[i])+1)
        copy(newPara, cplexParas[i])
        newPara[len(newPara)-1] = []byte("\r\n")
        newPara[len(newPara)-2] = []byte("};\r\n") //add a new line "}\n"
    }
    lastPara := paras[len(paras)-1] //add comment to last line(checksum)
    lastPara[len(lastPara)-1] = insSlashCmt(lastPara[len(lastPara)-1])
    newFile, err := os.Create(oFilename)
    if err != nil {
        log.Fatalln(err)
    }
    defer newFile.Close()
    var buffer bytes.Buffer
    for _, para := range paras {
        for _, line := range para {
            buffer.Write(line)
        }
    }
    resData := make([]byte, 0)
    resDataBuf := make([]byte, 1000)
    for {
        n, err := buffer.Read(resDataBuf)
        resData = append(resData, resDataBuf[:n]...)
        if err == io.EOF {
            break
        }
    }
    _, err = newFile.Write(resData)
    if err != nil {
        log.Fatalln(err)
    }
}

func insSlashCmt(line []byte) []byte {
    res := make([]byte, len(line)+2)
    copy(res[2:], line)
    res[0] = '/'
    res[1] = '/'
    return res
}
