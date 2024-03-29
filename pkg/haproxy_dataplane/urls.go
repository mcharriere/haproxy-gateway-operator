package haproxy_dataplane

const URL_VERSION_GET string = "%s/v2/services/haproxy/configuration/version"

const URL_TRANSACTION_START string = "%s/v2/services/haproxy/transactions?version=%d"
const URL_TRANSACTION_COMMIT string = "%s/v2/services/haproxy/transactions/%s"
const URL_TRANSACTION_DELETE string = "%s/v2/services/haproxy/transactions/%s"

const URL_ACL_ADD string = "%s/v2/services/haproxy/configuration/acls?transaction_id=%s&parent_name=%s&parent_type=frontend"
const URL_ACL_GET string = "%s/v2/services/haproxy/configuration/acls?parent_name=%s&parent_type=frontend"
const URL_ACL_REPLACE string = "%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=frontend"
const URL_ACL_DELETE string = "%s/v2/services/haproxy/configuration/acls/%d?transaction_id=%s&parent_name=%s&parent_type=frontend"

const URL_BACKEND_ADD string = "%s/v2/services/haproxy/configuration/backends?transaction_id=%s"
const URL_BACKEND_GET string = "%s/v2/services/haproxy/configuration/backends"
const URL_BACKEND_REPLACE string = "%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s"
const URL_BACKEND_DELETE string = "%s/v2/services/haproxy/configuration/backends/%s?transaction_id=%s"

const URL_SERVER_ADD string = "%s/v2/services/haproxy/configuration/servers?transaction_id=%s&backend=%s"
const URL_SERVER_GET string = "%s/v2/services/haproxy/configuration/servers?backend=%s"
const URL_SERVER_REPLACE string = "%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&backend=%s"
const URL_SERVER_DELETE string = "%s/v2/services/haproxy/configuration/servers/%s?transaction_id=%s&backend=%s"

const URL_RULE_ADD string = "%s/v2/services/haproxy/configuration/backend_switching_rules?transaction_id=%s&frontend=%s"
const URL_RULE_GET string = "%s/v2/services/haproxy/configuration/backend_switching_rules?frontend=%s"
const URL_RULE_REPLACE string = "%s/v2/services/haproxy/configuration/backend_switching_rules/%d?transaction_id=%s&frontend=%s"
const URL_RULE_DELETE string = "%s/v2/services/haproxy/configuration/backend_switching_rules/%d?transaction_id=%s&frontend=%s"
