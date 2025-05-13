// 使用 puppeteer-core
const puppeteer = require('puppeteer-core');
const fs = require('fs');
const TurndownService = require('turndown');

(async () => {
  const browser = await puppeteer.launch({
    executablePath: '/Applications/Microsoft Edge.app/Contents/MacOS/Microsoft Edge',
    headless: false,
    args: [
      '--proxy-server=http://127.0.0.1:7890', // 使用你的代理
    ]
  });

  const page = await browser.newPage();
  
  // 设置真实浏览器的 User-Agent
  await page.setUserAgent('Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36');

  // 增加超时时间为 60 秒
  await page.goto('https://v2ex.com/t/1019827', { timeout: 60000 });
  
  // 强制页面背景为白色
  await page.addStyleTag({ content: 'body { background-color: white !important; }' });

  // 提取主题内容区域
  const topicContent = await page.$eval('#Wrapper', el => el.innerHTML);

  // 转换为 Markdown
  const turndownService = new TurndownService();
  const markdown = turndownService.turndown(topicContent);

  // 保存为 Markdown 文件
  fs.writeFileSync('v2ex_page.md', markdown);
  console.log('页面主题内容已保存为 Markdown：v2ex_page.md');

  // 全页截图
  await page.screenshot({
    path: 'v2ex_full.png',
    fullPage: true
  });
  console.log('已保存完整页面截图：v2ex_full.png');

  await browser.close();
})();