package httpapi

import (
	"context"
	"encoding/json"
	iface "github.com/ipfs/interface-go-ipfs-core"

	"github.com/ipfs/go-cid"
	caopts "github.com/ipfs/interface-go-ipfs-core/options"
	"github.com/ipfs/interface-go-ipfs-core/path"
	"github.com/pkg/errors"
)

type PinAPI HttpApi

type pinRefKeyObject struct {
	Type string
}

type pinRefKeyList struct {
	Keys map[string]pinRefKeyObject
}

type pin struct {
	path path.Resolved
	typ  string
	err  error
}

func (p *pin) Err() error {
	return p.err
}

func (p *pin) Path() path.Resolved {
	return p.path
}

func (p *pin) Type() string {
	return p.typ
}

func (api *PinAPI) Add(ctx context.Context, p path.Path, opts ...caopts.PinAddOption) error {
	options, err := caopts.PinAddOptions(opts...)
	if err != nil {
		return err
	}

	return api.core().Request("pin/add", p.String()).
		Option("recursive", options.Recursive).Exec(ctx, nil)
}

func (api *PinAPI) Ls(ctx context.Context, opts ...caopts.PinLsOption) (<-chan iface.Pin, error) {
	options, err := caopts.PinLsOptions(opts...)
	if err != nil {
		return nil, err
	}

	var out pinRefKeyList
	err = api.core().Request("pin/ls").
		Option("type", options.Type).Exec(ctx, &out)
	if err != nil {
		return nil, err
	}

	pins := make(chan iface.Pin)
	go func(ch chan<- iface.Pin) {
		for hash, p := range out.Keys {
			c, err := cid.Parse(hash)
			if err != nil {
				return
			}
			ch <- &pin{typ: p.Type, path: path.IpldPath(c)}
		}
		close(ch)
	}(pins)
	return pins, nil
}

// IsPinned returns whether or not the given cid is pinned
// and an explanation of why its pinned
func (api *PinAPI) IsPinned(ctx context.Context, p path.Path, opts ...caopts.PinIsPinnedOption) (string, bool, error) {
	options, err := caopts.PinIsPinnedOptions(opts...)
	if err != nil {
		return "", false, err
	}
	var out pinRefKeyList
	err = api.core().Request("pin/ls").
		Option("type", options.WithType).
		Option("arg", p.String()).
		Exec(ctx, &out)
	if err != nil {
		return "", false, err
	}

	for hash := range out.Keys {
		_, err := cid.Parse(hash)
		if err != nil {
			return "", false, err
		}
		return hash, true, nil
	}
	return "", false, errors.New("pin path not found")
}

func (api *PinAPI) Rm(ctx context.Context, p path.Path, opts ...caopts.PinRmOption) error {
	options, err := caopts.PinRmOptions(opts...)
	if err != nil {
		return err
	}

	return api.core().Request("pin/rm", p.String()).
		Option("recursive", options.Recursive).
		Exec(ctx, nil)
}

func (api *PinAPI) Update(ctx context.Context, from path.Path, to path.Path, opts ...caopts.PinUpdateOption) error {
	options, err := caopts.PinUpdateOptions(opts...)
	if err != nil {
		return err
	}

	return api.core().Request("pin/update", from.String(), to.String()).
		Option("unpin", options.Unpin).Exec(ctx, nil)
}

type pinVerifyRes struct {
	ok       bool
	badNodes []iface.BadPinNode
}

func (r *pinVerifyRes) Ok() bool {
	return r.ok
}

func (r *pinVerifyRes) BadNodes() []iface.BadPinNode {
	return r.badNodes
}

type badNode struct {
	err error
	cid cid.Cid
}

func (n *badNode) Path() path.Resolved {
	return path.IpldPath(n.cid)
}

func (n *badNode) Err() error {
	return n.err
}

func (api *PinAPI) Verify(ctx context.Context) (<-chan iface.PinStatus, error) {
	resp, err := api.core().Request("pin/verify").Option("verbose", true).Send(ctx)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	res := make(chan iface.PinStatus)

	go func() {
		defer resp.Close()
		defer close(res)
		dec := json.NewDecoder(resp.Output)
		for {
			var out struct {
				Cid string
				Ok  bool

				BadNodes []struct {
					Cid string
					Err string
				}
			}
			if err := dec.Decode(&out); err != nil {
				return // todo: handle non io.EOF somehow
			}

			badNodes := make([]iface.BadPinNode, len(out.BadNodes))
			for i, n := range out.BadNodes {
				c, err := cid.Decode(n.Cid)
				if err != nil {
					badNodes[i] = &badNode{
						cid: c,
						err: err,
					}
					continue
				}

				if n.Err != "" {
					err = errors.New(n.Err)
				}
				badNodes[i] = &badNode{
					cid: c,
					err: err,
				}
			}

			select {
			case res <- &pinVerifyRes{
				ok: out.Ok,

				badNodes: badNodes,
			}:
			case <-ctx.Done():
				return
			}
		}
	}()

	return res, nil
}

func (api *PinAPI) core() *HttpApi {
	return (*HttpApi)(api)
}
