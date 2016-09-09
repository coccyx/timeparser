package timeparser

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeParser(t *testing.T) {
	n, err := time.Parse(time.RFC822, "25 May 80 12:00 CST")
	if err != nil {
		panic(err)
	}
	now := func() time.Time {
		return n
	}

	// Test Now
	tn, _ := TimeParserNow("now", now)
	assert.Equal(t, n, tn)

	// Test -1h
	x := n.Add(time.Duration(1) * time.Hour * -1)
	tn, _ = TimeParserNow("-1h", now)
	assert.Equal(t, x, tn)

	// Test -10s
	x = n.Add(time.Duration(10) * time.Second * -1)
	tn, _ = TimeParserNow("-10s", now)
	assert.Equal(t, x, tn)

	// Test -59m
	x = n.Add(time.Duration(59) * time.Minute * -1)
	tn, _ = TimeParserNow("-59m", now)
	assert.Equal(t, x, tn)

	// Test +1day
	x = n.AddDate(0, 0, 1)
	tn, _ = TimeParserNow("+1day", now)
	assert.Equal(t, x, tn)
}
