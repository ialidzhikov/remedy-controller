// +build !ignore_autogenerated

/*
Copyright (c) 2020 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	azure "github.wdf.sap.corp/kubernetes/remedy-controller/pkg/apis/azure"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*PublicIPAddress)(nil), (*azure.PublicIPAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PublicIPAddress_To_azure_PublicIPAddress(a.(*PublicIPAddress), b.(*azure.PublicIPAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.PublicIPAddress)(nil), (*PublicIPAddress)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_PublicIPAddress_To_v1alpha1_PublicIPAddress(a.(*azure.PublicIPAddress), b.(*PublicIPAddress), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PublicIPAddressList)(nil), (*azure.PublicIPAddressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PublicIPAddressList_To_azure_PublicIPAddressList(a.(*PublicIPAddressList), b.(*azure.PublicIPAddressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.PublicIPAddressList)(nil), (*PublicIPAddressList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_PublicIPAddressList_To_v1alpha1_PublicIPAddressList(a.(*azure.PublicIPAddressList), b.(*PublicIPAddressList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PublicIPAddressSpec)(nil), (*azure.PublicIPAddressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec(a.(*PublicIPAddressSpec), b.(*azure.PublicIPAddressSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.PublicIPAddressSpec)(nil), (*PublicIPAddressSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec(a.(*azure.PublicIPAddressSpec), b.(*PublicIPAddressSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PublicIPAddressStatus)(nil), (*azure.PublicIPAddressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus(a.(*PublicIPAddressStatus), b.(*azure.PublicIPAddressStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.PublicIPAddressStatus)(nil), (*PublicIPAddressStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus(a.(*azure.PublicIPAddressStatus), b.(*PublicIPAddressStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_PublicIPAddress_To_azure_PublicIPAddress(in *PublicIPAddress, out *azure.PublicIPAddress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_PublicIPAddress_To_azure_PublicIPAddress is an autogenerated conversion function.
func Convert_v1alpha1_PublicIPAddress_To_azure_PublicIPAddress(in *PublicIPAddress, out *azure.PublicIPAddress, s conversion.Scope) error {
	return autoConvert_v1alpha1_PublicIPAddress_To_azure_PublicIPAddress(in, out, s)
}

func autoConvert_azure_PublicIPAddress_To_v1alpha1_PublicIPAddress(in *azure.PublicIPAddress, out *PublicIPAddress, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_azure_PublicIPAddress_To_v1alpha1_PublicIPAddress is an autogenerated conversion function.
func Convert_azure_PublicIPAddress_To_v1alpha1_PublicIPAddress(in *azure.PublicIPAddress, out *PublicIPAddress, s conversion.Scope) error {
	return autoConvert_azure_PublicIPAddress_To_v1alpha1_PublicIPAddress(in, out, s)
}

func autoConvert_v1alpha1_PublicIPAddressList_To_azure_PublicIPAddressList(in *PublicIPAddressList, out *azure.PublicIPAddressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]azure.PublicIPAddress)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PublicIPAddressList_To_azure_PublicIPAddressList is an autogenerated conversion function.
func Convert_v1alpha1_PublicIPAddressList_To_azure_PublicIPAddressList(in *PublicIPAddressList, out *azure.PublicIPAddressList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PublicIPAddressList_To_azure_PublicIPAddressList(in, out, s)
}

func autoConvert_azure_PublicIPAddressList_To_v1alpha1_PublicIPAddressList(in *azure.PublicIPAddressList, out *PublicIPAddressList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PublicIPAddress)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_azure_PublicIPAddressList_To_v1alpha1_PublicIPAddressList is an autogenerated conversion function.
func Convert_azure_PublicIPAddressList_To_v1alpha1_PublicIPAddressList(in *azure.PublicIPAddressList, out *PublicIPAddressList, s conversion.Scope) error {
	return autoConvert_azure_PublicIPAddressList_To_v1alpha1_PublicIPAddressList(in, out, s)
}

func autoConvert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec(in *PublicIPAddressSpec, out *azure.PublicIPAddressSpec, s conversion.Scope) error {
	out.IPAddress = in.IPAddress
	return nil
}

// Convert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec is an autogenerated conversion function.
func Convert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec(in *PublicIPAddressSpec, out *azure.PublicIPAddressSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_PublicIPAddressSpec_To_azure_PublicIPAddressSpec(in, out, s)
}

func autoConvert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec(in *azure.PublicIPAddressSpec, out *PublicIPAddressSpec, s conversion.Scope) error {
	out.IPAddress = in.IPAddress
	return nil
}

// Convert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec is an autogenerated conversion function.
func Convert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec(in *azure.PublicIPAddressSpec, out *PublicIPAddressSpec, s conversion.Scope) error {
	return autoConvert_azure_PublicIPAddressSpec_To_v1alpha1_PublicIPAddressSpec(in, out, s)
}

func autoConvert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus(in *PublicIPAddressStatus, out *azure.PublicIPAddressStatus, s conversion.Scope) error {
	out.Exists = in.Exists
	out.ID = (*string)(unsafe.Pointer(in.ID))
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.ProvisioningState = (*string)(unsafe.Pointer(in.ProvisioningState))
	return nil
}

// Convert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus is an autogenerated conversion function.
func Convert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus(in *PublicIPAddressStatus, out *azure.PublicIPAddressStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_PublicIPAddressStatus_To_azure_PublicIPAddressStatus(in, out, s)
}

func autoConvert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus(in *azure.PublicIPAddressStatus, out *PublicIPAddressStatus, s conversion.Scope) error {
	out.Exists = in.Exists
	out.ID = (*string)(unsafe.Pointer(in.ID))
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.ProvisioningState = (*string)(unsafe.Pointer(in.ProvisioningState))
	return nil
}

// Convert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus is an autogenerated conversion function.
func Convert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus(in *azure.PublicIPAddressStatus, out *PublicIPAddressStatus, s conversion.Scope) error {
	return autoConvert_azure_PublicIPAddressStatus_To_v1alpha1_PublicIPAddressStatus(in, out, s)
}
