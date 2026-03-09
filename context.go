package logging

import "context"

// ctxKey is an unexported type used as the context key for logging fields.
type ctxKey struct{}

// ContextWithFields returns a copy of the parent context with the given fields stored.
// Multiple calls accumulate fields; later values overwrite earlier ones for the same key.
//
//	ctx = logging.ContextWithFields(ctx, logging.Fields{"request_id": "abc"})
//	logger.WithContext(ctx).Infof("request received")
func ContextWithFields(ctx context.Context, fields Fields) context.Context {
	existing := FieldsFromContext(ctx)
	merged := make(Fields, len(existing)+len(fields))
	for k, v := range existing {
		merged[k] = v
	}
	for k, v := range fields {
		merged[k] = v
	}
	return context.WithValue(ctx, ctxKey{}, merged)
}

// FieldsFromContext extracts logging fields from the context.
// Returns nil if no fields are stored.
func FieldsFromContext(ctx context.Context) Fields {
	if ctx == nil {
		return nil
	}
	if f, ok := ctx.Value(ctxKey{}).(Fields); ok {
		return f
	}
	return nil
}
