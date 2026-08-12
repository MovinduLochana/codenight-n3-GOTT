package main

import "fmt"

// TODO: Define UserRole type.

// TODO: Declare Guest, Member, Admin, Owner constants using iota.

// TODO: Implement GetPermission function.

func main() {
	// Loop over roles and print permissions.
	roles := []UserRole{Guest, Member, Admin, Owner}
	names := []string{"Guest", "Member", "Admin", "Owner"}

	for i, r := range roles {
		fmt.Printf("%s: %s\n", names[i], GetPermission(r))
	}
}
