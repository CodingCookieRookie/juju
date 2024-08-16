// Code generated by triggergen. DO NOT EDIT.

package triggers

import (
	"fmt"

	"github.com/juju/juju/core/database/schema"
)

// ChangeLogTriggersForModelConfig generates the triggers for the
// model_config table.
func ChangeLogTriggersForModelConfig(columnName string, namespaceID int) func() schema.Patch {
	return func() schema.Patch {
		return schema.MakePatch(fmt.Sprintf(`
-- insert trigger for ModelConfig
CREATE TRIGGER trg_log_model_config_insert
AFTER INSERT ON model_config FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (1, %[2]d, NEW.%[1]s, DATETIME('now'));
END;

-- update trigger for ModelConfig
CREATE TRIGGER trg_log_model_config_update
AFTER UPDATE ON model_config FOR EACH ROW
WHEN 
	NEW.value != OLD.value 
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (2, %[2]d, OLD.%[1]s, DATETIME('now'));
END;

-- delete trigger for ModelConfig
CREATE TRIGGER trg_log_model_config_delete
AFTER DELETE ON model_config FOR EACH ROW
BEGIN
    INSERT INTO change_log (edit_type_id, namespace_id, changed, created_at)
    VALUES (4, %[2]d, OLD.%[1]s, DATETIME('now'));
END;`, columnName, namespaceID))
	}
}
