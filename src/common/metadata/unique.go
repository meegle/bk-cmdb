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

package metadata

type ObjectUnique struct {
	ID        uint64       `json:"id" bson:"id"`
	ObjID     string       `json:"bk_obj_id" bson:"bk_obj_id"`
	MustCheck bool         `json:"must_check" bson:"must_check"`
	Keys      []UinqueKeys `json:"keys" bson:"keys"`
	Ispre     bool         `json:"ispre" bson:"ispre"`
	OwnerID   string       `json:"bk_supplier_account" bson:"bk_supplier_account"`
	LastTime  Time         `json:"last_time" bson:"last_time"`
}

type UinqueKeys struct {
	Kind string `json:"key_kind" bson:"key_kind"`
	ID   uint64 `json:"key_id" bson:"key_id"`
}

type CreateUniqueRequest struct {
	ObjID     string       `json:"bk_obj_id" bson:"bk_obj_id"`
	MustCheck bool         `json:"must_check" bson:"must_check"`
	Keys      []UinqueKeys `json:"keys" bson:"keys"`
}

type CreateUniqueResult struct {
	BaseResp
	Data RspID `json:"data"`
}

type UpdateUniqueRequest struct {
	ID        uint64       `json:"id" bson:"id"`
	ObjID     string       `json:"bk_obj_id" bson:"bk_obj_id"`
	MustCheck bool         `json:"must_check" bson:"must_check"`
	Keys      []UinqueKeys `json:"keys" bson:"keys"`
}

type UpdateUniqueResult struct {
	BaseResp
}

type DeleteUniqueRequest struct {
	ID    uint64 `json:"id"`
	ObjID string `json:"bk_obj_id"`
}

type DeleteUniqueResult struct {
	BaseResp
}

type SearchUniqueRequest struct {
	ObjID string `json:"bk_obj_id"`
}

type SearchUniqueResult struct {
	BaseResp
	Data []ObjectUnique `json:"data"`
}
