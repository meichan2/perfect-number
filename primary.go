package main

import(
   "fmt"
   "math"
   "time"
)

func is_prime(n int) bool {
    switch {
    case n < 2:
       return false
    case n == 2 || n == 3 || n == 5:
       return true
    case n & 1 == 0:
       return false
    case n % 2 == 0 || n % 3 == 0 || n % 5 == 0:
       return false
    }

    f := true
    v := 7
    m := int(math.Pow(float64(n), 0.5)) + 1
    for v < m {
        if n % v == 0 {
            return false
        }
        if f {
            v += 4
        } else {
            v += 2
        }
        f = !f
    }
    return true
}

func main() {
  start := time.Now()
  stnum:=1
  //ednum:=100000000
  stnum+=1
  var prnum [] int
  for k:=stnum;len(prnum)<10;k++{
    b:=float64(k)
    an:=math.Pow(2,b-1.0)
    a:=math.Pow(2,b)-1.0
    ia:=int(a)
    if is_prime(ia){
      fmt.Printf("%.0f\n",a*an)
      prnum=append(prnum,int(a*an))
    }else if k%10000000==0{
      end := time.Now()
      fmt.Printf("%d %f秒\n",k,(end.Sub(start)).Seconds())
    }
  }
  fmt.Println(prnum)
}
