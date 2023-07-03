package main

import (
  "math"
  "fmt"
)

type StockPoint struct {
  Price float64
  UpperAcceptableMean float64
  LowerAcceptedMean float64
}

func PrintPoint(point StockPoint) {
  fmt.Printf("At price $%f, upper $%f, lower $%f\n", point.Price, point.UpperAcceptableMean, point.LowerAcceptedMean)
}

const Buy = 1
const Sell = -1
const Keep = 0
func getStrategyFromPoint(stockPoint StockPoint) int {
  if stockPoint.Price < stockPoint.LowerAcceptedMean {
    return Buy
  } else if stockPoint.Price > stockPoint.UpperAcceptableMean {
    return Sell
  }
  return Keep
}

type Wallet struct {
  Cash float64
  Stocks int
}
func PrintWallet(wallet Wallet) {
  fmt.Printf("Client Wallet has $%f in cash and %d stocks\n", wallet.Cash, wallet.Stocks)
}

func ApplyStrategy(wallet Wallet, points []StockPoint) Wallet {
  for i, point := range points {
    fmt.Printf("------- day %d -------\n", i)
    PrintWallet(wallet)
    PrintPoint(point)
    switch(getStrategyFromPoint(point)) {
      case Buy:
        shouldBuyQuantity := math.Floor(wallet.Cash / point.Price)
        wallet.Cash -= shouldBuyQuantity * point.Price
        wallet.Stocks += int(shouldBuyQuantity)
        fmt.Printf("Did buy %f\n", shouldBuyQuantity)
      case Sell:
        shouldSellQuantity := wallet.Stocks / 3
        wallet.Cash += float64(shouldSellQuantity) * point.Price
        wallet.Stocks -= shouldSellQuantity
        fmt.Printf("Did sell %d\n", shouldSellQuantity)
    }
  }
  return wallet
}

func main() {
  PrintWallet(ApplyStrategy(Wallet{1000, 100}, []StockPoint{
    {12, 14.78, 12.51},
    {14.04, 14.01, 12.55},
    {12.42, 14.79, 12.85},
    {15, 14.95, 12.48},
    {35, 34.71, 25.24},
    {27.6, 34.91, 29.01},
    {35.30, 34.21, 30.33},
    {33.22, 40.52, 33.80},
    {41.92, 41.72, 34.83},
    {41.72, 41.53, 34.96},
    {36.66, 41.58, 37},
    {37.75, 41.70, 37.90},
    {41.86, 41.70, 37.30},
    {38.5, 42.43, 38.81}}))
}
