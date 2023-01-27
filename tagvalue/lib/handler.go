package tv

type TagValueHandler interface {
	// GetTagValue returns a struct to use for tag value marshalling and unmarshalling
	GetTagValue() (interface{}, error)

	// FromTagValue provides an unmarshalled tag-value
	FromTagValue(interface{}) error
}

type TagValuePrefix interface {
	// TagValuePrefix allows a particular tag-value struct to print any prefix, such as # --- Package: name ---
	TagValuePrefix() string
}

type ToValue interface {
	// ToTagValue converts the struct directly to a value string
	ToTagValue() (string, error)
}

type FromValue interface {
	// FromTagValue provides a string representation to convert to a struct
	FromTagValue(string) error
}
