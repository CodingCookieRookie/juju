// Copyright 2023 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package schema

import (
	"embed"
	"fmt"
	"sort"

	"github.com/juju/juju/core/database/schema"
	"github.com/juju/juju/domain/schema/model/triggers"
)

//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/storage-triggers.gen.go -package=triggers -tables=block_device,storage_attachment,storage_filesystem,storage_filesystem_attachment,storage_volume,storage_volume_attachment,storage_volume_attachment_plan
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/model-triggers.gen.go -package=triggers -tables=model_config
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/objectstore-triggers.gen.go -package=triggers -tables=object_store_metadata_path
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/secret-triggers.gen.go -package=triggers -tables=secret_metadata,secret_rotation,secret_revision,secret_revision_expire,secret_revision_obsolete,secret_revision,secret_reference,secret_deleted_value_ref
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/network-triggers.gen.go -package=triggers -tables=subnet
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/machine-triggers.gen.go -package=triggers -tables=machine,machine_lxd_profile
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/machine-cloud-instance-triggers.gen.go -package=triggers -tables=machine_cloud_instance
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/machine-requires-reboot-triggers.gen.go -package=triggers -tables=machine_requires_reboot
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/application-triggers.gen.go -package=triggers -tables=application,application_config_hash,charm,unit,application_scale,port_range
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/relation-triggers.gen.go -package=triggers -tables=relation_application_setting
//go:generate go run ./../../generate/triggergen -db=model -destination=./model/triggers/cleanup-triggers.gen.go -package=triggers -tables=removal

//go:embed model/sql/*.sql
var modelSchemaDir embed.FS

const (
	tableModelConfig tableNamespaceID = iota
	tableModelObjectStoreMetadata
	tableBlockDeviceMachine
	tableStorageAttachment
	tableFileSystem
	tableFileSystemAttachment
	tableVolume
	tableVolumeAttachment
	tableVolumeAttachmentPlan
	tableSecretMetadataAutoPrune
	tableSecretRotation
	tableSecretRevisionObsolete
	tableSecretRevisionExpire
	tableSecretRevision
	tableSecretReference
	tableSubnet
	tableMachine
	tableMachineLxdProfile
	tableMachineCloudInstance
	tableMachineRequireReboot
	tableCharm
	tableUnit
	tableApplicationScale
	tablePortRange
	tableSecretDeletedValueRef
	tableApplication
	tableRemoval
	tableApplicationConfigHash
	tableAgentVersion
	tableRelationApplicationSetting
)

