import asyncio
import time


# 模拟一个异步网页抓取任务
async def fetch_page(url):
    print(f"开始抓取: {url}")
    await asyncio.sleep(2)  # 模拟网络延迟
    print(f"完成抓取: {url}")
    return f"{url} 内容"


# 异步处理单个 URL 的任务
async def process_url(url):
    try:
        content = await fetch_page(url)
        # 可在此处添加解析、保存等操作
        return content
    except Exception as e:
        print(f"抓取失败 {url}: {e}")
        return None


# 主函数：并发执行多个任务
async def main(urls):
    tasks = [process_url(url) for url in urls]
    results = await asyncio.gather(*tasks)

    # 输出结果（可替换为写入文件或数据库）
    for url, result in zip(urls, results):
        if result:
            print(f"结果来自 {url}: {len(result)} 字节")


# 入口点
if __name__ == "__main__":
    # 示例网址列表
    URLS = [
        "https://www.lipijin.com/resume",
        "https://v2ex.com/t/1117140#reply4",
        "https://v2ex.com/t/1126201#reply3",
        "https://v2ex.com/t/1075954",
        "https://v2ex.com/t/654289"
    ]

    start_time = time.time()
    asyncio.run(main(URLS))
    end_time = time.time()

    print(f"\n总耗时: {end_time - start_time:.2f} 秒")