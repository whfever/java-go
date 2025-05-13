// 使用 puppeteer-core
const puppeteer = require('puppeteer-core');
const fs = require('fs');

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
  
  // 获取页面HTML内容
  const content = await page.content();
  
  // 保存为HTML文件
  fs.writeFileSync('/Users/sure/Documents/v2ex_page.html', content);
  console.log('页面已保存为 v2ex_page.html');
  
  // 截图（可选）
  await page.screenshot({ path: 'v2ex.png' });
  
  await browser.close();
})();