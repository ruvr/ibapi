package ibapi

import (
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/govalues/decimal"
)

// IbWrapper contain the funcs to handle the msg from TWS or Gateway
type IbWrapper interface {
	TickPrice(reqID int64, tickType int64, price float64, attrib TickAttrib)
	TickSize(reqID int64, tickType int64, size decimal.Decimal)
	OrderStatus(orderID int64, status string, filled decimal.Decimal, remaining decimal.Decimal, avgFillPrice float64, permID int64, parentID int64, lastFillPrice float64, clientID int64, whyHeld string, mktCapPrice float64)
	Error(reqID int64, errCode int64, errString string)
	OpenOrder(orderID int64, contract *Contract, order *Order, orderState *OrderState)
	UpdateAccountValue(tag string, val string, currency string, accName string)
	UpdatePortfolio(contract *Contract, position decimal.Decimal, marketPrice float64, marketValue float64, averageCost float64, unrealizedPNL float64, realizedPNL float64, accName string)
	UpdateAccountTime(accTime time.Time)
	NextValidID(reqID int64)
	ContractDetails(reqID int64, conDetails *ContractDetails)
	ExecDetails(reqID int64, contract *Contract, execution *Execution)
	UpdateMktDepth(reqID int64, position int64, operation int64, side int64, price float64, size decimal.Decimal)
	UpdateMktDepthL2(reqID int64, position int64, marketMaker string, operation int64, side int64, price float64, size decimal.Decimal, isSmartDepth bool)
	UpdateNewsBulletin(msgID int64, msgType int64, newsMessage string, originExchange string)
	ManagedAccounts(accountsList []string)
	ReceiveFA(faData int64, cxml string)
	HistoricalData(reqID int64, bar *BarData)
	HistoricalDataEnd(reqID int64, startDateStr string, endDateStr string)
	HistoricalDataUpdate(reqID int64, bar *BarData)
	BondContractDetails(reqID int64, conDetails *ContractDetails)
	ScannerParameters(xml string)
	ScannerData(reqID int64, rank int64, conDetails *ContractDetails, distance string, benchmark string, projection string, legs string)
	ScannerDataEnd(reqID int64)
	TickOptionComputation(reqID int64, tickType int64, tickAttrib int64, impliedVol float64, delta float64, optPrice float64, pvDiviedn float64, gamma float64, vega float64, theta float64, undPrice float64)
	TickGeneric(reqID int64, tickType int64, value float64)
	TickString(reqID int64, tickType int64, value string)
	TickEFP(reqID int64, tickType int64, basisPoints float64, formattedBasisPoints string, totalDividends float64, holdDays int64, futureLastTradeDate string, dividendImpact float64, dividendsToLastTradeDate float64)
	CurrentTime(t time.Time)
	RealtimeBar(reqID int64, time int64, open float64, high float64, low float64, close float64, volume decimal.Decimal, wap decimal.Decimal, count int64)
	FundamentalData(reqID int64, data string)
	ContractDetailsEnd(reqID int64)
	OpenOrderEnd()
	AccountDownloadEnd(accName string)
	ExecDetailsEnd(reqID int64)
	DeltaNeutralValidation(reqID int64, deltaNeutralContract DeltaNeutralContract)
	TickSnapshotEnd(reqID int64)
	MarketDataType(reqID int64, marketDataType int64)
	Position(account string, contract *Contract, position decimal.Decimal, avgCost float64)
	PositionEnd()
	AccountSummary(reqID int64, account string, tag string, value string, currency string)
	AccountSummaryEnd(reqID int64)
	VerifyMessageAPI(apiData string)
	VerifyCompleted(isSuccessful bool, err string)
	DisplayGroupList(reqID int64, groups string)
	DisplayGroupUpdated(reqID int64, contractInfo string)
	VerifyAndAuthMessageAPI(apiData string, xyzChallange string)
	VerifyAndAuthCompleted(isSuccessful bool, err string)
	PositionMulti(reqID int64, account string, modelCode string, contract *Contract, position decimal.Decimal, avgCost float64)
	PositionMultiEnd(reqID int64)
	AccountUpdateMulti(reqID int64, account string, modleCode string, tag string, value string, currency string)
	AccountUpdateMultiEnd(reqID int64)
	SecurityDefinitionOptionParameter(reqID int64, exchange string, underlyingContractID int64, tradingClass string, multiplier string, expirations []string, strikes []float64)
	SecurityDefinitionOptionParameterEnd(reqID int64)
	SoftDollarTiers(reqID int64, tiers []SoftDollarTier)
	FamilyCodes(famCodes []FamilyCode)
	SymbolSamples(reqID int64, contractDescriptions []ContractDescription)
	SmartComponents(reqID int64, smartComps []SmartComponent)
	TickReqParams(tickerID int64, minTick float64, bboExchange string, snapshotPermissions int64)
	MktDepthExchanges(depthMktDataDescriptions []DepthMktDataDescription)
	HeadTimestamp(reqID int64, headTimestamp string)
	TickNews(tickerID int64, timeStamp int64, providerCode string, articleID string, headline string, extraData string)
	NewsProviders(newsProviders []NewsProvider)
	NewsArticle(reqID int64, articleType int64, articleText string)
	HistoricalNews(reqID int64, time string, providerCode string, articleID string, headline string)
	HistoricalNewsEnd(reqID int64, hasMore bool)
	HistogramData(reqID int64, histogram []HistogramData)
	RerouteMktDataReq(reqID int64, contractID int64, exchange string)
	RerouteMktDepthReq(reqID int64, contractID int64, exchange string)
	MarketRule(marketRuleID int64, priceIncrements []PriceIncrement)
	Pnl(reqID int64, dailyPnL float64, unrealizedPnL float64, realizedPnL float64)
	PnlSingle(reqID int64, position decimal.Decimal, dailyPnL float64, unrealizedPnL float64, realizedPnL float64, value float64)
	HistoricalTicks(reqID int64, ticks []HistoricalTick, done bool)
	HistoricalTicksBidAsk(reqID int64, ticks []HistoricalTickBidAsk, done bool)
	HistoricalTicksLast(reqID int64, ticks []HistoricalTickLast, done bool)
	TickByTickAllLast(reqID int64, tickType int64, time int64, price float64, size decimal.Decimal, tickAttribLast TickAttribLast, exchange string, specialConditions string)
	TickByTickBidAsk(reqID int64, time int64, bidPrice float64, askPrice float64, bidSize decimal.Decimal, askSize decimal.Decimal, tickAttribBidAsk TickAttribBidAsk)
	TickByTickMidPoint(reqID int64, time int64, midPoint float64)
	OrderBound(reqID int64, apiClientID int64, apiOrderID int64)
	CompletedOrder(contract *Contract, order *Order, orderState *OrderState)
	CompletedOrdersEnd()
	CommissionReport(commissionReport CommissionReport)
	ConnectAck()
	ConnectionClosed()
	ReplaceFAEnd(reqID int64, text string)
	WshMetaData(reqID int64, dataJson string)
	WshEventData(reqID int64, dataJson string)
	UserInfo(reqID int64, whiteBrandingId string)
	HistoricalSchedule(reqID int64, startDateTime string, endDateTime string, timeZone string, sessions []HistoricalSession)
}

