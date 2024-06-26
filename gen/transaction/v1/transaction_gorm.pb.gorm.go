package transactionv1

import (
	context "context"
	fmt "fmt"
	gorm1 "github.com/infobloxopen/atlas-app-toolkit/v2/gorm"
	errors "github.com/infobloxopen/protoc-gen-gorm/errors"
	v1 "github.com/sandisuryadi36/sansan-dashboard/gen/user/v1"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	strings "strings"
	time "time"
)

type UserTransactionORM struct {
	CreatedAt         *time.Time
	DeletedAt         *time.Time
	Id                uint64 `gorm:"primaryKey;not null"`
	TransactionDate   *time.Time
	TransactionStatus string
	UpdatedAt         *time.Time
	User              *v1.UserORM `gorm:"foreignKey:UserId;references:Id"`
	UserId            *uint64
}

// TableName overrides the default tablename generated by GORM
func (UserTransactionORM) TableName() string {
	return "user_transactions"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *UserTransaction) ToORM(ctx context.Context) (UserTransactionORM, error) {
	to := UserTransactionORM{}
	var err error
	if prehook, ok := interface{}(m).(UserTransactionWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.User != nil {
		tempUser, err := m.User.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	to.TransactionStatus = m.TransactionStatus
	if m.TransactionDate != nil {
		t := m.TransactionDate.AsTime()
		to.TransactionDate = &t
	}
	if m.CreatedAt != nil {
		t := m.CreatedAt.AsTime()
		to.CreatedAt = &t
	}
	if m.UpdatedAt != nil {
		t := m.UpdatedAt.AsTime()
		to.UpdatedAt = &t
	}
	if m.DeletedAt != nil {
		t := m.DeletedAt.AsTime()
		to.DeletedAt = &t
	}
	if posthook, ok := interface{}(m).(UserTransactionWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *UserTransactionORM) ToPB(ctx context.Context) (UserTransaction, error) {
	to := UserTransaction{}
	var err error
	if prehook, ok := interface{}(m).(UserTransactionWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if m.User != nil {
		tempUser, err := m.User.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.User = &tempUser
	}
	to.TransactionStatus = m.TransactionStatus
	if m.TransactionDate != nil {
		to.TransactionDate = timestamppb.New(*m.TransactionDate)
	}
	if m.CreatedAt != nil {
		to.CreatedAt = timestamppb.New(*m.CreatedAt)
	}
	if m.UpdatedAt != nil {
		to.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}
	if m.DeletedAt != nil {
		to.DeletedAt = timestamppb.New(*m.DeletedAt)
	}
	if posthook, ok := interface{}(m).(UserTransactionWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type UserTransaction the arg will be the target, the caller the one being converted from

// UserTransactionBeforeToORM called before default ToORM code
type UserTransactionWithBeforeToORM interface {
	BeforeToORM(context.Context, *UserTransactionORM) error
}

// UserTransactionAfterToORM called after default ToORM code
type UserTransactionWithAfterToORM interface {
	AfterToORM(context.Context, *UserTransactionORM) error
}

// UserTransactionBeforeToPB called before default ToPB code
type UserTransactionWithBeforeToPB interface {
	BeforeToPB(context.Context, *UserTransaction) error
}

// UserTransactionAfterToPB called after default ToPB code
type UserTransactionWithAfterToPB interface {
	AfterToPB(context.Context, *UserTransaction) error
}

// DefaultCreateUserTransaction executes a basic gorm create call
func DefaultCreateUserTransaction(ctx context.Context, in *UserTransaction, db *gorm.DB) (*UserTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type UserTransactionORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm.DB) error
}

func DefaultReadUserTransaction(ctx context.Context, in *UserTransaction, db *gorm.DB) (*UserTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := UserTransactionORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(UserTransactionORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type UserTransactionORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm.DB) error
}

func DefaultDeleteUserTransaction(ctx context.Context, in *UserTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&UserTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type UserTransactionORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm.DB) error
}

func DefaultDeleteUserTransactionSet(ctx context.Context, in []*UserTransaction, db *gorm.DB) error {
	if in == nil {
		return errors.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&UserTransactionORM{})).(UserTransactionORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&UserTransactionORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&UserTransactionORM{})).(UserTransactionORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type UserTransactionORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*UserTransaction, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*UserTransaction, *gorm.DB) error
}

// DefaultStrictUpdateUserTransaction clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateUserTransaction(ctx context.Context, in *UserTransaction, db *gorm.DB) (*UserTransaction, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateUserTransaction")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	lockedRow := &UserTransactionORM{}
	db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow)
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Omit().Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, err
}

type UserTransactionORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm.DB) error
}

