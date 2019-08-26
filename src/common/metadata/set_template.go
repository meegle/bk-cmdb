/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package metadata

import "time"

// SetTemplate 集群模板
type SetTemplate struct {
	ID    int64  `field:"id" json:"id" bson:"id"`
	Name  string `field:"name" json:"name" bson:"name"`
	BizID int64  `field:"bk_biz_id" json:"bk_biz_id" bson:"bk_biz_id"`

	// 通用字段
	Creator         string    `field:"creator" json:"creator,omitempty" bson:"creator"`
	Modifier        string    `field:"modifier" json:"modifier,omitempty" bson:"modifier"`
	CreateTime      time.Time `field:"create_time" json:"create_time,omitempty" bson:"create_time"`
	LastTime        time.Time `field:"last_time" json:"last_time,omitempty" bson:"last_time"`
	SupplierAccount string    `field:"bk_supplier_account" json:"bk_supplier_account,omitempty" bson:"bk_supplier_account"`
}

// 拓扑模板与服务模板多对多关系, 记录拓扑模板的构成
type SetTemplateServiceTemplateRelation struct {
	SetTemplateID     int64 `json:"set_template_id"`
	ServiceTemplateID int64 `json:"service_template_id"`
}