// Wrapper is the default wrapper provided by this golang implement.
type Wrapper struct {
	orderID int64
}

func (w *Wrapper) GetNextOrderID() (i int64) {
	i = w.orderID
	atomic.AddInt64(&w.orderID, 1)
	return
}

func (w Wrapper) ConnectAck() {
	log.Info("<ConnectAck>...")
}

func (w Wrapper) ConnectionClosed() {
	log.Info("<ConnectionClosed>...")
}

func (w *Wrapper) NextValidID(reqID int64) {
	atomic.StoreInt64(&w.orderID, reqID)
	log.Info("<NextValidID>", slog.Int64("reqID", reqID))
}

func (w Wrapper) ManagedAccounts(accountsList []string) {
	log.Info("<ManagedAccounts>", slog.Any("accountList", accountsList))
}

func (w Wrapper) TickPrice(reqID int64, tickType int64, price float64, attrib TickAttrib) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickPrice>", slog.Int64("tickType", tickType), slog.Float64("price", price))
}

func (w Wrapper) UpdateAccountTime(accTime time.Time) {
	log.Info("<UpdateAccountTime>", slog.Time("accountTime", accTime))
}

func (w Wrapper) UpdateAccountValue(tag string, value string, currency string, account string) {
	log.Info("<UpdateAccountValue>", slog.String("tag", tag), slog.String("value", value), slog.String("currency", currency), slog.String("account", account))
}

