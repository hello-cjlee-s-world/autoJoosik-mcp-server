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
	Symbol string `json:"symbol" jsonschema:"종목 코드"`
	Amount string `json:"amount" jsonschema:"수량"`
}

type BuyResponse struct {
	Success string `json:"success"`
	Message string `json:"message"`
}

func GetStock(ctx context.Context, req *mcp.CallToolRequest, input GetStockInput) (
	*mcp.CallToolResult,
	GetStockOut,
	error,
) {
	result, err := stockapi.CallStockAPI()
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
		Description: "주식 데이터 조회",
	}, GetStock)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
