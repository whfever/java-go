// puppeteer-all.js
const puppeteer = require('puppeteer-core');
const fs = require('fs');
const TurndownService = require('turndown');

// 默认内置多个网址
const DEFAULT_URLS = [
  'https://v2ex.com/t/1019827',
  'https://v2ex.com/t/1090050',
  'https://v2ex.com/t/1051125','https://v2ex.com/t/1019827','https://v2ex.com/t/927121',
  'https://windypath.com/resume','https://www.v2ex.com/t/967464',
  'https://v2ex.com/t/1113398','https://www.v2ex.com/t/1082507',
  'https://www.v2ex.com/t/1113398','https://www.v2ex.com/t/1113398',
];

// 解析命令行参数
const args = process.argv.slice(2);

// 提取所有 --url 参数
const urls = [];
let i = 0;
while (i < args.length) {
  if (args[i] === '--url' && args[i + 1]) {
    urls.push(args[i + 1]);
    i += 2;
  } else {
    i++;
  }
}
const baseUrl = '/Users/sure/Documents/';
// 提取其他选项
const options = {
  pdf: args.includes('--pdf'),
  html: args.includes('--html'),
  md: args.includes('--md'),
  png: args.includes('--png')
};

// 使用命令行或默认 URL 列表
const targetUrls = urls.length > 0 ? urls : DEFAULT_URLS;

(async () => {
  const browser = await puppeteer.launch({
    executablePath: '/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge',
    headless: false,
    args: [
      '--proxy-server=http://127.0.0.1:7890', // 使用你的代理
    ]
  });

  for (const url of targetUrls) {
    const page = await browser.newPage();
    
    // 设置真实浏览器的 User-Agent
    await page.setUserAgent('Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36');

    // 增加超时时间为 60 秒
    await page.goto(url, { timeout: 60000 });
    
    // 强制页面背景为白色
    await page.addStyleTag({ content: 'body { background-color: white !important; }' });

    // 获取网页标题作为文件名的一部分
    let urlName = await page.evaluate(() => document.title);
    urlName = urlName.replace(/[^a-zA-Z0-9\u4e00-\u9fa5]/g, '_');
    urlName = urlName.length > 10 ? urlName.substring(0, 10) : urlName;
    urlName = `${baseUrl}${urlName}`;
    // 获取主题区域内容，如果不存在则回退到 body
    let topicContent;
    try {
      topicContent = await page.$eval('#Wrapper', el => el.innerHTML);
    } catch (e) {
      topicContent = await page.evaluate(() => document.body.innerHTML);
    }

    // 构建完整 HTML 页面
    const fullHTML = `<!DOCTYPE html>
<html lang="zh">
<head>
  <meta charset="UTF-8">
  <title>${urlName}</title>
  <style>body { background-color: white; }</style>
</head>
<body>
  <div id="content">
    ${topicContent}
  </div>
</body>
</html>`;

    // 转换为 Markdown 并保存
    if (options.md) {
      const turndownService = new TurndownService();
      const markdown = turndownService.turndown(topicContent);
      fs.writeFileSync(`${urlName}.md`, markdown);
      console.log(`Markdown saved: v2ex_page_${urlName}.md`);
    }

    // 保存为 HTML 文件
    if (options.html) {
      fs.writeFileSync(`${urlName}.html`, fullHTML);
      console.log(`HTML saved: v2ex_page_${urlName}.html`);
    }

    // 保存为 PDF
    if (options.pdf) {
      await page.pdf({
        path: `{urlName}.pdf`,
        format: 'A4',
        printBackground: true
      });
      console.log(`PDF saved: v2ex_page_${urlName}.pdf`);
    }

    // 全页截图
    if (options.png) {
      await page.screenshot({
        path: `${urlName}.png`,
        fullPage: true
      });
      console.log(`Screenshot saved: v2ex_full_${urlName}.png`);
    }
  }

  await browser.close();
})();