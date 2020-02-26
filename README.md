
# StockTerm

![demo-img](https://raw.githubusercontent.com/iamquang95/stockterm/master/resource/demo_screen.png)

StockTerm is a reactive stock portfolio observer.

The screen shot was taken on a bad day :( The 2019-nCoV is destroying the market.

## How to use
Just simply copy `sample.config.json` to `config.json` at the same level with `main.go`
Then run `go run *.go`

## Config file format
`watchingStocks` has to contains all your stocks (`portfolio` has to be the subset of `watchingStocks`)

The first four watching stocks will has last trading day price chart.
`portfolio` includes all your stocks with number of purchased with average purchase price.

## Dependencies 
[https://github.com/iamquang95/termui](https://github.com/iamquang95/termui) (This is my forked branch of [https://github.com/gizak](https://github.com/gizak) with some enhancement of line chart rendering)

Data is fetched from cafef and vndirect.

## License
[MIT](http://opensource.org/licenses/MIT)
