package smtp

type BodyType string

const (
	Body7Bit       BodyType = "7BIT"
	Body8BitMIME   BodyType = "8BITMIME"
	BodyBinaryMIME BodyType = "BINARYMIME"
)

// MailOptions contains custom arguments that were
// passed as an argument to the MAIL command.
type MailOptions struct {
	// Value of BODY= argument, 7BIT, 8BITMIME or BINARYMIME.
	Body BodyType

	// Size of the body. Can be 0 if not specified by client.
	Size int

	// TLS is required for the message transmission.
	//
	// The message should be rejected if it can't be transmitted
	// with TLS.
	RequireTLS bool

	// The message envelope or message header contains UTF-8-encoded strings.
	// This flag is set by SMTPUTF8-aware (RFC 6531) client.
	UTF8 bool

	// The authorization identity asserted by the message sender in decoded
	// form with angle brackets stripped.
	//
	// nil value indicates missing AUTH, non-nil empty string indicates
	// AUTH=<>.
	//
	// Defined in RFC 4954.
	Auth *string
}
