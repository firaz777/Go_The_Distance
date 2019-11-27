package main
// this is the import cycle of dependancies
import(
"bytes"
"os"
"fmt"
"io"
"log"
"time"
"encoding/binary"
"bufio"
"strings"
//"math"
//"sync"
//"runtime"
//"github.com/pkg/profile"
)
// A unit is a base for breaking up the binary input file to discreate useful units.
type Unit struct{
 pID , phylum, class, order, family,subfamily,tribe,genus,species,subspecies,country,projectcode,continercode []byte
 datasetcodes []byte
 Seq []byte
 start, stop int32
 pA, pT, pG, pC ,lat,lon float64
 bin int8
 }
 //this is the goal structure for final output
  type outputdist struct{
   name1, name2 []byte
   distance float64
  }
//the binary writer outputs units of size UnitSize bytes
const UnitSize = 1773
const check_length_1 = 100
const check_length_2 = 300
const threshold_1 = 0.85
const threshold_2 = 0.65
//diffrence in pecent compsition of (ie quick first check for if you want a specic match value, only above 90% match would have a threshold of 0.1)
const OverallThresh = 0.7
//This function reads in int32 from a 4 byte slice (ie. casting the bytes to a int)

func read_int32(data []byte) (ret int32) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}

func read_int8(data []byte) (ret int8) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}

//This function reads in float64 from a 8 byte slice (ie. casting bytes to a float)
func read_float64(data []byte) (ret float64) {
    buf := bytes.NewBuffer(data)
    binary.Read(buf, binary.LittleEndian, &ret)
    return
}
//this converts a slice of size UnitSize to a Unit 
func SlicetoUnit(seq []byte )Unit{
var unitmade = Unit{seq[0:32],seq[32:64],seq[64:96],seq[96:128],seq[128:160],seq[160:192],seq[192:224],seq[224:256],seq[256:288],seq[288:320],seq[320:352],seq[352:384],seq[384:416],seq[416:716],seq[716:1716],read_int32(seq[1716:1720]),read_int32(seq[1720:1724]),read_float64(seq[1724:1732]),read_float64(seq[1732:1740]),read_float64(seq[1740:1748]),read_float64(seq[1748:1756]),read_float64(seq[1756:1764]),read_float64(seq[1764:1772]),read_int8(seq[1772:1773])}
//fmt.Println(unitmade)
return unitmade
}

 
 // this function takes in 2 sequences and their respective unit indexs and returns the relitive p distance (need not check if
 // nucleotides are valid due to indexs found in binary wrting step)
 
 func pDist(seq1 []byte,seq2 []byte, seq1_start int32, seq1_end int32, seq2_start int32, seq2_end int32) float64{
  //this casting could likely be sped up by comparing strings char values instead of bytes (done)
  // seq1 := []byte(seq1)
  // seq2 := []byte(seq2)
  diff := 0.0
  count:= 0.0
  //fmt.Println(a1,a2,t1,t2,g1,g2,c1,c2)

  // if (math.Abs(a1-a2)+math.Abs(t1-t2)+math.Abs(g1-g2)+math.Abs(c1-c2)>OverallThresh){
  //   //fmt.Println("Hello world!")
  //   return(0.51)
  // }
  // try to find way to find indexs without if's
  last := seq1_end
  if seq2_end < seq1_end { last = seq2_end }
  first := seq1_start
  if seq2_start > seq1_start {first=seq2_start}

  for i:= first; i<last;i++{
    //if (seq1[i] == 65 || seq1[i] == 84 || seq1[i] == 67 || seq1[i] == 71) && (seq2[i] == 65 || seq2[i] == 84 || seq2[i] == 67 || seq2[i] == 71){
    //changed binary writer to change any invalid char into "-"(byte 45)
      if !(seq1[i] == 45 || seq2[i] == 45){
      count +=1
      if seq1[i] != seq2[i]{
        diff +=1
      }
    }
    // if count == check_length_1{
    //   dist := diff/count
    //   if dist > threshold_1{
    //     //fmt.Println("Hello world!")
    //     return(0.51)
    //   }
    // }

    //   if count == check_length_2{
    //   dist := diff/count
    //   if dist > threshold_2{
    //     //fmt.Println("Hello world!")
    //     return(0.51)
    //   }
    // }


  }
    dist := diff/count
    return(dist)
  }

  //only reason this is a function is so concurancy can be used to allow read and proccess simaltaniusly
  // func workergo(outSlice []Unit, subUnits []Unit /*, wg *sync.WaitGroup*/){
  //   //numcomp :=0
  //   for _,i:=range outSlice{
  //     //fmt.Print(len(subUnits))
  //       for _,j:= range subUnits{
  //         //_=j
  //         //fmt.Println(pDist(i.Seq,j.Seq,i.start,i.stop,j.start,j.stop,i.pA,i.pT,i.pG,i.pC,j.pA,j.pT,j.pG,j.pC))
  //         //pDist(i.Seq,j.Seq,i.start,i.stop,j.start,j.stop,i.pA,i.pT,i.pG,i.pC,j.pA,j.pT,j.pG,j.pC)
  //         output := outputdist{i.pID,j.pID,pDist(i.Seq,j.Seq,i.start,i.stop,j.start,j.stop,i.pA,i.pT,i.pG,i.pC,j.pA,j.pT,j.pG,j.pC)}
  //         //list = append(list, []string{string(output.name1),string(output.name2)})
  //         //numcomp+=1
  //         fmt.Println(string(output.name1),string(output.name2),output.distance)

          
  //         _=output
          
  //       }
  //     }
  //     //fmt.Print(numcomp)
  //     //fmt.Println(list)
  //     //signal routine is finished
  //     //wg.Done()
  // }



 
