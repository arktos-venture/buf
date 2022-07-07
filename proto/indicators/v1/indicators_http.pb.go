// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.1

package v1Indicators

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type IndicatorsHTTPServer interface {
	ADX(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	ATR(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	AVGPRICE(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	BBANDS(context.Context, *IndicatorRequest) (*IndicatorBBandsReply, error)
	CCI(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	DMI(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	EMA(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	LINEARREG_SLOPE(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	MACD(context.Context, *IndicatorRequest) (*IndicatorMacdReply, error)
	RSI(context.Context, *IndicatorRequest) (*IndicatorRsiReply, error)
	SAR(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	SMA(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	STDDEV(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	VAR(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	WILLR(context.Context, *IndicatorRequest) (*IndicatorReply, error)
	WMA(context.Context, *IndicatorRequest) (*IndicatorReply, error)
}

func RegisterIndicatorsHTTPServer(s *http.Server, srv IndicatorsHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/quotes/indicator/avgprice", _Indicators_AVGPRICE0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/sma", _Indicators_SMA0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/ema", _Indicators_EMA0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/wma", _Indicators_WMA0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/volatility", _Indicators_VAR0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/rsi", _Indicators_RSI0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/stddev", _Indicators_STDDEV0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/slope", _Indicators_LINEARREG_SLOPE0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/dmi", _Indicators_DMI0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/adx", _Indicators_ADX0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/macd", _Indicators_MACD0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/atr", _Indicators_ATR0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/cci", _Indicators_CCI0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/bbands", _Indicators_BBANDS0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/willr", _Indicators_WILLR0_HTTP_Handler(srv))
	r.POST("/v1/quotes/indicator/sar", _Indicators_SAR0_HTTP_Handler(srv))
}

func _Indicators_AVGPRICE0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/AVGPRICE")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AVGPRICE(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_SMA0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/SMA")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SMA(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_EMA0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/EMA")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EMA(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_WMA0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/WMA")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WMA(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_VAR0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/VAR")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.VAR(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_RSI0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/RSI")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RSI(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorRsiReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_STDDEV0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/STDDEV")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.STDDEV(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_LINEARREG_SLOPE0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/LINEARREG_SLOPE")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LINEARREG_SLOPE(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_DMI0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/DMI")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DMI(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_ADX0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/ADX")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ADX(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_MACD0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/MACD")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MACD(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorMacdReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_ATR0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/ATR")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ATR(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_CCI0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/CCI")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CCI(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_BBANDS0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/BBANDS")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BBANDS(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorBBandsReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_WILLR0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/WILLR")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.WILLR(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

func _Indicators_SAR0_HTTP_Handler(srv IndicatorsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IndicatorRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/indicators.v1.Indicators/SAR")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SAR(ctx, req.(*IndicatorRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IndicatorReply)
		return ctx.Result(200, reply)
	}
}

type IndicatorsHTTPClient interface {
	ADX(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	ATR(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	AVGPRICE(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	BBANDS(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorBBandsReply, err error)
	CCI(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	DMI(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	EMA(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	LINEARREG_SLOPE(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	MACD(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorMacdReply, err error)
	RSI(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorRsiReply, err error)
	SAR(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	SMA(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	STDDEV(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	VAR(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	WILLR(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
	WMA(ctx context.Context, req *IndicatorRequest, opts ...http.CallOption) (rsp *IndicatorReply, err error)
}

type IndicatorsHTTPClientImpl struct {
	cc *http.Client
}

func NewIndicatorsHTTPClient(client *http.Client) IndicatorsHTTPClient {
	return &IndicatorsHTTPClientImpl{client}
}

func (c *IndicatorsHTTPClientImpl) ADX(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/adx"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/ADX"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) ATR(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/atr"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/ATR"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) AVGPRICE(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/avgprice"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/AVGPRICE"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) BBANDS(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorBBandsReply, error) {
	var out IndicatorBBandsReply
	pattern := "/v1/quotes/indicator/bbands"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/BBANDS"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) CCI(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/cci"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/CCI"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) DMI(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/dmi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/DMI"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) EMA(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/ema"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/EMA"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) LINEARREG_SLOPE(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/slope"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/LINEARREG_SLOPE"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) MACD(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorMacdReply, error) {
	var out IndicatorMacdReply
	pattern := "/v1/quotes/indicator/macd"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/MACD"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) RSI(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorRsiReply, error) {
	var out IndicatorRsiReply
	pattern := "/v1/quotes/indicator/rsi"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/RSI"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) SAR(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/sar"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/SAR"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) SMA(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/sma"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/SMA"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) STDDEV(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/stddev"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/STDDEV"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) VAR(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/volatility"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/VAR"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) WILLR(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/willr"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/WILLR"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *IndicatorsHTTPClientImpl) WMA(ctx context.Context, in *IndicatorRequest, opts ...http.CallOption) (*IndicatorReply, error) {
	var out IndicatorReply
	pattern := "/v1/quotes/indicator/wma"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/indicators.v1.Indicators/WMA"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