// DefaultPatchUserTransaction executes a basic gorm update call with patch behavior
func DefaultPatchUserTransaction(ctx context.Context, in *UserTransaction, updateMask *field_mask.FieldMask, db *gorm.DB) (*UserTransaction, error) {
	if in == nil {
		return nil, errors.NilArgumentError
	}
	var pbObj UserTransaction
	var err error
	if hook, ok := interface{}(&pbObj).(UserTransactionWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadUserTransaction(ctx, &UserTransaction{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(UserTransactionWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskUserTransaction(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(UserTransactionWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateUserTransaction(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(UserTransactionWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type UserTransactionWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *UserTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *UserTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *UserTransaction, *field_mask.FieldMask, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *UserTransaction, *field_mask.FieldMask, *gorm.DB) error
}

// DefaultPatchSetUserTransaction executes a bulk gorm update call with patch behavior
func DefaultPatchSetUserTransaction(ctx context.Context, objects []*UserTransaction, updateMasks []*field_mask.FieldMask, db *gorm.DB) ([]*UserTransaction, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*UserTransaction, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchUserTransaction(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskUserTransaction patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskUserTransaction(ctx context.Context, patchee *UserTransaction, patcher *UserTransaction, updateMask *field_mask.FieldMask, prefix string, db *gorm.DB) (*UserTransaction, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors.NilArgumentError
	}
	var err error
	var updatedUser bool
	var updatedTransactionDate bool
	var updatedCreatedAt bool
	var updatedUpdatedAt bool
	var updatedDeletedAt bool
	for i, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if !updatedUser && strings.HasPrefix(f, prefix+"User.") {
			updatedUser = true
			if patcher.User == nil {
				patchee.User = nil
				continue
			}
			if patchee.User == nil {
				patchee.User = &v1.User{}
			}
			if o, err := v1.DefaultApplyFieldMaskUser(ctx, patchee.User, patcher.User, &field_mask.FieldMask{Paths: updateMask.Paths[i:]}, prefix+"User.", db); err != nil {
				return nil, err
			} else {
				patchee.User = o
			}
			continue
		}
		if f == prefix+"User" {
			updatedUser = true
			patchee.User = patcher.User
			continue
		}
		if f == prefix+"TransactionStatus" {
			patchee.TransactionStatus = patcher.TransactionStatus
			continue
		}
		if !updatedTransactionDate && strings.HasPrefix(f, prefix+"TransactionDate.") {
			if patcher.TransactionDate == nil {
				patchee.TransactionDate = nil
				continue
			}
			if patchee.TransactionDate == nil {
				patchee.TransactionDate = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"TransactionDate."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.TransactionDate, patchee.TransactionDate, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"TransactionDate" {
			updatedTransactionDate = true
			patchee.TransactionDate = patcher.TransactionDate
			continue
		}
		if !updatedCreatedAt && strings.HasPrefix(f, prefix+"CreatedAt.") {
			if patcher.CreatedAt == nil {
				patchee.CreatedAt = nil
				continue
			}
			if patchee.CreatedAt == nil {
				patchee.CreatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"CreatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.CreatedAt, patchee.CreatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"CreatedAt" {
			updatedCreatedAt = true
			patchee.CreatedAt = patcher.CreatedAt
			continue
		}
		if !updatedUpdatedAt && strings.HasPrefix(f, prefix+"UpdatedAt.") {
			if patcher.UpdatedAt == nil {
				patchee.UpdatedAt = nil
				continue
			}
			if patchee.UpdatedAt == nil {
				patchee.UpdatedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"UpdatedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.UpdatedAt, patchee.UpdatedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"UpdatedAt" {
			updatedUpdatedAt = true
			patchee.UpdatedAt = patcher.UpdatedAt
			continue
		}
		if !updatedDeletedAt && strings.HasPrefix(f, prefix+"DeletedAt.") {
			if patcher.DeletedAt == nil {
				patchee.DeletedAt = nil
				continue
			}
			if patchee.DeletedAt == nil {
				patchee.DeletedAt = &timestamppb.Timestamp{}
			}
			childMask := &field_mask.FieldMask{}
			for j := i; j < len(updateMask.Paths); j++ {
				if trimPath := strings.TrimPrefix(updateMask.Paths[j], prefix+"DeletedAt."); trimPath != updateMask.Paths[j] {
					childMask.Paths = append(childMask.Paths, trimPath)
				}
			}
			if err := gorm1.MergeWithMask(patcher.DeletedAt, patchee.DeletedAt, childMask); err != nil {
				return nil, nil
			}
		}
		if f == prefix+"DeletedAt" {
			updatedDeletedAt = true
			patchee.DeletedAt = patcher.DeletedAt
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListUserTransaction executes a gorm list call
func DefaultListUserTransaction(ctx context.Context, db *gorm.DB) ([]*UserTransaction, error) {
	in := UserTransaction{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []UserTransactionORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(UserTransactionORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*UserTransaction{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type UserTransactionORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm.DB) (*gorm.DB, error)
}
type UserTransactionORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm.DB, *[]UserTransactionORM) error
}
