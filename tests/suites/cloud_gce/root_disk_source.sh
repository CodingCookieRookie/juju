run_root_disk_source_default() {
	echo

	file="${TEST_DIR}/test-root-disk-source-default.log"

	ensure "test-root-disk-source-default" "${file}"

	# Deploy juju-qa-test without root-disk-source constraint — should default to pd-standard.
	juju deploy juju-qa-test --channel latest/edge
	wait_for_machine_agent_status "0" "started"

	instance_id="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["instance-id"]')"
	az="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["hardware"]' | tr ' ' '\n' | grep 'availability-zone' | cut -d= -f2)"

	boot_disk_type="$(gcloud compute disks describe "${instance_id}" --zone="${az}" --format="value(type.basename())")"
	if [ "${boot_disk_type}" != "pd-standard" ]; then
		echo "FAIL: expected default boot disk type pd-standard, got ${boot_disk_type}"
		return 1
	fi
	echo "OK: default boot disk type is pd-standard"

	destroy_model "test-root-disk-source-default"
}

run_root_disk_source_storage_pool() {
	echo

	file="${TEST_DIR}/test-root-disk-source-storage-pool.log"

	ensure "test-root-disk-source-storage-pool" "${file}"

	juju create-storage-pool ssd-gce gce disk-type=pd-ssd

	# Deploy juju-qa-test with root-disk-source=pd-ssd constraint - should be provisioned with pd-ssd.
	juju deploy juju-qa-test --channel latest/edge --constraints "root-disk-source=ssd-gce"
	wait_for_machine_agent_status "0" "started"

	# Verify the instance's boot disk is pd-ssd.
	instance_id="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["instance-id"]')"
	az="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["hardware"]' | tr ' ' '\n' | grep 'availability-zone' | cut -d= -f2)"

	boot_disk_type="$(gcloud compute disks describe "${instance_id}" --zone="${az}" --format="value(type.basename())")"
	if [ "${boot_disk_type}" != "pd-ssd" ]; then
		echo "FAIL: expected boot disk type pd-ssd, got ${boot_disk_type}"
		return 1
	fi
	echo "OK: boot disk type is pd-ssd"

	destroy_model "test-root-disk-source-storage-pool"
}

run_root_disk_source_storage_pool_named_local_ssd() {
	echo

	file="${TEST_DIR}/test-root-disk-source-storage-pool-named-local-ssd.log"

	ensure "test-root-disk-source-storage-pool-named-local-ssd" "${file}"

	juju create-storage-pool local-ssd gce disk-type=pd-ssd

	# Deploy juju-qa-test with root-disk-source=pd-ssd constraint - should be provisioned with pd-ssd.
	juju deploy juju-qa-test --channel latest/edge --constraints "root-disk-source=local-ssd"
	wait_for_machine_agent_status "0" "started"

	# Verify the instance's boot disk is pd-ssd.
	instance_id="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["instance-id"]')"
	az="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["hardware"]' | tr ' ' '\n' | grep 'availability-zone' | cut -d= -f2)"

	boot_disk_type="$(gcloud compute disks describe "${instance_id}" --zone="${az}" --format="value(type.basename())")"
	if [ "${boot_disk_type}" != "pd-ssd" ]; then
		echo "FAIL: expected boot disk type pd-ssd, got ${boot_disk_type}"
		return 1
	fi
	echo "OK: boot disk type is pd-ssd"

	destroy_model "test-root-disk-source-storage-pool-named-local-ssd"
}

run_root_disk_source_disk_type() {
	echo

	file="${TEST_DIR}/test-root-disk-source-disk-type.log"

	ensure "test-root-disk-source-disk-type" "${file}"

	# Deploy juju-qa-test with root-disk-source=pd-ssd constraint - should be provisioned with pd-ssd.
	juju deploy juju-qa-test --channel latest/edge --constraints "root-disk-source=pd-ssd"
	wait_for_machine_agent_status "0" "started"

	# Verify the instance's boot disk is pd-ssd.
	instance_id="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["instance-id"]')"
	az="$(juju show-machine 0 --format=yaml | yq -r '.machines["0"]["hardware"]' | tr ' ' '\n' | grep 'availability-zone' | cut -d= -f2)"

	boot_disk_type="$(gcloud compute disks describe "${instance_id}" --zone="${az}" --format="value(type.basename())")"
	if [ "${boot_disk_type}" != "pd-ssd" ]; then
		echo "FAIL: expected boot disk type pd-ssd, got ${boot_disk_type}"
		return 1
	fi
	echo "OK: boot disk type is pd-ssd"

	destroy_model "test-root-disk-source-disk-type"
}

run_root_disk_source_local() {
	echo

	file="${TEST_DIR}/test-root-disk-source-local.log"

	ensure "test-root-disk-source-local" "${file}"

	# Deploy juju-qa-test with local-ssd — deploy succeeds but machine provisioning should fail.
	juju deploy juju-qa-test --channel latest/edge --constraints "root-disk-source=local-ssd"

	echo "Waiting for status failure message indicating local-ssd is not valid..."
	if ( wait_for "local SSD disk storage not valid" '.machines["0"]["machine-status"]["message"]' ); then
		machine_msg=$(juju status --format=yaml | yq -r '.machines["0"]["machine-status"]["message"]')
		echo "OK: local-ssd correctly rejected with message: ${machine_msg}"
		destroy_model "test-root-disk-source-local"
		return 0
	else
		machine_msg=$(juju status --format=yaml | yq -r '.machines["0"]["machine-status"]["message"]')
		echo "FAIL: expected 'local SSD disk storage not valid' in machine status, got: ${machine_msg}"
		destroy_model "test-root-disk-source-local"
		return 1
	fi
}

run_root_disk_source_invalid() {
	echo

	file="${TEST_DIR}/test-root-disk-source-invalid.log"

	ensure "test-root-disk-source-invalid" "${file}"

	# Deploy juju-qa-test with an unknown disk type — deploy succeeds but machine provisioning should fail with "not valid".
	juju deploy juju-qa-test --channel latest/edge --constraints "root-disk-source=invalid-disk"

	echo "Waiting for status failure message indicating invalid disk type is not valid..."
	if ( wait_for 'invalid-disk' '.machines["0"]["machine-status"]["message"]' ); then
		machine_msg=$(juju status --format=yaml | yq -r '.machines["0"]["machine-status"]["message"]')
		echo "OK: unknown disk type correctly rejected with message: ${machine_msg}"
		destroy_model "test-root-disk-source-invalid"
		return 0
	else
		machine_msg=$(juju status --format=yaml | yq -r '.machines["0"]["machine-status"]["message"]')
		echo "FAIL: expected 'root disk source \"invalid-disk\" not valid' in machine status, got: ${machine_msg}"
		destroy_model "test-root-disk-source-invalid"
		return 1
	fi
}

test_root_disk_source() {
	if [ "$(skip 'test_root_disk_source')" ]; then
		echo "==> TEST SKIPPED: root disk source"
		return
	fi

	(
		set_verbosity

		cd .. || exit

		run "run_root_disk_source_default"
		run "run_root_disk_source_storage_pool"
		run "run_root_disk_source_disk_type"
		run "run_root_disk_source_local"
		run "run_root_disk_source_invalid"
	)
}
