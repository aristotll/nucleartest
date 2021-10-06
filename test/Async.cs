using System;
using System.Threading;
using System.Threading.Tasks;

namespace test
{
    class AsyncTest
    {
        public static async Task<int> AwaitFunc() {
        var sum = 0;
        await Task.Run(() => {
            for (int i = 0; i < 100; i++)
            {
                Thread.Sleep(500);  // 毫秒
                sum += i;
            }
        });

        return sum; 
    }

    public static async void Func() {
        Task<int> t = AwaitFunc();
        Console.WriteLine("ok");
        int res = await t;
        Console.WriteLine("res: " + res);
    }

    // static void Main(string[] args)
    // {
    //     Func();
    // }
    }
}