# 2023 noneland backend interview

## API Router

* `現貨`
  * [GET] 取得`現貨`交易規範 api/binance/spot/exchangeInfo
  * [GET] 取得`現貨`帳戶餘額 api/binance/spot/balance
  * [GET] 取得`現貨`帳戶轉入轉出紀錄 api/binance/spot/transfer/records
* `合約`
  * [GET] 取得`合約`交易規範 api/binance/futures/exchangeInfo
  * [GET] 取得`合約`帳戶餘額 api/binance/futures/balance

## 相關說明

* APIKEY 已埋入 header，待請購完成可直接打入指定 URL 串接測試
* 關於提到的前台、後台驗證機制，考量到安全性可採用 `JWT Token` 驗證，已實作 JWT 創建及驗證功能
  * 若需登入驗證及發送，後續優化可再補上 Login 相關驗證及返回 Token 邏輯
  * 避免 User 重複獲取 Token，應可採用緩存優化
* Limiter 已先使用 `go.uber.org/ratelimit` 實作並加入各 Router 中
  * 三方平台 exchangeInfo 雖不常變動，後續優化可再定時獲取設定熱更新
* 返回前端格式目前暫定如下：
```
  {
      errcode: 1,
      errmsg: "success",
      data: {
         "free": "10.12345"
      }
  }
```

## User 使用現貨 API 流量限制問題

* 假定有幾個情境是確定的：
  * 平台的 exchaneInfo Limiter 規則絕對優先，以能正常使用不被 ban 為原則
  * User 使用的報價 API Request 跟一般後台使用的現貨 API Request `可區分`，這樣就可以針對不同的來源設定幾組 Channel 來監聽
  * Server 接收到的 Request 通通要處理完畢，不採用逾時拋棄制
* 構想是這樣:
  * 先定義一組`介面`，讓各種實作的 method 可以依照不同的情境來被註冊使用
  * 設定普通、緊急的任務 channel，把可以被區分的 API 各自推入有緩衝的 channel
  * 定義普通、緊急任務的寬限值，這邊預設了一個 Default 的邏輯，緊急任務優先處理，普通任務累積一個值後分批處理
  * 使用 channel, atomic 設值、取值，保證了原子性與執行緒安全
  * 相關實作的範例在 internal\task\defaulttask.go，未能盡善盡美，有問題還請多包涵與指點