// ModelDDL is used to create model databases.
func ModelDDL() *schema.Schema {
	entries, err := modelSchemaDir.ReadDir("model/sql")
	if err != nil {
		panic(err)
	}

	var names []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		names = append(names, entry.Name())
	}

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	patches := make([]func() schema.Patch, len(names))
	for i, name := range names {
		data, err := modelSchemaDir.ReadFile(fmt.Sprintf("model/sql/%s", name))
		if err != nil {
			panic(err)
		}

		patches[i] = func() schema.Patch {
			return schema.MakePatch(string(data))
		}
	}

	// Changestream triggers.
	patches = append(patches,
		triggers.ChangeLogTriggersForBlockDevice("machine_uuid", tableBlockDeviceMachine),
		triggers.ChangeLogTriggersForModelConfig("key", tableModelConfig),
		triggers.ChangeLogTriggersForObjectStoreMetadataPath("path", tableModelObjectStoreMetadata),
		triggers.ChangeLogTriggersForStorageAttachment("storage_instance_uuid", tableStorageAttachment),
		triggers.ChangeLogTriggersForStorageFilesystem("uuid", tableFileSystem),
		triggers.ChangeLogTriggersForStorageFilesystemAttachment("uuid", tableFileSystemAttachment),
		triggers.ChangeLogTriggersForStorageVolume("uuid", tableVolume),
		triggers.ChangeLogTriggersForStorageVolumeAttachment("uuid", tableVolumeAttachment),
		triggers.ChangeLogTriggersForStorageVolumeAttachmentPlan("uuid", tableVolumeAttachmentPlan),
		triggers.ChangeLogTriggersForSecretMetadata("secret_id", tableSecretMetadataAutoPrune),
		triggers.ChangeLogTriggersForSecretRotation("secret_id", tableSecretRotation),
		triggers.ChangeLogTriggersForSecretRevisionObsolete("revision_uuid", tableSecretRevisionObsolete),
		triggers.ChangeLogTriggersForSecretRevisionExpire("revision_uuid", tableSecretRevisionExpire),
		triggers.ChangeLogTriggersForSecretRevision("uuid", tableSecretRevision),
		triggers.ChangeLogTriggersForSecretReference("secret_id", tableSecretReference),
		triggers.ChangeLogTriggersForSubnet("uuid", tableSubnet),
		triggers.ChangeLogTriggersForMachine("uuid", tableMachine),
		triggers.ChangeLogTriggersForMachineLxdProfile("machine_uuid", tableMachineLxdProfile),
		triggers.ChangeLogTriggersForMachineCloudInstance("machine_uuid", tableMachineCloudInstance),
		triggers.ChangeLogTriggersForMachineRequiresReboot("machine_uuid", tableMachineRequireReboot),
		triggers.ChangeLogTriggersForCharm("uuid", tableCharm),
		triggers.ChangeLogTriggersForUnit("uuid", tableUnit),
		triggers.ChangeLogTriggersForApplicationScale("application_uuid", tableApplicationScale),
		triggers.ChangeLogTriggersForPortRange("unit_uuid", tablePortRange),
		triggers.ChangeLogTriggersForSecretDeletedValueRef("revision_uuid", tableSecretDeletedValueRef),
		triggers.ChangeLogTriggersForApplication("uuid", tableApplication),
		triggers.ChangeLogTriggersForRemoval("uuid", tableRemoval),
		triggers.ChangeLogTriggersForApplicationConfigHash("application_uuid", tableApplicationConfigHash),
		triggers.ChangeLogTriggersForRelationApplicationSetting("relation_endpoint_uuid",
			tableRelationApplicationSetting),
	)

	// Generic triggers.
	patches = append(patches,
		triggersForImmutableTable("model", "", "model table is immutable, only insertions are allowed"),

		// The charm is unmodifiable.
		// There is a lot of assumptions in the code that the charm is immutable
		// from modification. This is a safety net to ensure that the charm is
		// not modified.
		triggersForUnmodifiableTable("charm_action", "charm_action table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_config", "charm_config table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_container_mount", "charm_container_mount table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_container", "charm_container table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_device", "charm_device table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_extra_binding", "charm_extra_binding table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_hash", "charm_hash table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_manifest_base", "charm_manifest base table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_metadata", "charm_metadata table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_relation", "charm_relation table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_resource", "charm_resource table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_storage", "charm_storage table is unmodifiable, only insertions and deletions are allowed"),
		triggersForUnmodifiableTable("charm_term", "charm_term table is unmodifiable, only insertions and deletions are allowed"),

		// Secret permissions do not allow subject or scope to be updated.
		triggerGuardForTable("secret_permission",
			"OLD.subject_type_id <> NEW.subject_type_id OR OLD.scope_uuid <> NEW.scope_uuid OR OLD.scope_type_id <> NEW.scope_type_id",
			"secret permission subjects and scopes must be identical",
		),

		triggerGuardForTable("sequence_charm_local",
			"OLD.reference_name = NEW.reference_name AND NEW.sequence <= OLD.sequence",
			"sequence number must monotonically increase",
		),
	)

	// For agent_version we only care if the single row is updated.
	// We emit the new target agent version.
	patches = append(patches, func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
CREATE TRIGGER trg_log_agent_version_update
AFTER UPDATE ON agent_version FOR EACH ROW
WHEN 
	NEW.target_version != OLD.target_version
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %d, NEW.target_version, DATETIME('now'));
END;`, tableAgentVersion))
	})

	modelSchema := schema.New()
	for _, fn := range patches {
		modelSchema.Add(fn())
	}
	return modelSchema
}
