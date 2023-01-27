// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package acoremodrandomsuffix

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

const (
	// AStrength is a Attribute of type Strength.
	AStrength Attribute = iota
	// AAgility is a Attribute of type Agility.
	AAgility
	// AIntellect is a Attribute of type Intellect.
	AIntellect
	// ASpirit is a Attribute of type Spirit.
	ASpirit
	// AStamina is a Attribute of type Stamina.
	AStamina
	// AAttackPower is a Attribute of type AttackPower.
	AAttackPower
	// ASpellPower is a Attribute of type SpellPower.
	ASpellPower
	// AHaste is a Attribute of type Haste.
	AHaste
	// AHit is a Attribute of type Hit.
	AHit
	// ACrit is a Attribute of type Crit.
	ACrit
	// AExpertise is a Attribute of type Expertise.
	AExpertise
	// ADefenseRating is a Attribute of type DefenseRating.
	ADefenseRating
	// ADodge is a Attribute of type Dodge.
	ADodge
	// AParry is a Attribute of type Parry.
	AParry
)

var ErrInvalidAttribute = fmt.Errorf("not a valid Attribute, try [%s]", strings.Join(_AttributeNames, ", "))

const _AttributeName = "StrengthAgilityIntellectSpiritStaminaAttackPowerSpellPowerHasteHitCritExpertiseDefenseRatingDodgeParry"

var _AttributeNames = []string{
	_AttributeName[0:8],
	_AttributeName[8:15],
	_AttributeName[15:24],
	_AttributeName[24:30],
	_AttributeName[30:37],
	_AttributeName[37:48],
	_AttributeName[48:58],
	_AttributeName[58:63],
	_AttributeName[63:66],
	_AttributeName[66:70],
	_AttributeName[70:79],
	_AttributeName[79:92],
	_AttributeName[92:97],
	_AttributeName[97:102],
}

// AttributeNames returns a list of possible string values of Attribute.
func AttributeNames() []string {
	tmp := make([]string, len(_AttributeNames))
	copy(tmp, _AttributeNames)
	return tmp
}

var _AttributeMap = map[Attribute]string{
	AStrength:      _AttributeName[0:8],
	AAgility:       _AttributeName[8:15],
	AIntellect:     _AttributeName[15:24],
	ASpirit:        _AttributeName[24:30],
	AStamina:       _AttributeName[30:37],
	AAttackPower:   _AttributeName[37:48],
	ASpellPower:    _AttributeName[48:58],
	AHaste:         _AttributeName[58:63],
	AHit:           _AttributeName[63:66],
	ACrit:          _AttributeName[66:70],
	AExpertise:     _AttributeName[70:79],
	ADefenseRating: _AttributeName[79:92],
	ADodge:         _AttributeName[92:97],
	AParry:         _AttributeName[97:102],
}

// String implements the Stringer interface.
func (x Attribute) String() string {
	if str, ok := _AttributeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("Attribute(%d)", x)
}

var _AttributeValue = map[string]Attribute{
	_AttributeName[0:8]:    AStrength,
	_AttributeName[8:15]:   AAgility,
	_AttributeName[15:24]:  AIntellect,
	_AttributeName[24:30]:  ASpirit,
	_AttributeName[30:37]:  AStamina,
	_AttributeName[37:48]:  AAttackPower,
	_AttributeName[48:58]:  ASpellPower,
	_AttributeName[58:63]:  AHaste,
	_AttributeName[63:66]:  AHit,
	_AttributeName[66:70]:  ACrit,
	_AttributeName[70:79]:  AExpertise,
	_AttributeName[79:92]:  ADefenseRating,
	_AttributeName[92:97]:  ADodge,
	_AttributeName[97:102]: AParry,
}

// ParseAttribute attempts to convert a string to a Attribute.
func ParseAttribute(name string) (Attribute, error) {
	if x, ok := _AttributeValue[name]; ok {
		return x, nil
	}
	return Attribute(0), fmt.Errorf("%s is %w", name, ErrInvalidAttribute)
}

// MarshalText implements the text marshaller method.
func (x Attribute) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *Attribute) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseAttribute(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

var errAttributeNilPtr = errors.New("value pointer is nil") // one per type for package clashes

// Scan implements the Scanner interface.
func (x *Attribute) Scan(value interface{}) (err error) {
	if value == nil {
		*x = Attribute(0)
		return
	}

	// A wider range of scannable types.
	// driver.Value values at the top of the list for expediency
	switch v := value.(type) {
	case int64:
		*x = Attribute(v)
	case string:
		*x, err = ParseAttribute(v)
	case []byte:
		*x, err = ParseAttribute(string(v))
	case Attribute:
		*x = v
	case int:
		*x = Attribute(v)
	case *Attribute:
		if v == nil {
			return errAttributeNilPtr
		}
		*x = *v
	case uint:
		*x = Attribute(v)
	case uint64:
		*x = Attribute(v)
	case *int:
		if v == nil {
			return errAttributeNilPtr
		}
		*x = Attribute(*v)
	case *int64:
		if v == nil {
			return errAttributeNilPtr
		}
		*x = Attribute(*v)
	case float64: // json marshals everything as a float64 if it's a number
		*x = Attribute(v)
	case *float64: // json marshals everything as a float64 if it's a number
		if v == nil {
			return errAttributeNilPtr
		}
		*x = Attribute(*v)
	case *uint:
		if v == nil {
			return errAttributeNilPtr
		}
		*x = Attribute(*v)
	case *uint64:
		if v == nil {
			return errAttributeNilPtr
		}
		*x = Attribute(*v)
	case *string:
		if v == nil {
			return errAttributeNilPtr
		}
		*x, err = ParseAttribute(*v)
	}

	return
}

// Value implements the driver Valuer interface.
func (x Attribute) Value() (driver.Value, error) {
	return x.String(), nil
}
