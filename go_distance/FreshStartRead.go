package main

import (
"bytes"
"fmt"
"encoding/binary"
"os"
"log"
"time"
"io"
"bufio"
)

const UnitSize = 1040








func pDist(seq1 [1000]byte,seq2 [1000]byte, seq1_start int32, seq1_end int32, seq2_start int32, seq2_end int32) float64{
  diff := 0.0
  count:= 0.0
  last := seq1_end
  if seq2_end < seq1_end { last = seq2_end }
  first := seq1_start
  if seq2_start > seq1_start {first=seq2_start}

  for i:= first; i<last;i++{

      if !(seq1[i] == 45 || seq2[i] == 45){
      count +=1
      if seq1[i] != seq2[i]{
        diff +=1
      }
    }
  }
    dist := diff/count
    return(dist)
  }









type Unit struct{
	PID [32]byte
	Seq [1000]byte
	Start,Stop int32

}

type Wrapper struct{
	ListOfUnits [10000]Unit
}




func readChunks(r io.Reader) error {
	myWrapper := Wrapper{}
    if _, ok := r.(*bufio.Reader); !ok {
        r = bufio.NewReader(r)
    }
    buf := make([]byte, 0, 10000*UnitSize)
    for {
        n, err := io.ReadFull(r, buf[:cap(buf)])
        bufr := bytes.NewReader(buf[:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            if err != io.ErrUnexpectedEOF {
                return err
            }
        }

        // Process buf
        myWrapper = Wrapper{}
        err2 := binary.Read(bufr, binary.LittleEndian, &myWrapper)
        // for _,i:=range myWrapper.ListOfUnits{
        //   fmt.Println(i.PID)
        // }
		if err2 != nil {
			fmt.Println("binary.Read failed:", err2)
		}
  }
  return nil
}







func main() {

   var path string
   fmt.Println("Please enter path of the .dat database:")
   fmt.Scan(&path)

   start := time.Now()
   file, err := os.Open(path)
   if err != nil {
    log.Fatal(err)
   }
   defer file.Close()
   //go:noescape
   err = readChunks(file)
    if err != nil {
        fmt.Println(err)
        return
    }
   

  elapsed := time.Since(start)
  fmt.Printf("Took: %s:", elapsed)
}

