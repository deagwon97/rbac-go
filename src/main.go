package main

import (
	"rbac-go/routes"
)

// @title           RBAC GO API
// @version         1.0
// @description     This is a RBAC server.
func main() {
	routes.Run(":8000")
}
