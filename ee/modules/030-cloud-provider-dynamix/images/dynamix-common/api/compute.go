/*
Copyright 2024 Flant JSC
Licensed under the Deckhouse Platform Enterprise Edition (EE) license. See https://github.com/deckhouse/deckhouse/blob/main/ee/LICENSE
*/

package api

import (
	"context"
	"fmt"
	"k8s.io/klog"
	"net"
	"strconv"

	decort "repository.basistech.ru/BASIS/decort-golang-sdk"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/compute"

	"dynamix-common/retry"
)

type ComputeService struct {
	client  *decort.DecortClient
	retryer retry.Retryer
}

func NewComputeService(client *decort.DecortClient) *ComputeService {
	return &ComputeService{
		client:  client,
		retryer: retry.NewRetryer(),
	}
}

func (c *ComputeService) GetVMByName(ctx context.Context, name string) (*compute.ItemCompute, error) {
	var vm *compute.ItemCompute

	err := c.retryer.Do(ctx, func() (bool, error) {
		computes, err := c.client.CloudAPI().Compute().List(ctx, compute.ListRequest{
			Name: name,
		})
		if err != nil {
			return false, err
		}

		if len(computes.Data) == 0 {
			return true, ErrNotFound
		}

		vm = &computes.Data[0]

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func (c *ComputeService) GetVMByID(ctx context.Context, id string) (*compute.ItemCompute, error) {
	computeID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse compute id [%s]: %v", id, err)
	}

	var vm *compute.ItemCompute

	err = c.retryer.WithAttempts(5).Do(ctx, func() (bool, error) {
		computes, err := c.client.CloudAPI().Compute().List(ctx, compute.ListRequest{
			ByID: computeID,
		})
		if err != nil {
			return false, err
		}

		if len(computes.Data) == 0 {
			return true, ErrNotFound
		}

		vm = &computes.Data[0]

		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return vm, nil
}

// GetVMIPAddresses return external and local IPv4
func (c *ComputeService) GetVMIPAddresses(vm *compute.ItemCompute) ([]string, []string, error) {
	var (
		externalIPs []string
		localIPs    []string
	)

	for _, vmInterface := range vm.Interfaces {
		ip := net.ParseIP(vmInterface.IPAddress)
		if ip == nil {
			klog.V(4).Infof("GetVMIPAddresses: invalid IP address [%v]", vmInterface.IPAddress)
			continue
		}

		// skip IPv6
		if ip.To4() == nil {
			klog.V(4).Infof("GetVMIPAddresses: ip [%v] skipped, not IPv4", ip.String())
			continue
		}

		//if vmInterface.NetType == "EXTNET" {
		//	klog.V(4).Infof("GetVMIPAddresses: externalIP [%v] ", vmInterface.IPAddress)
		//	externalIPs = append(externalIPs, vmInterface.IPAddress)
		//} else {
		//	klog.V(4).Infof("GetVMIPAddresses: internalIP [%v] ", vmInterface.IPAddress)
		//	localIPs = append(localIPs, vmInterface.IPAddress)
		//}

		// TODO: FIXME
		if vmInterface.NetType == "EXTNET" {
			klog.V(4).Infof("GetVMIPAddresses: externalIP [%v] ", vmInterface.IPAddress)
			externalIPs = append(externalIPs, vmInterface.IPAddress)
			klog.V(4).Infof("GetVMIPAddresses: internalIP [%v] ", vmInterface.IPAddress)
			localIPs = append(localIPs, vmInterface.IPAddress)
		}
	}

	return externalIPs, localIPs, nil
}
