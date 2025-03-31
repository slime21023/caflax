# Caflax

一個基於 Golang 與 Iris web framework 實現的輕量級靜態檔案伺服器，靈感來自 [Vercel/Serve](https://github.com/vercel/serve)。Caflax 專為快速在本地提供靜態網站或檔案而設計。

## 特點

- **快速** - 使用 Go 和 Iris 框架構建，高效能伺服靜態內容
- **SPA 模式** - 完美支援單頁應用程式，自動將路由導向 index.html
- **安全** - 自動添加安全相關標頭，保護您的應用程式
- **CORS 支援** - 可選開啟 Cross-Origin Resource Sharing
- **壓縮** - 支援 gzip 和 brotli 壓縮，減少傳輸大小
- **快取控制** - 可選禁用瀏覽器快取，方便開發階段使用
- **目錄列表** - 支援顯示目錄內容列表
- **簡單易用** - 直觀的命令行界面，適合各種開發場景

## 使用方式

### 基本用法

```bash
# 在當前目錄啟動伺服器
caflax

# 指定要伺服的目錄
caflax ./dist

# 指定埠號
caflax -p 8080

# 指定主機名和埠號
caflax -H 0.0.0.0 -p 8000 ./public
```

### 進階選項

```bash
# 啟用 SPA 模式 (所有不存在的路徑都會導向 index.html)
caflax --spa

# 啟用 CORS
caflax --cors

# 禁用瀏覽器快取
caflax --no-cache

# 啟用壓縮 (預設已啟用)
caflax --compress

# 安靜模式，減少日誌輸出
caflax --quiet

# 顯示版本資訊
caflax --version
```

### 完整幫助資訊

執行 `caflax --help` 查看完整的命令行選項說明。

## 適用場景

### 開發環境
- 本地開發靜態網站
- 展示前端專案
- 測試單頁應用程式
- 臨時分享靜態檔案或文檔
- 快速啟動一個本地伺服器

### 生產環境考量

Caflax 可用於下列生產環境場景，但需要注意一些限制：

**適合場景**：
- 小型內部網站或工具的靜態檔案伺服
- 中小流量的靜態內容網站
- 結合 CDN 的靜態資源分發節點
- 公司內部文件或知識庫託管

**使用建議**：
- 在生產環境中，建議在 Caflax 前方配置 Nginx/Caddy 等反向代理以提供 HTTPS 支援
- 對於高流量網站，考慮使用專業的靜態檔案託管服務
- 涉及敏感資料時，應增加額外的認證和授權層
- 搭配監控工具以追蹤伺服器性能和狀態