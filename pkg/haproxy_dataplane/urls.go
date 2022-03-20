package haproxy_dataplane

const URL_TRANSACTION_START string = "%s/v2/services/haproxy/transactions?version=%d"
const URL_TRANSACTION_COMMIT string = "%s/v2/services/haproxy/transactions/%s"
const URL_VERSION_GET string = "%s/v2/services/haproxy/configuration/version"
const URL_ACL_ADD string = "%s/v2/services/haproxy/configuration/acls?transaction_id=%s&parent_name=%s&parent_type=frontend"
const URL_ACL_GET string = "%s/v2/services/haproxy/configuration/acls?parent_name=%s&parent_type=frontend"
const URL_ACL_REPLACE string = "%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=frontend"
