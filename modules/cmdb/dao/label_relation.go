package dao

import "github.com/erda-project/erda/apistructs"

// LabelRelation 标签关联关系
type LabelRelation struct {
	BaseModel
	LabelID uint64                      `json:"label_id"` // 标签 id
	RefType apistructs.ProjectLabelType `json:"ref_type"` // 标签作用类型, eg: issue
	RefID   uint64                      `json:"ref_id"`   // 标签关联目标 id
}

// TableName 表名
func (LabelRelation) TableName() string {
	return "dice_label_relations"
}

// CreateLabelRelation 创建标签关联关系
func (client *DBClient) CreateLabelRelation(lr *LabelRelation) error {
	return client.Create(lr).Error
}

// DeleteLabelRelations 删除标签关联关系
func (client *DBClient) DeleteLabelRelations(refType apistructs.ProjectLabelType, refID uint64) error {
	return client.Where("ref_type = ?", refType).Where("ref_id = ?", refID).
		Delete(LabelRelation{}).Error
}

// DeleteLabelRelationsByLabel 根据 labelID 删除标签
func (client *DBClient) DeleteLabelRelationsByLabel(labelID uint64) error {
	return client.Where("label_id = ?", labelID).Delete(LabelRelation{}).Error
}

// GetLabelRelationsByRef 获取标签关联关系列表
func (client *DBClient) GetLabelRelationsByRef(refType apistructs.ProjectLabelType, refID uint64) ([]LabelRelation, error) {
	var lrs []LabelRelation
	if err := client.Where("ref_type = ?", refType).Where("ref_id = ?", refID).
		Find(&lrs).Error; err != nil {
		return nil, err
	}

	return lrs, nil
}

// GetLabelRelationsByLabels 获取标签关联关系列表
func (client *DBClient) GetLabelRelationsByLabels(refType apistructs.ProjectLabelType, labelIDs []uint64) ([]LabelRelation, error) {
	var lrs []LabelRelation
	if err := client.Debug().Where("ref_type = ?", refType).Where("label_id in (?)", labelIDs).
		Find(&lrs).Error; err != nil {
		return nil, err
	}

	return lrs, nil
}

// BatchGetIssueLabelIDMap 批量查询 issue label id
func (client *DBClient) BatchQueryIssueLabelIDMap(issueIDs []int64) (map[uint64][]uint64, error) {
	if len(issueIDs) == 0 {
		return nil, nil
	}
	var refs []LabelRelation
	sql := client.Where("`ref_type` = ?", "issue").Where("`ref_id` IN (?)", issueIDs).Find(&refs)
	if err := sql.Error; err != nil {
		return nil, err
	}
	if len(refs) == 0 {
		return nil, nil
	}
	// key: issueID, value: labelIDs
	m := make(map[uint64][]uint64, len(issueIDs))
	for _, ref := range refs {
		m[ref.RefID] = append(m[ref.RefID], ref.LabelID)
	}
	return m, nil
}