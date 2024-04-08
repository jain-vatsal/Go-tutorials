package main

import "errors"

/*
MUTATIONS
INSERT AN ELEMENT
m[key] = elem

GET AN ELEMENT
elem = m[key]

DELETE AN ELEMENT
delete(m, key)

CHECK IF A KEY EXISTS
elem, ok := m[key]

If key is in m, then ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type
*/

/*
ASSIGNMENT
It's important to keep up with privacy regulations and to respect our user's data. We need a function that will delete user records.

Complete the deleteIfNecessary function.

Check the scheduledForDeletion bool to determine if they are scheduled for deletion or not.

If the user doesn't exist in the map, return the error not found.
If they exist but aren't scheduled for deletion, return deleted as false with no errors.
If they exist and are scheduled for deletion, return deleted as true with no errors and delete their record from the map.
NOTE ON PASSING MAPS
Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we can make changes to the original, we don't have a copy.
*/

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	nm, ok := users[name]

	if !ok {
		return false, errors.New("Not found")
	}

	if !nm.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)

	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
