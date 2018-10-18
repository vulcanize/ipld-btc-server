// Copyright 2018 Vulcanize
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

package vat_fold

import (
	. "github.com/onsi/gomega"

	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/transformers/vat_fold"
)

type MockVatFoldRepository struct {
	createErr                       error
	markHeaderCheckedErr            error
	MarkHeaderCheckedPassedHeaderID int64
	missingHeaders                  []core.Header
	missingHeadersErr               error
	PassedStartingBlockNumber       int64
	PassedEndingBlockNumber         int64
	PassedHeaderID                  int64
	PassedModels                    []vat_fold.VatFoldModel
}

func (repository *MockVatFoldRepository) Create(headerID int64, models []vat_fold.VatFoldModel) error {
	repository.PassedHeaderID = headerID
	repository.PassedModels = models
	return repository.createErr
}

func (repository *MockVatFoldRepository) MarkHeaderChecked(headerID int64) error {
	repository.MarkHeaderCheckedPassedHeaderID = headerID
	return repository.markHeaderCheckedErr
}

func (repository *MockVatFoldRepository) MissingHeaders(startingBlockNumber, endingBlockNumber int64) ([]core.Header, error) {
	repository.PassedStartingBlockNumber = startingBlockNumber
	repository.PassedEndingBlockNumber = endingBlockNumber
	return repository.missingHeaders, repository.missingHeadersErr
}

func (repository *MockVatFoldRepository) SetMarkHeaderCheckedErr(e error) {
	repository.markHeaderCheckedErr = e
}

func (repository *MockVatFoldRepository) SetMissingHeadersErr(e error) {
	repository.missingHeadersErr = e
}

func (repository *MockVatFoldRepository) SetMissingHeaders(headers []core.Header) {
	repository.missingHeaders = headers
}

func (repository *MockVatFoldRepository) SetCreateError(e error) {
	repository.createErr = e
}

func (repository *MockVatFoldRepository) AssertMarkHeaderCheckedCalledWith(i int64) {
	Expect(repository.MarkHeaderCheckedPassedHeaderID).To(Equal(i))
}