// 使用 puppeteer-core
const puppeteer = require('puppeteer-core');
const fs = require('fs');
// 分片截图
// 独立滚动截图函数
const scrollAndScreenshot = async (page, fileNamePrefix = 'screenshot') => {
  const viewportHeight = await page.evaluate(() => window.innerHeight);
  const totalHeight = await page.evaluate(() => document.body.scrollHeight);
  let currentScroll = 0;
  let screenshotCount = 0;

  while (currentScroll < totalHeight) {
    await page.evaluate(_scrollY => window.scrollTo(0, _scrollY), currentScroll);
    await new Promise(r => setTimeout(r, 500)); // 替换 waitForTimeout

    await page.screenshot({
      path: `${fileNamePrefix}_${screenshotCount}.png`,
      fullPage: false
    });

    currentScroll += viewportHeight;
    screenshotCount++;
  }

  return screenshotCount;
};

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

  // 保存为 PDF
  await page.pdf({
    path: '/Users/sure/Documents/v2ex_page.pdf',
    format: 'A4',
    printBackground: true
  });
  console.log('页面已保存为 PDF：v2ex_page.pdf');

  // 执行滚动截图
  const count = await scrollAndScreenshot(page, 'v2ex_screenshot');
  console.log(`已保存 ${count} 张截图作为滚动页面内容`);

  await browser.close();
})();