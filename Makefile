dev_server_start:
	go build
	./devops-ui -server -consul_address=$(consul_address) \
		-database_user=$(database_user) -database_paswd=$(database_paswd) \
		-database_name=$(database_name) -database_debug=$(database_debug) \
		-database_address=$(database_address) \
		-gitlab_token=$(gitlab_token) -gitlab_base_url=$(gitlab_base_url) \
		-ucloud_publick_key=$(ucloud_publick_key) -ucloud_private_key=$(ucloud_private_key) \
		-log_level=$(log_level) \
		-upyun_bucket=$(upyun_bucket) -upyun_op=$(upyun_op) -upyun_passwd=$(upyun_passwd) -upyun_prefix=$(upyun_prefix) \
		-http_server_port=$(http_server_port) -log_server_port=$(log_server_port)
dev_worker_start:
	go build
	./devops-ui -worker -consul_address=$(consul_address) \
		-database_user=$(database_user) -database_paswd=$(database_paswd) \
		-database_name=$(database_name) -database_debug=$(database_debug) \
		-database_address=$(database_address) \
		-gitlab_token=$(gitlab_token) -gitlab_base_url=$(gitlab_base_url) \
		-ucloud_publick_key=$(ucloud_publick_key) -ucloud_private_key=$(ucloud_private_key) \
		-log_level=$(log_level) \
		-upyun_bucket=$(upyun_bucket) -upyun_op=$(upyun_op) -upyun_passwd=$(upyun_passwd) -upyun_prefix=$(upyun_prefix) \
		-http_server_port=$(http_server_port) -log_server_port=$(log_server_port)
server_start:
	./devops-ui -server -consul_address=$(consul_address) \
		-database_user=$(database_user) -database_paswd=$(database_paswd) \
		-database_name=$(database_name) -database_debug=$(database_debug) \
		-database_address=$(database_address) \
		-gitlab_token=$(gitlab_token) -gitlab_base_url=$(gitlab_base_url) \
		-ucloud_publick_key=$(ucloud_publick_key) -ucloud_private_key=$(ucloud_private_key) \
		-log_level=$(log_level) \
		-upyun_bucket=$(upyun_bucket) -upyun_op=$(upyun_op) -upyun_passwd=$(upyun_passwd) -upyun_prefix=$(upyun_prefix) \
		-http_server_port=$(http_server_port) -log_server_port=$(log_server_port)
worker_start:
	./devops-ui -worker -consul_address=$(consul_address) \
		-database_user=$(database_user) -database_paswd=$(database_paswd) \
		-database_name=$(database_name) -database_debug=$(database_debug) \
		-database_address=$(database_address) \
		-gitlab_token=$(gitlab_token) -gitlab_base_url=$(gitlab_base_url) \
		-ucloud_publick_key=$(ucloud_publick_key) -ucloud_private_key=$(ucloud_private_key) \
		-log_level=$(log_level) \
		-upyun_bucket=$(upyun_bucket) -upyun_op=$(upyun_op) -upyun_passwd=$(upyun_passwd) -upyun_prefix=$(upyun_prefix) \
		-http_server_port=$(http_server_port) -log_server_port=$(log_server_port)
release:
	cd ember-app-espire && yarn run build && cd ..
	cp ember-app-espire/dist/assets/ember-app-espire.css static/css/.
	cp ember-app-espire/dist/assets/vendor.css static/css/.
	cp ember-app-espire/dist/assets/semantic.min.css static/css/.
	cp -r ember-app-espire/dist/assets/themes static/css/.
	rm -r static/css/themes/github
	rm -r static/css/themes/material
	cp ember-app-espire/dist/assets/vendor.js static/js/.
	cp ember-app-espire/dist/assets/tablesort.js static/js/.
	cp ember-app-espire/dist/assets/ember-app-espire.js static/js/.
	cp ember-app-espire/dist/assets/semantic.min.js static/js/.
	env GOOS=linux GOARCH=amd64 packr build && git tag |tail -1 | awk '{print "tar -zcf linux_"$$0"_amd64.tar.gz devops-ui"}' | bash
	env GOOS=darwin GOARCH=amd64 packr build && git tag |tail -1 | awk '{print "tar -zcf darwin_"$$0"_amd64.tar.gz devops-ui"}' | bash
