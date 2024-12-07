// Code generated by ent, DO NOT EDIT.

package track

import (
	"affluo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldID, id))
}

// IPAddress applies equality check predicate on the "ip_address" field. It's identical to IPAddressEQ.
func IPAddress(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldIPAddress, v))
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldUserAgent, v))
}

// DeviceFingerprint applies equality check predicate on the "device_fingerprint" field. It's identical to DeviceFingerprintEQ.
func DeviceFingerprint(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDeviceFingerprint, v))
}

// Referrer applies equality check predicate on the "referrer" field. It's identical to ReferrerEQ.
func Referrer(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldReferrer, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldCreatedAt, v))
}

// IsUniqueClick applies equality check predicate on the "is_unique_click" field. It's identical to IsUniqueClickEQ.
func IsUniqueClick(v bool) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldIsUniqueClick, v))
}

// IPAddressEQ applies the EQ predicate on the "ip_address" field.
func IPAddressEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldIPAddress, v))
}

// IPAddressNEQ applies the NEQ predicate on the "ip_address" field.
func IPAddressNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldIPAddress, v))
}

// IPAddressIn applies the In predicate on the "ip_address" field.
func IPAddressIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldIPAddress, vs...))
}

// IPAddressNotIn applies the NotIn predicate on the "ip_address" field.
func IPAddressNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldIPAddress, vs...))
}

// IPAddressGT applies the GT predicate on the "ip_address" field.
func IPAddressGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldIPAddress, v))
}

// IPAddressGTE applies the GTE predicate on the "ip_address" field.
func IPAddressGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldIPAddress, v))
}

// IPAddressLT applies the LT predicate on the "ip_address" field.
func IPAddressLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldIPAddress, v))
}

// IPAddressLTE applies the LTE predicate on the "ip_address" field.
func IPAddressLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldIPAddress, v))
}

// IPAddressContains applies the Contains predicate on the "ip_address" field.
func IPAddressContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldIPAddress, v))
}

// IPAddressHasPrefix applies the HasPrefix predicate on the "ip_address" field.
func IPAddressHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldIPAddress, v))
}

// IPAddressHasSuffix applies the HasSuffix predicate on the "ip_address" field.
func IPAddressHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldIPAddress, v))
}

// IPAddressEqualFold applies the EqualFold predicate on the "ip_address" field.
func IPAddressEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldIPAddress, v))
}

// IPAddressContainsFold applies the ContainsFold predicate on the "ip_address" field.
func IPAddressContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldIPAddress, v))
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldUserAgent, v))
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldUserAgent, v))
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldUserAgent, vs...))
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldUserAgent, vs...))
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldUserAgent, v))
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldUserAgent, v))
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldUserAgent, v))
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldUserAgent, v))
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldUserAgent, v))
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldUserAgent, v))
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldUserAgent, v))
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldUserAgent, v))
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldUserAgent, v))
}

// DeviceFingerprintEQ applies the EQ predicate on the "device_fingerprint" field.
func DeviceFingerprintEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldDeviceFingerprint, v))
}

// DeviceFingerprintNEQ applies the NEQ predicate on the "device_fingerprint" field.
func DeviceFingerprintNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldDeviceFingerprint, v))
}

// DeviceFingerprintIn applies the In predicate on the "device_fingerprint" field.
func DeviceFingerprintIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldDeviceFingerprint, vs...))
}

// DeviceFingerprintNotIn applies the NotIn predicate on the "device_fingerprint" field.
func DeviceFingerprintNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldDeviceFingerprint, vs...))
}

// DeviceFingerprintGT applies the GT predicate on the "device_fingerprint" field.
func DeviceFingerprintGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldDeviceFingerprint, v))
}

// DeviceFingerprintGTE applies the GTE predicate on the "device_fingerprint" field.
func DeviceFingerprintGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldDeviceFingerprint, v))
}

// DeviceFingerprintLT applies the LT predicate on the "device_fingerprint" field.
func DeviceFingerprintLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldDeviceFingerprint, v))
}

// DeviceFingerprintLTE applies the LTE predicate on the "device_fingerprint" field.
func DeviceFingerprintLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldDeviceFingerprint, v))
}

// DeviceFingerprintContains applies the Contains predicate on the "device_fingerprint" field.
func DeviceFingerprintContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldDeviceFingerprint, v))
}

// DeviceFingerprintHasPrefix applies the HasPrefix predicate on the "device_fingerprint" field.
func DeviceFingerprintHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldDeviceFingerprint, v))
}

// DeviceFingerprintHasSuffix applies the HasSuffix predicate on the "device_fingerprint" field.
func DeviceFingerprintHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldDeviceFingerprint, v))
}

