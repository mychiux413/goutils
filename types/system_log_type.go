package t

type SystemLogType string

const (
	/*
		通常會放些重要的成功訊息並讓 QA 可以進行確認
		主要提供關鍵的狀態供快速判斷系統狀態
		1. service 啟動成功與否
		2. 關鍵的 transactions 被執行完成與否
		3. configuration
	*/
	SYSTEM_LOG_TYPE_INFO SystemLogType = "info"

	/*
		可以用來警示的錯誤，出現時可以提供分析的資訊
		1. 網路暫時中斷
		2. 轉換的型態可能造成資料遺失
		3. 定時執行的任務失敗
	*/
	SYSTEM_LOG_TYPE_WARN SystemLogType = "warn"

	/*
		一般的錯誤，出現了就需要思考並排時間解決
		1. 網路連線無預期中斷且無法連回
		2. 傳輸資料失敗
		3. CRUD database 失敗
		4. 諸如此類會影響功能的錯誤
	*/
	SYSTEM_LOG_TYPE_ERROR SystemLogType = "error"

	/*
		最緊急的錯誤，只要一出現此類的錯誤，半夜看到也要從床上跳起來立刻給 hotfix 那種。
		1. 會嚴重影響財損的錯誤
		2. 通常此錯誤發生就會讓服務暫時停止運作
	*/
	SYSTEM_LOG_TYPE_FATAL SystemLogType = "fatal"
)
