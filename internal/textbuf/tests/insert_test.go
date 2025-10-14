package textbuf_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/eu-ge-ne/toy2/internal/textbuf"
)

func TestInsertIntoTheEnd(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(buf.Count(), []byte("Lorem"))
	assert.Equal(t, "Lorem",
		buf.All())
	assert.Equal(t, 5, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" ipsum"))
	assert.Equal(t, "Lorem ipsum",
		buf.All())
	assert.Equal(t, 11, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" dolor"))
	assert.Equal(t, "Lorem ipsum dolor",
		buf.All())
	assert.Equal(t, 17, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" sit"))
	assert.Equal(t, "Lorem ipsum dolor sit",
		buf.All())
	assert.Equal(t, 21, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" amet,"))
	assert.Equal(t, "Lorem ipsum dolor sit amet,",
		buf.All())
	assert.Equal(t, 27, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" consectetur"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur",
		buf.All())
	assert.Equal(t, 39, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" adipiscing"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing",
		buf.All())
	assert.Equal(t, 50, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" elit,"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit,",
		buf.All())
	assert.Equal(t, 56, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" sed"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed",
		buf.All())
	assert.Equal(t, 60, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" do"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do",
		buf.All())
	assert.Equal(t, 63, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" eiusmod"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		buf.All())
	assert.Equal(t, 71, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" tempor"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
		buf.All())
	assert.Equal(t, 78, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" incididunt"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt",
		buf.All())
	assert.Equal(t, 89, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" ut"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut",
		buf.All())
	assert.Equal(t, 92, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" labore"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore",
		buf.All())
	assert.Equal(t, 99, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" et"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et",
		buf.All())
	assert.Equal(t, 102, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" dolore"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore",
		buf.All())
	assert.Equal(t, 109, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" magna"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna",
		buf.All())
	assert.Equal(t, 115, buf.Count())
	buf.Validate()

	buf.Insert(buf.Count(), []byte(" aliqua."))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertIntoTheBeginning(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, []byte(" aliqua."))
	assert.Equal(t, " aliqua.",
		buf.All())
	assert.Equal(t, 8, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" magna"))
	assert.Equal(t, " magna aliqua.",
		buf.All())
	assert.Equal(t, 14, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" dolore"))
	assert.Equal(t, " dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 21, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" et"))
	assert.Equal(t, " et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 24, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" labore"))
	assert.Equal(t, " labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 31, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" ut"))
	assert.Equal(t, " ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 34, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" incididunt"))
	assert.Equal(t, " incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 45, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" tempor"))
	assert.Equal(t, " tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 52, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" eiusmod"))
	assert.Equal(t, " eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 60, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" do"))
	assert.Equal(t, " do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 63, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" sed"))
	assert.Equal(t, " sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 67, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" elit,"))
	assert.Equal(t, " elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 73, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" adipiscing"))
	assert.Equal(t, " adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 84, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" consectetur"))
	assert.Equal(t, " consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 96, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" amet,"))
	assert.Equal(t, " amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 102, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" sit"))
	assert.Equal(t, " sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 106, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" dolor"))
	assert.Equal(t, " dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 112, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte(" ipsum"))
	assert.Equal(t, " ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 118, buf.Count())
	buf.Validate()

	buf.Insert(0, []byte("Lorem"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertSplittingNodes(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, []byte("Lorem aliqua."))
	assert.Equal(t, "Lorem aliqua.",
		buf.All())
	assert.Equal(t, 13, buf.Count())
	buf.Validate()

	buf.Insert(5, []byte(" ipsum magna"))
	assert.Equal(t, "Lorem ipsum magna aliqua.",
		buf.All())
	assert.Equal(t, 25, buf.Count())
	buf.Validate()

	buf.Insert(11, []byte(" dolor dolore"))
	assert.Equal(t, "Lorem ipsum dolor dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 38, buf.Count())
	buf.Validate()

	buf.Insert(17, []byte(" sit et"))
	assert.Equal(t, "Lorem ipsum dolor sit et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 45, buf.Count())
	buf.Validate()

	buf.Insert(21, []byte(" amet, labore"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 58, buf.Count())
	buf.Validate()

	buf.Insert(27, []byte(" consectetur ut"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 73, buf.Count())
	buf.Validate()

	buf.Insert(39, []byte(" adipiscing incididunt"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 95, buf.Count())
	buf.Validate()

	buf.Insert(50, []byte(" elit, tempor"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 108, buf.Count())
	buf.Validate()

	buf.Insert(56, []byte(" sed eiusmod"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 120, buf.Count())
	buf.Validate()

	buf.Insert(60, []byte(" do"))
	assert.Equal(t, "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		buf.All())
	assert.Equal(t, 123, buf.Count())
	buf.Validate()
}

func TestInsertAtTheNegativeIndex(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, []byte("ipsum"))
	assert.Equal(t, "ipsum", buf.All())
	assert.Equal(t, 5, buf.Count())
	buf.Validate()

	buf.Insert(-5, []byte(" "))
	assert.Equal(t, " ipsum", buf.All())
	assert.Equal(t, 6, buf.Count())
	buf.Validate()

	buf.Insert(-6, []byte("Lorem"))
	assert.Equal(t, "Lorem ipsum", buf.All())
	assert.Equal(t, 11, buf.Count())
	buf.Validate()
}

func TestInsertSplittingNodeWithFixup(t *testing.T) {
	buf := textbuf.New()

	buf.Insert(0, []byte("11"))
	buf.Insert(2, []byte("22"))

	buf.Insert(2, []byte("3"))
	buf.Insert(3, []byte("3"))

	buf.Insert(4, []byte("4"))
	buf.Insert(5, []byte("4"))

	assert.Equal(t, "11334422", buf.All())

	buf.Insert(4, []byte("-"))

	assert.Equal(t, "1133-4422", buf.All())
	buf.Validate()
}
