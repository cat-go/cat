# cat
cat client for golang
# 背景
公司使用cat做监控与链路追踪，官方golang的client已经多年不维护了，使用的话还是要自己封装
# 使用
```
go get github.com/cat-go/cat
```
# Quickstart
```
package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/cat-go/cat"
)

const TestType = "foo"

var wg = sync.WaitGroup{}

func init() {
	cat.DebugOn()
	cat.Init(&cat.Options{
    		AppId:      "cat-go",
    		Port:       2280,
    		HttpPort:   8080,
    		ServerAddr: "127.0.0.1",
    	})
}

// send transaction
func case1() {
	t := cat.NewTransaction(TestType, "test")
	defer t.Complete()

	if rand.Int31n(100) == 0 {
		t.SetStatus(cat.FAIL)
	}

	t.AddData("foo", "bar")

	t.NewEvent(TestType, "event-1")
	t.Complete()

	if rand.Int31n(100) == 0 {
		t.LogEvent(TestType, "event-2", cat.FAIL)
	} else {
		t.LogEvent(TestType, "event-2")
	}
	t.LogEvent(TestType, "event-3", cat.SUCCESS, "k=v")

	t.SetDurationStart(time.Now().Add(-5 * time.Second))
	t.SetTime(time.Now().Add(-5 * time.Second))
	t.SetDuration(time.Millisecond * 500)
}

// send completed transaction with duration
func case2() {
	cat.NewCompletedTransactionWithDuration(TestType, "completed", time.Second*24)
	cat.NewCompletedTransactionWithDuration(TestType, "completed-over-60s", time.Second*65)
}

// send event
func case3() {
	// way 1
	e := cat.NewEvent(TestType, "event-4")
	e.Complete()
	// way 2

	if rand.Int31n(100) == 0 {
		cat.LogEvent(TestType, "event-5", cat.FAIL)
	} else {
		cat.LogEvent(TestType, "event-5")
	}
	cat.LogEvent(TestType, "event-6", cat.SUCCESS, "foobar")
}

// send error with backtrace
func case4() {
	if rand.Int31n(100) == 0 {
		err := errors.New("error")
		cat.LogError(err)
	}
}

// send metric
func case5() {
	cat.LogMetricForCount("metric-1")
	cat.LogMetricForCount("metric-2", 3)
	cat.LogMetricForDuration("metric-3", 150*time.Millisecond)
	cat.NewMetricHelper("metric-4").Count(7)
	cat.NewMetricHelper("metric-5").Duration(time.Second)
}

func run(f func()) {
	defer wg.Done()

	for i := 0; i < 100000000; i++ {
		f()
		time.Sleep(time.Microsecond * 100)
	}
}

func start(f func()) {
	wg.Add(1)
	go run(f)
}

func main() {
	start(case1)
	start(case2)
	start(case3)
	start(case4)
	start(case5)

	wg.Wait()

	cat.Shutdown()
}
```
# cat的一些概念

#### 监控模型

CAT主要支持以下四种监控模型：

+  **Transaction**	  适合记录跨越系统边界的程序访问行为,比如远程调用，数据库调用，也适合执行时间较长的业务逻辑监控，Transaction用来记录一段代码的执行时间和次数
+  **Event**	   用来记录一件事发生的次数，比如记录系统异常，它和transaction相比缺少了时间的统计，开销比transaction要小
+  **Heartbeat**	表示程序内定期产生的统计信息, 如CPU利用率, 内存利用率, 连接池状态, 系统负载等
+  **Metric**	  用于记录业务指标、指标可能包含对一个指标记录次数、记录平均值、记录总和，业务指标最低统计粒度为1分钟

#### 主要功能
+  **Transaction报表** 监控一段代码运行情况：运行次数、QPS、错误次数、失败率、响应时间统计（平均影响时间、Tp分位值）等等。
+  **Event报表** 监控一段代码运行次数：例如记录程序中一个事件记录了多少次，错误了多少次。Event报表的整体结构与Transaction报表几乎一样，只缺少响应时间的统计。
+  **Problem报表**	Problem记录整个项目在运行过程中出现的问题，包括一些异常、错误、访问较长的行为。Problem报表是由logview存在的特征整合而成，方便用户定位问题。 来源：

```
1. 业务代码显示调用Cat.logError(e) API进行埋点，具体埋点说明可查看埋点文档。
2. 与LOG框架集成，会捕获log日志中有异常堆栈的exception日志。
3. long-url，表示Transaction打点URL的慢请求
4. long-sql，表示Transaction打点SQL的慢请求
5. long-service，表示Transaction打点Service或者PigeonService的慢请求
6. long-call，表示Transaction打点Call或者PigeonCall的慢请求
7. long-cache，表示Transaction打点Cache.开头的慢请求
```

+  **Heartbeat报表** 是CAT客户端，以一分钟为周期，定期向服务端汇报当前运行时候的一些状态。
+  **Business报表**	对应着业务指标，比如订单指标。与Transaction、Event、Problem不同，Business更偏向于宏观上的指标，另外三者偏向于微观代码的执行情况。

# 声明
代码根据官方v2版本调整
