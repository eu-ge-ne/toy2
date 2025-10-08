package textbuf

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/std"
)

func TestInsertIntoTheEnd(t *testing.T) {
	buf := New()

	buf.insert(buf.Count(), "Lorem")
	assert.Equal(t, "Lorem",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 5, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " ipsum")
	assert.Equal(t, "Lorem ipsum",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 11, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " dolor")
	assert.Equal(t, "Lorem ipsum dolor",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 17, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " sit")
	assert.Equal(t, "Lorem ipsum dolor sit",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 21, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " amet,")
	assert.Equal(t, "Lorem ipsum dolor sit amet,",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 27, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " consectetur")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 39, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " adipiscing")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 50, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " elit,")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit,",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 56, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " sed")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 60, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " do")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 63, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " eiusmod")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 71, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " tempor")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 78, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " incididunt")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 89, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " ut")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 92, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " labore")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 99, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " et")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 102, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " dolore")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 109, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " magna")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 115, buf.Count())
	buf.Validate()

	buf.insert(buf.Count(), " aliqua.")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertIntoTheBeginning(t *testing.T) {
	buf := New()

	buf.insert(0, " aliqua.")
	assert.Equal(t, " aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 8, buf.Count())
	buf.Validate()

	buf.insert(0, " magna")
	assert.Equal(t, " magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 14, buf.Count())
	buf.Validate()

	buf.insert(0, " dolore")
	assert.Equal(t, " dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 21, buf.Count())
	buf.Validate()

	buf.insert(0, " et")
	assert.Equal(t, " et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 24, buf.Count())
	buf.Validate()

	buf.insert(0, " labore")
	assert.Equal(t, " labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 31, buf.Count())
	buf.Validate()

	buf.insert(0, " ut")
	assert.Equal(t, " ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 34, buf.Count())
	buf.Validate()

	buf.insert(0, " incididunt")
	assert.Equal(t, " incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 45, buf.Count())
	buf.Validate()

	buf.insert(0, " tempor")
	assert.Equal(t, " tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 52, buf.Count())
	buf.Validate()

	buf.insert(0, " eiusmod")
	assert.Equal(t, " eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 60, buf.Count())
	buf.Validate()

	buf.insert(0, " do")
	assert.Equal(t, " do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 63, buf.Count())
	buf.Validate()

	buf.insert(0, " sed")
	assert.Equal(t, " sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 67, buf.Count())
	buf.Validate()

	buf.insert(0, " elit,")
	assert.Equal(t, " elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 73, buf.Count())
	buf.Validate()

	buf.insert(0, " adipiscing")
	assert.Equal(t, " adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 84, buf.Count())
	buf.Validate()

	buf.insert(0, " consectetur")
	assert.Equal(t, " consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 96, buf.Count())
	buf.Validate()

	buf.insert(0, " amet,")
	assert.Equal(t, " amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 102, buf.Count())
	buf.Validate()

	buf.insert(0, " sit")
	assert.Equal(t, " sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 106, buf.Count())
	buf.Validate()

	buf.insert(0, " dolor")
	assert.Equal(t, " dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 112, buf.Count())
	buf.Validate()

	buf.insert(0, " ipsum")
	assert.Equal(t, " ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 118, buf.Count())
	buf.Validate()

	buf.insert(0, "Lorem")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertSplittingNodes(t *testing.T) {
	buf := New()

	buf.insert(0, "Lorem aliqua.")
	assert.Equal(t, "Lorem aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 13, buf.Count())
	buf.Validate()

	buf.insert(5, " ipsum magna")
	assert.Equal(t, "Lorem ipsum magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 25, buf.Count())
	buf.Validate()

	buf.insert(11, " dolor dolore")
	assert.Equal(t, "Lorem ipsum dolor dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 38, buf.Count())
	buf.Validate()

	buf.insert(17, " sit et")
	assert.Equal(t, "Lorem ipsum dolor sit et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 45, buf.Count())
	buf.Validate()

	buf.insert(21, " amet, labore")
	assert.Equal(t, "Lorem ipsum dolor sit amet, labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 58, buf.Count())
	buf.Validate()

	buf.insert(27, " consectetur ut")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 73, buf.Count())
	buf.Validate()

	buf.insert(39, " adipiscing incididunt")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 95, buf.Count())
	buf.Validate()

	buf.insert(50, " elit, tempor")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 108, buf.Count())
	buf.Validate()

	buf.insert(56, " sed eiusmod")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 120, buf.Count())
	buf.Validate()

	buf.insert(60, " do")
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertAtTheNegativeIndex(t *testing.T) {
	buf := New()

	buf.insert(0, "ipsum")
	assert.Equal(t, "ipsum", std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 5, buf.Count())
	buf.Validate()

	buf.insert(-5, " ")
	assert.Equal(t, " ipsum", std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 6, buf.Count())
	buf.Validate()

	buf.insert(-6, "Lorem")
	assert.Equal(t, "Lorem ipsum", std.IterToStr(buf.Slice(0, math.MaxInt)))
	assert.Equal(t, 11, buf.Count())
	buf.Validate()
}

func TestInsertSplittingNodeWithFixup(t *testing.T) {
	buf := New()

	buf.insert(0, "11")
	buf.insert(2, "22")

	buf.insert(2, "3")
	buf.insert(3, "3")

	buf.insert(4, "4")
	buf.insert(5, "4")

	assert.Equal(t, "11334422", std.IterToStr(buf.Slice(0, math.MaxInt)))

	buf.insert(4, "-")

	assert.Equal(t, "1133-4422", std.IterToStr(buf.Slice(0, math.MaxInt)))
	buf.Validate()
}