func (w Wrapper) AccountDownloadEnd(accName string) {
	log.Info("<AccountDownloadEnd>", slog.String("accountName", accName))
}

func (w Wrapper) AccountUpdateMulti(reqID int64, account string, modelCode string, tag string, value string, currency string) {
	log.With(slog.Int64("reqID", reqID)).Info("<AccountUpdateMulti>",
		slog.String("account", account),
		slog.String("modelCode", modelCode),
		slog.String("tag", tag),
		slog.String("value", value),
		slog.String("curreny", currency),
	)
}

func (w Wrapper) AccountUpdateMultiEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<AccountUpdateMultiEnd>")
}

func (w Wrapper) AccountSummary(reqID int64, account string, tag string, value string, currency string) {
	log.With(slog.Int64("reqID", reqID)).Info("<AccountSummary>",
		slog.String("account", account),
		slog.String("tag", tag),
		slog.String("value", value),
		slog.String("curreny", currency),
	)

}

func (w Wrapper) AccountSummaryEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<AccountSummaryEnd>")
}

func (w Wrapper) VerifyMessageAPI(apiData string) {
	log.Info("<VerifyMessageAPI>", slog.String("apiData", apiData))
}

func (w Wrapper) VerifyCompleted(isSuccessful bool, err string) {
	log.Info("<VerifyCompleted>", slog.Bool("isSuccessful", isSuccessful), slog.String("error", err))
}

func (w Wrapper) VerifyAndAuthMessageAPI(apiData string, xyzChallange string) {
	log.Info("<VerifyMessageAPI>", slog.String("apiData", apiData), slog.String("xyzChallange", xyzChallange))
}

func (w Wrapper) VerifyAndAuthCompleted(isSuccessful bool, err string) {
	log.Info("<VerifyCompleted>", slog.Bool("isSuccessful", isSuccessful), slog.String("error", err))
}

func (w Wrapper) DisplayGroupList(reqID int64, groups string) {
	log.With(slog.Int64("reqID", reqID)).Info("<DisplayGroupList>", slog.String("groups", groups))
}

func (w Wrapper) DisplayGroupUpdated(reqID int64, contractInfo string) {
	log.With(slog.Int64("reqID", reqID)).Info("<DisplayGroupUpdated>", slog.String("contractInfo", contractInfo))
}

func (w Wrapper) PositionMulti(reqID int64, account string, modelCode string, contract *Contract, position decimal.Decimal, avgCost float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<PositionMulti>",
		slog.String("account", account),
		slog.String("modelCode", modelCode),
		slog.Any("contract", contract),
		slog.String("position", position.String()),
		slog.Float64("avgCost", avgCost),
	)
}

func (w Wrapper) PositionMultiEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<PositionMultiEnd>")
}

func (w Wrapper) UpdatePortfolio(contract *Contract, position decimal.Decimal, marketPrice float64, marketValue float64, averageCost float64, unrealizedPNL float64, realizedPNL float64, accName string) {
	log.Info("<UpdatePortfolio>",
		slog.String("localSymbol", contract.LocalSymbol),
		slog.String("position", position.String()),
		slog.Float64("marketPrice", marketPrice),
		slog.Float64("averageCost", averageCost),
		slog.Float64("unrealizedPNL", unrealizedPNL),
		slog.Float64("realizedPNL", realizedPNL),
		slog.String("accName", accName),
	)
}

func (w Wrapper) Position(account string, contract *Contract, position decimal.Decimal, avgCost float64) {
	log.Info("<UpdatePortfolio>",
		slog.String("account", account),
		slog.Any("contract", contract),
		slog.String("position", position.String()),
		slog.Float64("avgCost", avgCost),
	)
}

func (w Wrapper) PositionEnd() {
	log.Info("<PositionEnd>")
}

func (w Wrapper) Pnl(reqID int64, dailyPnL float64, unrealizedPnL float64, realizedPnL float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<PNL>",
		slog.Float64("dailyPnL", dailyPnL),
		slog.Float64("unrealizedPnL", unrealizedPnL),
		slog.Float64("realizedPnL", realizedPnL),
	)
}

