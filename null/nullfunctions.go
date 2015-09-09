package null

func (n *NullDouble) IsSet() (float64, bool) {
	if n == nil {
		return 0.0, false
	}

	return n.Value, n.Set
}

func (n *NullFloat) IsSet() (float32, bool) {
	if n == nil {
		return 0.0, false
	}

	return n.Value, n.Set
}

func (n *NullInt32) IsSet() (int32, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullInt64) IsSet() (int64, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullUint32) IsSet() (uint32, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullUint64) IsSet() (uint64, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSint32) IsSet() (int32, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSint64) IsSet() (int64, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullFixed32) IsSet() (uint32, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullFixed64) IsSet() (uint64, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSfixed32) IsSet() (int32, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSfixed64) IsSet() (int64, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullBool) IsSet() (bool, bool) {
	if n == nil {
		return false, false
	}

	return n.Value, n.Set
}

func (n *NullString) IsSet() (string, bool) {
	if n == nil {
		return "", false
	}

	return n.Value, n.Set
}

func (n *NullBytes) IsSet() ([]byte, bool) {
	if n == nil {
		return []byte{}, false
	}

	return n.Value, n.Set
}

// Nuller allows every one of these guys to be of Nuller type. Allows you to do batch work on them but you lose
// type safty
type Nuller interface {
	Interface() (interface{}, bool)
}

func (n *NullDouble) Interface() (interface{}, bool) {
	if n == nil {
		return 0.0, false
	}

	return n.Value, n.Set
}

func (n *NullFloat) Interface() (interface{}, bool) {
	if n == nil {
		return 0.0, false
	}

	return n.Value, n.Set
}

func (n *NullInt32) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullInt64) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullUint32) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullUint64) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSint32) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSint64) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullFixed32) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullFixed64) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSfixed32) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullSfixed64) Interface() (interface{}, bool) {
	if n == nil {
		return 0, false
	}

	return n.Value, n.Set
}

func (n *NullBool) Interface() (interface{}, bool) {
	if n == nil {
		return false, false
	}

	return n.Value, n.Set
}

func (n *NullString) Interface() (interface{}, bool) {
	if n == nil {
		return "", false
	}

	return n.Value, n.Set
}

func (n *NullBytes) Interface() (interface{}, bool) {
	if n == nil {
		return []byte{}, false
	}

	return n.Value, n.Set
}
