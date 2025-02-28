CREATE TABLE unit_resolve_kind (
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_unit_resolve_kind
ON unit_resolve_kind (name);

INSERT INTO unit_resolve_kind VALUES
(0, 'none'),
(1, 'retry-hooks'),
(2, 'no-hooks');

CREATE TABLE unit (
    uuid TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    life_id INT NOT NULL,
    application_uuid TEXT NOT NULL,
    net_node_uuid TEXT NOT NULL,
    -- Freshly created units will not have a charm URL set.
    charm_uuid TEXT,
    -- Resolve Kind starts out as None, is only set when a hook errors.
    resolve_kind_id INT NOT NULL DEFAULT 0,
    password_hash_algorithm_id TEXT,
    password_hash TEXT,
    CONSTRAINT fk_unit_life
    FOREIGN KEY (life_id)
    REFERENCES life (id),
    CONSTRAINT fk_unit_application
    FOREIGN KEY (application_uuid)
    REFERENCES application (uuid),
    CONSTRAINT fk_unit_net_node
    FOREIGN KEY (net_node_uuid)
    REFERENCES net_node (uuid),
    CONSTRAINT fk_unit_resolve_kind
    FOREIGN KEY (resolve_kind_id)
    REFERENCES unit_resolve_kind (id),
    CONSTRAINT fk_unit_charm
    FOREIGN KEY (charm_uuid)
    REFERENCES charm (uuid),
    CONSTRAINT fk_unit_password_hash_algorithm
    FOREIGN KEY (password_hash_algorithm_id)
    REFERENCES password_hash_algorithm (id)
);

CREATE UNIQUE INDEX idx_unit_name
ON unit (name);

CREATE INDEX idx_unit_application
ON unit (application_uuid);

CREATE INDEX idx_unit_net_node
ON unit (net_node_uuid);

-- unit_principal table is a table which is used to store the.
-- principal units for subordinate units.
CREATE TABLE unit_principal (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    principal_uuid TEXT NOT NULL,
    CONSTRAINT fk_unit_principal_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_unit_principal_principal
    FOREIGN KEY (principal_uuid)
    REFERENCES unit (uuid)
);

CREATE TABLE unit_agent (
    unit_uuid TEXT NOT NULL,
    url TEXT NOT NULL,
    version_major INT NOT NULL,
    version_minor INT NOT NULL,
    version_tag TEXT,
    version_patch INT NOT NULL,
    version_build INT,
    hash TEXT NOT NULL,
    hash_kind_id INT NOT NULL DEFAULT 0,
    binary_size INT NOT NULL,
    CONSTRAINT fk_unit_agent_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_unit_agent_hash_kind
    FOREIGN KEY (hash_kind_id)
    REFERENCES hash_kind (id),
    CONSTRAINT fk_object_store_metadata_path_unit
    FOREIGN KEY (url)
    REFERENCES object_store_metadata_path (path),
    PRIMARY KEY (unit_uuid, url)
);

CREATE TABLE unit_state (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    uniter_state TEXT,
    storage_state TEXT,
    secret_state TEXT,
    CONSTRAINT fk_unit_state_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid)
);

-- Local charm state stored upon hook commit with uniter state.
CREATE TABLE unit_state_charm (
    unit_uuid TEXT NOT NULL,
    "key" TEXT NOT NULL,
    value TEXT NOT NULL,
    PRIMARY KEY (unit_uuid, "key"),
    CONSTRAINT fk_unit_state_charm_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid)
);

-- Local relation state stored upon hook commit with uniter state.
CREATE TABLE unit_state_relation (
    unit_uuid TEXT NOT NULL,
    "key" TEXT NOT NULL,
    value TEXT NOT NULL,
    PRIMARY KEY (unit_uuid, "key"),
    CONSTRAINT fk_unit_state_relation_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid)
);

-- cloud containers belong to a k8s unit.
CREATE TABLE k8s_pod (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    -- provider_id comes from the provider, no FK.
    -- it represents the k8s pod UID.
    provider_id TEXT NOT NULL,
    CONSTRAINT fk_k8s_pod_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid)
);

CREATE UNIQUE INDEX idx_k8s_pod_provider
ON k8s_pod (provider_id);

CREATE TABLE k8s_pod_port (
    unit_uuid TEXT NOT NULL,
    port TEXT NOT NULL,
    CONSTRAINT fk_k8s_pod_port_k8s_pod
    FOREIGN KEY (unit_uuid)
    REFERENCES k8s_pod (unit_uuid),
    PRIMARY KEY (unit_uuid, port)
);

-- Status values for unit agents.
CREATE TABLE unit_agent_status_value (
    id INT PRIMARY KEY,
    status TEXT NOT NULL
);

INSERT INTO unit_agent_status_value VALUES
(0, 'allocating'),
(1, 'executing'),
(2, 'idle'),
(3, 'error'),
(4, 'failed'),
(5, 'lost'),
(6, 'rebooting');

-- Status values for cloud containers.
CREATE TABLE k8s_pod_status_value (
    id INT PRIMARY KEY,
    status TEXT NOT NULL
);

INSERT INTO k8s_pod_status_value VALUES
(0, 'waiting'),
(1, 'blocked'),
(2, 'running');

CREATE TABLE unit_agent_status (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    status_id INT NOT NULL,
    message TEXT,
    data TEXT,
    updated_at DATETIME,
    CONSTRAINT fk_unit_agent_status_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_unit_agent_status_status
    FOREIGN KEY (status_id)
    REFERENCES unit_agent_status_value (id)
);

CREATE TABLE unit_workload_status (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    status_id INT NOT NULL,
    message TEXT,
    data TEXT,
    updated_at DATETIME,
    CONSTRAINT fk_unit_workload_status_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_workload_status_value_status
    FOREIGN KEY (status_id)
    REFERENCES workload_status_value (id)
);

CREATE TABLE k8s_pod_status (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    status_id INT NOT NULL,
    message TEXT,
    data TEXT,
    updated_at DATETIME,
    CONSTRAINT fk_k8s_pod_status_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_k8s_pod_status_status
    FOREIGN KEY (status_id)
    REFERENCES k8s_pod_status_value (id)
);

CREATE TABLE unit_constraint (
    unit_uuid TEXT NOT NULL PRIMARY KEY,
    constraint_uuid TEXT,
    CONSTRAINT fk_unit_constraint_unit
    FOREIGN KEY (unit_uuid)
    REFERENCES unit (uuid),
    CONSTRAINT fk_unit_constraint_constraint
    FOREIGN KEY (constraint_uuid)
    REFERENCES "constraint" (uuid)
);

CREATE VIEW v_unit_constraint AS
SELECT
    uc.unit_uuid,
    c.arch,
    c.cpu_cores,
    c.cpu_power,
    c.mem,
    c.root_disk,
    c.root_disk_source,
    c.instance_role,
    c.instance_type,
    ctype.value AS container_type,
    c.virt_type,
    c.allocate_public_ip,
    c.image_id,
    ctag.tag,
    cspace.space AS space_name,
    cspace."exclude" AS space_exclude,
    czone.zone
FROM unit_constraint AS uc
JOIN "constraint" AS c ON ac.constraint_uuid = c.uuid
LEFT JOIN container_type AS ctype ON c.container_type_id = ctype.id
LEFT JOIN constraint_tag AS ctag ON c.uuid = ctag.constraint_uuid
LEFT JOIN constraint_space AS cspace ON c.uuid = cspace.constraint_uuid
LEFT JOIN constraint_zone AS czone ON c.uuid = czone.constraint_uuid;
