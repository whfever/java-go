// playwright-all.js
const { chromium } = require('playwright');
const fs = require('fs');
const TurndownService = require('turndown');

// 默认内置多个网址
const DEFAULT_URLS = [
  'https://v2ex.com/t/1019827'
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

// 提取其他选项
const options = {
  pdf: args.includes('--pdf'),
  html: args.includes('--html'),
  md: args.includes('--md')
};

// 使用命令行或默认 URL 列表
const targetUrls = urls.length > 0 ? urls : DEFAULT_URLS;

// Microsoft Edge 的安装路径（适用于 macOS）
const EDGE_PATH = '/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge';

(async () => {
  // 启动 Microsoft Edge 浏览器
  const browser = await chromium.launch({ 
    executablePath: EDGE_PATH, 
    headless: false 
  });
  const page = await browser.newPage();
  
  // 设置 User-Agent
  await page.setUserAgent('Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36');

  for (const url of targetUrls) {
    // 跳转到目标网址
    await page.goto(url, { waitUntil: 'networkidle' });

    // 获取网页标题作为文件名的一部分
    let urlName = await page.title();
    urlName = urlName.replace(/[^a-zA-Z0-9\u4e00-\u9fa5]/g, '_');

    // 获取主体内容
    let topicContent;
    try {
      topicContent = await page.locator('#Wrapper').innerHTML();
    } catch (e) {
      topicContent = await page.content();
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
      fs.writeFileSync(`playwright_page_${urlName}.md`, markdown);
      console.log(`Markdown saved: playwright_page_${urlName}.md`);
    }

    // 保存为 HTML 文件
    if (options.html) {
      fs.writeFileSync(`playwright_page_${urlName}.html`, fullHTML);
      console.log(`HTML saved: playwright_page_${urlName}.html`);
    }

    // 保存为 PDF
    if (options.pdf) {
      await page.pdf({
        path: `playwright_page_${urlName}.pdf`,
        format: 'A4',
        printBackground: true
      });
      console.log(`PDF saved: playwright_page_${urlName}.pdf`);
    }
  }

  await browser.close();
})();