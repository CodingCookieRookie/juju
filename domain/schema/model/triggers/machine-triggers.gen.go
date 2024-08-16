// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)

// ChangeLogTriggersForMachine generates the triggers for the
// machine table.
func ChangeLogTriggersForMachine(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for Machine
CREATE TRIGGER trg_log_machine_insert
AFTER INSERT ON machine FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for Machine
CREATE TRIGGER trg_log_machine_update
AFTER UPDATE ON machine FOR EACH ROW
WHEN 
	NEW.name != OLD.name OR
	NEW.net_node_uuid != OLD.net_node_uuid OR
	NEW.life_id != OLD.life_id OR
	(NEW.base != OLD.base OR (NEW.base IS NOT NULL AND OLD.base IS NULL) OR (NEW.base IS NULL AND OLD.base IS NOT NULL)) OR
	(NEW.nonce != OLD.nonce OR (NEW.nonce IS NOT NULL AND OLD.nonce IS NULL) OR (NEW.nonce IS NULL AND OLD.nonce IS NOT NULL)) OR
	(NEW.password_hash_algorithm_id != OLD.password_hash_algorithm_id OR (NEW.password_hash_algorithm_id IS NOT NULL AND OLD.password_hash_algorithm_id IS NULL) OR (NEW.password_hash_algorithm_id IS NULL AND OLD.password_hash_algorithm_id IS NOT NULL)) OR
	(NEW.password_hash != OLD.password_hash OR (NEW.password_hash IS NOT NULL AND OLD.password_hash IS NULL) OR (NEW.password_hash IS NULL AND OLD.password_hash IS NOT NULL)) OR
	(NEW.clean != OLD.clean OR (NEW.clean IS NOT NULL AND OLD.clean IS NULL) OR (NEW.clean IS NULL AND OLD.clean IS NOT NULL)) OR
	(NEW.force_destroyed != OLD.force_destroyed OR (NEW.force_destroyed IS NOT NULL AND OLD.force_destroyed IS NULL) OR (NEW.force_destroyed IS NULL AND OLD.force_destroyed IS NOT NULL)) OR
	(NEW.placement != OLD.placement OR (NEW.placement IS NOT NULL AND OLD.placement IS NULL) OR (NEW.placement IS NULL AND OLD.placement IS NOT NULL)) OR
	(NEW.agent_started_at != OLD.agent_started_at OR (NEW.agent_started_at IS NOT NULL AND OLD.agent_started_at IS NULL) OR (NEW.agent_started_at IS NULL AND OLD.agent_started_at IS NOT NULL)) OR
	(NEW.hostname != OLD.hostname OR (NEW.hostname IS NOT NULL AND OLD.hostname IS NULL) OR (NEW.hostname IS NULL AND OLD.hostname IS NOT NULL)) OR
	(NEW.is_controller != OLD.is_controller OR (NEW.is_controller IS NOT NULL AND OLD.is_controller IS NULL) OR (NEW.is_controller IS NULL AND OLD.is_controller IS NOT NULL)) 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for Machine
CREATE TRIGGER trg_log_machine_delete
AFTER DELETE ON machine FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
