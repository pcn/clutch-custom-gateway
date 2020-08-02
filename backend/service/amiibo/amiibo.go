package amiibo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"

	"github.com/lyft/clutch/backend/service"
	amiibov1 "github.com/pcn/clutch-custom-gateway/backend/api/amiibo/v1"
)

const Name = "clutch.service.amiibo"

const apiHost = "https://www.amiiboapi.com"

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	return &client{http: &http.Client{}}, nil
}

type Client interface {
	GetAmiibo(ctx context.Context, name string) ([]*amiibov1.Amiibo, error)
}

type client struct {
	http *http.Client
}

type RawAmiibo struct {
	Character    string `json:"character"`
	AmiiboSeries string `json:"amiiboSeries"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Type         string `json:"type"`
}

func (r RawAmiibo) toProto() *amiibov1.Amiibo {
	t := strings.ToUpper(r.Type)
	return &amiibov1.Amiibo{
		Name:         r.Name,
		AmiiboSeries: r.AmiiboSeries,
		ImageUrl:     r.Image,
		Character:    r.Character,
		Type:         amiibov1.Amiibo_Type(amiibov1.Amiibo_type_value[t]),
	}
}

func charactersFromJSON(data []byte) ([]*amiibov1.Amiibo, error) {
	raw := &RawResponse{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}
	ret := make([]*amiibov1.Amiibo, len(raw.Amiibo))
	for i, a := range raw.Amiibo {
		ret[i] = a.toProto()
	}
	return ret, nil
}

func (c *client) GetAmiibo(ctx context.Context, name string) ([]*amiibov1.Amiibo, error) {
	url := fmt.Sprintf("%s/api/amiibo?character=%s", apiHost, name)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, status.Error(service.CodeFromHTTPStatus(resp.StatusCode), string(body))
	}
	return charactersFromJSON(body)
}