func (w Wrapper) PnlSingle(reqID int64, position decimal.Decimal, dailyPnL float64, unrealizedPnL float64, realizedPnL float64, value float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<PNLSingle>",
		slog.Float64("dailyPnL", dailyPnL),
		slog.Float64("unrealizedPnL", unrealizedPnL),
		slog.Float64("realizedPnL", realizedPnL),
		slog.Float64("value", value),
	)
}

func (w Wrapper) OpenOrder(orderID int64, contract *Contract, order *Order, orderState *OrderState) {
	log.With(slog.Int64("orderID", orderID)).Info("<OpenOrder>",
		slog.Any("contract", contract),
		slog.Any("order", order),
		slog.Any("orderState", orderState),
	)
}

func (w Wrapper) OpenOrderEnd() {
	log.Info("<OpenOrderEnd>")

}

func (w Wrapper) OrderStatus(orderID int64, status string, filled decimal.Decimal, remaining decimal.Decimal, avgFillPrice float64, permID int64, parentID int64, lastFillPrice float64, clientID int64, whyHeld string, mktCapPrice float64) {
	log.With(slog.Int64("orderID", orderID)).Info("<OrderStatus>",
		slog.String("status", status),
		slog.String("filled", filled.String()),
		slog.String("remaining", remaining.String()),
		slog.Float64("avgFillPrice", avgFillPrice),
	)
}

func (w Wrapper) ExecDetails(reqID int64, contract *Contract, execution *Execution) {
	log.With(slog.Int64("reqID", reqID)).Info("<ExecDetails>",
		slog.Any("contract", contract),
		slog.Any("execution", execution),
	)
}

func (w Wrapper) ExecDetailsEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<ExecDetailsEnd>")
}

func (w Wrapper) DeltaNeutralValidation(reqID int64, deltaNeutralContract DeltaNeutralContract) {
	log.With(slog.Int64("reqID", reqID)).Info("<DeltaNeutralValidation>",
		slog.Any("deltaNeutralContract", deltaNeutralContract),
	)
}

func (w Wrapper) CommissionReport(commissionReport CommissionReport) {
	log.Info("<CommissionReport>",
		slog.Any("commissionReport", commissionReport),
	)
}

func (w Wrapper) OrderBound(reqID int64, apiClientID int64, apiOrderID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<OrderBound>",
		slog.Int64("apiClientID", apiClientID),
		slog.Int64("apiOrderID", apiOrderID),
	)
}

func (w Wrapper) ContractDetails(reqID int64, conDetails *ContractDetails) {
	log.With(slog.Int64("reqID", reqID)).Info("<ContractDetails>",
		slog.Any("conDetails", conDetails),
	)
}

func (w Wrapper) ContractDetailsEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<ContractDetailsEnd>")
}

func (w Wrapper) BondContractDetails(reqID int64, conDetails *ContractDetails) {
	log.With(slog.Int64("reqID", reqID)).Info("<BondContractDetails>",
		slog.Any("conDetails", conDetails),
	)
}

func (w Wrapper) SymbolSamples(reqID int64, contractDescriptions []ContractDescription) {
	log.With(slog.Int64("reqID", reqID)).Info("<SymbolSamples>",
		slog.Any("contractDescriptions", contractDescriptions),
	)
}

func (w Wrapper) SmartComponents(reqID int64, smartComps []SmartComponent) {
	log.With(slog.Int64("reqID", reqID)).Info("<SmartComponents>",
		slog.Any("smartComps", smartComps),
	)
}

func (w Wrapper) MarketRule(marketRuleID int64, priceIncrements []PriceIncrement) {
	log.Info("<MarketRule>",
		slog.Int64("marketRuleID", marketRuleID),
		slog.Any("priceIncrements", priceIncrements),
	)
}

func (w Wrapper) RealtimeBar(reqID int64, time int64, open float64, high float64, low float64, close float64, volume decimal.Decimal, wap decimal.Decimal, count int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<RealtimeBar>",
		slog.Int64("time", time),
		slog.Float64("open", open),
		slog.Float64("high", high),
		slog.Float64("low", low),
		slog.Float64("close", close),
		slog.String("volume", volume.String()),
		slog.String("wap", wap.String()),
		slog.Int64("count", count),
	)
}

func (w Wrapper) HistoricalData(reqID int64, bar *BarData) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalData>",
		slog.Any("bar", bar),
	)
}

