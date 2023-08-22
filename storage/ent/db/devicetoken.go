// Code generated by ent, DO NOT EDIT.

package db

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dexidp/dex/storage/ent/db/devicetoken"
)

// DeviceToken is the model entity for the DeviceToken schema.
type DeviceToken struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DeviceCode holds the value of the "device_code" field.
	DeviceCode string `json:"device_code,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Token holds the value of the "token" field.
	Token *[]byte `json:"token,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry time.Time `json:"expiry,omitempty"`
	// LastRequest holds the value of the "last_request" field.
	LastRequest time.Time `json:"last_request,omitempty"`
	// PollInterval holds the value of the "poll_interval" field.
	PollInterval int `json:"poll_interval,omitempty"`
	// CodeChallenge holds the value of the "code_challenge" field.
	CodeChallenge string `json:"code_challenge,omitempty"`
	// CodeChallengeMethod holds the value of the "code_challenge_method" field.
	CodeChallengeMethod string `json:"code_challenge_method,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DeviceToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case devicetoken.FieldToken:
			values[i] = new([]byte)
		case devicetoken.FieldID, devicetoken.FieldPollInterval:
			values[i] = new(sql.NullInt64)
		case devicetoken.FieldDeviceCode, devicetoken.FieldStatus, devicetoken.FieldCodeChallenge, devicetoken.FieldCodeChallengeMethod:
			values[i] = new(sql.NullString)
		case devicetoken.FieldExpiry, devicetoken.FieldLastRequest:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type DeviceToken", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DeviceToken fields.
func (dt *DeviceToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case devicetoken.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dt.ID = int(value.Int64)
		case devicetoken.FieldDeviceCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device_code", values[i])
			} else if value.Valid {
				dt.DeviceCode = value.String
			}
		case devicetoken.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				dt.Status = value.String
			}
		case devicetoken.FieldToken:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value != nil {
				dt.Token = value
			}
		case devicetoken.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				dt.Expiry = value.Time
			}
		case devicetoken.FieldLastRequest:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_request", values[i])
			} else if value.Valid {
				dt.LastRequest = value.Time
			}
		case devicetoken.FieldPollInterval:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field poll_interval", values[i])
			} else if value.Valid {
				dt.PollInterval = int(value.Int64)
			}
		case devicetoken.FieldCodeChallenge:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge", values[i])
			} else if value.Valid {
				dt.CodeChallenge = value.String
			}
		case devicetoken.FieldCodeChallengeMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code_challenge_method", values[i])
			} else if value.Valid {
				dt.CodeChallengeMethod = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this DeviceToken.
// Note that you need to call DeviceToken.Unwrap() before calling this method if this DeviceToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (dt *DeviceToken) Update() *DeviceTokenUpdateOne {
	return (&DeviceTokenClient{config: dt.config}).UpdateOne(dt)
}

// Unwrap unwraps the DeviceToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dt *DeviceToken) Unwrap() *DeviceToken {
	_tx, ok := dt.config.driver.(*txDriver)
	if !ok {
		panic("db: DeviceToken is not a transactional entity")
	}
	dt.config.driver = _tx.drv
	return dt
}

// String implements the fmt.Stringer.
func (dt *DeviceToken) String() string {
	var builder strings.Builder
	builder.WriteString("DeviceToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dt.ID))
	builder.WriteString("device_code=")
	builder.WriteString(dt.DeviceCode)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(dt.Status)
	builder.WriteString(", ")
	if v := dt.Token; v != nil {
		builder.WriteString("token=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("expiry=")
	builder.WriteString(dt.Expiry.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_request=")
	builder.WriteString(dt.LastRequest.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("poll_interval=")
	builder.WriteString(fmt.Sprintf("%v", dt.PollInterval))
	builder.WriteString(", ")
	builder.WriteString("code_challenge=")
	builder.WriteString(dt.CodeChallenge)
	builder.WriteString(", ")
	builder.WriteString("code_challenge_method=")
	builder.WriteString(dt.CodeChallengeMethod)
	builder.WriteByte(')')
	return builder.String()
}

// DeviceTokens is a parsable slice of DeviceToken.
type DeviceTokens []*DeviceToken

func (dt DeviceTokens) config(cfg config) {
	for _i := range dt {
		dt[_i].config = cfg
	}
}
