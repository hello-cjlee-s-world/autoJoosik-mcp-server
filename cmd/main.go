package main

import (
	"autoJoosik-mcp-server/internal/stockapi"
	"autoJoosik-mcp-server/pkg/logger"
	//"autoJoosik-mcp-server/pkg/properties"
	"context"
	//"github.com/alexflint/go-arg"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"log"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

type GetStockInput struct {
}

type GetStockOut struct {
	Data    string `json:"data" jsonschema:"데이터"`
	Mapping string `json:"mapping" jsonschema:"컬럼명"`
}

type BuyStockInput struct {
	StkCd string `json:"StkCd" jsonschema:"종목 코드"`
	Qty   int    `json:"Qty" jsonschema:"수량"`
}

type BuyStockOut struct {
	Message string `json:"message"`
	StkCd   string `json:"stkCd"`
	Qty     string `json:"qty"`
	Error   string `json:"error"`
}

type SellStockInput struct {
	StkCd string `json:"StkCd" jsonschema:"종목 코드"`
	Qty   int    `json:"Qty" jsonschema:"수량"`
}

type SellStockOut struct {
	Message string `json:"message"`
	StkCd   string `json:"stkCd"`
	Qty     string `json:"qty"`
	Error   string `json:"error"`
}

func GetStock(ctx context.Context, req *mcp.CallToolRequest, input GetStockInput) (
	*mcp.CallToolResult,
	GetStockOut,
	error,
) {
	result, err := stockapi.CallStockInfoAPI()
	if err != nil {
		return nil, GetStockOut{}, err
	}

	output := GetStockOut{
		Data:    result.Data,
		Mapping: result.Mapping,
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: result.Data,
			},
		},
	}, output, nil
}

func BuyStock(ctx context.Context, req *mcp.CallToolRequest, input BuyStockInput) (
	*mcp.CallToolResult,
	BuyStockOut,
	error,
) {
	// 에이전트가 보낸 값
	stkCd := input.StkCd
	Qty := input.Qty

	result, err := stockapi.CallStockBuyAPI(stkCd, Qty)
	if err != nil {
		return nil, BuyStockOut{}, err
	}

	output := BuyStockOut{
		Message: result.Message,
		StkCd:   result.StkCd,
		Qty:     result.Qty,
		Error:   result.Error,
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: result.Message,
			},
		},
	}, output, nil
}

func SellStock(ctx context.Context, req *mcp.CallToolRequest, input SellStockInput) (
	*mcp.CallToolResult,
	SellStockOut,
	error,
) {
	// 에이전트가 보낸 값
	stkCd := input.StkCd
	Qty := input.Qty

	result, err := stockapi.CallStockSellAPI(stkCd, Qty)
	if err != nil {
		return nil, SellStockOut{}, err
	}

	output := SellStockOut{
		Message: result.Message,
		StkCd:   result.StkCd,
		Qty:     result.Qty,
		Error:   result.Error,
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: result.Message,
			},
		},
	}, output, nil
}
func SayHi(ctx context.Context, req *mcp.CallToolRequest, input Input) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	output := Output{
		Greeting: "안녕",
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "..."},
		},
	}, output, nil
}

type Args struct {
	Config string `arg:"-c, --config" help:"Configuration file"`
}

//var props *properties.PropertiesInfo

func main() {
	//// 설정 불러오기
	//var args Args
	//arg.MustParse(&args)
	//props = properties.GetInstance()
	//
	//if args.Config == "" {
	//	props.Init("internal/config/autoJoosik_market_data_fetcher_conf.yml")
	//} else {
	//	props.Init(args.Config)
	//}

	// logger 초기화
	logger.LoggerInit(logger.LoggerConfig{
		Level:         "debug",
		Filename:      "./logs/app.log",
		MaxSize:       10,
		MaxBackups:    10,
		MaxAge:        10,
		Compress:      false,
		ConsoleOutput: true,
	})
	logger.Info("server info") //"port", props.Server.Port,

	// Create a server with a single tool.
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "인사"}, SayHi)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "get_stock",
		Description: "내가 관심 있는 주식 데이터를 조회한다.",
	}, GetStock)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "buy_stock",
		Description: "주식 종목 코드와 수량을 받아 매수 요청을 보낸다.",
	}, BuyStock)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "sell_stock",
		Description: "주식 종목 코드와 수량을 받아 매매 요청을 보낸다.",
	}, SellStock)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
