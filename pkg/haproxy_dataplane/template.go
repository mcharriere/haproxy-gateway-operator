package haproxy_dataplane

var RAW_CONFIG = `
global
  daemon
  maxconn 50000

userlist dataplaneapi
  user admin insecure-password securePassword

program api
  command /usr/bin/dataplaneapi --host 0.0.0.0 --port 5555 --haproxy-bin /usr/sbin/haproxy --config-file /usr/local/etc/haproxy/haproxy.cfg --reload-cmd "kill -SIGUSR2 1" --reload-delay 5 --userlist dataplaneapi 
  no option start-on-reload

defaults
  mode http
  timeout connect 50s
  timeout client 900s
  timeout server 900s

frontend http
  bind *:80
  default_backend default
{{ range .Frontends }}
  acl {{ .Name }} hdr(host) -i {{ .Host }}
  use_backend {{ .Backend }} if {{ .Name }}
{{ end }}

backend default
  balance roundrobin
  # option httpchk GET /v1/sys/health
  server black_page blank.org:443 maxconn 5000 weight 1

{{ range .Backends }}
backend {{ .Name }}
  balance roundrobin
{{- range .Servers }}
  server {{ .Name }} {{ .Address }}:{{ .Port }}
{{- end }}
{{ end }}
`
