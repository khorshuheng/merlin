// Copyright 2020 The Merlin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"strconv"
	"time"

	"github.com/gojek/merlin/mlp"
)

const (
	ModelTypePyFunc     = "pyfunc"
	ModelTypeTensorflow = "tensorflow"
	ModelTypeXgboost    = "xgboost"
	ModelTypeSkLearn    = "sklearn"
	ModelTypePyTorch    = "pytorch"
	ModelTypeOnnx       = "onnx"
	ModelTypePyFuncV2   = "pyfunc_v2"
)

type Id int

func (id Id) String() string {
	return strconv.Itoa(int(id))
}

func ParseId(id string) (Id, error) {
	if parsed, err := strconv.Atoi(id); err != nil {
		return -1, err
	} else {
		return Id(parsed), nil
	}
}

type CreatedUpdated struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	Id           Id          `json:"id"`
	Name         string      `json:"name" validate:"required,min=3,max=25,subdomain_rfc1123"`
	ProjectId    Id          `json:"project_id"`
	Project      mlp.Project `json:"-" gorm:"-"`
	ExperimentId Id          `json:"mlflow_experiment_id" gorm:"column:mlflow_experiment_id"`
	Type         string      `json:"type" gorm:"type"`
	MlflowUrl    string      `json:"mlflow_url" gorm:"-"`

	Endpoints []*ModelEndpoint `json:"endpoints" gorm:"foreignkey:ModelId;"`

	CreatedUpdated
}