//The main exsicuted part of the program
func main() {
  //var SliceOfUnits []Unit

  //runtime.GOMAXPROCS(1)
  //create waitgroup for routine
  //var wg sync.WaitGroup
  //counttest := 0
  //max := 0.0
  // CPU profiling by default
  //defer profile.Start().Stop()
  var paths string
   fmt.Println("Please enter path of the .dat database:")
   fmt.Scan(&paths)

  var paths2 string
   fmt.Println("Please enter path of compare file in .fas or .fasta format:")
   fmt.Scan(&paths2)

  //starts timer
  start := time.Now()

  //opens smaller comare file and adds units to a list
    file2, err := os.Open(paths2)
  if err != nil {
    log.Fatalf("failed opening file: %s", err)
  }
  scanner := bufio.NewScanner(file2)
  scanner.Split(bufio.ScanLines)
  // the output list of subunits
  var subUnits []Unit
  //an indivdual Unit structre for holding info
  var subUnit Unit
  //a simple line counter
  var linecounter int = 0
  //read line by line for smaller compare file since so small comparitivly
  for scanner.Scan() {
    if linecounter %2 == 0{
      stringSlice := strings.Split(scanner.Text(),"|")
       subUnit.pID = []byte(stringSlice[0])
       subUnit.phylum = []byte(stringSlice[1])
       subUnit.class = []byte(stringSlice[2])
      linecounter += 1
    }else{
      var start int
       var end int
      subUnit.Seq = []byte(scanner.Text())
      chars :=[]byte(scanner.Text())
      //fmt.Println(chars)
      //finds start index
      for i, c:=range scanner.Text(){
        if (c==65 || c==67 || c==84 ||c==71) {
                    start = i
                    break
                }
      }
     // finds end index
      for i:= len(chars)-1; i>=0; i--{
        if (chars[i]==65 || chars[i]==67 || chars[i]==84 || chars[i]==71) {
                    end = i
                    break
                } 
      }
      length:= float64(end - start)
      countA:=0.0
      countT:=0.0
      countG:=0.0
      countC:=0.0
      for c:=start; c<end;c ++{
        if (chars[c]==65){
          countA +=1
        }
        if (chars[c]==84){
          countT +=1
        }
        if (chars[c]==71){
          countG +=1
        }
        if (chars[c]==67){
          countC +=1
        }
      }
      //fmt.Println(countA,countT,countC,countG)
      subUnit.pA = (countA/length)
      subUnit.pT = (countT/length)
      subUnit.pG = (countG/length)
      subUnit.pC = (countC/length)
      subUnit.start=int32(start)
      subUnit.stop=int32(end)

      //fmt.Println(subUnit)
      //subUnits is a list of compare files units
      subUnits = append(subUnits,subUnit)
      linecounter +=1
    } 
  }
  
  
  file2.Close() 

   
  
  
  //var outSlice []Unit
  var count int = 0
  //countabove :=0
  file, err := os.Open(paths)
  if err != nil {
    log.Fatal(err)
  }

  defer file.Close()
    //reads 1 Units at a time
  buffer := make([]byte,1*UnitSize)
  for {

    if _, err := file.Read(buffer); err == io.EOF {
      break
    }

    // creates units from the 10000 unit buffer 
    var divided []byte
    for i:=0; i<len(buffer); i+= UnitSize{
      end := i + UnitSize

      if end > len(buffer) {
        end = len(buffer)
      }

      divided =  buffer[i:end]
      //fmt.Println(divided)
    }
    //fmt.Println(divided)
    //fmt.Print(buffer)

    unit := (SlicetoUnit(divided))
    //SliceOfUnits = append(SliceOfUnits,unit)
    //fmt.Println(unit)
    for _,i:= range (subUnits){
      output := outputdist{unit.pID,i.pID,pDist(unit.Seq,i.Seq,unit.start,unit.stop,i.start,i.stop)}
      // fmt.Println(string(output.name1),string(output.name2),output.distance)
      // fmt.Println(i.bin,unit.bin)
      _=output 
      count +=1
      // if output.distance > 0.5{
      //   countabove +=1
      // }
    }

    // if len(SliceOfUnits)==10000{
    //  for _,i:= range SliceOfUnits{
    //   fmt.Println(i.pID)
    //  }
    //   SliceOfUnits = nil
    // }
    //testslice = append(testslice,divided[0])

    //outSlice = append(outSlice,SlicetoUnit(divided[0]))
    


    // //outSlice= append(outSlice,SlicetoUnit(buffer))
    // //outSlice can only ever reach a size of 10k units
    // if len(outSlice) == 10000{
    //   fmt.Println(outSlice)
    //   //Proccess outSlice here
    //   //must still modify for loop to do thresholds
    //   //adds a routine to waitgroup
    //   //wg.Add(1)
    //   //run the routine
    //   workergo(outSlice,subUnits/*,&wg*/) //add go to function for concurancy
    //   //clear outslice
    //    outSlice = nil
    //   }
       
      
     }
     //runs on the last outslice
     //wg.Add(1)
     // if outSlice != nil{
     // workergo(outSlice,subUnits,&wg) //add go to function for concurancy
     // outSlice = nil}

     //holds all go routines untill finished all routines
  //wg.Wait()
  //fmt.Println(counttest)
  elapsed := time.Since(start)
  fmt.Printf("Took: %s to do:", elapsed)
  fmt.Println(count, "comparisons")


  }