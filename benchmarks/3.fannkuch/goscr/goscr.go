//package goscr
//import "./fannkuch"
package main
import "./fannkuch"
import "fmt"
func main(){
    fmt.Println(Fannkuch(4))
}

var (                                                                                                                  
    NCHUNKS = 720
    CHUNKSZ = 0
    NTASKS  = 0
)

var Fact []int

type Ctx_M struct {
    res, chk int
}

// type Ctx_FannkuchRecursive_Source interface {                                                                       
func (ctx *Ctx_M) Recv_NewWorker_FannkuchRecursive_Result(r fannkuch.Result) {
    if (*ctx).res < r.MaxFlips {
        (*ctx).res = r.MaxFlips
    }
    (*ctx).chk += r.Checksum
}
func (ctx *Ctx_M) Recv_NewWorker_FannkuchRecursive_Result_2(r fannkuch.Result){
    if (*ctx).res < r.MaxFlips {
        (*ctx).res = r.MaxFlips
    }
    (*ctx).chk += r.Checksum
}
func (_ *Ctx_M) End(){ }

type Ctx_W struct {
    n, idxMin, idxMax int
    p, pp, count []int
}

func computeResult(ctx *Ctx_W) fannkuch.Result { 
    
    // first permutation
    for i := 0; i < (*ctx).n; i++ {
        (*ctx).p[i] = i
    }
    for i, idx := (*ctx).n-1, (*ctx).idxMin; i > 0; i-- {
        d := idx / Fact[i]
        (*ctx).count[i] = d
        idx = idx % Fact[i]
    
        copy((*ctx).pp, (*ctx).p)
        for j := 0; j <= i; j++ {
            if j+d <= i {
                (*ctx).p[j] = (*ctx).pp[j+d]
            } else {
                (*ctx).p[j] = (*ctx).pp[j+d-i-1]
            }
        }
    }
    
    maxFlips := 1
    checkSum := 0

    for idx, sign := (*ctx).idxMin, true; ; sign = !sign {
    
        // count flips
        first := (*ctx).p[0]
        if first != 0 {
            flips := 1
            if (*ctx).p[first] != 0 {
                copy((*ctx).pp, (*ctx).p)
                p0 := first
                for {
                    flips++
                    for i, j := 1, p0-1; i < j; i, j = i+1, j-1 {
                        (*ctx).pp[i], (*ctx).pp[j] = (*ctx).pp[j], (*ctx).pp[i]
                    }
                    t := (*ctx).pp[p0]
                    (*ctx).pp[p0] = p0
                    p0 = t
                    if (*ctx).pp[p0] == 0 {
                        break
                    }
                }
            }
            if maxFlips < flips {
                maxFlips = flips
            }
            if sign {
                checkSum += flips
            } else {
                checkSum -= flips
            }
        }
    
        if idx++; idx == (*ctx).idxMax {
            break
        }
    
        // next permutation
        if sign {
            (*ctx).p[0], (*ctx).p[1] = (*ctx).p[1], first
        } else {
            (*ctx).p[1], (*ctx).p[2] = (*ctx).p[2], (*ctx).p[1]
            for k := 2; ; k++ {
                if (*ctx).count[k]++; (*ctx).count[k] <= k {
                    break
                }
                (*ctx).count[k] = 0
                for j := 0; j <= k; j++ {
                    (*ctx).p[j] = (*ctx).p[j+1]
                }
                (*ctx).p[k+1] = first
                first = (*ctx).p[0]
            }
        }
    }
    return fannkuch.Result{checkSum, maxFlips}
}

//type Ctx_FannkuchRecursive_NewWorker interface {
func (ctx *Ctx_W) Choice_NewWorker_FannkuchRecursive_() fannkuch.Select_NewWorker {
    (*ctx).idxMax = (*ctx).idxMin + CHUNKSZ
    if (*ctx).idxMax < Fact[(*ctx).n] {
        return fannkuch.Invite_Source_FannkuchRecursive{}
    }
    (*ctx).idxMax = Fact[(*ctx).n]
    return computeResult(ctx)
}
func (ctx *Ctx_W) Init_NewWorker_FannkuchRecursive_Ctx() fannkuch.Ctx_FannkuchRecursive_NewWorker {
    p := make([]int, (*ctx).n)
    pp := make([]int, (*ctx).n)
    count := make([]int, (*ctx).n)        
    x := Ctx_W{(*ctx).n, (*ctx).idxMax, 0, p, pp, count}
    return &x
}
func (ctx *Ctx_W) Send_Source_FannkuchRecursive_Result() fannkuch.Result { 
    return computeResult(ctx)
}
func (_ *Ctx_W) End() { }

func Fannkuch(n int) (int, int) {
    Fact = make([]int, n+1) 
    Fact[0] = 1 
    for i := 1; i < len(Fact); i++ { 
        Fact[i] = Fact[i-1] * i 
    } 
    
    CHUNKSZ = (Fact[n] + NCHUNKS - 1) / NCHUNKS 
    CHUNKSZ += CHUNKSZ % 2 
    NTASKS = (Fact[n] + CHUNKSZ - 1) / CHUNKSZ 
    ctx_m := Ctx_M{0, 0}
    p := make([]int, n)
    pp := make([]int, n)
    count := make([]int, n)        
    ctx_w := Ctx_W{n, 0, 0, p, pp, count}
    fannkuch.Start(&ctx_m, &ctx_w)
    return ctx_m.res, ctx_m.chk

}
