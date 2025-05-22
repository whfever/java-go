import argparse
import asyncio
import os
import re
from urllib.parse import urlparse

from playwright.async_api import async_playwright


# 工具函数：清理文件名
def sanitize_filename(title):
    return re.sub(r'[^a-zA-Z0-9\u4e00-\u9fa5]', '_', title)


# 工具函数：提取域名
def extract_domain(url):
    parsed = urlparse(url)
    return parsed.netloc.replace("www.", "")


# 处理 HTML 输出
async def handle_html(page, url, base_path):
    content = await page.eval_on_selector("#Wrapper", "el => el.innerHTML") or \
              await page.eval_on_selector("body", "el => el.innerHTML")
    full_html = f"""<!DOCTYPE html>
<html lang="zh">
<head>
  <meta charset="UTF-8">
  <title>{base_path}</title>
  <style>body {{ background-color: white; }}</style>
</head>
<body>
{content}
</body>
</html>"""
    filename = f"{base_path}.html"
    with open(filename, 'w', encoding='utf-8') as f:
        f.write(full_html)
    print(f"HTML saved: {filename}")


# 处理 Markdown 输出
async def handle_md(page, url, base_path):
    from markdownify import markdownify
    content = await page.eval_on_selector("#Wrapper", "el => el.innerHTML") or \
              await page.eval_on_selector("body", "el => el.innerHTML")
    md_content = markdownify(content)
    filename = f"{base_path}.md"
    with open(filename, 'w', encoding='utf-8') as f:
        f.write(md_content)
    print(f"Markdown saved: {filename}")


# 处理 PDF 输出
async def handle_pdf(page, base_path):
    output_path = f"{base_path}.pdf"
    await page.pdf(path=output_path, format="A4", print_background=True)
    print(f"PDF saved: {output_path}")


# 处理截图输出
async def handle_png(page, base_path):
    output_path = f"{base_path}.png"
    await page.screenshot(path=output_path, full_page=True)
    print(f"PNG saved: {output_path}")


# 主流程：处理单个 URL
async def process_url(url, options):
    try:
        async with async_playwright() as p:
            browser = await p.chromium.launch(headless=False,
                                              proxy={"server": options.get("proxy")},
                                              executable_path=options.get("browser_executable"))
            page = await browser.new_page()
            await page.set_user_agent(options["user_agent"])
            await page.goto(url, timeout=60000)

            # 设置页面背景色
            await page.add_init_script("""
                document.body.style.backgroundColor = "white";
            """)

            domain = extract_domain(url)
            title = await page.title()
            safe_title = sanitize_filename(title)
            base_name = f"{options['output_dir']}/{domain}_{safe_title}"

            # 执行指定的导出操作
            if options["html"]:
                await handle_html(page, url, base_name)
            if options["md"]:
                await handle_md(page, url, base_name)
            if options["pdf"]:
                await handle_pdf(page, base_name)
            if options["png"]:
                await handle_png(page, base_name)

            await browser.close()

    except Exception as e:
        print(f"Error processing {url}: {e}")


# 命令行解析与主入口
def main():
    parser = argparse.ArgumentParser(description="网页抓取工具，支持多种输出格式")
    parser.add_argument("--url", action="append", help="添加要抓取的URL")
    parser.add_argument("--pdf", action="store_true", help="保存为PDF")
    parser.add_argument("--html", action="store_true", help="保存为HTML")
    parser.add_argument("--md", action="store_true", help="保存为Markdown")
    parser.add_argument("--png", action="store_true", help="保存截图")
    parser.add_argument("--proxy", default="http://127.0.0.1:7890", help="代理地址")
    parser.add_argument("--user-agent", default="Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
                        help="自定义User-Agent")
    parser.add_argument("--edge", help="指定Microsoft Edge路径")

    args = parser.parse_args()

    DEFAULT_URLS = [
        "https://www.lipijin.com/resume",
        "https://v2ex.com/t/1117140#reply4",
        "https://v2ex.com/t/1126201#reply3",
    ]

    OUTPUT_DIR = "./output"
    os.makedirs(OUTPUT_DIR, exist_ok=True)

    options = {
        "urls": args.url or DEFAULT_URLS,
        "proxy": args.proxy,
        "user_agent": args.user_agent,
        "browser_executable": args.edge,
        "output_dir": OUTPUT_DIR,
        "pdf": args.pdf,
        "html": args.html,
        "md": args.md,
        "png": args.png
    }

    tasks = [process_url(url, options) for url in options["urls"]]
    asyncio.run(asyncio.gather(*tasks))


if __name__ == "__main__":
    main()