// DeviceFingerprintEqualFold applies the EqualFold predicate on the "device_fingerprint" field.
func DeviceFingerprintEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldDeviceFingerprint, v))
}

// DeviceFingerprintContainsFold applies the ContainsFold predicate on the "device_fingerprint" field.
func DeviceFingerprintContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldDeviceFingerprint, v))
}

// ReferrerEQ applies the EQ predicate on the "referrer" field.
func ReferrerEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldReferrer, v))
}

// ReferrerNEQ applies the NEQ predicate on the "referrer" field.
func ReferrerNEQ(v string) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldReferrer, v))
}

// ReferrerIn applies the In predicate on the "referrer" field.
func ReferrerIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldReferrer, vs...))
}

// ReferrerNotIn applies the NotIn predicate on the "referrer" field.
func ReferrerNotIn(vs ...string) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldReferrer, vs...))
}

// ReferrerGT applies the GT predicate on the "referrer" field.
func ReferrerGT(v string) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldReferrer, v))
}

// ReferrerGTE applies the GTE predicate on the "referrer" field.
func ReferrerGTE(v string) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldReferrer, v))
}

// ReferrerLT applies the LT predicate on the "referrer" field.
func ReferrerLT(v string) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldReferrer, v))
}

// ReferrerLTE applies the LTE predicate on the "referrer" field.
func ReferrerLTE(v string) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldReferrer, v))
}

// ReferrerContains applies the Contains predicate on the "referrer" field.
func ReferrerContains(v string) predicate.Track {
	return predicate.Track(sql.FieldContains(FieldReferrer, v))
}

// ReferrerHasPrefix applies the HasPrefix predicate on the "referrer" field.
func ReferrerHasPrefix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasPrefix(FieldReferrer, v))
}

// ReferrerHasSuffix applies the HasSuffix predicate on the "referrer" field.
func ReferrerHasSuffix(v string) predicate.Track {
	return predicate.Track(sql.FieldHasSuffix(FieldReferrer, v))
}

// ReferrerIsNil applies the IsNil predicate on the "referrer" field.
func ReferrerIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldReferrer))
}

// ReferrerNotNil applies the NotNil predicate on the "referrer" field.
func ReferrerNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldReferrer))
}

// ReferrerEqualFold applies the EqualFold predicate on the "referrer" field.
func ReferrerEqualFold(v string) predicate.Track {
	return predicate.Track(sql.FieldEqualFold(FieldReferrer, v))
}

// ReferrerContainsFold applies the ContainsFold predicate on the "referrer" field.
func ReferrerContainsFold(v string) predicate.Track {
	return predicate.Track(sql.FieldContainsFold(FieldReferrer, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldType, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldStatus, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Track {
	return predicate.Track(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Track {
	return predicate.Track(sql.FieldLTE(FieldCreatedAt, v))
}

// IsUniqueClickEQ applies the EQ predicate on the "is_unique_click" field.
func IsUniqueClickEQ(v bool) predicate.Track {
	return predicate.Track(sql.FieldEQ(FieldIsUniqueClick, v))
}

// IsUniqueClickNEQ applies the NEQ predicate on the "is_unique_click" field.
func IsUniqueClickNEQ(v bool) predicate.Track {
	return predicate.Track(sql.FieldNEQ(FieldIsUniqueClick, v))
}

// AdditionalMetadataIsNil applies the IsNil predicate on the "additional_metadata" field.
func AdditionalMetadataIsNil() predicate.Track {
	return predicate.Track(sql.FieldIsNull(FieldAdditionalMetadata))
}

// AdditionalMetadataNotNil applies the NotNil predicate on the "additional_metadata" field.
func AdditionalMetadataNotNil() predicate.Track {
	return predicate.Track(sql.FieldNotNull(FieldAdditionalMetadata))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCampaign applies the HasEdge predicate on the "campaign" edge.
func HasCampaign() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CampaignTable, CampaignColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCampaignWith applies the HasEdge predicate on the "campaign" edge with a given conditions (other predicates).
func HasCampaignWith(preds ...predicate.Campaign) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newCampaignStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLink applies the HasEdge predicate on the "link" edge.
func HasLink() predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LinkTable, LinkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLinkWith applies the HasEdge predicate on the "link" edge with a given conditions (other predicates).
func HasLinkWith(preds ...predicate.CampaignLink) predicate.Track {
	return predicate.Track(func(s *sql.Selector) {
		step := newLinkStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Track) predicate.Track {
	return predicate.Track(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Track) predicate.Track {
	return predicate.Track(sql.NotPredicates(p))
}
