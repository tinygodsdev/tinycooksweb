package handler

import (
	"context"
	"html/template"

	"github.com/jfyne/live"
)

type (
	NotFoundInstance struct {
		*CommonInstance
	}
)

func (ins *NotFoundInstance) withError(err error) *NotFoundInstance {
	ins.Error = err
	return ins
}

// must be present in all instances
func (ins *NotFoundInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewNotFoundInstance(s live.Socket) *NotFoundInstance {
	m, ok := s.Assigns().(*NotFoundInstance)
	if !ok {
		return &NotFoundInstance{
			CommonInstance: h.NewCommon(s, view404),
		}
	}

	return m
}

func (h *Handler) NotFound() live.Handler {
	t := template.Must(template.New("base.layout.html").Funcs(funcMap).ParseFiles(
		h.t+"base.layout.html",
		h.t+"page.404.html",
	))

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewNotFoundInstance // NB: make sure constructor is correct
		// SAFE TO COPY

		// update locale logic
		lvh.HandleParams(func(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
			instance := constructor(s)
			instance.SetLocale(p.String(paramLocale))
			err := instance.updateForLocale(ctx, s, h)
			if err != nil {
				return nil, err
			}
			return instance, nil
		})

		lvh.HandleError(func(ctx context.Context, err error) {
			h.HandleError(ctx, err)
		})
		// SAFE TO COPY END
	}
	// COMMON BLOCK END

	lvh.HandleMount(func(ctx context.Context, s live.Socket) (interface{}, error) {
		instance := h.NewNotFoundInstance(s)
		instance.fromContext(ctx)
		return instance, nil
	})

	return lvh
}
