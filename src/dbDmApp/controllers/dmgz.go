package dmgz

func Sqrt(x float64) float64 {
    n:=0.0
    for i:=0;i<1000;i++ {
        n-= (n*n-x)/(2*x)
    }
    return n
}
