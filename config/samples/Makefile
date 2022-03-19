.PHONY: cleanup apply

cleanup:
	kubectl delete -f haproxy-opeartor_v1alpha1_haproxyconfig.yaml
	kubectl delete -f haproxy-opeartor_v1alpha1_route.yaml

all:
	kubectl apply -f haproxy-opeartor_v1alpha1_haproxyconfig.yaml
	kubectl apply -f haproxy-opeartor_v1alpha1_route.yaml

config:
	kubectl apply -f haproxy-opeartor_v1alpha1_haproxyconfig.yaml

route:
	kubectl apply -f haproxy-opeartor_v1alpha1_route.yaml

