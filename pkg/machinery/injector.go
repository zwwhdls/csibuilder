package machinery

import "csibuilder/pkg/model"

// injector is used to inject certain fields to file templates.
type injector struct {
	config   *model.Config
	resource *model.Resource
}

// injectInto injects fields from the universe into the builder
func (i injector) injectInto(builder Builder) {
	// Inject project configuration
	// Inject resource
	if i.resource != nil {
		if builderWithResource, hasResource := builder.(HasResource); hasResource {
			builderWithResource.InjectResource(i.resource)
		}
	}
	if builderWithRepository, hasRepository := builder.(HasRepository); hasRepository {
		builderWithRepository.InjectRepository(i.config.GetRepository())
	}
}
