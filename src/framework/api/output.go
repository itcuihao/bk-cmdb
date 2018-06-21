/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"configcenter/src/framework/common"
	"configcenter/src/framework/core/output/module/inst"
	"configcenter/src/framework/core/output/module/model"
	"io"
	"strings"
)

// GetModel get the model
func GetModel(supplierAccount, classificationID, objID string) (model.Model, error) {
	return mgr.OutputerMgr.GetModel(supplierAccount, classificationID, objID)
}

// CreateClassification create a new classification
func CreateClassification(name string) model.Classification {
	return mgr.OutputerMgr.CreateClassification(name)
}

// FindClassificationsLikeName find a array of the classification by the name
func FindClassificationsLikeName(name string) (model.ClassificationIterator, error) {
	return mgr.OutputerMgr.FindClassificationsLikeName(name)
}

// FindClassificationsByCondition find a array of the classification by the condition
func FindClassificationsByCondition(condition common.Condition) (model.ClassificationIterator, error) {
	return mgr.OutputerMgr.FindClassificationsByCondition(condition)
}

// CreateBusiness create a new Business object
func CreateBusiness(supplierAccount string) (*BusinessWrapper, error) {

	targetModel, err := GetModel(supplierAccount, "bk_organization", "biz")
	if nil != err {
		return nil, err
	}
	businessInst := mgr.OutputerMgr.InstOperation().CreateBusinessInst(targetModel)
	wrapper := &BusinessWrapper{business: businessInst}
	//wrapper.SetSupplierAccount(supplierAccount)
	return wrapper, err
}

// FindBusinessLikeName find all insts by the name
func FindBusinessLikeName(supplierAccount, businessName string) (*BusinessIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_organization", "biz")
	if nil != err {
		return nil, err
	}

	iter, err := mgr.OutputerMgr.InstOperation().FindBusinessLikeName(targetModel, businessName)
	return &BusinessIteratorWrapper{business: iter}, err
}

// FindBusinessByCondition find all insts by the condition
func FindBusinessByCondition(supplierAccount string, cond common.Condition) (*BusinessIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_organization", "biz")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindBusinessByCondition(targetModel, cond)
	return &BusinessIteratorWrapper{business: iter}, err
}

// CreateSet create a new set object
func CreateSet(supplierAccount string) (*SetWrapper, error) {

	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "set")
	if nil != err {
		return nil, err
	}

	setInst := mgr.OutputerMgr.InstOperation().CreateSetInst(targetModel)
	return &SetWrapper{set: setInst}, err
}

// FindSetLikeName find all insts by the name
func FindSetLikeName(supplierAccount, setName string) (*SetIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "set")
	if nil != err {
		return nil, err
	}

	iter, err := mgr.OutputerMgr.InstOperation().FindSetsLikeName(targetModel, setName)
	return &SetIteratorWrapper{set: iter}, err
}

// FindSetByCondition find all insts by the condition
func FindSetByCondition(supplierAccount string, cond common.Condition) (*SetIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "set")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindSetsByCondition(targetModel, cond)
	return &SetIteratorWrapper{set: iter}, err
}

// CreateModule create a new module object
func CreateModule(supplierAccount string) (*ModuleWrapper, error) {

	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "module")
	if nil != err {
		return nil, err
	}

	moduleInst := mgr.OutputerMgr.InstOperation().CreateModuleInst(targetModel)
	return &ModuleWrapper{module: moduleInst}, err
}

// FindModuleLikeName find all insts by the name
func FindModuleLikeName(supplierAccount, moduleName string) (*ModuleIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "module")
	if nil != err {
		return nil, err
	}

	iter, err := mgr.OutputerMgr.InstOperation().FindModulesLikeName(targetModel, moduleName)
	return &ModuleIteratorWrapper{module: iter}, err
}

// FindModuleByCondition find all insts by the condition
func FindModuleByCondition(supplierAccount string, cond common.Condition) (*ModuleIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_biz_topo", "module")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindModulesByCondition(targetModel, cond)
	return &ModuleIteratorWrapper{module: iter}, err
}

// CreateHost create a new host object
func CreateHost(supplierAccount string) (*HostWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "host")
	if nil != err {
		return nil, err
	}

	hostInst := mgr.OutputerMgr.InstOperation().CreateHostInst(targetModel)
	return &HostWrapper{host: hostInst}, err
}

// FindHostLikeName find all insts by the name
func FindHostLikeName(supplierAccount, hostName string) (*HostIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "host")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindHostsLikeName(targetModel, hostName)
	return &HostIteratorWrapper{host: iter}, err
}

// FindHostByCondition find all insts by the condition
func FindHostByCondition(supplierAccount string, cond common.Condition) (*HostIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "host")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindHostsByCondition(targetModel, cond)
	return &HostIteratorWrapper{host: iter}, err
}

// CreatePlat create a new plat object
func CreatePlat(supplierAccount string) (*PlatWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "plat")
	if nil != err {
		return nil, err
	}

	platInst := mgr.OutputerMgr.InstOperation().CreatePlatInst(targetModel)
	platInst.SetValue(fieldSupplierAccount, supplierAccount)
	platInst.SetValue(fieldObjectID, plat)
	return &PlatWrapper{plat: platInst}, err
}

// GetPlatID get the plat id
func GetPlatID(supplierAccount, platName string) (int64, error) {

	cond := CreateCondition().Field(fieldSupplierAccount).Eq(supplierAccount).Field(fieldPlatName).In([]string{platName})
	iter, iterErr := FindPlatByCondition(supplierAccount, cond)
	if nil != iterErr {
		if strings.Contains(iterErr.Error(), "empty") {
			return -1, io.EOF
		}
		return -1, iterErr
	}

	platID := -1
	err := iter.ForEach(func(plat *PlatWrapper) error {
		id, err := plat.GetID()
		if nil != err {
			return err
		}
		platID = id
		return nil
	})

	if nil != err {
		return -1, err
	}

	if -1 == platID {
		return -1, io.EOF
	}
	return int64(platID), nil
}

// FindPlatLikeName find all insts by the name
func FindPlatLikeName(supplierAccount, platName string) (*PlatIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "plat")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindPlatsLikeName(targetModel, platName)
	return &PlatIteratorWrapper{plat: iter}, err
}

// FindPlatByCondition find all insts by the condition
func FindPlatByCondition(supplierAccount string, cond common.Condition) (*PlatIteratorWrapper, error) {
	targetModel, err := GetModel(supplierAccount, "bk_host_manage", "plat")
	if nil != err {
		return nil, err
	}
	iter, err := mgr.OutputerMgr.InstOperation().FindPlatsByCondition(targetModel, cond)
	return &PlatIteratorWrapper{plat: iter}, err
}

// CreateCommonInst create a common inst object
func CreateCommonInst(target model.Model) (inst.CommonInstInterface, error) {
	return mgr.OutputerMgr.InstOperation().CreateCommonInst(target), nil
}

// FindInstsLikeName find all insts by the name
func FindInstsLikeName(target model.Model, instName string) (inst.Iterator, error) {
	return mgr.OutputerMgr.InstOperation().FindCommonInstLikeName(target, instName)
}

// FindInstsByCondition find all insts by the condition
func FindInstsByCondition(target model.Model, cond common.Condition) (inst.Iterator, error) {
	return mgr.OutputerMgr.InstOperation().FindCommonInstByCondition(target, cond)
}
