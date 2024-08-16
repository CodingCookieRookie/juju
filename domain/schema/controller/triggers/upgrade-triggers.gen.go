// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)

// ChangeLogTriggersForUpgradeInfo generates the triggers for the
// upgrade_info table.
func ChangeLogTriggersForUpgradeInfo(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for UpgradeInfo
CREATE TRIGGER trg_log_upgrade_info_insert
AFTER INSERT ON upgrade_info FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for UpgradeInfo
CREATE TRIGGER trg_log_upgrade_info_update
AFTER UPDATE ON upgrade_info FOR EACH ROW
WHEN 
	NEW.previous_version != OLD.previous_version OR
	NEW.target_version != OLD.target_version OR
	NEW.state_type_id != OLD.state_type_id 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for UpgradeInfo
CREATE TRIGGER trg_log_upgrade_info_delete
AFTER DELETE ON upgrade_info FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}

// ChangeLogTriggersForUpgradeInfoControllerNode generates the triggers for the
// upgrade_info_controller_node table.
func ChangeLogTriggersForUpgradeInfoControllerNode(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for UpgradeInfoControllerNode
CREATE TRIGGER trg_log_upgrade_info_controller_node_insert
AFTER INSERT ON upgrade_info_controller_node FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for UpgradeInfoControllerNode
CREATE TRIGGER trg_log_upgrade_info_controller_node_update
AFTER UPDATE ON upgrade_info_controller_node FOR EACH ROW
WHEN 
	NEW.controller_node_id != OLD.controller_node_id OR
	NEW.upgrade_info_uuid != OLD.upgrade_info_uuid OR
	(NEW.node_upgrade_completed_at != OLD.node_upgrade_completed_at OR (NEW.node_upgrade_completed_at IS NOT NULL AND OLD.node_upgrade_completed_at IS NULL) OR (NEW.node_upgrade_completed_at IS NULL AND OLD.node_upgrade_completed_at IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for UpgradeInfoControllerNode
CREATE TRIGGER trg_log_upgrade_info_controller_node_delete
AFTER DELETE ON upgrade_info_controller_node FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
