using System;
using System.Threading.Tasks;
using System.Threading;

namespace Test
{
    class Async1
    {
        // 异步任务，获取 url 的 resp 内容
        async Task<byte[]> HttpGet(string url) {
            Console.WriteLine("[HttpGet] http get start");
            byte[] resp = {1, 2, 3};
            await Task.Run(() => {
                Thread.Sleep(2000);
                Console.WriteLine("[HttpGet] http get {0} done!", url);
            });
            Console.WriteLine("[HttpGet] http get end");
            return resp;
        }

        // 对 resp 进行处理
        async void HandlerResponse(string url) {
            // 需要等待 HttpGet 执行完成
            var resp = await HttpGet(url);
            Console.WriteLine("[HandlerResponse] get the resp from HttpGet");
            Task.Run(() => {
                Console.WriteLine(resp);
                Console.WriteLine("[HandlerResponse] handler resp done!");
            });
        }

        void DoOtherWork() {
            Console.WriteLine("[otherWork] do some thing....");
        }

        static void Main(string[] args)
        {
            var a = new Async1();
            //Console.WriteLine("[main] thread start");

            // HandlerResponse 是一个异步任务
            a.HandlerResponse("www.baidu.com");
            
            //Console.WriteLine("[main] thread end");

            // 其他任务，这里不会被阻塞，因为上面的异步任务没有 await
            a.DoOtherWork();

            // 等待线程全部执行完成（不知道 C# 中如何优雅等待）
            Thread.Sleep(10000);
        }
    }
}