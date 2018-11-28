package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

// When CPU profiling is enabled, the Go program stops about 100 times per second(1秒钟进行100次采样)
// sample consisting of the program counters on the currently executing goroutine's stack.
// [采样内容为正在运行的协程的堆栈信息，记录了每次程序运行时所在的函数, 它的父函数等]
/*
(pprof) top10
Total: 2525 samples(每秒钟采样100次，程序总共运行了25秒多)
第一列：the number of samples in which the function was running(说明在堆栈的最底部)
        the top10 output is sorted by this sample count.
第二列: a percentage of total samples
第三列：第三列第三行的值 == 11.8 + 10.6 + 9.9
第四列：The fourth and fifth columns show the number of samples in which the function appeared
        either running or waiting for a called function to return.
     298  11.8%  11.8%      345  13.7% runtime.mapaccess1_fast64
     268  10.6%  22.4%     2124  84.1% main.FindLoops
     251   9.9%  32.4%      451  17.9% scanblock
     178   7.0%  39.4%      351  13.9% hash_insert
     131   5.2%  44.6%      158   6.3% sweepspan
     119   4.7%  49.3%      350  13.9% main.DFS
      96   3.8%  53.1%       98   3.9% flushptrbuf
      95   3.8%  56.9%       95   3.8% runtime.aeshash64
      95   3.8%  60.6%      101   4.0% runtime.settype_flush
      88   3.5%  64.1%      988  39.1% runtime.mallocgc
*/

/*
but each stack sample only includes the bottom 100 stack frames: 解释main.main函数的第五列为啥不是100%
(pprof) top5 -cum
Total: 2525 samples
       0   0.0%   0.0%     2144  84.9% gosched0
       0   0.0%   0.0%     2144  84.9% main.main
       0   0.0%   0.0%     2144  84.9% runtime.main
       0   0.0%   0.0%     2124  84.1% main.FindHavlakLoops
     268  10.6%  10.6%     2124  84.1% main.FindLoops
(pprof) top5 -cum
*/

/*
	The stack trace samples contain more interesting data about function call relationships than the text listings can show

*/
// there is a row for each function that appeared in a sample. 继而可以统计每个函数出现的频次
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// TODO things
}