func (w Wrapper) HistoricalDataEnd(reqID int64, startDateStr string, endDateStr string) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalDataEnd>",
		slog.String("startDate", startDateStr),
		slog.String("endDate", endDateStr),
	)
}

func (w Wrapper) HistoricalDataUpdate(reqID int64, bar *BarData) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalDataUpdate>",
		slog.Any("bar", bar),
	)
}

func (w Wrapper) HeadTimestamp(reqID int64, headTimestamp string) {
	log.With(slog.Int64("reqID", reqID)).Info("<HeadTimestamp>",
		slog.String("headTimestamp", headTimestamp),
	)
}

func (w Wrapper) HistoricalTicks(reqID int64, ticks []HistoricalTick, done bool) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalTicks>",
		slog.Any("ticks", ticks),
		slog.Bool("done", done),
	)
}

func (w Wrapper) HistoricalTicksBidAsk(reqID int64, ticks []HistoricalTickBidAsk, done bool) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalTicksBidAsk>",
		slog.Any("ticks", ticks),
		slog.Bool("done", done),
	)
}

func (w Wrapper) HistoricalTicksLast(reqID int64, ticks []HistoricalTickLast, done bool) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalTicksLast>",
		slog.Any("ticks", ticks),
		slog.Bool("done", done),
	)
}

func (w Wrapper) TickSize(reqID int64, tickType int64, size decimal.Decimal) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickSize>",
		slog.Int64("tickType", tickType),
		slog.String("size", size.String()),
	)
}

func (w Wrapper) TickSnapshotEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickSnapshotEnd>")
}

func (w Wrapper) MarketDataType(reqID int64, marketDataType int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<MarketDataType>",
		slog.Int64("marketDataType", marketDataType),
	)
}

func (w Wrapper) TickByTickAllLast(reqID int64, tickType int64, time int64, price float64, size decimal.Decimal, tickAttribLast TickAttribLast, exchange string, specialConditions string) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickByTickAllLast>",
		slog.Int64("tickType", tickType),
		slog.Int64("time", time),
		slog.Float64("price", price),
		slog.String("size", size.String()),
	)
}

func (w Wrapper) TickByTickBidAsk(reqID int64, time int64, bidPrice float64, askPrice float64, bidSize decimal.Decimal, askSize decimal.Decimal, tickAttribBidAsk TickAttribBidAsk) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickByTickBidAsk>",
		slog.Int64("time", time),
		slog.Float64("bidPrice", bidPrice),
		slog.Float64("askPrice", askPrice),
		slog.String("bidPrice", bidSize.String()),
		slog.String("askPrice", askSize.String()),
	)
}

func (w Wrapper) TickByTickMidPoint(reqID int64, time int64, midPoint float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickByTickMidPoint>",
		slog.Int64("time", time),
		slog.Float64("midPoint", midPoint),
	)
}

func (w Wrapper) TickString(reqID int64, tickType int64, value string) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickString>",
		slog.Int64("tickType", tickType),
		slog.String("value", value),
	)
}

func (w Wrapper) TickGeneric(reqID int64, tickType int64, value float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickGeneric>",
		slog.Int64("tickType", tickType),
		slog.Float64("value", value),
	)
}

func (w Wrapper) TickEFP(reqID int64, tickType int64, basisPoints float64, formattedBasisPoints string, totalDividends float64, holdDays int64, futureLastTradeDate string, dividendImpact float64, dividendsToLastTradeDate float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickEFP>",
		slog.Int64("tickType", tickType),
		slog.Float64("basisPoints", basisPoints),
	)
}

func (w Wrapper) TickReqParams(reqID int64, minTick float64, bboExchange string, snapshotPermissions int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickReqParams>",
		slog.Float64("minTick", minTick),
		slog.String("bboExchange", bboExchange),
		slog.Int64("snapshotPermissions", snapshotPermissions),
	)
}
func (w Wrapper) MktDepthExchanges(depthMktDataDescriptions []DepthMktDataDescription) {
	log.Info("<MktDepthExchanges>",
		slog.Any("depthMktDataDescriptions", depthMktDataDescriptions),
	)
}

