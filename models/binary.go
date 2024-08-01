/*
 * Religare - An Instrumental Trans Communication solution for communicating with paranormal entities
 *
 * Copyright (C) 2024 Raony Benjamim
 * Check the LICENSE.md file for more information regarding the code license
 */
package models

type Binary int

const (
	Zero Binary = 0
	One  Binary = 1
)

func (b Binary) String() string {
	if b == Zero {
		return "0"
	}
	return "1"
}
