# Portfolio Tracker API

This project provides a RESTful API to manage a portfolio of stocks, allowing users to perform operations like adding trades, fetching portfolio holdings, and calculating cumulative returns.

Features
Add Trades: Users can add buy or sell trades to their portfolio.

Fetch Holdings: Get a list of all the securities in the portfolio with their quantities and average buy prices.

Calculate Returns: Calculate the cumulative returns of the portfolio based on a specified current price.

API Endpoints
1. Add Trade
Endpoint: POST /trade
Adds a new trade (either buy or sell) to the portfolio.

Request Body:
json
Copy
{
    "ticker": "TCS",
    "price": 1833.45,
    "quantity": 5,
    "type": "Buy"
}
ticker: The ticker symbol of the security.

price: The price at which the security was bought or sold.

quantity: The number of shares bought or sold.

type: The type of trade, either Buy or Sell.

Response:
json
Copy
{
    "message": "Trade Added"
}
2. Get Returns
Endpoint: GET /returns?price={currentPrice}
Calculates the cumulative returns of the portfolio based on the current price of the securities.

Query Parameters:
price: The current price of the security.

Example Request:
bash
Copy
curl --location 'http://localhost:8090/returns?price=120'
Response:
json
Copy
{
    "returns": 2000
}
3. Get Holdings
Endpoint: GET /holdings
Fetches the current holdings in the portfolio, including the ticker symbol, quantity, and average buy price of each security.

Example Request:
bash
Copy
curl http://localhost:8090/holdings
Response:
json
Copy
{
  "holdings": [
    {
      "ticket_symbol": "TCS",
      "quantity": 5,
      "average_buy_price": 1833.45
    },
    {
      "ticket_symbol": "WIPRO",
      "quantity": 10,
      "average_buy_price": 319.25
    }
  ]
}
Setup
Prerequisites
Go: Ensure that Go is installed. You can download Go from here.

Gin: This project uses the Gin web framework. You can install it by running the following command:

bash
Copy
go get -u github.com/gin-gonic/gin
Installing and Running
Clone the repository:

bash
Copy
git clone https://github.com/yourusername/portfolio-tracker.git
cd portfolio-tracker
Install dependencies:

bash
Copy
go mod tidy
Run the server:

bash
Copy
go run main.go
The API will be available at http://localhost:8090.