package artifactory

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/jfrog/terraform-provider-artifactory/v6/pkg/utils"
)

func resourceArtifactoryLocalCargoRepository() *schema.Resource {
	const packageType = "cargo"

	var cargoLocalSchema = utils.MergeSchema(baseLocalRepoSchema, map[string]*schema.Schema{
		"anonymous_access": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: `(Optional) Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option. Default value is 'false'.`,
		},
	}, repoLayoutRefSchema("local", packageType), compressionFormats)

	type CargoLocalRepo struct {
		LocalRepositoryBaseParams
		AnonymousAccess bool `json:"cargoAnonymousAccess"`
	}

	var unPackLocalCargoRepository = func(data *schema.ResourceData) (interface{}, string, error) {
		d := &utils.ResourceData{ResourceData: data}
		repo := CargoLocalRepo{
			LocalRepositoryBaseParams: unpackBaseRepo("local", data, packageType),
			AnonymousAccess:           d.GetBool("anonymous_access", false),
		}

		return repo, repo.Id(), nil
	}

	return mkResourceSchema(cargoLocalSchema, defaultPacker(cargoLocalSchema), unPackLocalCargoRepository, func() interface{} {
		return &CargoLocalRepo{
			LocalRepositoryBaseParams: LocalRepositoryBaseParams{
				PackageType: packageType,
				Rclass:      "local",
			},
		}
	})
}
