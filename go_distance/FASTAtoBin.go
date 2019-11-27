package main

import (
    "bufio"
    "log"
    "os"
    "strings"
    "encoding/binary"
    "fmt"
)

 type Unit struct{
 lable0 , lable1, lable2 [32]byte
 Seq [1000]byte
 start, end int32
 pA, pT, pG, pC float64
 }

func main() {
    var paths string
   fmt.Println("Please enter path of the .fas or . fasta file you wish to civert to .dat:")
   fmt.Scan(&paths)

	var count int = 0
	var stringSlice []string

    var paths2 string
   fmt.Println("Name your output file:")
   fmt.Scan(&paths2)

    file, err := os.Open(paths)

    f,_ := os.Create(paths2)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    //adjust the capacity to your need (max characters in line)
	const maxCapacity = 4000  
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

    for scanner.Scan() {
    	 
        //fmt.Println(scanner.Text())
        if count%2 == 0 {
        	//parse header
        	stringSlice = strings.Split(scanner.Text(),"|")
            
        	

        	count += 1
        }else{
        	var l0, l1, l2 [32]byte
    	 	var seq [1000]byte
            var start, end int32

        	//parse seq
        	stringSlice = append(stringSlice,scanner.Text())
            

        	//fmt.Printf("%v", stringSlice)
        	copy(l0[:],stringSlice[0])
        	copy(l1[:],stringSlice[1])
        	copy(l2[:],stringSlice[2])
            //could be 3 or 
            if len(stringSlice) == 5{
        	copy(seq[:],stringSlice[4])
            }else{
                copy(seq[:],stringSlice[3])
            }

            countA := 0.0
            countT := 0.0
            countG := 0.0
            countC :=0.0

            for i, c:= range seq{
                if c==97{
                    seq[i]=65
                }
                if c==99{
                    seq[i]=67
                }
                if c==116{
                    seq[i]=84
                }
                if c==103{
                    seq[i]=71
                }
                if (c!=65 && c!=67 && c!=84 && c!=71){
                    seq[i]=45
                }
                if (c==65){
                    countA +=1
                }
                if (c==84){
                    countT +=1
                }
                if (c==67){
                    countC +=1
                }
                if (c==71){
                    countG +=1
                }
            }
            startpos:=0
            endpos :=0
            for i, c:= range seq{
                if (c==65 || c==67 || c==84 ||c==71) {
                    startpos = i
                    break
                }
            }       
            for i:= 999; i>=0; i--{
                if (seq[i]==65 || seq[i]==67 || seq[i]==84 || seq[i]==71) {
                    endpos = i
                    break
                } 
            }
           start = int32(startpos)
           end=int32(endpos)
           length := float64(endpos-startpos)

           pA:= countA/(countA+countT+countV+countG)
           pT:= countT/(countA+countT+countV+countG)
           pG:= countG/(countA+countT+countV+countG)
           pC:= countC/(countA+countT+countV+countG)

        	seqUnit := Unit{l0 , l1, l2 , seq,start, end, pA, pT, pG, pC}
        	binary.Write(f, binary.LittleEndian, &seqUnit)
        	count +=1
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}