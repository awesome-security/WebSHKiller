package main

import "os"

var (
	OPENSTACK_URL      = os.Getenv("WEBSHDD_OPENSTACK_URL")
	OPENSTACK_USERNAME = os.Getenv("WEBSHDD_OPENSTACK_USERNAME")
	OPENSTACK_PASSWORD = os.Getenv("WEBSHDD_OPENSTACK_PASSWORD")
)
