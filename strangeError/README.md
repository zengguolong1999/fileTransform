# command-line-arguments
./test.go:50:13: invalid line number: 

This error appears when I ran "go run test.go". However, line 50 is a comment line. I tried delete the colon and the error disappears. It is also ok if I append some other content after I deleted the colon.
Why would this happen? Will a comment make a difference to the result of compiling?