/*
Returns the order book.

tickerId -  the request's identifier
position -  the order book's row being updated
operation - how to refresh the row:

	0 = insert (insert this new order into the row identified by 'position')
	1 = update (update the existing order in the row identified by 'position')
	2 = delete (delete the existing order at the row identified by 'position').

side -  0 for ask, 1 for bid
price - the order's price
size -  the order's size
*/
func (w Wrapper) UpdateMktDepth(reqID int64, position int64, operation int64, side int64, price float64, size decimal.Decimal) {
	log.With(slog.Int64("reqID", reqID)).Info("<UpdateMktDepth>",
		slog.Int64("position", position),
		slog.Int64("operation", operation),
		slog.Int64("side", side),
		slog.Float64("price", price),
		slog.String("size", size.String()),
	)
}

func (w Wrapper) UpdateMktDepthL2(reqID int64, position int64, marketMaker string, operation int64, side int64, price float64, size decimal.Decimal, isSmartDepth bool) {
	log.With(slog.Int64("reqID", reqID)).Info("<UpdateMktDepthL2>",
		slog.Int64("position", position),
		slog.String("marketMaker", marketMaker),
		slog.Int64("operation", operation),
		slog.Int64("side", side),
		slog.Float64("price", price),
		slog.String("size", size.String()),
		slog.Bool("isSmartDepth", isSmartDepth),
	)
}

func (w Wrapper) TickOptionComputation(reqID int64, tickType int64, tickAttrib int64, impliedVol float64, delta float64, optPrice float64, pvDiviedn float64, gamma float64, vega float64, theta float64, undPrice float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<TickOptionComputation>",
		slog.Int64("tickType", tickType),
		slog.Int64("tickAttrib", tickAttrib),
		slog.Float64("impliedVol", impliedVol),
		slog.Float64("delta", delta),
		slog.Float64("optPrice", optPrice),
		slog.Float64("pvDiviedn", pvDiviedn),
		slog.Float64("gamma", gamma),
		slog.Float64("vega", vega),
		slog.Float64("theta", theta),
		slog.Float64("undPrice", undPrice),
	)
}

func (w Wrapper) FundamentalData(reqID int64, data string) {
	log.With(slog.Int64("reqID", reqID)).Info("<FundamentalData>",
		slog.String("data", data),
	)
}

func (w Wrapper) ScannerParameters(xml string) {
	log.Info("<ScannerParameters>",
		slog.String("xml", xml),
	)

}

func (w Wrapper) ScannerData(reqID int64, rank int64, conDetails *ContractDetails, distance string, benchmark string, projection string, legs string) {
	log.With(slog.Int64("reqID", reqID)).Info("<ScannerData>",
		slog.Int64("rank", rank),
		slog.Any("conDetails", conDetails),
		slog.String("distance", distance),
		slog.String("benchmark", benchmark),
		slog.String("projection", projection),
		slog.String("legs", legs),
	)
}

func (w Wrapper) ScannerDataEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<ScannerDataEnd>")
}

func (w Wrapper) HistogramData(reqID int64, histogram []HistogramData) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistogramData>",
		slog.Any("histogram", histogram),
	)
}

func (w Wrapper) RerouteMktDataReq(reqID int64, contractID int64, exchange string) {
	log.With(slog.Int64("reqID", reqID)).Info("<RerouteMktDataReq>",
		slog.Int64("contractID", contractID),
		slog.String("exchange", exchange),
	)
}

func (w Wrapper) RerouteMktDepthReq(reqID int64, contractID int64, exchange string) {
	log.With(slog.Int64("reqID", reqID)).Info("<RerouteMktDepthReq>",
		slog.Int64("contractID", contractID),
		slog.String("exchange", exchange),
	)
}

func (w Wrapper) SecurityDefinitionOptionParameter(reqID int64, exchange string, underlyingContractID int64, tradingClass string, multiplier string, expirations []string, strikes []float64) {
	log.With(slog.Int64("reqID", reqID)).Info("<SecurityDefinitionOptionParameter>",
		slog.String("exchange", exchange),
		slog.Int64("underlyingContractID", underlyingContractID),
		slog.String("tradingClass", tradingClass),
		slog.String("multiplier", multiplier),
		slog.Any("expirations", expirations),
		slog.Any("strikes", strikes),
	)
}

func (w Wrapper) SecurityDefinitionOptionParameterEnd(reqID int64) {
	log.With(slog.Int64("reqID", reqID)).Info("<SecurityDefinitionOptionParameterEnd>")
}

