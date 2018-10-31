// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import "os"

// Perm is the interface that groups file permission control methods.
type Perm interface {
	GetOwner(string) (string, error)
	GetGroup(string) (string, error)
	GetMode(string) (os.FileMode, error)

	ChOwner(string, string) error
	ChGroup(string, string) error
	ChMode(string, os.FileMode) error
}

// SimplePerm implements Perm that simply returns file owner and group by Get* methods.
// Set* methods nothing to do.
type SimplePerm struct {
	owner, group string
}

// NewSimplePerm creates and initializes a new SimplePerm.
func NewSimplePerm(owner, group string) *SimplePerm {
	return &SimplePerm{
		owner: owner,
		group: group,
	}
}

// GetOwner returns pre-setted owner.
func (s *SimplePerm) GetOwner(string) (string, error) {
	return s.owner, nil
}

// GetGroup returns pre-setted group.
func (s *SimplePerm) GetGroup(string) (string, error) {
	return s.group, nil
}

// GetMode returns as it is os.ModePerm.
func (s *SimplePerm) GetMode(string) (os.FileMode, error) {
	return os.ModePerm, nil
}

// ChOwner is nothing to do.
func (s *SimplePerm) ChOwner(string, string) error {
	return nil
}

// ChGroup is nothing to do.
func (s *SimplePerm) ChGroup(string, string) error {
	return nil
}

// ChMode is nothing to do.
func (s *SimplePerm) ChMode(string, os.FileMode) error {
	return nil
}
