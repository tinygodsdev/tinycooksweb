package handler

import (
	"context"

	"github.com/jfyne/live"
)

type (
	PrivacyInstance struct {
		*CommonInstance
	}
)

// must be present in all instances
func (ins *PrivacyInstance) updateForLocale(ctx context.Context, s live.Socket, h *Handler) error {
	return nil
}

func (h *Handler) NewPrivacyInstance(s live.Socket) *PrivacyInstance {
	m, ok := s.Assigns().(*PrivacyInstance)
	if !ok {
		return &PrivacyInstance{
			CommonInstance: h.NewCommon(s, viewPrivacy),
		}
	}

	return m
}

func (h *Handler) Privacy() live.Handler {
	t := h.template("base.layout.html", "privacy")

	lvh := live.NewHandler(live.WithTemplateRenderer(t))
	// COMMON BLOCK START
	// this logic must be present in all handlers
	{
		constructor := h.NewPrivacyInstance // NB: make sure constructor is correct
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
		instance := h.NewPrivacyInstance(s)
		instance.fromContext(ctx)
		return instance, nil
	})

	return lvh
}
