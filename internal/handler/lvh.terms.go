package handler

import (
	"context"

	"github.com/jfyne/live"
)

type (
	TermsInstance struct {
		*CommonInstance
	}
)

// must be present in all instances
func (ins *TermsInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewTermsInstance(s live.Socket) *TermsInstance {
	m, ok := s.Assigns().(*TermsInstance)
	if !ok {
		return &TermsInstance{
			CommonInstance: h.NewCommon(s, viewTerms),
		}
	}

	return m
}

func (h *Handler) Terms() live.Handler {
	t := h.template("base.layout.html", "page.terms.html")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewTermsInstance // NB: make sure constructor is correct
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
		instance := h.NewTermsInstance(s)
		instance.fromContext(ctx)
		return instance, nil
	})

	return lvh
}
