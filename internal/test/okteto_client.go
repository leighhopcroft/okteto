// Copyright 2021 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"context"

	"github.com/okteto/okteto/pkg/types"
)

type FakeOktetoClientProvider struct {
	UserContext *types.UserContext
	Err         error
}

func NewFakeOktetoClientProvider(userContext *types.UserContext, err error) *FakeOktetoClientProvider {
	return &FakeOktetoClientProvider{UserContext: userContext, Err: err}
}

func (f FakeOktetoClientProvider) NewOktetoUserClient() (types.UserInterface, error) {
	return FakeUserClient{UserContext: f.UserContext, err: f.Err}, nil
}

type FakeUserClient struct {
	UserContext *types.UserContext
	err         error
}

// GetUserContext get user context
func (f FakeUserClient) GetUserContext(ctx context.Context) (*types.UserContext, error) {
	return f.UserContext, f.err
}
