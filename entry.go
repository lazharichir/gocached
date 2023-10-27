package gocached

import "time"

type Entry[K comparable, V any] struct {
	Key     K
	Value   V
	Written time.Time
	Options EntryOptions
}

func (entry *Entry[K, V]) IsOutdated() bool {
	if entry.IsPastExpiryDate() || entry.IsPastTTL() {
		return true
	}
	return false
}

func (entry *Entry[K, V]) IsPastExpiryDate() bool {
	return entry.Options.ExpiryDate != nil && time.Now().After(*entry.Options.ExpiryDate)
}

func (entry *Entry[K, V]) IsPastTTL() bool {
	return entry.Options.TTL != nil && time.Now().Sub(entry.Written) > *entry.Options.TTL
}

type EntryOptions struct {
	ExpiryDate *time.Time
	TTL        *time.Duration
}

type EntryFn func(*EntryOptions)

func makeEntryOptions(opts []EntryFn) EntryOptions {
	options := EntryOptions{}
	for _, optFn := range opts {
		optFn(&options)
	}
	return options
}

func WithTTL(ttl time.Duration) EntryFn {
	return func(opts *EntryOptions) {
		opts.TTL = &ttl
	}
}

func WithExpiryDate(expiryDate time.Time) EntryFn {
	return func(opts *EntryOptions) {
		opts.ExpiryDate = &expiryDate
	}
}
