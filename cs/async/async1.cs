using System;
using System.Threading.Tasks;

namespace AsyncBreakfast
{
    // These classes are intentionally empty for the purpose of this example. 
    // They are simply marker classes for the purpose of demonstration, contain 
    // no properties, and serve no other purpose.
    internal class Bacon { }    // 培根
    internal class Coffee { } // 咖啡
    internal class Egg { } // 鸡蛋
    internal class Juice { } // 果汁
    internal class Toast { } // 吐司

    class Program
    {
        static void Main(string[] args)
        {
            var now = DateTime.Now.ToLongTimeString().ToString();
            // 倒咖啡
            Coffee cup = PourCoffee();
            Console.WriteLine($"[{now}]咖啡准备好了");

            now = DateTime.Now.ToLongTimeString().ToString();
            // 炒鸡蛋
            Egg eggs = FryEggs(2);
            Console.WriteLine($"[{now}]鸡蛋准备好了");

            // 炒培根
            Bacon bacon = FryBacon(3);
            Console.WriteLine("培根准备好了");

            // 吐司面包
            Toast toast = ToastBread(2);
            // 涂黄油
            ApplyButter(toast);
            // 涂果酱
            ApplyJam(toast);
            Console.WriteLine("吐司准备好了");

            // 倒橙汁
            Juice oj = PourOJ();
            Console.WriteLine("橙汁准备好了");
            Console.WriteLine("早餐准备好了！");
        }

        private static Juice PourOJ()
        {
            Console.WriteLine("倒橙汁");
            return new Juice();
        }

        private static void ApplyJam(Toast toast) =>
            Console.WriteLine("把果酱涂在吐司上");

        private static void ApplyButter(Toast toast) =>
            Console.WriteLine("将黄油涂在吐司上");

        private static Toast ToastBread(int slices)
        {
            for (int slice = 0; slice < slices; slice++)
            {
                Console.WriteLine("将一片面包放入烤面包机");
            }
            Console.WriteLine("开始烤...");
            Task.Delay(3000).Wait();
            Console.WriteLine("从烤面包机中取出吐司");

            return new Toast();
        }

        private static Bacon FryBacon(int slices)
        {
            Console.WriteLine($"将 {slices} 片培根放入锅中");
            Console.WriteLine("烹饪培根的第一面...");
            Task.Delay(3000).Wait();
            for (int slice = 0; slice < slices; slice++)
            {
                Console.WriteLine("翻转一片培根");
            }
            Console.WriteLine("煮培根的第二面……");
            Task.Delay(3000).Wait();
            Console.WriteLine("把培根放在盘子里");

            return new Bacon();
        }

        // 炒鸡蛋
        private static Egg FryEggs(int howMany)
        {
            Console.WriteLine("加热鸡蛋盘...");
            Task.Delay(3000).Wait();
            Console.WriteLine($"敲碎 {howMany} 个鸡蛋");
            Console.WriteLine("煮鸡蛋...");
            Task.Delay(3000).Wait();
            Console.WriteLine("煮鸡蛋...");

            return new Egg();
        }

        private static Coffee PourCoffee()
        {
            Console.WriteLine("倒咖啡");
            return new Coffee();
        }
    }
}
