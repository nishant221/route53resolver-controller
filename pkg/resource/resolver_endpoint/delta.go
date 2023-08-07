// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package resolver_endpoint

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.CreatorRequestID, b.ko.Spec.CreatorRequestID) {
		delta.Add("Spec.CreatorRequestID", a.ko.Spec.CreatorRequestID, b.ko.Spec.CreatorRequestID)
	} else if a.ko.Spec.CreatorRequestID != nil && b.ko.Spec.CreatorRequestID != nil {
		if *a.ko.Spec.CreatorRequestID != *b.ko.Spec.CreatorRequestID {
			delta.Add("Spec.CreatorRequestID", a.ko.Spec.CreatorRequestID, b.ko.Spec.CreatorRequestID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Direction, b.ko.Spec.Direction) {
		delta.Add("Spec.Direction", a.ko.Spec.Direction, b.ko.Spec.Direction)
	} else if a.ko.Spec.Direction != nil && b.ko.Spec.Direction != nil {
		if *a.ko.Spec.Direction != *b.ko.Spec.Direction {
			delta.Add("Spec.Direction", a.ko.Spec.Direction, b.ko.Spec.Direction)
		}
	}
	if !reflect.DeepEqual(a.ko.Spec.IPAddresses, b.ko.Spec.IPAddresses) {
		delta.Add("Spec.IPAddresses", a.ko.Spec.IPAddresses, b.ko.Spec.IPAddresses)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Name, b.ko.Spec.Name) {
		delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
	} else if a.ko.Spec.Name != nil && b.ko.Spec.Name != nil {
		if *a.ko.Spec.Name != *b.ko.Spec.Name {
			delta.Add("Spec.Name", a.ko.Spec.Name, b.ko.Spec.Name)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResolverEndpointType, b.ko.Spec.ResolverEndpointType) {
		delta.Add("Spec.ResolverEndpointType", a.ko.Spec.ResolverEndpointType, b.ko.Spec.ResolverEndpointType)
	} else if a.ko.Spec.ResolverEndpointType != nil && b.ko.Spec.ResolverEndpointType != nil {
		if *a.ko.Spec.ResolverEndpointType != *b.ko.Spec.ResolverEndpointType {
			delta.Add("Spec.ResolverEndpointType", a.ko.Spec.ResolverEndpointType, b.ko.Spec.ResolverEndpointType)
		}
	}
	if !ackcompare.SliceStringPEqual(a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs) {
		delta.Add("Spec.SecurityGroupIDs", a.ko.Spec.SecurityGroupIDs, b.ko.Spec.SecurityGroupIDs)
	}
	if !reflect.DeepEqual(a.ko.Spec.Tags, b.ko.Spec.Tags) {
		delta.Add("Spec.Tags", a.ko.Spec.Tags, b.ko.Spec.Tags)
	}

	return delta
}