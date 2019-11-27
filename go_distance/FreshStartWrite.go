package main

import (
"fmt"
"encoding/binary"
"os"
"log"
"bufio"
"strings"
"time"
)


//size of a unit is 1724 bytes
type Unit struct{
	pID [32]byte//, phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode [32]byte
	//datasetcodes [300]byte
	seq [1000]byte
	start,stop int32

}

type Suplemet struct{
   pID,phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode [32]byte
   datasetcodes [300]byte
}


type Wrapper struct{
	ListOfUnits [10000]Unit
}

type WrapperSuplement struct{
  ListofSuplements [10000]Suplemet
}



func main(){
	var stringSlice []string
	var listOfUnits [10000]Unit
	var listOfUnitsEmpty [10000]Unit
  var ListofSuplements [10000]Suplemet
  var ListofSuplementsEmpty [10000]Suplemet

	//inputs of path and output file name
   var path string
   fmt.Println("Please enter path of the .fas or .fasta file you wish to convert to .dat: ")
   fmt.Scan(&path)
   var path2 string
   fmt.Println("Name your output file: ")
   fmt.Scan(&path2)
   var path3 string
   fmt.Println("Name your supplementry data file:")
   fmt.Scan(&path3)

   start := time.Now()


   //opens specified fasta file
   file, err := os.Open(path)
   //creates output file
   f,_ := os.Create(path2)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
  f2,_ := os.Create(path3)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


    //creates a scanner to read the fasta file
    scanner := bufio.NewScanner(file)
    //adjust the capacity to your need (max characters in line)
	const maxCapacity = 4000  
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)


	//starts creating units and appending to unit list until length is 10000
	count :=0
	for scanner.Scan() {


		//Writes the list to binary
		 if ((count/2)==9999){
		 	NewWraper :=Wrapper{listOfUnits}
      NewWrapperSupplement :=WrapperSuplement{ListofSuplements}
		 	binary.Write(f, binary.LittleEndian, &NewWraper)
      binary.Write(f2,binary.LittleEndian, &NewWrapperSupplement)
		 	listOfUnits = listOfUnitsEmpty
      ListofSuplements = ListofSuplementsEmpty
		 	//fmt.Println(listOfUnits[0].pID[0])
			count=0
		} 
    	 

        if count%2 == 0 {
        	//parse header
        	stringSlice = strings.Split(scanner.Text(),"|")
        	count += 1

        }else{
        		//Declares Nessasary variables to create struct
        		var pID , phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode [32]byte
        		var datasetcodes [300]byte
        		var seq [1000]byte
        		var start, end int32

        		stringSlice = append(stringSlice,scanner.Text())

        		copy(pID[:],stringSlice[0])
        		copy(phylum[:],stringSlice[1])
        		copy(class[:],stringSlice[2])
        		copy(order[:],stringSlice[3])
        		copy(family[:],stringSlice[4])
        		copy(subfamily[:],stringSlice[5])
        		copy(tribe[:],stringSlice[6])
        		copy(genus[:],stringSlice[7])
        		copy(species[:],stringSlice[8])
        		copy(subspecies[:],stringSlice[9])
        		copy(country[:],stringSlice[10])
        		copy(projectcode[:],stringSlice[11])
        		copy(continercode[:],stringSlice[12])
        		copy(datasetcodes[:],stringSlice[12])
        		copy(seq[:],stringSlice[len(stringSlice)-1])

        		//Ensures Data integrity of the sequence for the read script in the next step
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


           		seqUnit := Unit{pID /*, phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode,datasetcodes*/,seq,start, end}
              SupUnit := Suplemet{pID,phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode,datasetcodes}
           		listOfUnits[count/2]=seqUnit
              ListofSuplements[count/2]=SupUnit
           		count+=1
        }
    }
    //fmt.Println(listOfUnits[0].pID[0])
    if (listOfUnits[0].pID[0] != 0){
    NewWraper :=Wrapper{listOfUnits}
    NewWrapperSupplement := WrapperSuplement{ListofSuplements}
	  binary.Write(f, binary.LittleEndian, &NewWraper)
    binary.Write(f2,binary.LittleEndian,&NewWrapperSupplement)
}

  elapsed := time.Since(start)
  fmt.Printf("Took: %s:", elapsed)
}