func (w Wrapper) SoftDollarTiers(reqID int64, tiers []SoftDollarTier) {
	log.With(slog.Int64("reqID", reqID)).Info("<SoftDollarTiers>",
		slog.Any("tiers", tiers),
	)
}

func (w Wrapper) FamilyCodes(famCodes []FamilyCode) {
	log.Info("<FamilyCodes>",
		slog.Any("famCodes", famCodes),
	)
}

func (w Wrapper) NewsProviders(newsProviders []NewsProvider) {
	log.Info("<NewsProviders>",
		slog.Any("newsProviders", newsProviders),
	)
}

func (w Wrapper) TickNews(tickerID int64, timeStamp int64, providerCode string, articleID string, headline string, extraData string) {
	log.With(slog.Int64("tickerID", tickerID)).Info("<TickNews>",
		slog.Int64("timeStamp", timeStamp),
		slog.String("providerCode", providerCode),
		slog.String("articleID", articleID),
		slog.String("headline", headline),
		slog.String("extraData", extraData),
	)
}

func (w Wrapper) NewsArticle(reqID int64, articleType int64, articleText string) {
	log.With(slog.Int64("reqID", reqID)).Info("<NewsArticle>",
		slog.Int64("articleType", articleType),
		slog.String("articleText", articleText),
	)
}

func (w Wrapper) HistoricalNews(reqID int64, time string, providerCode string, articleID string, headline string) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalNews>",
		slog.String("time", time),
		slog.String("providerCode", providerCode),
		slog.String("articleID", articleID),
		slog.String("headline", headline),
	)
}

func (w Wrapper) HistoricalNewsEnd(reqID int64, hasMore bool) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalNewsEnd>",
		slog.Bool("hasMore", hasMore),
	)
}

func (w Wrapper) UpdateNewsBulletin(msgID int64, msgType int64, newsMessage string, originExch string) {
	log.With(slog.Int64("msgID", msgID)).Info("<UpdateNewsBulletin>",
		slog.Int64("msgType", msgType),
		slog.String("newsMessage", newsMessage),
		slog.String("originExch", originExch),
	)
}

func (w Wrapper) ReceiveFA(faData int64, cxml string) {
	log.Info("<ReceiveFA>",
		slog.Int64("faData", faData),
		slog.String("cxml", cxml),
	)
}

func (w Wrapper) CurrentTime(t time.Time) {
	log.Info("<CurrentTime>",
		slog.Time("time", t),
	)
}

func (w Wrapper) Error(reqID int64, errCode int64, errString string) {
	log.With(slog.Int64("reqID", reqID)).Info("<Error>",
		slog.Int64("errCode", errCode),
		slog.String("errString", errString),
	)
}

func (w Wrapper) CompletedOrder(contract *Contract, order *Order, orderState *OrderState) {
	log.Info("<CompletedOrder>",
		slog.Any("contract", contract),
		slog.Any("order", order),
		slog.Any("orderState", orderState),
	)
}

func (w Wrapper) CompletedOrdersEnd() {
	log.Info("<CompletedOrdersEnd>")
}

func (w Wrapper) ReplaceFAEnd(reqID int64, text string) {
	log.With(slog.Int64("reqID", reqID)).Info("<ReplaceFAEnd>", slog.String("text", text))
}

func (w Wrapper) WshMetaData(reqID int64, dataJson string) {
	log.With(slog.Int64("reqID", reqID)).Info("<WshMetaData>", slog.String("dataJson", dataJson))
}
func (w Wrapper) WshEventData(reqID int64, dataJson string) {
	log.With(slog.Int64("reqID", reqID)).Info("<WshEventData>", slog.String("dataJson", dataJson))
}

func (w Wrapper) UserInfo(reqID int64, whiteBrandingId string) {
	log.With(slog.Int64("reqID", reqID)).Info("<UserInfo>", slog.String("whiteBrandingId", whiteBrandingId))
}

func (w Wrapper) HistoricalSchedule(reqID int64, startDateTime string, endDateTime string, timeZone string, sessions []HistoricalSession) {
	log.With(slog.Int64("reqID", reqID)).Info("<HistoricalSchedule>", slog.String("startDateTime", startDateTime), slog.String("endDateTime", endDateTime), slog.String("timeZone", timeZone))
}
