package repository

import (
	"context"

	"github.com/shutterbase/shutterbase/ent"
	"github.com/shutterbase/shutterbase/ent/imagetag"
	"github.com/shutterbase/shutterbase/ent/project"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func GetImageTags(ctx context.Context, projectId uuid.UUID, paginationParameters *PaginationParameters) ([]*ent.ImageTag, int, error) {

	conditions :=
		imagetag.And(
			imagetag.HasProjectWith(project.ID(projectId)),
			imagetag.Or(
				imagetag.DescriptionContains(paginationParameters.Search),
				imagetag.NameContains(paginationParameters.Search),
			),
		)

	items, err := databaseClient.ImageTag.Query().
		Limit(paginationParameters.Limit).
		Offset(paginationParameters.Offset).
		Where(conditions).
		All(ctx)
	if err != nil {
		log.Info().Err(err).Msg("Error getting image tags")
		return nil, 0, err
	}

	count, err := databaseClient.ImageTag.Query().Where(conditions).Count(ctx)
	if err != nil {
		return nil, 0, err
	}

	return items, count, err
}

func GetImageTag(ctx context.Context, id uuid.UUID) (*ent.ImageTag, error) {
	item, err := databaseClient.ImageTag.Query().Where(imagetag.ID(id)).WithCreatedBy().WithModifiedBy().Only(ctx)
	if err != nil {
		log.Info().Err(err).Msg("Error getting image tag")
	}
	return item, err
}

func DeleteImageTag(ctx context.Context, id uuid.UUID) error {
	err := databaseClient.ImageTag.DeleteOneID(id).Exec(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error deleting image tag")
	}
	return err
}