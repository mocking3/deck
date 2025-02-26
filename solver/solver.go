package solver

import (
	"github.com/hbagdi/deck/crud"
	"github.com/hbagdi/deck/diff"
	cruds "github.com/hbagdi/deck/solver/kong"
	drycrud "github.com/hbagdi/deck/solver/kong/dry"
	"github.com/hbagdi/go-kong/kong"
	"github.com/pkg/errors"
)

// Solve generates a diff and walks the graph.
func Solve(doneCh chan struct{}, syncer *diff.Syncer,
	client *kong.Client, dry bool) []error {
	var r *crud.Registry
	var err error
	if dry {
		r, err = buildDryRegistry(client)
	} else {
		r, err = buildRegistry(client)
	}
	if err != nil {
		return append([]error{}, errors.Wrapf(err, "cannot build registry"))
	}

	return syncer.Run(doneCh, 10, func(e diff.Event) (crud.Arg, error) {
		return r.Do(e.Kind, e.Op, e)
	})
}

func buildDryRegistry(client *kong.Client) (*crud.Registry, error) {
	var r crud.Registry
	err := r.Register("service", &drycrud.ServiceCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'service' crud")
	}
	err = r.Register("route", &drycrud.RouteCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'route' crud")
	}
	err = r.Register("upstream", &drycrud.UpstreamCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'upstream' crud")
	}
	err = r.Register("target", &drycrud.TargetCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'target' crud")
	}
	err = r.Register("certificate", &drycrud.CertificateCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'certificate' crud")
	}
	err = r.Register("plugin", &drycrud.PluginCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'plugin' crud")
	}
	err = r.Register("consumer", &drycrud.ConsumerCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'consumer' crud")
	}
	err = r.Register("key-auth", &drycrud.KeyAuthCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'key-auth' crud")
	}
	err = r.Register("hmac-auth", &drycrud.HMACAuthCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'hmac-auth' crud")
	}
	err = r.Register("jwt-auth", &drycrud.JWTAuthCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'jwt-auth' crud")
	}
	err = r.Register("basic-auth", &drycrud.BasicAuthCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'basic-auth' crud")
	}
	err = r.Register("acl-group", &drycrud.ACLGroupCRUD{})
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'acl-group' crud")
	}
	return &r, nil
}

func buildRegistry(client *kong.Client) (*crud.Registry, error) {
	var r crud.Registry
	var err error
	service, err := cruds.NewServiceCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a service CRUD")
	}
	err = r.Register("service", service)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'service' crud")
	}
	route, err := cruds.NewRouteCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a route CRUD")
	}
	err = r.Register("route", route)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'route' crud")
	}
	upstream, err := cruds.NewUpstreamCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a upstream CRUD")
	}
	err = r.Register("upstream", upstream)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'upstream' crud")
	}
	target, err := cruds.NewTargetCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a target CRUD")
	}
	err = r.Register("target", target)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'target' crud")
	}
	certificate, err := cruds.NewCertificateCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a certificate CRUD")
	}
	err = r.Register("certificate", certificate)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'certificate' crud")
	}
	plugin, err := cruds.NewPluginCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a plugin CRUD")
	}
	err = r.Register("plugin", plugin)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'plugin' crud")
	}
	consumer, err := cruds.NewConsumerCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a 'consumer' CRUD")
	}
	err = r.Register("consumer", consumer)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'consumer' crud")
	}
	keyAuth, err := cruds.NewKeyAuthCRUD(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating a 'key-auth' CRUD")
	}
	err = r.Register("key-auth", keyAuth)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'key-auth' crud")
	}
	hmacAuth, err := cruds.NewHMACAuthCRUD(client)
	if err != nil {
		return nil, errors.Wrapf(err, "creating 'hmac-auth' crud")
	}
	err = r.Register("hmac-auth", hmacAuth)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'hmac-auth' crud")
	}
	jwtAuth, err := cruds.NewJWTAuthCRUD(client)
	if err != nil {
		return nil, errors.Wrapf(err, "creating 'jwt-auth' crud")
	}
	err = r.Register("jwt-auth", jwtAuth)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'jwt-auth' crud")
	}
	basicAuth, err := cruds.NewBasicAuthCRUD(client)
	if err != nil {
		return nil, errors.Wrapf(err, "creating 'basic-auth' crud")
	}
	err = r.Register("basic-auth", basicAuth)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'basic-auth' crud")
	}
	aclGroups, err := cruds.NewACLGroupCRUD(client)
	if err != nil {
		return nil, errors.Wrapf(err, "creating 'acl' crud")
	}
	err = r.Register("acl-group", aclGroups)
	if err != nil {
		return nil, errors.Wrapf(err, "registering 'acl-group' crud")
	}
	return &r, nil
}
